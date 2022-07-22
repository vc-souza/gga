package ds

import (
	"errors"
	"fmt"
)

var ErrUndefOp = errors.New("undefined operation")

var ErrDirected = WrapErr(ErrUndefOp, "directed graph")

var ErrUndirected = WrapErr(ErrUndefOp, "undirected graph")

var ErrDoesNotExist = errors.New("does not exist")

var ErrNoVtx = WrapErr(ErrDoesNotExist, "vertex")

var ErrNoEdge = WrapErr(ErrDoesNotExist, "edge")

var ErrNoRevEdge = WrapErr(ErrDoesNotExist, "reverse edge")

var ErrExists = errors.New("already exists")

var ErrNilArg = errors.New("nil argument")

var ErrInvLoop = errors.New("invalid loop")

var ErrInvType = errors.New("invalid type")

// TODO: docs
func WrapErr(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}
