package xml

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

// MarshalHTTP writes a proper xml document to w in a "compliant" (lol) way
func MarshalHTTP(w http.ResponseWriter, obj interface{}) error {
	io.WriteString(w, strings.Trim(xml.Header, "\n"))
	return xml.NewEncoder(w).Encode(obj)
}
