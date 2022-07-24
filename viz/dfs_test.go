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

			ut.Nil(t, err)

			fst, tps, err := algo.DFS(g, true)

			ut.Nil(t, err)

			tvCount := 0
			rvCount := 0
			teCount := 0
			feCount := 0
			beCount := 0
			ceCount := 0

			vi := NewDFSViz(g, fst, tps, nil)

			vi.OnTreeVertex = func(int, algo.DFNode) { tvCount++ }
			vi.OnRootVertex = func(int, algo.DFNode) { rvCount++ }

			vi.OnTreeEdge = func(int, int) { teCount++ }
			vi.OnForwardEdge = func(int, int) { feCount++ }
			vi.OnBackEdge = func(int, int) { beCount++ }
			vi.OnCrossEdge = func(int, int) { ceCount++ }

			ExportViz(vi, ut.DummyWriter{})

			ut.Equal(t, tc.expectTV, tvCount)
			ut.Equal(t, tc.expectRV, rvCount)
			ut.Equal(t, tc.expectTE, teCount)
			ut.Equal(t, tc.expectFE, feCount)
			ut.Equal(t, tc.expectBE, beCount)
			ut.Equal(t, tc.expectCE, ceCount)
		})
	}
}
