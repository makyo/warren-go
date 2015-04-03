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
	render.HTML(200, "static/front", map[string]interface{}{
		"CSRF":    h.session.Values["_csrf_token"],
		"Title":   "Welcome",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request, render render.Render) {
	render.HTML(200, "static/about", map[string]interface{}{
		"Title":   "About",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}
