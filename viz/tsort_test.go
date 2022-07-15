package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSortViz(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UDGDress)

	ut.AssertEqual(t, true, err == nil)

	ord, err := algo.TSort(g)

	ut.AssertEqual(t, true, err == nil)

	vi := NewTSortViz(g, ord, nil)

	eNotExistsCount := 0
	eExitsCount := 0
	eCount := 0
	vCount := 0

	vi.OnVertexRank = func(*ds.GraphVertex[ds.Text], int) {
		vCount++
	}

	vi.OnOrderEdge = func(d *ds.GraphEdge[ds.Text], b bool) {
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