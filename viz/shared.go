package viz

import "github.com/vc-souza/gga/ds"

/*
AlgoViz holds default formatting arguments, to be passed to an exporter.
It can be used by other *Viz types that need such functionality.
*/
type AlgoViz struct {
	DefaultGraphFmt  ds.FmtAttrs
	DefaultVertexFmt ds.FmtAttrs
	DefaultEdgeFmt   ds.FmtAttrs
}
