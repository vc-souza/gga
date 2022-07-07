package ds

/*
	An Item implementation can be used as satellite data for an item in a gga data structure.
	The main feature of an Item is being able to provide a label for easy identification.
	Some use cases would be logging and the generation of data visualizations using tools like Graphviz.
*/
type Item interface {
	Label() string
}

// TODO: attr
type FormattingAttrs map[string]string

// TODO: docs
func (f FormattingAttrs) Accept(v FormattingVisitor) {
	v.VisitFormattingAttrs(f)
}

// TODO: docs
type Formattable struct {
	// TODO: docs
	FmtAttrs FormattingAttrs
}

// TODO: docs
func (f *Formattable) SetFormatting(attrs map[string]string) {
	f.FmtAttrs = attrs
}
