package ds

import (
	"sort"
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

var size int = 100

func TestByEdgeWeight(t *testing.T) {
	edges := make([]GE, size)

	for i := 0; i < size; i++ {
		edges[i] = GE{Wt: float64(100 - i - 1)}
	}

	sort.Sort(ByEdgeWeight(edges))

	for i := 0; i < size; i++ {
		ut.Equal(t, float64(i), edges[i].Wt)
	}
}
