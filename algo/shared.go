package algo

import "github.com/vc-souza/gga/ds"

const (
	ColorWhite = iota
	ColorGray
)

// TODO: docs
type EdgeTypes[V ds.Item] struct {
	Forward []*ds.GraphEdge[V]
	Back    []*ds.GraphEdge[V]
	Cross   []*ds.GraphEdge[V]
}
