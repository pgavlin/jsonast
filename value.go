package jsonast

// Value is a generic node in a JSON AST.
type Value interface {
	// IsString returns true if the data at this node is a string
	IsString() bool
	// IsNumber returns true if the data at this node is numeric
	IsNumber() bool
	// IsBool returns true if the data at this node is a bool
	IsBool() bool
	// IsObject returns true if the data at this node is an object
	IsObject() bool
	// IsArray returns true if the data at this node is an array
	IsArray() bool
	// IsNull returns true if the data at this node is null
	IsNull() bool
}

type valueImpl struct{}

func (v valueImpl) IsString() bool { return false }
func (v valueImpl) IsNumber() bool { return false }
func (v valueImpl) IsBool() bool   { return false }
func (v valueImpl) IsObject() bool { return false }
func (v valueImpl) IsArray() bool  { return false }
func (v valueImpl) IsNull() bool   { return false }
