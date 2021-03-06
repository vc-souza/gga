package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSortViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDress)

	ut.Nil(t, err)

	ord, err := algo.TSort(g)

	ut.Nil(t, err)

	vi := NewTSortViz(g, ord, nil)

	eNotExistsCount := 0
	eExitsCount := 0
	eCount := 0
	vCount := 0

	vi.OnVertexRank = func(int, int) {
		vCount++
	}

	vi.OnOrderEdge = func(_ int, _ int, _ int, b bool) {
		if b {
			eExitsCount++
		} else {
			eNotExistsCount++
		}

		eCount++
	}

	err = ExportViz(vi, ut.DummyWriter{})

	ut.Nil(t, err)

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, g.VertexCount()-1, eCount)
	ut.Equal(t, 4, eExitsCount)
	ut.Equal(t, 4, eNotExistsCount)
}
