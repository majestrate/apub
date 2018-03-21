package hostmeta

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	// write xml header
	io.WriteString(w, strings.Trim(xml.Header, "\n"))
	// encode xml document
	enc := xml.NewEncoder(w)
	enc.Encode(&metainfo)
}
