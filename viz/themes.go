package viz

import "github.com/vc-souza/gga/ds"

var Themes = struct {
	LightBreeze LightBreezeTheme
}{}

// A Theme implementation is capable of setting default formatting for a graph, its vertices and edges.
type Theme interface {
	SetGraphFmt(ds.FAttrs)
	SetVertexFmt(ds.FAttrs)
	SetEdgeFmt(ds.FAttrs)
}

// SetTheme sets the default formatting of an exporter using a Theme.
func SetTheme[V ds.Item](e *Exporter[V], t Theme) {
	if t == nil {
		return
	}

	e.DefaultGraphFmt = make(ds.FAttrs)
	t.SetGraphFmt(e.DefaultGraphFmt)

	e.DefaultVertexFmt = make(ds.FAttrs)
	t.SetVertexFmt(e.DefaultVertexFmt)

	e.DefaultEdgeFmt = make(ds.FAttrs)
	t.SetEdgeFmt(e.DefaultEdgeFmt)
}

type LightBreezeTheme struct{}

func (t LightBreezeTheme) SetGraphFmt(attrs ds.FAttrs) {
	attrs["bgcolor"] = "#ffffff"
	attrs["layout"] = "dot"
	attrs["nodesep"] = "0.8"
	attrs["ranksep"] = "0.5"
	attrs["pad"] = "0.2"
}

func (t LightBreezeTheme) SetVertexFmt(attrs ds.FAttrs) {
	attrs["shape"] = "Mrecord"
	attrs["style"] = "filled"
	attrs["fillcolor"] = "#7289da"
	attrs["fontcolor"] = "#ffffff"
	attrs["color"] = "#ffffff"
	attrs["penwidth"] = "1.1"
}

func (t LightBreezeTheme) SetEdgeFmt(attrs ds.FAttrs) {
	attrs["penwidth"] = "0.9"
	attrs["arrowsize"] = "0.8"
}
