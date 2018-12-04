package handler

import (
	"fmt"
	"net/http"
)

// IndexHandler to process gitlab events
type IndexHandler struct {
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Bot</h1><div>use /webhook for your gitlab trigger</div>")
}
