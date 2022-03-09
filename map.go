package json

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Map provides methods to either navigate through the content of a JSON object or propagate any error that has occurred
type Map struct {
	m   map[string]interface{}
	err error
}

// MapFromBytes creates an Map from a []byte
func MapFromBytes(jsonBytes []byte) Map {
	var m Map
	err := json.Unmarshal(jsonBytes, &m.m)
	if err != nil {
		return Map{nil, err}
	}
	return m
}

// MapFromReader creates an Map from an io.Reader
func MapFromReader(r io.Reader) Map {
	var m Map
	err := json.NewDecoder(r).Decode(&m.m)
	if err != nil {
		return Map{nil, err}
	}
	return m
}

// NewMap creates an Map from a []interface{}
func NewMap(m map[string]interface{}) Map {
	return Map{m, nil}
}

func (m Map) panicOnErr() {
	if m.err != nil {
		panic(m.err)
	}
}

// MustHas returns true if the map contains the key `key`, or panics if there was an error
func (m Map) MustHas(key string) bool {
	m.panicOnErr()
	_, ok := m.m[key]
	return ok
}

// MustString returns the value of a string at key `key` in the map, or panics if there was an error
func (m Map) MustString(key string) string {
	m.panicOnErr()
	s, err := m.String(key)
	if err != nil {
		panic(err)
	}
	return s
}

// MustFloat returns the value of a float at key `key` in the map, or panics if there was an error
func (m Map) MustFloat(key string) float64 {
	m.panicOnErr()
	f, err := m.Float(key)
	if err != nil {
		panic(err)
	}
	return f
}

// MustBool returns the value of a bool at key `key` in the map, or panics if there was an error
func (m Map) MustBool(key string) bool {
	m.panicOnErr()
	b, err := m.Bool(key)
	if err != nil {
		panic(err)
	}
	return b
}

// Err returns any error that was found up to this point
func (m Map) Err() error {
	return m.err
}

func (m Map) validateKey(key string) error {
	if !m.MustHas(key) {
		validKeys := make([]string, 0, len(m.m))
		for k, _ := range m.m {
			validKeys = append(validKeys, k)
		}
		return fmt.Errorf("key not found: %s, valid keys: %+v", key, strings.Join(validKeys, ", "))
	}
	return nil
}

// String returns the value of a string at key `key` in the map, or a propagated error
func (m Map) String(key string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	if err := m.validateKey(key); err != nil {
		return "", err
	}
	s, ok := m.m[key].(string)
	if !ok {
		return "", fmt.Errorf("item with key %s was a %T, not a string", key, m.m[key])
	}
	return s, nil
}

// Float returns the value of a float at key `key` in the map, or a propagated error
func (m Map) Float(key string) (float64, error) {
	if m.err != nil {
		return 0.0, m.err
	}
	if err := m.validateKey(key); err != nil {
		return 0.0, err
	}
	f, ok := m.m[key].(float64)
	if !ok {
		return 0.0, fmt.Errorf("item with key %s was a %T, not a float", key, m.m[key])
	}
	return f, nil
}

// Bool returns the value of a bool at key `key` in the map, or a propagated error
func (m Map) Bool(key string) (bool, error) {
	if m.err != nil {
		return false, m.err
	}
	if err := m.validateKey(key); err != nil {
		return false, err
	}
	b, ok := m.m[key].(bool)
	if !ok {
		return false, fmt.Errorf("item with key %s was a %T, not a bool", key, m.m[key])
	}
	return b, nil
}

// Map returns the map found at key `key` in the map; errors will be propagated
func (m Map) Map(key string) Map {
	if m.err != nil {
		return m
	}
	if err := m.validateKey(key); err != nil {
		return Map{nil, err}
	}
	result, ok := m.m[key].(map[string]interface{})
	if !ok {
		return Map{nil, fmt.Errorf("item with key %s was a %T, not a map", key, m.m[key])}
	}
	return Map{result, nil}
}

// Array returns the array found at key `key` in the map; errors will be propagated
func (m Map) Array(key string) Array {
	if m.err != nil {
		return Array{nil, m.err}
	}
	if err := m.validateKey(key); err != nil {
		return Array{nil, err}
	}
	a, ok := m.m[key].([]interface{})
	if !ok {
		return Array{nil, fmt.Errorf("item with key %s was a %T, not an array", key, m.m[key])}
	}
	return Array{a, nil}
}

// Bytes returns the serialized value into a slice of bytes, or a propagated error
func (m Map) Bytes() ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// Has returns true if the map contains the key `key`, or a propagated error
func (m Map) Has(key string) (bool, error) {
	if m.err != nil {
		return false, m.err
	}
	_, ok := m.m[key]
	return ok, nil
}

// MustBytes returns the serialized value into a slice of bytes, or panics if there was an error
func (m Map) MustBytes() []byte {
	if m.err != nil {
		panic(m.err)
	}
	result, err := json.Marshal(m.m)
	if err != nil {
		panic(err)
	}
	return result
}

// Inner returns the `[]interface{}` which this `Array` represents, or a propagated error
func (m Map) Inner() (map[string]interface{}, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.m, nil
}

// MustInner returns the `[]interface{}` which this `Array` represents, or panics if there was an error
func (m Map) MustInner() map[string]interface{} {
	if m.err != nil {
		panic(m.err)
	}
	return m.m
}