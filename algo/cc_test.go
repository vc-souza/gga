package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

var ccCases = []struct {
	desc string
	algo CCAlgo[ds.Text]
}{
	{
		desc: "DFS",
		algo: CCDFS[ds.Text],
	},
	{
		desc: "Union-Find",
		algo: CCUnionFind[ds.Text],
	},
}

func TestCC_directed(t *testing.T) {
	for _, tc := range ccCases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.Parse(ut.UDGDeps)

			ut.Nil(t, err)

			_, err = tc.algo(g)

			ut.NotNil(t, err)
			ut.True(t, errors.Is(err, ds.ErrUndefOp))
		})
	}
}

func TestCC_undirected(t *testing.T) {
	expect := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 1,
		"f": 1,
		"g": 1,
		"h": 2,
		"i": 2,
		"j": 3,
	}

	for _, tc := range ccCases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.Parse(ut.UUGDisc)

			ut.Nil(t, err)

			ccs, err := tc.algo(g)

			ut.Nil(t, err)

			sets := map[string]int{}

			for i, cc := range ccs {
				for _, v := range cc {
					sets[v.Label()] = i
				}
			}

			for k, cc := range expect {
				ut.Equal(t, cc, sets[k])
			}
		})
	}
}
