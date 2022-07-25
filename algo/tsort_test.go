package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSort_directed(t *testing.T) {
	g, idx, err := ds.Parse(ut.UDGDress)

	ut.Nil(t, err)

	expect := []int{
		idx("shirt"),
		idx("tie"),
		idx("watch"),
		idx("socks"),
		idx("undershorts"),
		idx("pants"),
		idx("belt"),
		idx("jacket"),
		idx("shoes"),
	}

	ord, err := TSort(g)

	ut.Nil(t, err)

	for i, v := range ord {
		ut.Equal(t, expect[i], v)
	}
}

func TestTSort_undirected(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGSimple)

	ut.Nil(t, err)

	_, err = TSort(g)

	ut.NotNil(t, err)
	ut.True(t, errors.Is(err, ds.ErrUndefOp))
}
