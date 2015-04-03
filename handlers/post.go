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

// Display a single post.
func (h *Handlers) DisplayPost(w http.ResponseWriter, r *http.Request, l *log.Logger, params martini.Params, render render.Render) {
	entity, err := models.GetEntity(params["entityId"], h.db)
	if err != nil {
		l.Printf("Could not fetch entity: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if entity.Id.Hex() == "" {
		h.NotFound(w, r, render)
		return
	}
	render.HTML(http.StatusOK, "post/displayPost", map[string]interface{}{
		"Title":   entity.Title,
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"CSRF":    h.session.Values["_csrf_token"],
		"Entity":  entity,
		"IsOwner": entity.BelongsToUser(h.user.Model),
		"Content": template.HTML(entity.RenderedContent),
	})
}

// Remove a post from the database.
func (h *Handlers) DeletePost(w http.ResponseWriter, r *http.Request, render render.Render, l *log.Logger) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	entity, err := models.GetEntity(r.FormValue("entityId"), h.db)
	if err != nil {
		l.Printf("Could not fetch entity: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	if entity.Id.Hex() == "" {
		h.NotFound(w, r, render)
		return
	}
	if !entity.BelongsToUser(h.user.Model) {
		h.Forbidden(w, r, render)
		return
	}
	err = entity.Delete(h.db)
	if err != nil {
		l.Printf("Could not delete entity: %+v\n", err)
		h.InternalServerError(w, r, render)
		return
	}
	h.session.AddFlash(NewFlash("Post deleted!", "success"))
	h.session.Save(r, w)
	http.Redirect(w, r, fmt.Sprintf("/~%s", h.user.Model.Username), http.StatusSeeOther)
}

func (h *Handlers) DisplaySharePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Share a post with one's followers.
func (h *Handlers) SharePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Display the form for creating a post.
func (h *Handlers) DisplayCreatePost(w http.ResponseWriter, r *http.Request, render render.Render) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	render.HTML(http.StatusOK, "post/displayCreatePost", map[string]interface{}{
		"Title":   "Create post",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
		"CSRF":    h.session.Values["_csrf_token"],
	})
}

// Create a post in the database.
func (h *Handlers) CreatePost(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	if !h.user.IsAuthenticated {
		h.session.AddFlash(NewFlash("Please log in to continue", "warning"))
		h.session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	entity := models.NewEntity(
		r.FormValue("contentType"),
		h.user.Model.Username,
		h.user.Model.Username,
		false, // Shared?
		r.FormValue("title"),
		r.FormValue("content"),
	)
	err := entity.Save(h.db, h.esConn)
	if err != nil {
		l.Print(err.Error())
		http.Error(w, "Could not save post", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/%s", entity.Id.Hex()), http.StatusSeeOther)
}
