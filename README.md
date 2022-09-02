# Better Parsing of Unstructured JSON in Go
___

## Goal

The intent of **luna** is to create a simpler interface to work with when parsing JSON in Go without having to
create a bunch of structs with tags.

Features include:

 * An interface that propagates errors until you finally extract a value to use
 * More detailed error messages that give valuable debugging information (such as other keys in the map when a key
is not found)
 * Equivalent versions of value-retrieving functions to panic on error (but allow nesting results inside function calls)

## Motivation

Parsing unstructured JSON is plain ugly in Go.  Let's say you have the below JSON  and you want to pull out
the first person's score:

```go
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
            "deleted": false,
        }
    ]
}`)
```

If you didn't want to go through the hassle of creating structs that represent each of the entities in the JSON, you
would have to do something like this:

```go
var dataMap map[string]interface{}
err := json.Unmarshal(data, &dataMap)
if err != nil {
	return err
}
peopleObj, ok := dataMap["people"] // peopleObj is an `interface{}`
if !ok {
	return fmt.Errorf("people key not found")
}
peopleArray, ok := peopleObj.([]interface{})
if !ok {
	return fmt.Errorf("people should be a []interface{}, but is a %T", peopleArray)
}
if len(peopleArray) == 0 {
	return fmt.Errorf("no people found")
}
firstPerson := peopleObj[0] // firstPerson is an `interface{}`
firstPersonMap, ok := firstPerson.(map[string]interface{})
if !ok {
	return fmt.Errorf("first person should be a map[string]interface{}, but is a %T", firstPerson)
}
firstPersonScore, ok := firstPersonMap["score"] // firstPersonScore is an `interface{}`
if !ok {
	return fmt.Errorf("score not found")
}
score, ok := firstPersonScore.(float64)
if !ok {
	return fmt.Errorf("score should be a float64, but is a %T", score)
}
```

Phew!  That was a lot of code to get one little value out of a tiny JSON.  Is it possible to write this more concisely
but retain the all the validation that guards us against the data not being in the expected format?

## Simple Code Without Sacrifice

Using this library, here's how you would pull out the same value and still include all the useful validation done
above...

```go
score, err := luna.MapFromBytes(data).Array("people").Map(0).Float("score")
if err != nil {
    return err
}
fmt.Printf("Alice's score: %f\n", score)
```

```
Alice's score: 89.5
```

But what if we got the key wrong and used `"grade"` instead of `"score"`?
```go
score, err := luna.MapFromBytes(data).Array("people").Map(0).Float("grade")
if err != nil {
    fmt.Printf("Uh oh! %s\n", err)
}
```

```
Uh oh! key 'grade' not found at path $['people'][0], valid keys: name, score, friends, deleted
```

Well that's much cleaner!  And we didn't have to do annoying type casting from `interface{}` or error checking along the
way. Just check once when you pull out the value and if any of the errors that we checked for in the old code happened,
they will be propagated to the final call to get the values out of the JSON.  As a bonus, the error message
conveniently contains the full path to where the error happened along with the valid keys at that path!

As another example, showing the power of the error propagation, what if we got the `people` key wrong and used `folks`?

```go
score, err := luna.MapFromBytes(data).Array("folks").Map(0).Float("grade")
if err != nil {
    fmt.Printf("Uh oh! %s\n", err)
}
```

```
Uh oh! key 'folks' not found at path $, valid keys: people
```

Even though the error occurred earlier on in the chained calls, we still see the original error!

### Wow! How did that work?

Actually, it's not that complicated.  Whenever you call `.Array(...)` or `.Map(...)` with an invalid index or key, it
returns a struct that holds on to that error and finally returns it when you call one of the other functions
like `Float(...)`. Meanwhile, it also builds up and keeps track of the path so that if any error does happen, you'll
know exactly where to look!

### So what's the full interface?

[Check it out](https://pkg.go.dev/github.com/mikecoop83/luna)...
