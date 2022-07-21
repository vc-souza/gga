package ds

import (
	"errors"
	"fmt"
)

var ErrUndefOp = errors.New("undefined operation for this data structure")

var ErrNotExists = errors.New("element does not exist")

var ErrVtxNotExists = fmt.Errorf("vertex: %w", ErrNotExists)

var ErrEdgeNotExists = fmt.Errorf("edge: %w", ErrNotExists)

var ErrRevEdgeNotExists = fmt.Errorf("reverse edge: %w", ErrNotExists)

var ErrExists = errors.New("element already exists")

var ErrNilArg = errors.New("received nil argument")

var ErrInvalidLoop = errors.New("invalid loop")

var ErrInvalidType = errors.New("invalid element type")
