package json

import (
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
	"io"
	"net/http"
)

type Marshaler = easyjson.Marshaler

var MarshalToWriter = easyjson.MarshalToWriter

func MarshalHTTP(w http.ResponseWriter, obj Marshaler) error {
	var jw jwriter.Writer
	obj.MarshalEasyJSON(&jw)
	if jw.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, jw.Error.Error())
		return jw.Error
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", jw.Size()))
	_, err := jw.DumpTo(w)
	return err
}
