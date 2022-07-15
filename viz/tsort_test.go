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

	count := 0

	vi.OnVertexRank = func(*ds.GraphVertex[ds.Text], int) {
		count++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.AssertEqual(t, g.VertexCount(), count)
}
