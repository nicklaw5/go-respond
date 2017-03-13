# go-respond

A Go package for handling common HTTP JSON responses.

[![Build Status](https://travis-ci.org/nicklaw5/go-respond.svg?branch=master)](https://travis-ci.org/nicklaw5/go-respond) [![Coverage Status](https://coveralls.io/repos/github/nicklaw5/go-respond/badge.svg)](https://coveralls.io/github/nicklaw5/go-respond)

## Usage

HTTP handlers can be messy, uncomfortable to work with, and requires more than just a few lines of code to accomplish a basic JSON response. Take the below example as one case in point.

This simple query returns a `200 OK` response, with `{"success":true}` as the body and also sets the appropriate `application/json` header:

```go
package main

import (
	"net/http"
)

type Response struct {
    Success bool `json: success`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        response, err := json.Marshal(&Response{true})
        if err != nil {
            panic(err)
        }

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        w.Write(response)
	})

	http.ListenAndServe(":8080", nil)
}
```

The goal of `go-respond` is to take this kind of grunt work out handlers. By using `go-respond` the above example can written more succinctly as:

```go
package main

import (
	"net/http"

	res "github.com/nicklaw5/go-respond"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res.NewResponse(w).Ok(nil)
	})

	http.ListenAndServe(":8080", nil)
}
```

# License

This package is distributed under the terms of the [MIT](LICENSE) License.
