// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in 
// the LICENSE file.

package handlers

import (
	"net/http"
)

func (h *Handlers) ListAll(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) ListFollowing(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handlers) ListFriends(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
