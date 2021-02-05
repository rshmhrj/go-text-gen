## Go Text Gen
This goal of this project is to easily generate text strings for use in other projects

### Requirements
- Go v1.15

### fourOhFour
Example usage of `fourOhFour` in a RESTful API using `gorilla/mux & gorilla/handlers` and `rs/zerolog`
```go
// NotFound handler
func (h *handler) NotFound(rw http.ResponseWriter, r *http.Request) {
hlog.FromRequest(r).Info().Msg("404 not found")
response := util.ResponseMessage{
Status:  util.StatusError,
Code:    http.StatusNotFound,
Message: fourOhFour.Generate(),
}
util.ReturnError(response, rw)
}
```

Example Result(s):
```bash
{ "status":"error", "code":404, "message":"Page not found" }
{ "status":"error", "code":404, "message":"Either we broke something or you cannot type" }
{ "status":"error", "code":404, "message":"202 + 202 = ?" }
```

**If you have any great 404 messages, please add and submit an MR!**


### names
Example usage of `names` to randomly assign a name to a [linked list](https://gitlab.com/rshmhrj/data-structs/lists) (similar to Docker & Kubernetes auto naming)
```go
// Initialize an empty linked list
func NewLinkedList(data ... interface{}) *linkedList {
return &linkedList{
name: names.NewGenerator().GenerateShort(),
...
}

l := lists.NewLinkedList("A", "B", "C")
fmt.Printf("%v\n", l)
```

Example Results for `GenerateShort()`:
```bash
the.much.game
as.busy.duty
last.home.blow
ill.fat.web
out.warm.map
not.west.shop
fast.high.till
none.last.roll
only.glad.rush
also.home.lock
true.only.bat
far.big.file
```

Example Results for `Generate()`:
```bash
well.inner.disaster
elsewhere.each.hair
recently.global.farm
nearly.late.chain
essentially.sudden.brain
here.obvious.airport
strange.unable.literature
each.electrical.clue
fresh.inner.housing
often.serious.relative
away.political.confidence
initially.much.classroom
```

### Future Ideas
- Pure random alphanumeric strings (similar to [random.org](https://www.random.org/strings/))
- Baby names
- Fantasy character names

**If you have any additional ideas, [please submit](https://gitlab.com/rshmhrj/go-text-gen/-/issues/new)!**
