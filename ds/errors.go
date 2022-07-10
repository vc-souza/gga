package ds

import "errors"

var ErrUndefOp = errors.New("undefined operation for this data structure")

var ErrNotExists = errors.New("element does not exist")

var ErrNilArg = errors.New("received nil argument")

var ErrLoop = errors.New("invalid loop")
