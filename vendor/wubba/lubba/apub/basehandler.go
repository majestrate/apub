package apub

import (
	"io"
	"net/http"
	"strings"
)

type UserHandler interface {
	RoutePath() string
	ServeUser(UserInfo, http.ResponseWriter, *http.Request)
}

type BaseHandler struct {
	Finder      func(string) (UserInfo, error)
	UserHandler UserHandler
}

func (h *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Finder == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if h.UserHandler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	route := h.UserHandler.RoutePath()
	username := r.URL.Path[len(route):]
	for strings.HasPrefix(username, "/") {
		username = username[1:]
	}
	u, err := h.Finder(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	if u == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.UserHandler.ServeUser(u, w, r)

}
