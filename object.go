package jsonast

// Object is the object JSON value
type Object interface {
	Value

	Attributes() []Attribute

	At(key string) (Value, bool)
}

// Attribute represents a single key-value pair in an object
type Attribute struct {
	Key   string
	Value Value
}

type objectImpl struct {
	valueImpl
	attributes []Attribute
}

func newObject(attributes []Attribute) Object {
	return objectImpl{
		attributes: attributes,
	}
}

func (o objectImpl) IsObject() bool {
	return true
}

func (o objectImpl) Attributes() []Attribute {
	return o.attributes
}

func (o objectImpl) At(key string) (Value, bool) {
	for _, attr := range o.attributes {
		if attr.Key == key {
			return attr.Value, true
		}
	}
	return nil, false
}
