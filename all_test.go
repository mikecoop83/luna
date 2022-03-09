package json

import (
	ejson "encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var simpleJson = []byte(`{
	"object": {
		"strKey": "strValue",
		"boolKey": true,
		"floatKey": 1.21,
		"arrayObjKey": [
			{
				"val": 1
			},
			{
				"val": 2
			}
		],
		"nestedArrayFloatKey": [
			[
				3, 4.1
			]
		],
		"arrayStrKey": [
			"str1", "str2"
		]
	}
}`)

func TestNestedArrayFloatValue(t *testing.T) {
	m := MapFromBytes(simpleJson)
	i, err := m.Map("object").Array("nestedArrayFloatKey").Array(0).Float(1)
	require.NoError(t, err)
	require.Equal(t, 4.1, i)
}

func TestNestedArrayFloatBadIndex(t *testing.T) {
	m := MapFromBytes(simpleJson)
	i, err := m.Map("object").Array("nestedArrayFloatKey").Array(1).Float(2)
	require.Error(t, err)
	require.Equal(t, 0.0, i)
}

func TestArrayLen(t *testing.T) {
	m := MapFromBytes(simpleJson)
	l, err := m.Map("object").Array("arrayStrKey").Len()
	require.NoError(t, err)
	require.Equal(t, 2, l)
}

func TestArrayMustLen(t *testing.T) {
	m := MapFromBytes(simpleJson)
	l := m.Map("object").Array("arrayStrKey").MustLen()
	require.Equal(t, 2, l)
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
	score, err := func() (float64, error) {
		var dataMap map[string]interface{}
		err := ejson.Unmarshal(data, &dataMap)
		if err != nil {
			return 0.0, err
		}
		peopleObj, ok := dataMap["people"] // peopleObj is a `[]interface{}`
		if !ok {
			return 0.0, errors.New("people key not found")
		}
		peopleArray, ok := peopleObj.([]interface{})
		if len(peopleArray) == 0 {
			return 0.0, errors.New("no people found")
		}
		firstPerson := peopleArray[0] // firstPerson is an `interface{}`
		firstPersonMap, ok := firstPerson.(map[string]interface{})
		if !ok {
			return 0.0, fmt.Errorf("first person should be a map[string]interface{}, but is a %T", firstPerson)
		}
		firstPersonScore, ok := firstPersonMap["score"] // firstPersonScore is an `interface{}`
		if !ok {
			return 0.0, fmt.Errorf("score not found")
		}
		score, ok := firstPersonScore.(float64)
		if !ok {
			return 0.0, fmt.Errorf("score should be a float64, but is a %T", score)
		}
		return score, nil
	}()
	require.NoError(t, err)
	require.Equal(t, 89.5, score)

	score, err = MapFromBytes(data).Array("people").Map(0).Float("score")
	require.NoError(t, err)
	require.Equal(t, score, 89.5)

	score, err = MapFromBytes(data).Array("entries").Map(0).Float("score")
	require.Error(t, err)
	// make sure it provides the bad key
	require.Contains(t, err.Error(), "entries")
	// and the good one too
	require.Contains(t, err.Error(), "people")
}
