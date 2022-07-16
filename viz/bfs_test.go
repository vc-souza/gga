package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestBFSViz(t *testing.T) {
	cases := []struct {
		desc     string
		input    string
		src      string
		expectUV int
		expectTV int
		expectSV int
		expectTE int
	}{
		{
			desc:     "graph",
			input:    ut.UUGSimple,
			src:      "u",
			expectUV: 0,
			expectTV: 8,
			expectSV: 1,
			expectTE: 7,
		},
		{
			desc:     "digraph",
			input:    ut.UDGSimple,
			src:      "3",
			expectUV: 1,
			expectTV: 5,
			expectSV: 1,
			expectTE: 4,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, vars, err := ds.NewTextParser().Parse(tc.input)

			ut.AssertEQ(t, true, err == nil)

			src := vars[tc.src]

			tree, err := algo.BFS(g, src)

			ut.AssertEQ(t, true, err == nil)

			uvCount := 0
			tvCount := 0
			svCount := 0
			teCount := 0

			vi := NewBFSViz(g, tree, src, nil)

			vi.OnUnVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFNode[ds.Text]) { uvCount++ }

			vi.OnTreeVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFNode[ds.Text]) { tvCount++ }

			vi.OnSourceVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFNode[ds.Text]) { svCount++ }

			vi.OnTreeEdge = func(d *ds.GraphEdge[ds.Text]) { teCount++ }

			ExportViz[ds.Text](vi, ut.DummyWriter{})

			ut.AssertEQ(t, tc.expectUV, uvCount)
			ut.AssertEQ(t, tc.expectTV, tvCount)
			ut.AssertEQ(t, tc.expectSV, svCount)
			ut.AssertEQ(t, tc.expectTE, teCount)
		})
	}
}
