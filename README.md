# Router

A simple router written in Go for APIs and servers of Eodomius projects

## Installation

```bash
go get github.com/eodomius/router
```

## Usage

### Create a router

You can create a new router with the `New` function :

```go
var router = router.New()
```

### Handle a request

You can handle a request with the [supported methods](#supported-methods) :

- First parameter is the path of the request

- Second parameter is the function to call when the request is received
  - First parameter is the responce
  - Second parameter is the request
  - Third parameter is a [Result](#result) structure

```go
  router.Get("/test/{id}", func(w http.ResponseWriter, r *http.Request, result *Result){
    w.Write([]byte("Get : Test ID : " + result.Params["id"]))
  })
  router.Post("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
    w.Write([]byte("Post : Test"))
  })
  router.Patch("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
    w.Write([]byte("Patch : Test"))
  })
  router.Put("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
    w.Write([]byte("Put : Test"))
  })
  router.Delete("/test", func(w http.ResponseWriter, r *http.Request, result *Result){
    w.Write([]byte("Delete : Test"))
  })
```

### Middleware

You can add middleware to the router with the `Use` function :

- This function takes a function as parameter :
  - First parameter is the responce
  - Second parameter is the request

```go
router.Use(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Middleware"))
})
```

## Supported methods

- `Get` : handle GET method
- `Head` : handle HEAD method
- `Connect` : handle CONNECT method
- `Trace` : handle TRACE method
- `Options` : handle OPTIONS method
- `Connect` : handle CONNECT method
- `Post` : handle POST method
- `Patch` : handle PATCH method
- `Put` : handle PUT method
- `Delete` : handle DELETE method

## Typedef

### Result

Result is a structure that contains parameters.

```go
type Result struct {
  Params map[string]string // [paramName]paramValue
}
```
