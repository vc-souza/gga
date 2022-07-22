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
			input:    ut.UUGSimple,
			expectTV: 8,
			expectRV: 1,
			expectTE: 14,
			expectFE: 0,
			expectBE: 6,
			expectCE: 0,
		},
		{
			desc:     "digraph",
			input:    ut.UDGSimple,
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
			g, _, err := ds.Parse(tc.input)

			ut.AssertNil(t, err)

			fst, tps, err := algo.DFS(g, true)

			ut.AssertNil(t, err)

			tvCount := 0
			rvCount := 0
			teCount := 0
			feCount := 0
			beCount := 0
			ceCount := 0

			vi := NewDFSViz(g, fst, tps, nil)

			vi.OnTreeVertex = func(*ds.GV[ds.Text], *algo.DFNode[ds.Text]) { tvCount++ }
			vi.OnRootVertex = func(*ds.GV[ds.Text], *algo.DFNode[ds.Text]) { rvCount++ }

			vi.OnTreeEdge = func(*ds.GE[ds.Text]) { teCount++ }
			vi.OnForwardEdge = func(*ds.GE[ds.Text]) { feCount++ }
			vi.OnBackEdge = func(*ds.GE[ds.Text]) { beCount++ }
			vi.OnCrossEdge = func(*ds.GE[ds.Text]) { ceCount++ }

			ExportViz[ds.Text](vi, ut.DummyWriter{})

			ut.AssertEqual(t, tc.expectTV, tvCount)
			ut.AssertEqual(t, tc.expectRV, rvCount)
			ut.AssertEqual(t, tc.expectTE, teCount)
			ut.AssertEqual(t, tc.expectFE, feCount)
			ut.AssertEqual(t, tc.expectBE, beCount)
			ut.AssertEqual(t, tc.expectCE, ceCount)
		})
	}
}
