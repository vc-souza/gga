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
