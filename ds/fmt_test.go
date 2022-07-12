package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSetFmtAttr(t *testing.T) {
	f := Formattable{}

	f.SetFmtAttr("a", "b")

	ut.AssertEqual(t, "b", f.Fmt["a"])
}

func TestResetFmt_empty(t *testing.T) {
	f := Formattable{}

	ut.AssertEqual(t, true, f.Fmt == nil)

	f.ResetFmt()

	ut.AssertEqual(t, true, f.Fmt == nil)
}

func TestResetFmt_not_empty(t *testing.T) {
	f := Formattable{}

	f.SetFmtAttr("a", "b")

	ut.AssertEqual(t, 1, len(f.Fmt))

	f.ResetFmt()

	ut.AssertEqual(t, 0, len(f.Fmt))
}
