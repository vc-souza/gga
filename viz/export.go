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
type Exporter struct {
	Lines []string
	Extra []string

	DefaultGraphFmt  ds.FAttrs
	DefaultVertexFmt ds.FAttrs
	DefaultEdgeFmt   ds.FAttrs

	UndirectedArrow string
	DirectedArrow   string
}

// NewExporter creates an initialized Exporter.
func NewExporter() *Exporter {
	res := Exporter{}

	res.UndirectedArrow = "--"
	res.DirectedArrow = "->"

	return &res
}

func (d *Exporter) add(s ...string) {
	d.Lines = append(d.Lines, s...)
}

func (d *Exporter) addDefault(pfx string, attrs ds.FAttrs) {
	if len(attrs) == 0 {
		return
	}

	d.add(fmt.Sprintf("%s%s", pfx, DotAttrs(attrs)))
}

func (d *Exporter) addDefaults() {
	d.addDefault("graph", d.DefaultGraphFmt)
	d.addDefault("node", d.DefaultVertexFmt)
	d.addDefault("edge", d.DefaultEdgeFmt)
}

/*
AddExtra adds extra lines to the end of the DOT file, which is a feature
needed by more complex visualizations that need to control more precisely
how a graph is rendered, like adding invisible, structural edges, etc.
*/
func (d *Exporter) AddExtra(s ...string) {
	d.Extra = append(d.Extra, s...)
}

// Export writes the data it has accumulated to an io.Writer.
func (d *Exporter) Export(g ds.G, w io.Writer) {
	g.Accept(d)

	s := strings.Join(d.Lines, "\n")
	r := strings.NewReader(s)

	io.Copy(w, r)
}

func (d *Exporter) VisitGraphStart(g ds.G) {
	var start string

	if g.Directed() {
		start = "strict digraph {"
	} else {
		start = "strict graph {"
	}

	d.add(start)
	d.addDefaults()
}

func (d *Exporter) VisitGraphEnd(ds.G) {
	if len(d.Extra) != 0 {
		d.add(d.Extra...)
	}

	d.add("}\n")
}

func (d *Exporter) VisitVertex(g ds.G, v int) {
	vtx := g.V[v]

	d.add(fmt.Sprintf(
		"%s%s",
		Quoted(vtx.Item),
		DotAttrs(vtx.F),
	))
}

func (d *Exporter) VisitEdge(g ds.G, v int, e int) {
	var op string

	if g.Directed() {
		op = d.DirectedArrow
	} else {
		op = d.UndirectedArrow
	}

	edge := g.V[v].E[e]
	attrs := ds.FAttrs{}

	for k, v := range edge.F {
		attrs[k] = v
	}

	if edge.Wt != 0 {
		label, ok := attrs["label"]

		if ok {
			label += " "
		}

		attrs["label"] = fmt.Sprintf("%s%.2f", label, edge.Wt)
	}

	d.add(fmt.Sprintf(
		"%s %s %s%s",
		Quoted(g.V[edge.Src].Item),
		op,
		Quoted(g.V[edge.Dst].Item),
		DotAttrs(attrs),
	))
}

/*
DotAttrs converts an object holding formatting attributes to its DOT language equivalent.
Full list of DOT attributes can be found here: https://graphviz.org/doc/info/attrs.html.
*/
func DotAttrs(f ds.FAttrs) string {
	if len(f) == 0 {
		return ""
	}

	s := []string{" ["}

	for k, v := range f {
		s = append(s, fmt.Sprintf(`%s="%s"`, k, v))
	}

	s = append(s, "]")

	return strings.Join(s, " ")
}

// ResetGraphFmt resets custom formatting attributes for every vertex and edge of a graph.
func ResetGraphFmt(g *ds.G) {
	for i := range g.V {
		g.V[i].ResetFmt()

		for j := range g.V[i].E {
			g.V[i].E[j].ResetFmt()
		}
	}
}

// Snapshot implements a shorthand for the quick export of a graph, using a theme.
func Snapshot(g ds.G, w io.Writer, t Theme) {
	ex := NewExporter()
	SetTheme(ex, t)
	ex.Export(g, w)
}

// Quoted adds quotes to the label of a ds.Item, which is useful for labels containing special characters.
func Quoted(i ds.Item) string {
	return fmt.Sprintf(`"%s"`, i.Label())
}
