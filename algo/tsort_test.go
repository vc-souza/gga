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

	g, vars, err := ds.NewTextParser().Parse(ut.UDGDress)

	ut.Equal(t, true, err == nil)

	ord, err := TSort(g)

	ut.Equal(t, true, err == nil)

	for i, v := range ord {
		ut.Equal(t, vars[expect[i]], v)
	}
}

func TestTSort_undirected(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UUGSimple)

	ut.Equal(t, true, err == nil)

	_, err = TSort(g)

	ut.Equal(t, true, err != nil)
	ut.Equal(t, true, errors.Is(err, ds.ErrUndefOp))
}
