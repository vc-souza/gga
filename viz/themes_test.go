package viz

import (
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

type testTheme struct{}

func (t testTheme) SetGraphFmt(attrs ds.FmtAttrs) {
	attrs["gattr"] = "test"
}

func (t testTheme) SetVertexFmt(attrs ds.FmtAttrs) {
	attrs["vattr"] = "test"
}

func (t testTheme) SetEdgeFmt(attrs ds.FmtAttrs) {
	attrs["eattr"] = "test"
}

func TestSetTheme(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UDGSimple)

	ut.AssertEqual(t, true, err == nil)

	e := NewExporter(g)

	e.DefaultGraphFmt = make(ds.FmtAttrs)
	e.DefaultGraphFmt["gattr"] = "init"
	ut.AssertEqual(t, "init", e.DefaultGraphFmt["gattr"])

	e.DefaultVertexFmt = make(ds.FmtAttrs)
	e.DefaultVertexFmt["vattr"] = "init"
	ut.AssertEqual(t, "init", e.DefaultVertexFmt["vattr"])

	e.DefaultEdgeFmt = make(ds.FmtAttrs)
	e.DefaultEdgeFmt["eattr"] = "init"
	ut.AssertEqual(t, "init", e.DefaultEdgeFmt["eattr"])

	SetTheme(e, testTheme{})

	ut.AssertEqual(t, "test", e.DefaultGraphFmt["gattr"])
	ut.AssertEqual(t, "test", e.DefaultVertexFmt["vattr"])
	ut.AssertEqual(t, "test", e.DefaultEdgeFmt["eattr"])
}

func TestThemes(t *testing.T) {
	cases := []struct {
		desc    string
		theme   Theme
		gEdited bool
		vEdited bool
		eEdited bool
	}{
		{
			desc:    "Light Breeze",
			theme:   Themes.LightBreeze,
			gEdited: true,
			vEdited: true,
			eEdited: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g, _, err := ds.NewTextParser().Parse(ut.UDGSimple)

			ut.AssertEqual(t, true, err == nil)

			ex := NewExporter(g)

			SetTheme(ex, tc.theme)

			ut.AssertEqual(t, tc.gEdited, len(ex.DefaultGraphFmt) >= 0)
			ut.AssertEqual(t, tc.vEdited, len(ex.DefaultVertexFmt) >= 0)
			ut.AssertEqual(t, tc.eEdited, len(ex.DefaultEdgeFmt) >= 0)
		})
	}
}
