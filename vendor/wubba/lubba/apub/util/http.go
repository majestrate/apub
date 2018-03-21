package util

import (
	"net/http"
)

func WantsJSON(r *http.Request) bool {
	// TODO: verify this is correct
	return r.Header.Get("Accept") == "application/jrd+json"
}
