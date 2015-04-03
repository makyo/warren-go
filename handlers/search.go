package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/martini-contrib/render"

	"github.com/warren-community/warren/models"
)

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	q := r.FormValue("q")
	if q == "" {
		render.HTML(http.StatusOK, "search/displaySearch", map[string]interface{}{
			"CSRF":    h.session.Values["_csrf_token"],
			"Title":   "Search",
			"User":    h.user,
			"Flashes": h.flashes(r, w),
		})
		return
	}
	searchJson := fmt.Sprintf(`{
		"query": {
			"simple_query_string": {
				"query": "%s",
				"analyzer": "snowball",
				"fields": ["IndexedContent", "Title", "Tags"],
				"default_operator": "and"
			}
		}
	}`, template.JSEscapeString(q))
	res, err := h.esConn.Search("warren", "entity", nil, searchJson)
	if err != nil {
		l.Printf("Error conducting search: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	var hits []models.Entity
	for _, hit := range res.Hits.Hits {
		entity, err := models.GetEntity(hit.Id, h.db)
		if err != nil {
			l.Printf("Error fetching entity: %+v\n", err)
			h.InternalServerError(w, r, render)
			return
		}
		hits = append(hits, entity)
	}
	render.HTML(http.StatusOK, "search/searchResults", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   fmt.Sprintf("Search results for: %s", q),
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"Res":     res,
		"Hits":    hits,
		"Q":       q,
	})
}
