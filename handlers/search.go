package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/martini-contrib/render"
)

func (h *Handlers) DisplaySearch(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	if r.FormValue("q") != "" {
		h.SearchResults(w, r, l)
		return
	}
	render.HTML(200, "search/displaySearch", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Search",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) SearchResults(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	searchJson := fmt.Sprintf(`{
		"query": {
			"term": {
				"indexedContent": "%s"
			}
		}
	}`, r.FormValue("q"))
	res, err := h.esConn.Search("warren", "entity", nil, searchJson)
	if err != nil {
		l.Print(err.Error())
		http.Error(w, "Error conducting search", http.StatusInternalServerError)
		return
	}
	l.Print(res.Hits.Hits)
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
