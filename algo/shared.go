package algo

import "github.com/vc-souza/gga/ds"

// EdgeTypes stores edges classified by a graph algorithm.
type EdgeTypes[T ds.Item] struct {
	Forward []*ds.GE[T]
	Back    []*ds.GE[T]
	Cross   []*ds.GE[T]
}

// min returns the minimum value between its integer inputs.
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
