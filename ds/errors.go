package ds

import (
	"errors"
	"fmt"
)

var ErrUndefOp = errors.New("undefined operation for this data structure")

var ErrNotExists = errors.New("element does not exist")

var ErrExists = errors.New("element already exists")

var ErrNilArg = errors.New("received nil argument")

var ErrInvalidLoop = errors.New("invalid loop")

var ErrInvalidType = errors.New("invalid element type")

type ErrInvalidSer struct {
	Reason error
}

func (e ErrInvalidSer) Error() string {
	return fmt.Sprintf("invalid serialization: %s", e.Reason.Error())
}

func (e ErrInvalidSer) Unwrap() error {
	return e.Reason
}
