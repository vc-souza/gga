package ds

/*
A GraphVisitor implements the Visitor pattern (https://en.wikipedia.org/wiki/Visitor_pattern) for gga graphs.
A Visitor declares methods that are executed at specific points during the traversal of a data structure.
This way, multiple behaviors can be attached to the data structure without having to modify it directly.
*/
type GraphVisitor[V Item] interface {
	// VisitGraphStart is called at the beginning of the graph visit.
	VisitGraphStart(g *G[V])

	// VisitGraphEnd is called at the end of the graph visit.
	VisitGraphEnd(g *G[V])

	// VisitVertex is called when visiting a graph vertex.
	VisitVertex(v *GV[V])

	// VisitEdge is called when visiting a graph edge.
	VisitEdge(e *GE[V])
}
