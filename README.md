# go-json  - Easier Parsing of Unstructured JSON in Go
___

## Goal

The intent of this library is to create an easier interface to work with when parsing JSON in Go without having to
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
	return errors.New("people key not found")
}
peopleArray, ok := peopleObj.([]interface{})
if !ok {
	return errors.New("people should be a []interface{}, but is a %T", peopleArray)
}
if len(peopleArray) == 0 {
	return errors.New("no people found")
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
but retain the all the validation that guards us against the data not being in the expected format.

## Let's see how it works now...

For the happy day scenario, let's get parse the JSON, and get the score of the first person in the people array...

```go
score, err := json.MapFromBytes(data).Array("people").Map(0).Float("score")
if err != nil {
    // no error in this case, but nice to only check once at the very end!
}
fmt.Printf("Alice's score: %f\n", score)
```

```
Alice's score: 89.5
```

But what if we got the name of the array wrong and used `"entries"` instead of `"people"`?
```go
score, err := json.MapFromBytes(data).Array("entries").Map(0).Float("score")
if err != nil {
    fmt.Printf("Uh oh! %s\n", err)
}
```

```
Uh oh! key not found: entries, valid keys: people
```

Well that's much cleaner!  And we didn't have to do annoying type casting from `interface{}` or error checking along the
way. Just check once when you pull out the value and if any of the errors that we checked for in the old code happened,
they will be propagated to the final call to get the values out of the JSON.

### Wow! How did that work?

Actually, it's not that complicated.  Whenever you call `.Array(...)` or `.Map(...)` with an invalid index or key, it
returns a struct that holds on to that error and finally returns it when you call one of the other functions
like `Float(...)`.

### But I really hate checking for errors... now what?

OK, what if you're confident the value is there, need to send it to another function, and are willing to let it panic
if something went wrong?  That's what the `Must*` versions of all the functions that return errors are for!

```go
processScore(json.MapFromBytes(data).Array("people").Map(0).MustFloat("score"))
```

### So what's the full interface?

[Check it out](DOCUMENTATION.md)...
