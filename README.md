# go-json  - Easier Parsing of Unstructured JSON in Go
___

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

In the happy day scenario...

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

But what if we got the name of the array wrong?
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

OK, what if you're confident the value is there, need to send it to another function, and are willing to let it panic
if something went wrong?  We have `Must*` versions of all the functions that return errors to do that. 

```go
    processScore(json.MapFromBytes(data).Array("people").Map(0).MustFloat("score"))
```

So what's the full interface?  [Check it out](https://github.com/mikecoop83/go-json/blob/main/interface.go)...
