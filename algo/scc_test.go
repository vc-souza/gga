package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCCKosaraju_directed(t *testing.T) {
	expected := map[string]int{
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
	}

	g, _, err := ds.NewTextParser().Parse(ut.UDGDeps)

	ut.AssertEqual(t, true, err == nil)

	sccs, err := SCCKosaraju(g)

	ut.AssertEqual(t, true, err == nil)

	sets := map[string]int{}

	for i, scc := range sccs {
		for _, v := range scc {
			sets[v.Label()] = i + 1
		}
	}

	for k, cc := range expected {
		ut.AssertEqual(t, cc, sets[k])
	}
}

func TestSCCKosaraju_undirected(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UUGSimple)

	ut.AssertEqual(t, true, err == nil)

	_, err = SCCKosaraju(g)

	ut.AssertEqual(t, true, err != nil)
	ut.AssertEqual(t, true, errors.Is(err, ds.ErrUndefOp))
}
