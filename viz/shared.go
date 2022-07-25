package viz

import (
	"io"

	"github.com/vc-souza/gga/ds"
)

/*
ThemedGraphViz contains data that is useful for graph algorithm visualizations.
When embedded, it also provides a good part of the AlgoViz interface for free.
*/
type ThemedGraphViz struct {
	Graph *ds.G
	Extra []string
	Theme Theme
}

func (v *ThemedGraphViz) GetGraph() *ds.G {
	return v.Graph
}

func (v *ThemedGraphViz) GetExtra() []string {
	return v.Extra
}

func (v *ThemedGraphViz) GetTheme() Theme {
	return v.Theme
}

/*
An AlgoViz implementer can traverse the results of a graph algorithm,
provide its input graph, and also support theming.
*/
type AlgoViz interface {
	GetGraph() *ds.G
	GetExtra() []string
	GetTheme() Theme
	Traverse() error
}

// ExportViz guides the execution of an AlgoViz implementation and then export its results.
func ExportViz(vi AlgoViz, w io.Writer) error {
	ex := NewExporter()

	ResetGraphFmt(vi.GetGraph())
	SetTheme(ex, vi.GetTheme())

	if err := vi.Traverse(); err != nil {
		return err
	}

	if len(vi.GetExtra()) != 0 {
		ex.AddExtra(vi.GetExtra()...)
	}

	ex.Export(vi.GetGraph(), w)

	return nil
}
