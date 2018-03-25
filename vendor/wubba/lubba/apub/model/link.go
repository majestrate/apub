package model

import "wubba/lubba/apub/json"

type Link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr,omitempty"`
}

type linkAP struct {
	Context   string `json:"@context"`
	Type      string `json:"type"`
	Href      string `json:"href"`
	MediaType string `json:"mediaType"`
	Name      string `json:"name"`
}

func (l *Link) APType() string {
	return LinkType
}

func (l *Link) UnmarshalJSON(data []byte) error {
	var ap linkAP
	err := json.Unmarshal(data, &ap)
	if err == nil {
		l.Href = ap.Href
		l.Type = ap.MediaType
		l.Rel = ap.Name
	}
	return err
}

func (l Link) MarshalJSON() ([]byte, error) {
	return json.Marshal(linkAP{
		Context:   ActivityStreams,
		Type:      LinkType,
		Href:      l.Href,
		MediaType: l.Type,
		Name:      l.Rel,
	})
}
