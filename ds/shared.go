package ds

/*
	An Item implementation can be used as satellite data for an item in a gga data structure.
	The main feature of an Item is being able to provide a label for easy identification.
	Some use cases would be logging and the generation of data visualizations using tools like Graphviz.
*/
type Item interface {
	Label() string
}

// FormattingAttrs holds flat formatting attributes for any type of item in a gga data structure.
type FormattingAttrs map[string]string

// Accept accepts a formatting visitor, and guides its execution using double-dispatching.
func (f FormattingAttrs) Accept(v FormattingVisitor) {
	v.VisitFormattingAttrs(f)
}

// A Formattable holds formatting attributes, and can accept them.
type Formattable struct {
	FmtAttrs FormattingAttrs
}

// SetFormatting sets the current formatting attributes of the Formattable.
func (f *Formattable) SetFormatting(attrs map[string]string) {
	f.FmtAttrs = attrs
}
