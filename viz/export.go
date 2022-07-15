package viz

import (
	"fmt"
	"io"
	"strings"

	"github.com/vc-souza/gga/ds"
)

/*
Exporter implements the ds.GraphVisitor interface in order to traverse a ds.Graph
and build a sequence of lines in the DOT language. After a successful visit, these
lines can then be exported to an io.Writer by calling its Export method.

Full specification of the DOT language, by Graphviz can be found here:
https://graphviz.org/doc/info/lang.html
*/
type Exporter[V ds.Item] struct {
	Graph *ds.Graph[V]
	Lines []string
	Extra []string

	DefaultGraphFmt  ds.FmtAttrs
	DefaultVertexFmt ds.FmtAttrs
	DefaultEdgeFmt   ds.FmtAttrs

	UndirectedArrow string
	DirectedArrow   string
}

// NewExporter creates an initialized Exporter.
func NewExporter[V ds.Item](graph *ds.Graph[V]) *Exporter[V] {
	res := Exporter[V]{}

	res.Graph = graph

	res.UndirectedArrow = "--"
	res.DirectedArrow = "->"

	return &res
}

func (d *Exporter[V]) add(s ...string) {
	d.Lines = append(d.Lines, s...)
}

func (d *Exporter[V]) addDefault(pfx string, attrs ds.FmtAttrs) {
	if len(attrs) == 0 {
		return
	}

	d.add(fmt.Sprintf("%s %s", pfx, DotAttrs(attrs)))
}

func (d *Exporter[V]) addDefaults() {
	d.addDefault("graph", d.DefaultGraphFmt)
	d.addDefault("node", d.DefaultVertexFmt)
	d.addDefault("edge", d.DefaultEdgeFmt)
}

/*
AddExtra adds extra lines to the end of the DOT file, which is a feature
needed by more complex visualizations that need to control more precisely
how a graph is rendered, like adding invisible, structural edges, etc.
*/
func (d *Exporter[V]) AddExtra(s ...string) {
	d.Extra = append(d.Extra, s...)
}

// Export writes the data it has accumulated to an io.Writer.
func (d *Exporter[V]) Export(w io.Writer) {
	d.Graph.Accept(d)

	s := strings.Join(d.Lines, "\n")
	r := strings.NewReader(s)

	io.Copy(w, r)
}

func (d *Exporter[V]) VisitGraphStart(g *ds.Graph[V]) {
	var start string

	if g.Directed() {
		start = "strict digraph {"
	} else {
		start = "strict graph {"
	}

	d.add(start)
	d.addDefaults()
}

func (d *Exporter[V]) VisitGraphEnd(g *ds.Graph[V]) {
	if len(d.Extra) != 0 {
		d.add(d.Extra...)
	}

	d.add("}\n")
}

func (d *Exporter[V]) VisitVertex(v *ds.GraphVertex[V]) {
	var line string

	if len(v.Fmt) == 0 {
		line = Quoted(v.Val)
	} else {
		line = fmt.Sprintf("%s %s", Quoted(v.Val), DotAttrs(v.Fmt))
	}

	d.add(line)
}

func (d *Exporter[V]) VisitEdge(e *ds.GraphEdge[V]) {
	var line string
	var op string

	if d.Graph.Directed() {
		op = d.DirectedArrow
	} else {
		op = d.UndirectedArrow
	}

	rel := fmt.Sprintf("%s %s %s", Quoted(e.Src), op, Quoted(e.Dst))

	if e.Wt != 0 {
		e.AppendFmtAttr("label", fmt.Sprintf(" %.2f", e.Wt))
	}

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

// ResetGraphFmt resets custom formatting attributes for every vertex and edge of a graph.
func ResetGraphFmt[V ds.Item](g *ds.Graph[V]) {
	for _, vtx := range g.Verts {
		vtx.ResetFmt()
	}

	for _, es := range g.Adj {
		for _, e := range es {
			e.ResetFmt()
		}
	}
}

// Snapshot implements a shorthand for the quick export of a graph, using a theme.
func Snapshot[V ds.Item](g *ds.Graph[V], w io.Writer, t Theme) {
	ex := NewExporter(g)
	SetTheme(ex, t)
	ex.Export(w)
}

// Quoted adds quotes to the label of a ds.Item, which is useful for labels containing special characters.
func Quoted[V ds.Item](v *V) string {
	return fmt.Sprintf(`"%s"`, (*v).Label())
}
