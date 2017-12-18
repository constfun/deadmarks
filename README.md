My Hello World for Go.

Parse the bookmarks file from Firefox and check for dead urls.

```
go run main.go bookmarks.json
```

TODO: 
* Use go routines and compare with `time`
* Can we use `tcp` and cut loading the response short after HTTP code?
* Can we stream bookmark structures while Unmarshalling? Premature, but for fun.
* Check go style guide.