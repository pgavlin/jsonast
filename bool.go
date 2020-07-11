package jsonast

// Bool is a JSON bool type
type Bool interface {
	Value

	True() bool
}

type boolImpl struct {
	valueImpl
	val bool
}

func newBool(b bool) Bool {
	return boolImpl{
		val: b,
	}
}

func (b boolImpl) IsBool() bool {
	return true
}

func (b boolImpl) True() bool {
	return b.val
}
