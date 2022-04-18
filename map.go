package luna

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Map provides methods to either navigate through the content of a JSON object or propagate any error that has occurred
type Map struct {
	m    map[string]interface{}
	path path
	err  error
}

func newMapAtRoot() Map {
	return Map{path: "$"}
}

// MapFromBytes creates a Map from a []byte
func MapFromBytes(jsonBytes []byte) Map {
	m := newMapAtRoot()
	err := json.Unmarshal(jsonBytes, &m.m)
	if err != nil {
		return Map{nil, "$", err}
	}
	return m
}

// MapFromReader creates a Map from an io.Reader
func MapFromReader(r io.Reader) Map {
	m := newMapAtRoot()
	err := json.NewDecoder(r).Decode(&m.m)
	if err != nil {
		return Map{nil, "$", err}
	}
	return m
}

// NewMap creates a Map from a map[string]interface{}
func NewMap(m map[string]interface{}) Map {
	return Map{m, "$", nil}
}

func (m Map) panicOnErr() {
	if m.err != nil {
		panic(m.err)
	}
}

// Err returns any error that was found up to this point
func (m Map) Err() error {
	return m.err
}

func (m Map) validateKey(key string) error {
	hasKey, err := m.Has(key)
	if err != nil {
		return err
	}
	if !hasKey {
		validKeys := make([]string, 0, len(m.m))
		for k, _ := range m.m {
			validKeys = append(validKeys, k)
		}
		return fmt.Errorf("key '%s' not found at path %s, valid keys: %+v", key, m.path, strings.Join(validKeys, ", "))
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
		return "", fmt.Errorf("item with key '%s' at path %s was a %T, not a string", key, m.path, m.m[key])
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
		return Map{nil, m.path, err}
	}
	currPath := m.path.appendKey(key)
	result, ok := m.m[key].(map[string]interface{})
	if !ok {
		return Map{nil, m.path, fmt.Errorf("item at path %s was a %T, not a map", currPath, m.m[key])}
	}
	return Map{result, currPath, nil}
}

// Array returns the array found at key `key` in the map; errors will be propagated
func (m Map) Array(key string) Array {
	if m.err != nil {
		return Array{nil, m.path, m.err}
	}
	if err := m.validateKey(key); err != nil {
		return Array{nil, m.path, err}
	}
	currPath := m.path.appendKey(key)
	a, ok := m.m[key].([]interface{})
	if !ok {
		return Array{nil, m.path, fmt.Errorf("item at path %s was a %T, not an array", currPath, m.m[key])}
	}
	return Array{a, currPath, nil}
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

// Inner returns the `map[string]interface{}` which this `Map` represents, or a propagated error
func (m Map) Inner() (map[string]interface{}, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.m, nil
}
