package ds

/*
	An Item implementation can be used as satellite data for an item in a gga data structure.
	The main feature of an Item is being able to provide a label for easy identification.
	Some use cases would be logging and the generation of data visualizations using tools like Graphviz.
*/
type Item interface {
	Label() string
}

// FmtAttrs holds flat formatting attributes for any type of item in a gga data structure.
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
