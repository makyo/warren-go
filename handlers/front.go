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
	render.HTML(200, "front/front", map[string]interface{}{
		"Title":   "Welcome",
		"User":    h.user,
		"Flashes": h.flashes(r, w),
	})
}
