// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"net/http"

	"github.com/martini-contrib/render"
)

// Serve the front-page of the site.
func (h *Handlers) Front(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(http.StatusOK, "static/front", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Welcome",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(http.StatusNotFound, "static/notFound", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Not found",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) InternalServerError(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(http.StatusInternalServerError, "static/internalServerError", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Internal server error",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) BadRequest(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(http.StatusBadRequest, "static/badRequest", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Bad request",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) Forbidden(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(http.StatusForbidden, "static/forbidden", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Forbidden!",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(http.StatusOK, "static/about", map[string]interface{}{
		"Title":   "About",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}
