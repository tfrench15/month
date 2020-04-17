package main

import (
	"net/http"

	"github.com/tfrench15/month/pkg/month"
)

func main() {
	srv := month.NewServer()

	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		panic(err)
	}
}
