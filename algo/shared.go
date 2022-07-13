package algo

import "github.com/vc-souza/gga/ds"

const (
	ColorWhite = iota
	ColorGray
	ColorBlack
)

// TODO: docs
type EdgeTypes[V ds.Item] struct {
	Forward []*ds.GraphEdge[V]
	Cross   []*ds.GraphEdge[V]
	Back    []*ds.GraphEdge[V]
}
