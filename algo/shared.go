package algo

import "github.com/vc-souza/gga/ds"

// EdgeTypes stores edges classified by a graph algorithm.
type EdgeTypes struct {
	Forward []ds.GEIdx
	Back    []ds.GEIdx
	Cross   []ds.GEIdx
}

// min returns the minimum value between its integer inputs.
func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
