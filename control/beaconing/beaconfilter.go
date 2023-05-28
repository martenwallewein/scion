package beaconing

import (
	"github.com/scionproto/scion/control/beacon"
	"github.com/scionproto/scion/control/ifstate"
	"github.com/scionproto/scion/pkg/addr"
	"github.com/scionproto/scion/pkg/segment"
)

/**
 * Static beacon filter to allow transit/peering relations for cre Ases, too
**/

func filterPropagateBeacons(local addr.IA, intf *ifstate.Interface, beacons []beacon.Beacon) []beacon.Beacon {

	switch local.AS().String() {
	case "50":
		return filterPropagateBeaconsAS50(intf, beacons)
		// ...
	}

	return beacons
}

func filterPropagateBeaconsAS50(intf *ifstate.Interface, beacons []beacon.Beacon) []beacon.Beacon {
	filteredBeacons := make([]beacon.Beacon, 0)
	for _, b := range beacons {
		conflictFound := false
		for _, as := range b.Segment.ASEntries {
			// Avoid transit between AS60 and AS70
			if isFromAS(&as, "60") && isToAS(intf, "70") {
				conflictFound = true
				break
			}
			if isFromAS(&as, "70") && isToAS(intf, "60") {
				conflictFound = true
				break
			}
		}

		if !conflictFound {
			filteredBeacons = append(filteredBeacons, b)
		}
	}

	return filteredBeacons
}

func isToAS(intf *ifstate.Interface, asn string) bool {
	return intf.TopoInfo().IA.AS().String() == asn
}

func isFromAS(as *segment.ASEntry, asn string) bool {
	return as.Local.AS().String() == asn
}
