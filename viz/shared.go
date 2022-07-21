package viz

import (
	"io"

	"github.com/vc-souza/gga/ds"
)

/*
ThemedGraphViz contains data that is useful for graph algorithm visualizations.
When embedded, it also provides a good part of the AlgoViz interface for free.
*/
type ThemedGraphViz[T ds.Item] struct {
	Graph *ds.G[T]
	Extra []string
	Theme Theme
}

func (v *ThemedGraphViz[T]) GetGraph() *ds.G[T] {
	return v.Graph
}

func (v *ThemedGraphViz[T]) GetExtra() []string {
	return v.Extra
}

func (v *ThemedGraphViz[T]) GetTheme() Theme {
	return v.Theme
}

/*
An AlgoViz implementer can traverse the results of a graph algorithm,
provide its input graph, and also support theming.
*/
type AlgoViz[T ds.Item] interface {
	GetGraph() *ds.G[T]
	GetExtra() []string
	GetTheme() Theme
	Traverse() error
}

// ExportViz guides the execution of an AlgoViz implementation and then export its results.
func ExportViz[T ds.Item](vi AlgoViz[T], w io.Writer) error {
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
