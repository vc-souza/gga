package algo

import (
	"errors"
	"strings"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func groupTag(ig *ds.Group[ds.Text]) string {
	s := make([]string, 0, len(ig.Items))

	for _, item := range ig.Items {
		s = append(s, item.Label())
	}

	return strings.Join(s, ",")
}

func TestGSCC_directed(t *testing.T) {
	expect := []string{
		"w,v,s",
		"z,x",
		"y,t,q",
		"u",
		"r",
	}

	g, _, err := ds.Parse(ut.UDGDeps)

	ut.Equal(t, true, err == nil)

	gscc, _, err := GSCC(g)

	ut.Equal(t, true, err == nil)

	ut.Equal(t, 5, gscc.VertexCount())
	ut.Equal(t, 5, gscc.EdgeCount())

	for i := range expect {
		ut.Equal(t, expect[i], groupTag(gscc.V[i].Ptr))
	}
}

func TestGSCC_undirected(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGSimple)

	ut.Equal(t, true, err == nil)

	_, _, err = GSCC(g)

	ut.Equal(t, true, err != nil)
	ut.Equal(t, true, errors.Is(err, ds.ErrUndefOp))
}
