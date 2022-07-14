package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestDFSViz(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		expectTV int
		expectRV int
		expectTE int
		expectFE int
		expectBE int
		expectCE int
	}{
		{
			desc:     "graph",
			input:    ut.BasicUUG,
			expectTV: 8,
			expectRV: 1,
			expectTE: 7,
			expectFE: 0,
			expectBE: 6,
			expectCE: 0,
		},
		{
			desc:     "digraph",
			input:    ut.BasicUDG,
			expectTV: 6,
			expectRV: 2,
			expectTE: 4,
			expectFE: 1,
			expectBE: 2,
			expectCE: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.NewTextParser().Parse(tc.input)

			ut.AssertEqual(t, true, err == nil)

			fst, tps, err := algo.DFS(g)

			ut.AssertEqual(t, true, err == nil)

			tvCount := 0
			rvCount := 0
			teCount := 0
			feCount := 0
			beCount := 0
			ceCount := 0

			vi := NewDFSViz(g, fst, tps)

			vi.OnTreeVertex = func(*ds.GraphVertex[ds.Text], *algo.DFSNode[ds.Text]) { tvCount++ }
			vi.OnRootVertex = func(*ds.GraphVertex[ds.Text], *algo.DFSNode[ds.Text]) { rvCount++ }

			vi.OnTreeEdge = func(*ds.GraphEdge[ds.Text]) { teCount++ }
			vi.OnForwardEdge = func(*ds.GraphEdge[ds.Text]) { feCount++ }
			vi.OnBackEdge = func(*ds.GraphEdge[ds.Text]) { beCount++ }
			vi.OnCrossEdge = func(*ds.GraphEdge[ds.Text]) { ceCount++ }

			vi.Export(ut.DummyWriter{})

			ut.AssertEqual(t, tc.expectTV, tvCount)
			ut.AssertEqual(t, tc.expectRV, rvCount)
			ut.AssertEqual(t, tc.expectTE, teCount)
			ut.AssertEqual(t, tc.expectFE, feCount)
			ut.AssertEqual(t, tc.expectBE, beCount)
			ut.AssertEqual(t, tc.expectCE, ceCount)
		})
	}
}