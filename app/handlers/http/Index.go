package http

import (
	"fmt"
	"net/http"
)

func HTTPIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
