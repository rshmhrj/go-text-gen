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