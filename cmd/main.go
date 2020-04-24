package main

import (
	"net/http"

	"github.com/tfrench15/month/pkg/month"
)

func main() {
	srv := month.NewServer()

	if err := http.ListenAndServe(":8080", srv); err != nil {
		panic(err)
	}
}
