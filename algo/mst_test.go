package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

type expectedMSTEdge struct {
	src string
	dst string
	wt  float64
}

var mstCases = []struct {
	desc   string
	algo   MSTAlgo[ds.Text]
	expect []expectedMSTEdge
}{
	{
		desc: "Kruskal",
		algo: MSTKruskal[ds.Text],
		expect: []expectedMSTEdge{
			{"g", "h", 1},
			{"c", "i", 2},
			{"f", "g", 2},
			{"a", "b", 4},
			{"c", "f", 4},
			{"c", "d", 7},
			{"a", "h", 8},
			{"d", "e", 9},
		},
	},
	{
		desc: "Prim",
		algo: MSTPrim[ds.Text],
		expect: []expectedMSTEdge{
			{"a", "b", 4},
			{"a", "h", 8},
			{"h", "g", 1},
			{"g", "f", 2},
			{"f", "c", 4},
			{"c", "i", 2},
			{"c", "d", 7},
			{"d", "e", 9},
		},
	},
}

func TestMST_directed(t *testing.T) {
	for _, tc := range mstCases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.Parse(ut.UDGDeps)

			ut.Equal(t, true, err == nil)

			_, err = tc.algo(g)

			ut.Equal(t, true, err != nil)
			ut.Equal(t, true, errors.Is(err, ds.ErrUndefOp))
		})
	}
}

func TestMST_undirected(t *testing.T) {
	for _, tc := range mstCases {
		t.Run(tc.desc, func(t *testing.T) {
			g, vars, err := ds.Parse(ut.WUGSimple)

			ut.Equal(t, true, err == nil)

			mst, err := tc.algo(g)

			ut.Equal(t, true, err == nil)

			ut.Equal(t, g.VertexCount()-1, len(mst))

			for i := 0; i < len(mst); i++ {
				ut.Equal(t, vars[tc.expect[i].src], mst[i].Src)
				ut.Equal(t, vars[tc.expect[i].dst], mst[i].Dst)
				ut.Equal(t, tc.expect[i].wt, mst[i].Wt)
			}
		})
	}
}
