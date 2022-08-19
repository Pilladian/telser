package main

import (
	"fmt"
	"net/http"
)

func rootRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello There!</h1><p>This is the main page.</p>")
}
