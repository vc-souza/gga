package ds

/*
A GraphVisitor implements the Visitor pattern (https://en.wikipedia.org/wiki/Visitor_pattern) for gga graphs.
A Visitor declares methods that are executed at specific points during the traversal of a data structure.
This way, multiple behaviors can be attached to the data structure without having to modify it directly.
*/
type GraphVisitor[V Item] interface {
	VisitGraphStart(g *Graph[V])
	VisitGraphEnd(g *Graph[V])
	VisitVertex(v *GraphVertex[V])
	VisitEdge(e *GraphEdge[V])
}
