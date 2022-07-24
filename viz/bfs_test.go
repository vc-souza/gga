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
			expectTE: 14,
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
			g, idx, err := ds.Parse(tc.input)

			ut.Nil(t, err)

			src := idx(tc.src)

			tree, err := algo.BFS(g, src)

			ut.Nil(t, err)

			uvCount := 0
			tvCount := 0
			svCount := 0
			teCount := 0

			vi := NewBFSViz(g, tree, src, nil)

			vi.OnUnVertex = func(int, algo.BFNode) { uvCount++ }

			vi.OnTreeVertex = func(int, algo.BFNode) { tvCount++ }

			vi.OnSourceVertex = func(int, algo.BFNode) { svCount++ }

			vi.OnTreeEdge = func(int, int) { teCount++ }

			ExportViz(vi, ut.DummyWriter{})

			ut.Equal(t, tc.expectUV, uvCount)
			ut.Equal(t, tc.expectTV, tvCount)
			ut.Equal(t, tc.expectSV, svCount)
			ut.Equal(t, tc.expectTE, teCount)
		})
	}
}
