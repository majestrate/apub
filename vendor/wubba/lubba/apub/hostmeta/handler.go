package hostmeta

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Handler struct {
	Hostname string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/xrd+xml; charset=utf-8")
	buff := new(bytes.Buffer)
	io.WriteString(buff, strings.Trim(xml.Header, "\n"))
	fmt.Fprintf(buff, `<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0"><Link rel="lrdd" template="https://%s/.well-known/webfinger?resource={uri}" type="application/xrd+xml" /></XRD>`, h.Hostname)
	io.Copy(w, buff)
}
