package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSetFmtAttr(t *testing.T) {
	f := Formattable{}

	f.SetFmtAttr("a", "b")

	ut.Equal(t, "b", f.Fmt["a"])
}

func TestAppendFmtAttr_exists(t *testing.T) {
	f := Formattable{}

	f.SetFmtAttr("a", "b")
	f.AppendFmtAttr("a", "b")

	ut.Equal(t, "bb", f.Fmt["a"])
}

func TestAppendFmtAttr_does_not_exist(t *testing.T) {
	f := Formattable{}

	f.AppendFmtAttr("a", "b")

	ut.Equal(t, "b", f.Fmt["a"])
}

func TestResetFmt_empty(t *testing.T) {
	f := Formattable{}

	ut.Equal(t, true, f.Fmt == nil)

	f.ResetFmt()

	ut.Equal(t, true, f.Fmt == nil)
}

func TestResetFmt_not_empty(t *testing.T) {
	f := Formattable{}

	f.SetFmtAttr("a", "b")

	ut.Equal(t, 1, len(f.Fmt))

	f.ResetFmt()

	ut.Equal(t, 0, len(f.Fmt))
}
