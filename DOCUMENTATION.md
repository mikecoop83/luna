# json

## Types

### type [Array](/array.go#L10)

`type Array struct { ... }`

Array provides methods to either navigate through the content of a JSON array or propagate any error that has occurred

#### func (Array) [Array](/array.go#L70)

`func (a Array) Array(idx int) Array`

Array returns the array found at index `idx` in the array; errors will be propagated

#### func (Array) [Bool](/array.go#L131)

`func (a Array) Bool(idx int) (bool, error)`

Bool returns the value of a bool at index `idx` in the array, or a propagated error

#### func (Array) [Bytes](/array.go#L147)

`func (a Array) Bytes() ([]byte, error)`

Bytes returns the serialized value into a slice of bytes, or a propagated error

#### func (Array) [Err](/array.go#L223)

`func (a Array) Err() error`

Err returns any error that was found up to this point

#### func (Array) [Float](/array.go#L115)

`func (a Array) Float(idx int) (float64, error)`

Float returns the value of a float at index `idx` in the array, or a propagated error

#### func (Array) [Len](/array.go#L179)

`func (a Array) Len() (int, error)`

Len returns the length of the array, or a propagated error

#### func (Array) [Map](/array.go#L49)

`func (a Array) Map(idx int) Map`

Map returns the map found at index `idx` in the array; errors will be propagated

#### func (Array) [MustBool](/array.go#L211)

`func (a Array) MustBool(idx int) bool`

MustBool returns the value of a bool at index `idx` in the array, or panics if there was an error

#### func (Array) [MustBytes](/array.go#L159)

`func (a Array) MustBytes() []byte`

MustBytes returns the serialized value into a slice of bytes, or panics if there was an error

#### func (Array) [MustFloat](/array.go#L199)

`func (a Array) MustFloat(idx int) float64`

MustFloat returns the value of a float at index `idx` in the array, or panics if there was an error

#### func (Array) [MustInner](/array.go#L171)

`func (a Array) MustInner() []interface{ ... }`

MustInner returns the `[]interface{}` which this `Array` represents, or panics if there was an error

#### func (Array) [MustLen](/array.go#L91)

`func (a Array) MustLen() int`

MustLen returns the length of the array, or panics if there was an error

#### func (Array) [MustString](/array.go#L187)

`func (a Array) MustString(idx int) string`

MustString returns the value of a string at index `idx` in the array, or panics if there was an error

#### func (Array) [String](/array.go#L99)

`func (a Array) String(idx int) (string, error)`

String returns the value of a string at index `idx` in the array, or a propagated error

### type [Map](/map.go#L11)

`type Map struct { ... }`

Map provides methods to either navigate through the content of a JSON object or propagate any error that has occurred

#### func (Map) [Array](/map.go#L167)

`func (m Map) Array(key string) Array`

Array returns the array found at key `key` in the map; errors will be propagated

#### func (Map) [Bool](/map.go#L136)

`func (m Map) Bool(key string) (bool, error)`

Bool returns the value of a bool at key `key` in the map, or a propagated error

#### func (Map) [Bytes](/map.go#L183)

`func (m Map) Bytes() ([]byte, error)`

Bytes returns the serialized value into a slice of bytes, or a propagated error

#### func (Map) [Err](/map.go#L90)

`func (m Map) Err() error`

Err returns any error that was found up to this point

#### func (Map) [Float](/map.go#L121)

`func (m Map) Float(key string) (float64, error)`

Float returns the value of a float at key `key` in the map, or a propagated error

#### func (Map) [Has](/map.go#L195)

`func (m Map) Has(key string) (bool, error)`

Has returns true if the map contains the key `key`, or a propagated error

#### func (Map) [Inner](/map.go#L216)

`func (m Map) Inner() (map[string]interface{ ... }, error)`

Inner returns the `[]interface{}` which this `Array` represents, or a propagated error

#### func (Map) [Map](/map.go#L151)

`func (m Map) Map(key string) Map`

Map returns the map found at key `key` in the map; errors will be propagated

#### func (Map) [MustBool](/map.go#L80)

`func (m Map) MustBool(key string) bool`

MustBool returns the value of a bool at key `key` in the map, or panics if there was an error

#### func (Map) [MustBytes](/map.go#L204)

`func (m Map) MustBytes() []byte`

MustBytes returns the serialized value into a slice of bytes, or panics if there was an error

#### func (Map) [MustFloat](/map.go#L70)

`func (m Map) MustFloat(key string) float64`

MustFloat returns the value of a float at key `key` in the map, or panics if there was an error

#### func (Map) [MustHas](/map.go#L53)

`func (m Map) MustHas(key string) bool`

MustHas returns true if the map contains the key `key`, or panics if there was an error

#### func (Map) [MustInner](/map.go#L224)

`func (m Map) MustInner() map[string]interface{ ... }`

MustInner returns the `[]interface{}` which this `Array` represents, or panics if there was an error

#### func (Map) [MustString](/map.go#L60)

`func (m Map) MustString(key string) string`

MustString returns the value of a string at key `key` in the map, or panics if there was an error

#### func (Map) [String](/map.go#L106)

`func (m Map) String(key string) (string, error)`

String returns the value of a string at key `key` in the map, or a propagated error

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
