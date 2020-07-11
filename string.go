package jsonast

import (
	"fmt"
)

// String is a string JSON value
type String interface {
	fmt.Stringer
	Value
}

type stringImpl struct {
	valueImpl
	str string
}

func newString(str string) String {
	return stringImpl{
		str: str,
	}
}

func (s stringImpl) IsString() bool {
	return true
}

func (s stringImpl) String() string {
	return s.str
}
