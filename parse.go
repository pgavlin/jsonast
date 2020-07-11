package jsonast

import (
	"encoding/json"
	"fmt"
	"io"
)

func parseArray(decoder *json.Decoder) (Value, error) {
	var elements []Value
	for decoder.More() {
		v, err := parseValue(decoder)
		if err != nil {
			return nil, err
		}
		elements = append(elements, v)
	}
	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	return newArray(elements), nil
}

func parseObject(decoder *json.Decoder) (Value, error) {
	var attributes []Attribute
	for decoder.More() {
		k, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		v, err := parseValue(decoder)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, Attribute{
			Key:   k.(string),
			Value: v,
		})
	}
	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	return newObject(attributes), nil
}

func parseValue(decoder *json.Decoder) (Value, error) {
	tok, err := decoder.Token()
	if err != nil {
		return nil, err
	}

	switch tok := tok.(type) {
	case json.Delim:
		switch tok {
		case '[':
			return parseArray(decoder)
		case '{':
			return parseObject(decoder)
		default:
			panic(fmt.Errorf("unexpected delimiter '%v'", tok))
		}
	case bool:
		return newBool(tok), nil
	case float64:
		return newNumber(tok), nil
	case string:
		return newString(tok), nil
	case nil:
		return Null, nil
	default:
		panic(fmt.Errorf("unexpected token %v of type %T", tok, tok))
	}

}

// Parse parses JSON from a Reader into a Value.
func Parse(r io.Reader) (Value, error) {
	return parseValue(json.NewDecoder(r))
}
