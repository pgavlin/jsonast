package jsonast

// Array is an array JSON value
type Array interface {
	Value

	// Elts returns all of the elements in the array
	Elts() []Value
}

type arrayImpl struct {
	valueImpl
	arr []Value
}

func newArray(elts []Value) Array {
	return arrayImpl{
		arr: elts,
	}
}

func (v arrayImpl) IsArray() bool {
	return true
}

func (v arrayImpl) Elts() []Value {
	return v.arr
}
