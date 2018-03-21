package hostmeta

import (
	"fmt"
	"net/http"
	"wubba/lubba/apub/xml"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var metainfo MetaInfo
	// set template
	metainfo.Template = fmt.Sprintf("https://%s/.well-known/webfinger?resource={uri}", r.Host)

	w.Header().Set("Content-Type", MimeType+"; charset=utf-8")
	// encode xml document
	xml.MarshalHTTP(w, &metainfo)
}
