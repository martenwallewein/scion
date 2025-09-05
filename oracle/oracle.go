package oracle

import (
	"github.com/gin-gonic/gin"
	"github.com/scionproto/scion/pkg/addr"
	"github.com/scionproto/scion/private/segment/segfetcher"
)

type Oracle struct {
	Addr string
}

func New(addr string) *Oracle {
	return &Oracle{Addr: addr}
}

type SegmentsFilterRequest struct {
	Src      addr.IA             `json:"src" binding:"required"`
	Dst      addr.IA             `json:"dst" binding:"required"`
	Segments segfetcher.Segments `json:"segments" binding:"required"`
}

type SegmentsFilterResponse struct {
	Src      addr.IA             `json:"src" binding:"required"`
	Dst      addr.IA             `json:"dst" binding:"required"`
	Segments segfetcher.Segments `json:"segments" binding:"required"`
}

func (o *Oracle) Run() error {
	r := gin.Default()

	r.POST("/segments", func(c *gin.Context) {

		var req SegmentsFilterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		response := SegmentsFilterResponse{
			Src:      req.Src,
			Dst:      req.Dst,
			Segments: req.Segments, // Here you would apply your filtering logic
		}
		c.JSON(200, response)
	})

	return r.Run(o.Addr)
}
