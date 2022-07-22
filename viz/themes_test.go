package viz

import (
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

type testTheme struct{}

func (t testTheme) SetGraphFmt(attrs ds.FAttrs) {
	attrs["gattr"] = "test"
}

func (t testTheme) SetVertexFmt(attrs ds.FAttrs) {
	attrs["vattr"] = "test"
}

func (t testTheme) SetEdgeFmt(attrs ds.FAttrs) {
	attrs["eattr"] = "test"
}

func TestSetTheme(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGSimple)

	ut.AssertNil(t, err)

	e := NewExporter(g)

	e.DefaultGraphFmt = make(ds.FAttrs)
	e.DefaultGraphFmt["gattr"] = "init"
	ut.AssertEqual(t, "init", e.DefaultGraphFmt["gattr"])

	e.DefaultVertexFmt = make(ds.FAttrs)
	e.DefaultVertexFmt["vattr"] = "init"
	ut.AssertEqual(t, "init", e.DefaultVertexFmt["vattr"])

	e.DefaultEdgeFmt = make(ds.FAttrs)
	e.DefaultEdgeFmt["eattr"] = "init"
	ut.AssertEqual(t, "init", e.DefaultEdgeFmt["eattr"])

	SetTheme(e, testTheme{})

	ut.AssertEqual(t, "test", e.DefaultGraphFmt["gattr"])
	ut.AssertEqual(t, "test", e.DefaultVertexFmt["vattr"])
	ut.AssertEqual(t, "test", e.DefaultEdgeFmt["eattr"])
}
