package ds

//  FAttrs holds flat formatting attributes for any type of item in a gga data structure.
type FAttrs map[string]string

// A Formattable holds formatting attributes, and can accept them.
type Formattable struct {
	F FAttrs
}

// SetFmtAttr records the formatting (attribute, value) pair.
func (f *Formattable) SetFmtAttr(k, v string) {
	if f.F == nil {
		f.F = make(FAttrs)
	}

	f.F[k] = v
}

/*
AppendFmtAttr appends a new value to an existing formatting attribute.
If the attribute hasn't been set yet, it is then set using the value.
*/
func (f *Formattable) AppendFmtAttr(k, v string) {
	old, ok := f.F[k]

	if ok {
		f.SetFmtAttr(k, old+v)
	} else {
		f.SetFmtAttr(k, v)
	}
}

// ResetFmt resets the current formatting attributes, if any.
func (f *Formattable) ResetFmt() {
	if f.F == nil {
		return
	}

	f.F = make(FAttrs)
}
