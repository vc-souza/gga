package ds

// A GraphVisitor implements generic graph visitor.
type GraphVisitor[V Item] interface {
	VisitGraphStart(g *Graph[V])
	VisitGraphEnd(g *Graph[V])
	VisitVertex(v *GraphVertex[V])
	VisitEdge(e *GraphEdge[V])
}
