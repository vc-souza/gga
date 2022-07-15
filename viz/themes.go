package viz

import "github.com/vc-souza/gga/ds"

// TODO: docs
type Theme[V ds.Item] interface {
	SetGraphFmt(*Exporter[V])
	SetVertexFmt(*Exporter[V])
	SetEdgeFmt(*Exporter[V])
}

// TODO: docs
func SetTheme[V ds.Item](e *Exporter[V], t Theme[V]) {
	if t == nil {
		return
	}

	t.SetGraphFmt(e)
	t.SetVertexFmt(e)
	t.SetEdgeFmt(e)
}

type LightBreezeTheme[V ds.Item] struct{}

func (t LightBreezeTheme[V]) SetGraphFmt(e *Exporter[V]) {
	e.DefaultGraphFmt["bgcolor"] = "#ffffff"
	e.DefaultGraphFmt["layout"] = "dot"
	e.DefaultGraphFmt["nodesep"] = "0.8"
	e.DefaultGraphFmt["ranksep"] = "0.5"
	e.DefaultGraphFmt["pad"] = "0.2"
}

func (t LightBreezeTheme[V]) SetVertexFmt(e *Exporter[V]) {
	e.DefaultVertexFmt["shape"] = "Mrecord"
	e.DefaultVertexFmt["style"] = "filled"
	e.DefaultVertexFmt["fillcolor"] = "#7289da"
	e.DefaultVertexFmt["fontcolor"] = "#ffffff"
	e.DefaultVertexFmt["color"] = "#ffffff"
	e.DefaultVertexFmt["penwidth"] = "1.1"
}

func (t LightBreezeTheme[V]) SetEdgeFmt(e *Exporter[V]) {
	e.DefaultEdgeFmt["penwidth"] = "0.9"
	e.DefaultEdgeFmt["arrowsize"] = "0.8"
}
