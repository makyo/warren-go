package handlers

import (
	"net/http"
)

func (h *Handlers) Front(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "front.tmpl")
}
