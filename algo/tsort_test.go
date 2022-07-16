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

	ut.AssertEQ(t, true, err == nil)

	ord, err := TSort(g)

	ut.AssertEQ(t, true, err == nil)

	for i, e := 0, ord.Front(); e != nil; i, e = i+1, e.Next() {
		val, ok := e.Value.(*ds.Text)

		ut.AssertEQ(t, true, ok)

		ut.AssertEQ(t, vars[expect[i]], val)
	}
}

func TestTSort_undirected(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UUGSimple)

	ut.AssertEQ(t, true, err == nil)

	_, err = TSort(g)

	ut.AssertEQ(t, true, err != nil)
	ut.AssertEQ(t, true, errors.Is(err, ds.ErrUndefOp))
}
