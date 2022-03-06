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

## For Example...

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
}
```

```
Uh oh! key not found: entries, valid keys: people
```

Well that's convenient!  And we didn't have to do annoying type casing from `interface{}` or error checking along the way.
Just check once when you pull out the value and if any error happened, it will be propagated to the final call to get the
values out of the JSON.

### Wow! How did that work?

Actually, it's not that complicated.  Whenever you call `.Array(...)` or `.Map(...)` with an invalid index or key, it
returns a struct that holds on to that error and finally returns it when you call one of the other functions
like `Float(...)`.

### But I really hate checking for errors... now what?

OK, what if you're confident the value is there, need to send it to another function, and are willing to let it panic
if something went wrong?  We have `Must*` versions of all the functions that return errors to satisfy your very confident
style of programming.

```go
    processScore(json.MapFromBytes(data).Array("people").Map(0).MustFloat("score"))
```

### So what's the full interface?

[Check it out](https://github.com/mikecoop83/go-json/blob/main/interface.go)...
