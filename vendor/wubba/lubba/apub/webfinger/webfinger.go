package webfinger

import (
	"net/http"
	"wubba/lubba/apub"
	"wubba/lubba/apub/util"
	"wubba/lubba/apub/xml"
)

type WebFinger struct {
	Finder func(string) (apub.UserInfo, error)
}

func (wf *WebFinger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if wf.Finder == nil {
		// no finder
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resource := r.URL.Query().Get("resource")
	u, err := wf.Finder(resource)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if u == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// set CORS header
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if util.WantsJSON(r) {
		w.Header().Set("Content-Type", JSONType+"; encoding=utf-8")
		// TODO: implement
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		// fallback to xml
		w.Header().Set("Content-Type", XMLType+"; encoding=utf-8")
		xml.MarshalHTTP(w, &XRD{
			NS:      XMLNS,
			Subject: u.WebFingerSubject(),
			Alias:   u.WebFingerAlias(),
			Links:   u.WebFingerLinks(),
		})
	}
}
