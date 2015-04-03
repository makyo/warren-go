// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/warren-community/warren/models"
)

func (h *Handlers) DisplayTag(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger, params martini.Params) {
	searchJson := fmt.Sprintf(`{
		"query": {
			"match": {
				"Tags": "%s"
			}
		}
	}`, template.JSEscapeString(params["tag"]))
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
	render.HTML(http.StatusOK, "posts/displayTag", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   fmt.Sprintf("Posts tagged %s", params["tag"]),
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"Hits":    hits,
	})
}

// List posts from those the user is following.
func (h *Handlers) ListFollowing(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// List posts from those with whom the user is friends.
func (h *Handlers) ListFriends(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
