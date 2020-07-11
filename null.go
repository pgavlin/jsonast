package jsonast

// Null is a null JSON value
var Null Value = null{}

type null struct {
	valueImpl
}

func (null) IsNull() bool {
	return true
}
