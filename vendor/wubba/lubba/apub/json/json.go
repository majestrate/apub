package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Marshaler = json.Marshaler
type Unmarshaler = json.Unmarshaler

var Marshal = json.Marshal
var Unmarshal = json.Unmarshal
var NewDecoder = json.NewDecoder
var NewEncoder = json.NewEncoder

func MarshalHTTP(w http.ResponseWriter, obj Marshaler) error {
	buff := new(Buffer)
	enc := NewEncoder(buff)
	err := enc.Encode(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return err
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buff.Len()))
	_, err = io.Copy(w, buff)
	return err
}
