package algo

import "github.com/vc-souza/gga/ds"

const (
	ColorWhite = iota
	ColorGray
)

// EdgeTypes stores edges classified by a graph algorithm.
type EdgeTypes[V ds.Item] struct {
	Forward []*ds.GraphEdge[V]
	Back    []*ds.GraphEdge[V]
	Cross   []*ds.GraphEdge[V]
}
