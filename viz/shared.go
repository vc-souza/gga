package viz

import (
	"io"

	"github.com/vc-souza/gga/ds"
)

/*
ThemedGraphViz contains data that is useful for graph algorithm visualizations.
When embedded, it also provides a good part of the AlgoViz interface for free.
*/
type ThemedGraphViz[V ds.Item] struct {
	Graph *ds.G[V]
	Extra []string
	Theme Theme
}

func (v *ThemedGraphViz[V]) GetGraph() *ds.G[V] {
	return v.Graph
}

func (v *ThemedGraphViz[V]) GetExtra() []string {
	return v.Extra
}

func (v *ThemedGraphViz[V]) GetTheme() Theme {
	return v.Theme
}

/*
An AlgoViz implementer can traverse the results of a graph algorithm,
provide its input graph, and also support theming.
*/
type AlgoViz[V ds.Item] interface {
	GetGraph() *ds.G[V]
	GetExtra() []string
	GetTheme() Theme
	Traverse() error
}

// ExportViz guides the execution of an AlgoViz implementation and then export its results.
func ExportViz[V ds.Item](vi AlgoViz[V], w io.Writer) error {
	ex := NewExporter(vi.GetGraph())

	ResetGraphFmt(vi.GetGraph())
	SetTheme(ex, vi.GetTheme())

	if err := vi.Traverse(); err != nil {
		return err
	}

	if len(vi.GetExtra()) != 0 {
		ex.AddExtra(vi.GetExtra()...)
	}

	ex.Export(w)

	return nil
}
