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
		algo   SCCAlgorithm
		expect map[string]int
	}{
		{
			desc: "Kosaraju",
			algo: SCCKosaraju[ds.Text],
			expect: map[string]int{
				"q": 3,
				"r": 1,
				"s": 5,
				"t": 3,
				"u": 2,
				"v": 5,
				"x": 4,
				"y": 3,
				"w": 5,
				"z": 4,
			},
		},
		{
			desc: "Tarjan",
			algo: SCCTarjan[ds.Text],
			expect: map[string]int{
				"q": 3,
				"r": 5,
				"s": 1,
				"t": 3,
				"u": 4,
				"v": 1,
				"x": 2,
				"y": 3,
				"w": 1,
				"z": 2,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.NewTextParser().Parse(ut.UDGDeps)

			ut.AssertEqual(t, true, err == nil)

			sccs, err := tc.algo(g)

			ut.AssertEqual(t, true, err == nil)

			sets := map[string]int{}

			for i, scc := range sccs {
				for _, v := range scc {
					sets[v.Label()] = i + 1
				}
			}

			for k, cc := range tc.expect {
				ut.AssertEqual(t, cc, sets[k])
			}
		})
	}
}

func TestSCC_undirected(t *testing.T) {
	cases := []struct {
		desc string
		algo SCCAlgorithm
	}{
		{
			desc: "Kosaraju",
			algo: SCCKosaraju[ds.Text],
		},
		{
			desc: "Tarjan",
			algo: SCCTarjan[ds.Text],
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.NewTextParser().Parse(ut.UUGSimple)

			ut.AssertEqual(t, true, err == nil)

			_, err = tc.algo(g)

			ut.AssertEqual(t, true, err != nil)
			ut.AssertEqual(t, true, errors.Is(err, ds.ErrUndefOp))
		})
	}
}
