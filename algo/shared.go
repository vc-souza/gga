package algo

import "github.com/vc-souza/gga/ds"

// EdgeTypes stores edges classified by a graph algorithm.
type EdgeTypes[V ds.Item] struct {
	Forward []*ds.GE[V]
	Back    []*ds.GE[V]
	Cross   []*ds.GE[V]
}

/*
iDFS stores attributes used by any algorithm that is based
on an iterative version of a DFS (Depth-First Search).
*/
type iDFS struct {
	visited bool
	next    int
}

// min returns the minimum value between its integer inputs.
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
