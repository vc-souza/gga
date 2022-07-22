package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSort_directed(t *testing.T) {
	expect := []string{
		"shirt",
		"tie",
		"watch",
		"socks",
		"undershorts",
		"pants",
		"belt",
		"jacket",
		"shoes",
	}

	g, vars, err := ds.Parse(ut.UDGDress)

	ut.AssertNil(t, err)

	ord, err := TSort(g)

	ut.AssertNil(t, err)

	for i, v := range ord {
		ut.AssertEqual(t, vars[expect[i]], v)
	}
}

func TestTSort_undirected(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGSimple)

	ut.AssertNil(t, err)

	_, err = TSort(g)

	ut.AssertNotNil(t, err)
	ut.AssertTrue(t, errors.Is(err, ds.ErrUndefOp))
}
