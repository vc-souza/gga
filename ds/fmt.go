package ds

//  FmtAttrs holds flat formatting attributes for any type of item in a gga data structure.
type FmtAttrs map[string]string

// A Formattable holds formatting attributes, and can accept them.
type Formattable struct {
	Fmt FmtAttrs
}

// SetFmtAttr records the formatting (attribute, value) pair.
func (f *Formattable) SetFmtAttr(k, v string) {
	if f.Fmt == nil {
		f.Fmt = make(FmtAttrs)
	}

	f.Fmt[k] = v
}

/*
AppendFmtAttr appends a new value to an existing formatting attribute.
If the attribute hasn't been set yet, it is then set using the value.
*/
func (f *Formattable) AppendFmtAttr(k, v string) {
	old, ok := f.Fmt[k]

	if ok {
		f.SetFmtAttr(k, old+v)
	} else {
		f.SetFmtAttr(k, v)
	}
}

// ResetFmt resets the current formatting attributes, if any.
func (f *Formattable) ResetFmt() {
	if f.Fmt == nil {
		return
	}

	f.Fmt = make(FmtAttrs)
}
