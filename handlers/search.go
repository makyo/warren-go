package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handlers) SearchResults(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	searchJson := fmt.Sprintf(`{
		"query": {
			"term": {
				"indexedContent": "%s"
			}
		}
	}`, r.FormValue("searchTerm"))
	res, err := h.esConn.Search("warren", "entity", nil, searchJson)
	if err != nil {
		l.Print(err.Error())
		http.Error(w, "Error conducting search", http.StatusInternalServerError)
		return
	}
	l.Print(res.Hits.Hits)
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
