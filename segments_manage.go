package segment

import "github.com/myfantasy/mft"

// Manage - segment
type Manage interface {
	InSegment(key int64) (ok bool, err mft.Error)
	AddSegment(sg Segment) (err mft.Error)
	CutSegment(sg Segment) (err mft.Error)
}

// ManageRouter - segment router
type ManageRouter interface {
	AddSegmentToCluster(cluster string, sg Segment) (err mft.Error)
	CutSegmentFromCluster(cluster string, sg Segment) (err mft.Error)
	GetClustersBySegment(key int64) (clusters []string, err mft.Error)
}
