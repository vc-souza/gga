package ds

/*
A GraphVisitor implements the Visitor pattern (https://en.wikipedia.org/wiki/Visitor_pattern) for gga graphs.
A Visitor declares methods that are executed at specific points during the traversal of a data structure.
This way, multiple behaviors can be attached to the data structure without having to modify it directly.
*/
type GraphVisitor[T Item] interface {
	// VisitGraphStart is called at the beginning of the graph visit.
	VisitGraphStart(g *G[T])

	// VisitGraphEnd is called at the end of the graph visit.
	VisitGraphEnd(g *G[T])

	// VisitVertex is called when visiting a graph vertex.
	VisitVertex(v *GV[T])

	// VisitEdge is called when visiting a graph edge.
	VisitEdge(e *GE[T])
}

// TODO: docs
type GraphVisitor2 interface {
	// VisitGraphStart is called at the beginning of the graph visit.
	VisitGraphStart(g *G2)

	// VisitGraphEnd is called at the end of the graph visit.
	VisitGraphEnd(g *G2)

	// VisitVertex is called when visiting a graph vertex.
	VisitVertex(v *GV2)

	// VisitEdge is called when visiting a graph edge.
	VisitEdge(e *GE2)
}
