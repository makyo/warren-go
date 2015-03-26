// Copyright 2015 The Warren Authors
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.

package handlers

import (
	"encoding/gob"
	"net/http"
)

func init() {
	gob.Register(Flash{})
}

/* Flash handling ---------------------------------------------------------- */
// Store a flash message with a CSS class for bootstrap.
type Flash struct {
	Class   string
	Message string
}

// Creates a new flash item to be stored in the session, with an optional CSS
// class, which can be "success", "info", "warning", or "danger".
func NewFlash(message string, class ...string) Flash {
	if len(class) == 0 {
		class = []string{"info"}
	}
	return Flash{
		Class:   class[0],
		Message: message,
	}
}

// Retrieve flashes, wiping them from the session if they exist.
func (h *Handlers) flashes(r *http.Request, w http.ResponseWriter) []interface{} {
	flashes := h.session.Flashes()
	if len(flashes) > 0 {
		h.session.Save(r, w)
	}
	return flashes
}
