package jsonast

// Number is the numeric JSON value
type Number interface {
	Value

	Float64() float64
}

type number struct {
	valueImpl
	f float64
}

func newNumber(f float64) Number {
	return number{
		f: f,
	}
}

func (n number) IsNumber() bool {
	return true
}

func (n number) Float64() float64 {
	return n.f
}
