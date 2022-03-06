package json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var simpleJson = []byte(`{
	"object": {
		"strKey": "strValue",
		"boolKey": true,
		"intKey": 42,
		"floatKey": 1.21,
		"arrayObjKey": [
			{
				"val": 1
			},
			{
				"val": 2
			}
		],
		"nestedArrayKey": [
			[
				3, 4, 5
			]
		],
		"arrayStrKey": [
			"str1",
			"str2"
		]
	}
}`)

func TestNestedArrayValue(t *testing.T) {
	m := MapFromBytes(simpleJson)
	i, err := m.Map("object").Array("nestedArrayKey").Array(0).Float(2)
	require.NoError(t, err)
	require.Equal(t, 5.0, i)
}

func TestNestedArrayBadIndex(t *testing.T) {
	m := MapFromBytes(simpleJson)
	i, err := m.Map("object").Array("nestedArrayKey").Array(1).Float(2)
	require.Error(t, err)
	require.Equal(t, 0.0, i)
}

func TestMissingArrayLen(t *testing.T) {
	m := MapFromBytes(simpleJson)
	notFound := m.Array("missing")
	require.Error(t, notFound.Err())
	l, err := notFound.Len()
	require.Error(t, err)
	require.Equal(t, 0, l)
}

func TestMissingArrayMustLen(t *testing.T) {
	m := MapFromBytes(simpleJson)
	notFound := m.Array("missing")
	require.Error(t, notFound.Err())
	require.Panics(t, func() {
		notFound.MustLen()
	})
}

func TestMissingMapError(t *testing.T) {
	m := MapFromBytes(simpleJson)
	notFound := m.Map("missing")
	require.Error(t, notFound.Err())
}

func TestMissingMapStringError(t *testing.T) {
	m := MapFromBytes(simpleJson)
	s, err := m.Map("missing").String("strKey")
	require.Error(t, err)
	require.Equal(t, "", s)
}

func TestString(t *testing.T) {
	m := MapFromBytes(simpleJson)
	s, err := m.Map("object").String("strKey")
	require.NoError(t, err)
	require.Equal(t, "strValue", s)
}

func TestReadme(t *testing.T) {
	data := []byte(`{
    "people": [
        {
            "name": "alice",
            "score": 89.5,
            "friends": [ "bob" ],
            "deleted": false
        },
        {
            "name": "bob",
            "score": 75.5,
            "friends": [],
            "deleted": false
        }
    ]
}`)
	score, err := MapFromBytes(data).Array("people").Map(0).Float("score")
	require.NoError(t, err)
	require.Equal(t, score, 89.5)

	score, err = MapFromBytes(data).Array("entries").Map(0).Float("score")
	require.Error(t, err)
	// make sure it provides the bad key
	require.Contains(t, err.Error(), "entries")
	// and the good one too
	require.Contains(t, err.Error(), "people")
}
