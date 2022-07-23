package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCC_directed(t *testing.T) {
	cases := []struct {
		desc   string
		algo   SCCAlgo
		expect map[string]int
	}{
		{
			desc: "Kosaraju",
			algo: SCCKosaraju,
			expect: map[string]int{
				"q": 2,
				"r": 0,
				"s": 4,
				"t": 2,
				"u": 1,
				"v": 4,
				"x": 3,
				"y": 2,
				"w": 4,
				"z": 3,
			},
		},
		{
			desc: "Tarjan",
			algo: SCCTarjan,
			expect: map[string]int{
				"q": 2,
				"r": 4,
				"s": 0,
				"t": 2,
				"u": 3,
				"v": 0,
				"x": 1,
				"y": 2,
				"w": 0,
				"z": 1,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, idx, err := ds.Parse(ut.UDGDeps)

			ut.Nil(t, err)

			sccs, err := tc.algo(g)

			ut.Nil(t, err)

			sets := map[int]int{}

			for i, scc := range sccs {
				for _, v := range scc {
					sets[v] = i
				}
			}

			for k, cc := range tc.expect {
				ut.Equal(t, cc, sets[idx(k)])
			}
		})
	}
}

func TestSCC_undirected(t *testing.T) {
	cases := []struct {
		desc string
		algo SCCAlgo
	}{
		{
			desc: "Kosaraju",
			algo: SCCKosaraju,
		},
		{
			desc: "Tarjan",
			algo: SCCTarjan,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.Parse(ut.UUGSimple)

			ut.Nil(t, err)

			_, err = tc.algo(g)

			ut.NotNil(t, err)
			ut.True(t, errors.Is(err, ds.ErrUndefOp))
		})
	}
}
