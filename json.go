package json

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func MapFromBytes(jsonBytes []byte) Map {
	var m valueMap
	err := json.Unmarshal(jsonBytes, &m.m)
	if err != nil {
		return errorMap{err}
	}
	return m
}

func MapFromReader(r io.Reader) Map {
	var m valueMap
	err := json.NewDecoder(r).Decode(&m.m)
	if err != nil {
		return errorMap{err}
	}
	return m
}

func ArrayFromBytes(jsonBytes []byte) Array {
	var a valueArray
	err := json.Unmarshal(jsonBytes, &a.a)
	if err != nil {
		return errorArray{err}
	}
	return a
}

func ArrayFromReader(r io.Reader) Array {
	var a valueArray
	err := json.NewDecoder(r).Decode(&a.a)
	if err != nil {
		return errorArray{err}
	}
	return a
}

func NewMap(m map[string]interface{}) Map {
	return valueMap{m}
}

func NewArray(a []interface{}) Array {
	return valueArray{a}
}

type valueArray struct {
	a []interface{}
}

func (va valueArray) validateIndex(idx int) error {
	if idx < 0 || idx >= len(va.a) {
		return fmt.Errorf("invalid index: %d; it should be between 0 and %d", idx, len(va.a)-1)
	}
	return nil
}

func (va valueArray) Map(idx int) Map {
	err := va.validateIndex(idx)
	if err != nil {
		return errorMap{err}
	}
	m, ok := va.a[idx].(map[string]interface{})
	if !ok {
		return errorMap{
			err: fmt.Errorf("item at index %d was a %T, not a map", idx, va.a[idx]),
		}
	}
	return valueMap{m}
}

func (va valueArray) Array(idx int) Array {
	err := va.validateIndex(idx)
	if err != nil {
		return errorArray{err}
	}
	a, ok := va.a[idx].([]interface{})
	if !ok {
		return errorArray{
			err: fmt.Errorf("item at index %d was a %T, not an array", idx, va.a[idx]),
		}
	}
	return valueArray{a}
}

func (va valueArray) MustLen() int {
	return len(va.a)
}

func (va valueArray) Items() []interface{} {
	return va.a
}

func (va valueArray) String(idx int) (string, error) {
	err := va.validateIndex(idx)
	if err != nil {
		return "", err
	}
	s, ok := va.a[idx].(string)
	if !ok {
		return "", fmt.Errorf("item at index %d was a %T, not a string", idx, va.a[idx])
	}
	return s, nil
}

func (va valueArray) Float(idx int) (float64, error) {
	err := va.validateIndex(idx)
	if err != nil {
		return 0.0, err
	}
	f, ok := va.a[idx].(float64)
	if !ok {
		return 0.0, fmt.Errorf("item at index %d was a %T, not a float", idx, va.a[idx])
	}
	return f, nil
}

func (va valueArray) Int(idx int) (int, error) {
	f, err := va.Float(idx)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func (va valueArray) Bool(idx int) (bool, error) {
	err := va.validateIndex(idx)
	if err != nil {
		return false, err
	}
	b, ok := va.a[idx].(bool)
	if !ok {
		return false, fmt.Errorf("item at index %d was a %T, not a bool", idx, va.a[idx])
	}
	return b, nil
}

func (va valueArray) Bytes() ([]byte, error) {
	buf, err := json.Marshal(va)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (va valueArray) MustBytes() []byte {
	result, err := json.Marshal(va.a)
	if err != nil {
		panic(err)
	}
	return result
}

func (va valueArray) MustInner() []interface{} {
	return va.a
}

func (va valueArray) Len() (int, error) {
	return len(va.a), nil
}

func (va valueArray) MustString(idx int) string {
	s, err := va.String(idx)
	if err != nil {
		panic(err)
	}
	return s
}

func (va valueArray) MustFloat(idx int) float64 {
	f, err := va.Float(idx)
	if err != nil {
		panic(err)
	}
	return f
}

func (va valueArray) MustInt(idx int) int {
	i, err := va.Int(idx)
	if err != nil {
		panic(err)
	}
	return i
}

func (va valueArray) MustBool(idx int) bool {
	b, err := va.Bool(idx)
	if err != nil {
		panic(err)
	}
	return b
}

func (va valueArray) Err() error {
	return nil
}

type valueMap struct {
	m map[string]interface{}
}

func (vm valueMap) MustHas(key string) bool {
	_, ok := vm.m[key]
	return ok
}

func (vm valueMap) MustString(key string) string {
	s, err := vm.String(key)
	if err != nil {
		panic(err)
	}
	return s
}

func (vm valueMap) MustFloat(key string) float64 {
	f, err := vm.Float(key)
	if err != nil {
		panic(err)
	}
	return f
}

func (vm valueMap) MustInt(key string) int {
	i, err := vm.Int(key)
	if err != nil {
		panic(err)
	}
	return i
}

func (vm valueMap) MustBool(key string) bool {
	b, err := vm.Bool(key)
	if err != nil {
		panic(err)
	}
	return b
}

func (vm valueMap) Err() error {
	return nil
}

func (vm valueMap) validateKey(key string) error {
	if !vm.MustHas(key) {
		validKeys := make([]string, 0, len(vm.m))
		for k, _ := range vm.m {
			validKeys = append(validKeys, k)
		}
		return fmt.Errorf("key not found: %s, valid keys: %+v", key, strings.Join(validKeys, ", "))
	}
	return nil
}

func (vm valueMap) String(key string) (string, error) {
	if err := vm.validateKey(key); err != nil {
		return "", err
	}
	s, ok := vm.m[key].(string)
	if !ok {
		return "", fmt.Errorf("item with key %s was a %T, not a string", key, vm.m[key])
	}
	return s, nil
}

func (vm valueMap) Float(key string) (float64, error) {
	if err := vm.validateKey(key); err != nil {
		return 0.0, err
	}
	f, ok := vm.m[key].(float64)
	if !ok {
		return 0.0, fmt.Errorf("item with key %s was a %T, not a float", key, vm.m[key])
	}
	return f, nil
}

func (vm valueMap) Int(key string) (int, error) {
	f, err := vm.Float(key)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func (vm valueMap) Bool(key string) (bool, error) {
	if err := vm.validateKey(key); err != nil {
		return false, err
	}
	b, ok := vm.m[key].(bool)
	if !ok {
		return false, fmt.Errorf("item with key %s was a %T, not a bool", key, vm.m[key])
	}
	return b, nil
}

func (vm valueMap) Map(key string) Map {
	if err := vm.validateKey(key); err != nil {
		return errorMap{err}
	}
	m, ok := vm.m[key].(map[string]interface{})
	if !ok {
		return errorMap{fmt.Errorf("item with key %s was a %T, not a map", key, vm.m[key])}
	}
	return valueMap{m}
}

func (vm valueMap) Array(key string) Array {
	if err := vm.validateKey(key); err != nil {
		return errorArray{err}
	}
	a, ok := vm.m[key].([]interface{})
	if !ok {
		return errorArray{fmt.Errorf("item with key %s was a %T, not an array", key, vm.m[key])}
	}
	return valueArray{a}
}

func (vm valueMap) Bytes() ([]byte, error) {
	buf, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (vm valueMap) Has(key string) (bool, error) {
	_, ok := vm.m[key]
	return ok, nil
}

func (vm valueMap) MustBytes() []byte {
	result, err := json.Marshal(vm.m)
	if err != nil {
		panic(err)
	}
	return result
}

func (vm valueMap) Inner() (map[string]interface{}, error) {
	return vm.m, nil
}

func (vm valueMap) MustInner() map[string]interface{} {
	return vm.m
}

type errorMap struct {
	err error
}

func (em errorMap) MustHas(key string) bool {
	panic(em.err)
}

func (em errorMap) MustBytes() []byte {
	panic(em.err)
}

func (em errorMap) MustInner() map[string]interface{} {
	panic(em.err)
}

func (em errorMap) Inner() (map[string]interface{}, error) {
	return nil, em.err
}

func (em errorMap) MustString(_ string) string {
	panic(em.err)
}

func (em errorMap) MustFloat(_ string) float64 {
	panic(em.err)
}

func (em errorMap) MustInt(_ string) int {
	panic(em.err)
}

func (em errorMap) MustBool(_ string) bool {
	panic(em.err)
}

func (em errorMap) Err() error {
	return em.err
}

func (em errorMap) Bytes() ([]byte, error) {
	return nil, em.err
}

func (em errorMap) Has(_ string) (bool, error) {
	return false, em.err
}

func (em errorMap) String(_ string) (string, error) {
	return "", em.err
}

func (em errorMap) Float(_ string) (float64, error) {
	return 0.0, em.err
}

func (em errorMap) Int(_ string) (int, error) {
	return 0, em.err
}

func (em errorMap) Bool(_ string) (bool, error) {
	return false, em.err
}

func (em errorMap) Map(_ string) Map {
	return em
}

func (em errorMap) Array(_ string) Array {
	return errorArray{
		em.err,
	}
}

type errorArray struct {
	err error
}

func (ea errorArray) MustString(_ int) string {
	panic(ea.err)
}

func (ea errorArray) MustFloat(_ int) float64 {
	panic(ea.err)
}

func (ea errorArray) MustInt(_ int) int {
	panic(ea.err)
}

func (ea errorArray) MustBool(_ int) bool {
	panic(ea.err)
}

func (ea errorArray) Err() error {
	return ea.err
}

func (ea errorArray) Items() []interface{} {
	return nil
}

func (ea errorArray) Bytes() ([]byte, error) {
	return nil, ea.err
}

func (ea errorArray) MustLen() int {
	panic(ea.err)
}

func (ea errorArray) Len() (int, error) {
	return 0, ea.err
}

func (ea errorArray) String(_ int) (string, error) {
	return "", ea.err
}

func (ea errorArray) Float(_ int) (float64, error) {
	return 0.0, ea.err
}

func (ea errorArray) Int(_ int) (int, error) {
	return 0, ea.err
}

func (ea errorArray) Bool(_ int) (bool, error) {
	return false, ea.err
}

func (ea errorArray) Map(_ int) Map {
	return errorMap{
		ea.err,
	}
}

func (ea errorArray) Array(_ int) Array {
	return ea
}

func (ea errorArray) Inner() ([]interface{}, error) {
	return nil, ea.err
}

func (va valueArray) Inner() ([]interface{}, error) {
	return va.a, nil
}

func (ea errorArray) MustBytes() []byte {
	panic(ea.err)
}

func (ea errorArray) MustInner() []interface{} {
	panic(ea.err)
}
