# go-respond

A Go package for handling common HTTP JSON responses.

[![Build Status](https://travis-ci.org/nicklaw5/go-respond.svg?branch=master)](https://travis-ci.org/nicklaw5/go-respond) [![Coverage Status](https://coveralls.io/repos/github/nicklaw5/go-respond/badge.svg)](https://coveralls.io/github/nicklaw5/go-respond)

## Installation

```bash
go get github.com/nicklaw5/go-respond
```

## Usage

The goal of `go-respond` is to take most of grunt work out preparing your JSON response. Here's a simple example:

```go
package main

import (
	"net/http"

	resp "github.com/nicklaw5/go-respond"
)

type response struct {
	Success bool `json:"success"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			Ok(&response{true})
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
| 409 | Conflict() |
| 422 | UnprocessableEntity() |
| 500 | InternalServerError() |

## Handling Errors

Your best option for handling errors that may occur when marshalling the JSON response, is to use [Negroni's Recovery middleware](https://github.com/urfave/negroni#recovery). Here's an example:

```go
package main

import (
  "net/http"

  "github.com/urfave/negroni"
  resp "github.com/nicklaw5/go-respond"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	  users := []User{
		  {1, "Billy", "billy@example.com"},
		  {2, "Joan", "joan@example.com"},
	  }

	  resp.NewResponse(w).
		  Created(users)
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

# License

This package is distributed under the terms of the [MIT](LICENSE) License.
