package ds

// TODO: docs
type FormattingVisitor interface {
	// TODO: docs
	VisitFormattingAttrs(fmt FormattingAttrs)
}

// TODO: docs
type GraphVisitor[V Item] interface {
	FormattingVisitor

	// TODO: docs
	VisitGraphStart(g *Graph[V])

	// TODO: docs
	VisitGraphEnd(g *Graph[V])

	// TODO: docs
	VisitVertex(v *GraphVertex[V])

	// TODO: docs
	VisitEdge(e *GraphEdge[V])
}
