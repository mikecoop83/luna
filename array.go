package json

import (
	"encoding/json"
	"fmt"
	"io"
)

// Array provides methods to either navigate through the content of a JSON array or propagate any error that has occurred
type Array struct {
	a    []interface{}
	path path
	err  error
}

// ArrayFromBytes creates an Array from a []byte
func ArrayFromBytes(jsonBytes []byte) Array {
	var a Array
	err := json.Unmarshal(jsonBytes, &a.a)
	if err != nil {
		return Array{nil, "$", err}
	}
	return a
}

// ArrayFromReader creates an Array from an io.Reader
func ArrayFromReader(r io.Reader) Array {
	var a Array
	err := json.NewDecoder(r).Decode(&a.a)
	if err != nil {
		return Array{nil, "$", err}
	}
	return a
}

// NewArray creates an Array from a []interface{}
func NewArray(a []interface{}) Array {
	return Array{a, "$", nil}
}

func (a Array) validateIndex(idx int) error {
	if idx < 0 || idx >= len(a.a) {
		return fmt.Errorf("invalid index: %d; it should be between 0 and %d", idx, len(a.a)-1)
	}
	return nil
}

// Map returns the map found at index `idx` in the array; errors will be propagated
func (a Array) Map(idx int) Map {
	if a.err != nil {
		return Map{nil, a.path, a.err}
	}
	err := a.validateIndex(idx)
	if err != nil {
		return Map{nil, a.path, err}
	}
	currPath := a.path.appendIndex(idx)
	m, ok := a.a[idx].(map[string]interface{})
	if !ok {
		return Map{
			nil,
			a.path,
			fmt.Errorf("item at path %s was a %T, not a map", currPath, a.a[idx]),
		}
	}
	return Map{m, currPath, nil}
}

// Array returns the array found at index `idx` in the array; errors will be propagated
func (a Array) Array(idx int) Array {
	if a.err != nil {
		return a
	}
	err := a.validateIndex(idx)
	if err != nil {
		return Array{nil, a.path, err}
	}
	currPath := a.path.appendIndex(idx)
	result, ok := a.a[idx].([]interface{})
	if !ok {
		return Array{
			nil,
			a.path,
			fmt.Errorf("item at path %s was a %T, not an array", currPath, a.a[idx]),
		}
	}
	return Array{result, currPath, nil}
}

// MustLen returns the length of the array, or panics if there was an error
func (a Array) MustLen() int {
	if a.err != nil {
		panic(a.err)
	}
	return len(a.a)
}

// String returns the value of a string at index `idx` in the array, or a propagated error
func (a Array) String(idx int) (string, error) {
	if a.err != nil {
		return "", a.err
	}
	err := a.validateIndex(idx)
	if err != nil {
		return "", err
	}
	s, ok := a.a[idx].(string)
	if !ok {
		return "", fmt.Errorf("item at index %d was a %T, not a string", idx, a.a[idx])
	}
	return s, nil
}

// Float returns the value of a float at index `idx` in the array, or a propagated error
func (a Array) Float(idx int) (float64, error) {
	if a.err != nil {
		return 0.0, a.err
	}
	err := a.validateIndex(idx)
	if err != nil {
		return 0.0, err
	}
	f, ok := a.a[idx].(float64)
	if !ok {
		return 0.0, fmt.Errorf("item at index %d was a %T, not a float", idx, a.a[idx])
	}
	return f, nil
}

// Bool returns the value of a bool at index `idx` in the array, or a propagated error
func (a Array) Bool(idx int) (bool, error) {
	if a.err != nil {
		return false, a.err
	}
	err := a.validateIndex(idx)
	if err != nil {
		return false, err
	}
	b, ok := a.a[idx].(bool)
	if !ok {
		return false, fmt.Errorf("item at index %d was a %T, not a bool", idx, a.a[idx])
	}
	return b, nil
}

// Bytes returns the serialized value into a slice of bytes, or a propagated error
func (a Array) Bytes() ([]byte, error) {
	if a.err != nil {
		return nil, a.err
	}
	buf, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// MustBytes returns the serialized value into a slice of bytes, or panics if there was an error
func (a Array) MustBytes() []byte {
	if a.err != nil {
		panic(a.err)
	}
	result, err := json.Marshal(a.a)
	if err != nil {
		panic(err)
	}
	return result
}

// MustInner returns the `[]interface{}` which this `Array` represents, or panics if there was an error
func (a Array) MustInner() []interface{} {
	if a.err != nil {
		panic(a.err)
	}
	return a.a
}

// Len returns the length of the array, or a propagated error
func (a Array) Len() (int, error) {
	if a.err != nil {
		return 0, a.err
	}
	return len(a.a), nil
}

// MustString returns the value of a string at index `idx` in the array, or panics if there was an error
func (a Array) MustString(idx int) string {
	if a.err != nil {
		panic(a.err)
	}
	s, err := a.String(idx)
	if err != nil {
		panic(err)
	}
	return s
}

// MustFloat returns the value of a float at index `idx` in the array, or panics if there was an error
func (a Array) MustFloat(idx int) float64 {
	if a.err != nil {
		panic(a.err)
	}
	f, err := a.Float(idx)
	if err != nil {
		panic(err)
	}
	return f
}

// MustBool returns the value of a bool at index `idx` in the array, or panics if there was an error
func (a Array) MustBool(idx int) bool {
	if a.err != nil {
		panic(a.err)
	}
	b, err := a.Bool(idx)
	if err != nil {
		panic(err)
	}
	return b
}

// Err returns any error that was found up to this point
func (a Array) Err() error {
	return a.err
}
