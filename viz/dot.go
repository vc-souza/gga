package viz

import (
	"fmt"
	"io"
	"strings"

	"github.com/vc-souza/gga/ds"
)

const (
	DigraphArrow = "->"
	GraphArrow   = "--"
)

/*
	DotExporter implements the ds.GraphVisitor interface in order to traverse a ds.Graph
	and build a sequence of lines in the DOT language. After a successful visit, these
	lines can then be exported to an io.Writer by calling its Export method.

	Full specification of the DOT language, by Graphviz can be found here:
	https://graphviz.org/doc/info/lang.html
*/
type DotExporter[V ds.Item] struct {
	Graph *ds.Graph[V]
	Lines []string

	DefaultGraphFmt  ds.FmtAttrs
	DefaultVertexFmt ds.FmtAttrs
	DefaultEdgeFmt   ds.FmtAttrs
}

// NewDotExporter creates an initialized DotExporter.
func NewDotExporter[V ds.Item](graph *ds.Graph[V]) *DotExporter[V] {
	res := DotExporter[V]{}

	res.Graph = graph
	res.Lines = []string{}

	return &res
}

func (d *DotExporter[V]) add(s string) {
	d.Lines = append(d.Lines, s)
}

func (d *DotExporter[V]) addDefault(pfx string, attrs ds.FmtAttrs) {
	if len(attrs) == 0 {
		return
	}

	d.add(fmt.Sprintf("%s %s", pfx, DotAttrs(attrs)))
}

func (d *DotExporter[V]) addDefaults() {
	d.addDefault("graph", d.DefaultGraphFmt)
	d.addDefault("node", d.DefaultVertexFmt)
	d.addDefault("edge", d.DefaultEdgeFmt)
}

// Export writes the data it has accumulated to an io.Writer.
func (d *DotExporter[V]) Export(w io.Writer) {
	io.Copy(w, strings.NewReader(strings.Join(d.Lines, "\n")))
}

func (d *DotExporter[V]) VisitGraphStart(g *ds.Graph[V]) {
	var start string

	if g.Directed() {
		start = "strict digraph {"
	} else {
		start = "strict graph {"
	}

	d.add(start)
	d.addDefaults()
}

func (d *DotExporter[V]) VisitGraphEnd(g *ds.Graph[V]) {
	d.add("}\n")
}

func (d *DotExporter[V]) VisitVertex(v *ds.GraphVertex[V]) {
	var line string

	if len(v.Fmt) == 0 {
		line = v.Label()
	} else {
		line = fmt.Sprintf("%s %s", v.Label(), DotAttrs(v.Fmt))
	}

	d.add(line)
}

func (d *DotExporter[V]) VisitEdge(e *ds.GraphEdge[V]) {
	var line string
	var op string

	if d.Graph.Directed() {
		op = DigraphArrow
	} else {
		op = GraphArrow
	}

	rel := fmt.Sprintf("%s %s %s", (*e.Src).Label(), op, (*e.Dst).Label())

	if len(e.Fmt) == 0 {
		line = rel
	} else {
		line = fmt.Sprintf("%s %s", rel, DotAttrs(e.Fmt))
	}

	d.add(line)
}

/*
DotAttrs converts an object holding formatting attributes to its DOT language equivalent.
Full list of DOT attributes can be found here: https://graphviz.org/doc/info/attrs.html.
*/
func DotAttrs(f ds.FmtAttrs) string {
	if len(f) == 0 {
		return ""
	}

	s := []string{"["}

	for k, v := range f {
		s = append(s, fmt.Sprintf(`%s="%s"`, k, v))
	}

	s = append(s, "]")

	return strings.Join(s, " ")
}
