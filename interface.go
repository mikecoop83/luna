package json

type Serializer interface {
	Bytes() ([]byte, error)
	MustBytes() []byte
}

type Array interface {
	Serializer

	// Len returns the length of the array, or a propagated error
	Len() (int, error)
	// String returns the value of a string at index `idx` in the array, or a propagated error
	String(idx int) (string, error)
	// Float returns the value of a float at index `idx` in the array, or a propagated error
	Float(idx int) (float64, error)
	// Int returns the value of an int at index `idx` in the array, or a propagated error
	Int(idx int) (int, error)
	// Bool returns the value of a bool at index `idx` in the array, or a propagated error
	Bool(idx int) (bool, error)
	// Inner returns the `[]interface{}` which this `Array` represents, or a propagated error
	Inner() ([]interface{}, error)
	// Map returns the map found at index `idx` in the array; errors will be propagated
	Map(idx int) Map
	// Array returns the array found at index `idx` in the array; errors will be propagated
	Array(idx int) Array

	// MustLen returns the length of the array, or panics if there was an error
	MustLen() int
	// MustString returns the value of a string at index `idx` in the array, or panics if there was an error
	MustString(idx int) string
	// MustFloat returns the value of a float at index `idx` in the array, or panics if there was an error
	MustFloat(idx int) float64
	// MustInt returns the value of an int at index `idx` in the array, or panics if there was an error
	MustInt(idx int) int
	// MustBool returns the value of a bool at index `idx` in the array, or panics if there was an error
	MustBool(idx int) bool
	// MustInner returns the `[]interface{}` which this `Array` represents, or panics if there was an error
	MustInner() []interface{}

	// Err returns any error that was found up to this point
	Err() error
}

type Map interface {
	Serializer

	// Has returns true if the map contains the key `key`, or a propagated error
	Has(key string) (bool, error)
	// String returns the value of a string at key `key` in the map, or a propagated error
	String(key string) (string, error)
	// Float returns the value of a float at key `key` in the map, or a propagated error
	Float(key string) (float64, error)
	// Int returns the value of an int at key `key` in the map, or a propagated error
	Int(key string) (int, error)
	// Bool returns the value of a bool at key `key` in the map, or a propagated error
	Bool(key string) (bool, error)
	// Inner returns the `[]interface{}` which this `Array` represents, or a propagated error
	Inner() (map[string]interface{}, error)
	// Map returns the map found at key `key` in the map; errors will be propagated
	Map(key string) Map
	// Array returns the array found at key `key` in the map; errors will be propagated
	Array(key string) Array

	// MustHas returns true if the map contains the key `key`, or panics if there was an error
	MustHas(key string) bool
	// MustString returns the value of a string at key `key` in the map, or panics if there was an error
	MustString(key string) string
	// MustFloat returns the value of a float at key `key` in the map, or panics if there was an error
	MustFloat(key string) float64
	// MustInt returns the value of an int at key `key` in the map, or panics if there was an error
	MustInt(key string) int
	// MustBool returns the value of a bool at key `key` in the map, or panics if there was an error
	MustBool(key string) bool
	// MustInner returns the `[]interface{}` which this `Array` represents, or panics if there was an error
	MustInner() map[string]interface{}

	// Err returns any error that was found up to this point
	Err() error
}
