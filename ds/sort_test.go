package ds

import (
	"sort"
	"testing"

	"github.com/vc-souza/gga/internal/testutils"
)

var size int = 100

func TestByEdgeWeight(t *testing.T) {
	edges := make([]*GE[Text], size)

	for i := 0; i < size; i++ {
		edges[i] = &GE[Text]{Wt: float64(100 - i - 1)}
	}

	sort.Sort(ByEdgeWeight[Text](edges))

	for i := 0; i < size; i++ {
		testutils.Equal(t, float64(i), edges[i].Wt)
	}
}
