package viz

import (
	"fmt"
	"io"
	"strings"

	"github.com/vc-souza/gga/ds"
)

type DotExporter[V ds.Item] struct {
	Graph *ds.Graph[V]
	Lines []string
}

// NewDotExporter creates an initialized DotExporter.
func NewDotExporter[V ds.Item](graph *ds.Graph[V]) *DotExporter[V] {
	res := DotExporter[V]{}

	res.Graph = graph
	res.Lines = []string{}

	return &res
}

// Attrs serializes a ds.FormattingAttrs value into a string.
func (d *DotExporter[V]) Attrs(f ds.FormattingAttrs) string {
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

	d.Lines = append(d.Lines, start)

	if len(g.FmtAttrs) == 0 {
		return
	}

	d.Lines = append(d.Lines, fmt.Sprintf("graph %s", d.Attrs(g.FmtAttrs)))
}

func (d *DotExporter[V]) VisitGraphEnd(g *ds.Graph[V]) {
	d.Lines = append(d.Lines, "}\n")
}

func (d *DotExporter[V]) VisitVertex(v *ds.GraphVertex[V]) {
	var line string

	if len(v.FmtAttrs) == 0 {
		line = v.Label()
	} else {
		line = fmt.Sprintf("%s %s", v.Label(), d.Attrs(v.FmtAttrs))
	}

	d.Lines = append(d.Lines, line)
}

func (d *DotExporter[V]) VisitEdge(e *ds.GraphEdge[V]) {
	var op string
	var line string

	if d.Graph.Directed() {
		op = "->"
	} else {
		op = "--"
	}

	rel := fmt.Sprintf("%s %s %s", (*e.Src).Label(), op, (*e.Dst).Label())

	if len(e.FmtAttrs) == 0 {
		line = rel
	} else {
		line = fmt.Sprintf("%s %s", rel, d.Attrs(e.FmtAttrs))
	}

	d.Lines = append(d.Lines, line)
}
