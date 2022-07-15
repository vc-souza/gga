package viz

import (
	"io"

	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type ThemedGraphViz[V ds.Item] struct {
	Graph *ds.Graph[V]
	Theme Theme
}

func (v *ThemedGraphViz[V]) GetGraph() *ds.Graph[V] {
	return v.Graph
}

func (v *ThemedGraphViz[V]) GetTheme() Theme {
	return v.Theme
}

// TODO: docs
type AlgoViz[V ds.Item] interface {
	GetGraph() *ds.Graph[V]
	GetTheme() Theme
	Traverse() error
}

// TODO: docs
func ExportViz[V ds.Item](vi AlgoViz[V], w io.Writer) error {
	ex := NewExporter(vi.GetGraph())

	ResetGraphFmt(vi.GetGraph())
	SetTheme(ex, vi.GetTheme())

	if err := vi.Traverse(); err != nil {
		return err
	}

	ex.Export(w)

	return nil
}
