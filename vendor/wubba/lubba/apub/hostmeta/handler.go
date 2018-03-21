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
	hostname string
}

func NewHandler(hostname string) *Handler {
	return &Handler{
		hostname: hostname,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xrd+xml; charset=utf-8")
	buff := new(bytes.Buffer)
	io.WriteString(buff, strings.Trim(xml.Header, "\n"))
	fmt.Fprintf(buff, `<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0"><Link rel="lrdd" template="https://%s/.well-known/webfinger?resource={uri}" type="application/xrd+xml" /></XRD>`, h.hostname)
	io.Copy(w, buff)
}
