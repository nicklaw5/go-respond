# go-respond

A Go package for handling common HTTP JSON responses.

[![GoDoc](https://godoc.org/github.com/nicklaw5/go-respond?status.svg)](https://godoc.org/github.com/nicklaw5/go-respond)
[![Build Status](https://travis-ci.org/nicklaw5/go-respond.svg?branch=master)](https://travis-ci.org/nicklaw5/go-respond)
[![Coverage Status](https://coveralls.io/repos/github/nicklaw5/go-respond/badge.svg)](https://coveralls.io/github/nicklaw5/go-respond)
[![Go Report Card](https://goreportcard.com/badge/github.com/nicklaw5/go-respond)](https://goreportcard.com/report/github.com/nicklaw5/go-respond)

## Installation

```bash
go get github.com/nicklaw5/go-respond
```

## Usage

The goal of `go-respond` is to take most of the grunt work out preparing your JSON response. Here's a simple example:

```go
package main

import (
    "net/http"

    resp "github.com/nicklaw5/go-respond"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        users := []User{
            {1, "Billy", "billy@example.com"},
            {2, "Joan", "joan@example.com"},
        }

        resp.NewResponse(w).Ok(users)
    })

    http.ListenAndServe(":8080", nil)
}
```

## Response Methods

| Response Code | Method Name |
| :---------- | :------------ |
| 200 | Ok() |
| 201 | Created() |
| 202 | Accepted() |
| 204 | NoContent() |
| 400 | BadRequest() |
| 401 | Unauthorized() |
| 403 | Forbidden() |
| 404 | NotFound() |
| 405 | MethodNotAllowed() |
| 406 | NotAcceptable() |
| 409 | Conflict() |
| 411 | LengthRequired() |
| 412 | PreconditionFailed() |
| 413 | RequestEntityTooLarge() |
| 415 | UnsupportedMediaType() |
| 422 | UnprocessableEntity() |
| 500 | InternalServerError() |
| 501 | NotImplemented() |
| 502 | BadGateway() |
| 503 | ServiceUnavailable() |
| 504 | GatewayTimeout() |

See [here](https://httpstatuses.com/) for a complete list of HTTP responses, along with an explanation of each.

Please submit a PR if you want to add to this list. Only the most common response types have been included.

## To Long, Don't Write

Sometimes you don't need to return a specific content-message but don't want the response body to be empty.
In this case you can use the `DefaultMessage()` for responding with json containing the default message for the corresponding status code.

```go
package main

import (
    "net/http"
    resp "github.com/nicklaw5/go-respond"
)

func main() {
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        // ...
        if !authenticated {
            resp.NewResponse(w).DefaultMessage().
                Unauthorized(nil)
        }
        // ...
    })
    http.ListenAndServe(":8080", nil)
}
```

Would respond with `{"status":401,"message":"Unauthorized"}`

## Handling Errors

The best option for handling errors that may occur while marshalling the JSON response, is to use [Negroni's Recovery middleware](https://github.com/urfave/negroni#recovery). Here's an example:

```go
package main

import (
    "net/http"

    "github.com/urfave/negroni"
    resp "github.com/nicklaw5/go-respond"
)

type Response struct {
    Success bool `json:"success"`
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        resp.NewResponse(w).Ok(&Response{true})
    })

    n := negroni.New()
    recovery := negroni.NewRecovery()
    recovery.ErrorHandlerFunc = func(error interface{}) {
        // do something with the unexpected error
    }

    n.Use(recovery)
    n.UseHandler(mux)

    http.ListenAndServe(":8080", n)
}
```

## License

This package is distributed under the terms of the [MIT](LICENSE) License.
