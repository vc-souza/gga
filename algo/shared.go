package algo

import "github.com/vc-souza/gga/ds"

// EdgeTypes stores edges classified by a graph algorithm.
type EdgeTypes[V ds.Item] struct {
	Forward []*ds.GraphEdge[V]
	Back    []*ds.GraphEdge[V]
	Cross   []*ds.GraphEdge[V]
}

// Min returns the minimum value between its integer inputs.
func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
