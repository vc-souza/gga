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

	ut.AssertNil(t, err)

	gscc, _, err := GSCC(g)

	ut.AssertNil(t, err)

	ut.AssertEqual(t, 5, gscc.VertexCount())
	ut.AssertEqual(t, 5, gscc.EdgeCount())

	for i := range expect {
		ut.AssertEqual(t, expect[i], groupTag(gscc.V[i].Ptr))
	}
}

func TestGSCC_undirected(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGSimple)

	ut.AssertNil(t, err)

	_, _, err = GSCC(g)

	ut.AssertNotNil(t, err)
	ut.AssertTrue(t, errors.Is(err, ds.ErrUndefOp))
}
