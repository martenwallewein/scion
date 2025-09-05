// Copyright 2020 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package segreq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scionproto/scion/oracle"
	"github.com/scionproto/scion/pkg/addr"
	"github.com/scionproto/scion/pkg/log"
	"github.com/scionproto/scion/pkg/private/serrors"
	seg "github.com/scionproto/scion/pkg/segment"
	"github.com/scionproto/scion/private/segment/segfetcher"
)

// ForwardingLookup handles path segment lookup requests in a non-core AS. If
// segments are missing, the request is forwarded to the respective core ASes.
// It should only be used in a non-core AS.
type ForwardingLookup struct {
	LocalIA     addr.IA
	CoreChecker CoreChecker
	Fetcher     *segfetcher.Fetcher
	Expander    WildcardExpander
}

// LookupSegments looks up the segments for the given request
//
//   - requests for up segment are answered directly, from the local DB
//   - down and core segments are forwarded to the responsible core ASes,
//     and results are cached
func (f ForwardingLookup) LookupSegments(ctx context.Context, src,
	dst addr.IA) (segfetcher.Segments, error) {

	segType, err := f.classify(ctx, src, dst)
	if err != nil {
		return nil, err
	}

	reqs, err := f.Expander.ExpandSrcWildcard(ctx,
		segfetcher.Request{
			Src:     src,
			Dst:     dst,
			SegType: segType,
		},
	)
	if err != nil {
		return nil, serrors.Wrap("expanding wildcard request", err)
	}
	segments, err := f.Fetcher.Fetch(ctx, reqs, false)
	if err != nil {
		for _, seq := range segments {
			fmt.Println("segment:", seq)
			fmt.Println("  info:", seq.Segment.Info)
			fmt.Println("  Entries:", seq.Segment.ASEntries)
		}
	}

	// Call http to localhost:8181/segments with src,dst,segments
	log.Debug("LookupSegments", "src", src, "dst", dst, "type", segType,
		"num_segs", len(segments), "err", err)

	client := http.Client{}
	url := "http://localhost:8281/segments"
	reqBody := oracle.SegmentsFilterRequest{
		Src:      src,
		Dst:      dst,
		Segments: segments,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Error("Failed to marshal request body", "err", err)
		return segments, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Error("Failed to create HTTP request", "err", err)
		return segments, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Failed to send HTTP request", "err", err)
		return segments, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Received non-OK response from oracle", "status", resp.Status)
		return segments, serrors.New("non-OK response from oracle", "status", resp.Status)
	}

	log.Debug("Successfully sent segments to oracle", "num_segs", len(segments))

	// Parse response from body
	respBody := oracle.SegmentsFilterResponse{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		log.Error("Failed to decode response body", "err", err)
		return segments, err
	}

	log.Debug("Received filtered segments from oracle", "num_segs", len(respBody.Segments))
	return respBody.Segments, err
}

// classify validates the request and determines the segment type for the request
func (f ForwardingLookup) classify(ctx context.Context,
	src, dst addr.IA) (seg.Type, error) {

	if src.ISD() == 0 || dst.ISD() == 0 {
		return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
			"src", src, "dst", dst, "reason", "zero ISD src or dst")

	}
	if dst == f.LocalIA {
		// this could be an otherwise valid request, but probably the requester switched Src and Dst
		return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
			"src", src, "dst", dst, "reason", "dst is local AS, confusion?")

	}
	srcCore, err := f.CoreChecker.IsCore(ctx, src)
	if err != nil {
		return 0, err
	}
	dstCore, err := f.CoreChecker.IsCore(ctx, dst)
	if err != nil {
		return 0, err
	}
	switch {
	case srcCore && dstCore:
		// core
		if src.ISD() != f.LocalIA.ISD() {
			return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
				"src", src, "dst", dst, "reason", "core segment request src ISD not local ISD")

		}
		return seg.TypeCore, nil
	case srcCore:
		// down
		if src.ISD() != dst.ISD() {
			return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
				"src", src, "dst", dst, "reason", "down segment request src/dst in different ISD")

		}
		return seg.TypeDown, nil
	case dstCore:
		// up
		if src != f.LocalIA {
			return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
				"src", src, "dst", dst, "reason", "up segment request src not local AS")

		}
		if dst.ISD() != f.LocalIA.ISD() {
			return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
				"src", src, "dst", dst, "reason", "up segment request dst in different ISD")

		}
		return seg.TypeUp, nil
	default:
		return 0, serrors.JoinNoStack(segfetcher.ErrInvalidRequest, nil,
			"src", src, "dst", dst, "reason", "non-core src & dst")

	}
}
