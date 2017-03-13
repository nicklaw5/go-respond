package main

import (
	"net/http"

	resp "github.com/nicklaw5/go-respond"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Ok(nil)
	})
	http.ListenAndServe(":8080", nil)
}
