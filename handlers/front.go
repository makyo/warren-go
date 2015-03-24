// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in 
// the LICENSE file.

package handlers

import (
	"net/http"
)

func (h *Handlers) Front(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "front.tmpl")
}
