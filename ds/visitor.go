package ds

// TODO: docs
type GraphVisitor[V Item] interface {
	// TODO: docs
	VisitGraphStart(g *Graph[V])

	// TODO: docs
	VisitGraphEnd(g *Graph[V])

	// TODO: docs
	VisitVertex(v *GraphVertex[V])

	// TODO: docs
	VisitEdge(e *GraphEdge[V])

	// TODO: docs
	VisitFormattingAttrs(fmt FormattingAttrs)
}
