package ds

// A FormattingVisitor implements an action on a FormattingAttrs.
type FormattingVisitor interface {
	VisitFormattingAttrs(fmt FormattingAttrs)
}

// A GraphVisitor implements generic graph visiting.
type GraphVisitor[V Item] interface {
	FormattingVisitor

	VisitGraphStart(g *Graph[V])
	VisitGraphEnd(g *Graph[V])
	VisitVertex(v *GraphVertex[V])
	VisitEdge(e *GraphEdge[V])
}
