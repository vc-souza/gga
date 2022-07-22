package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSortViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDress)

	ut.AssertNil(t, err)

	ord, err := algo.TSort(g)

	ut.AssertNil(t, err)

	vi := NewTSortViz(g, ord, nil)

	eNotExistsCount := 0
	eExitsCount := 0
	eCount := 0
	vCount := 0

	vi.OnVertexRank = func(*ds.GV[ds.Text], int) {
		vCount++
	}

	vi.OnOrderEdge = func(_ *ds.GE[ds.Text], b bool) {
		if b {
			eExitsCount++
		} else {
			eNotExistsCount++
		}

		eCount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.AssertEqual(t, g.VertexCount(), vCount)
	ut.AssertEqual(t, g.VertexCount()-1, eCount)
	ut.AssertEqual(t, 4, eExitsCount)
	ut.AssertEqual(t, 4, eNotExistsCount)
}
