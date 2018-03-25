package model

import "time"
import "wubba/lubba/apub/json"

type XMLContent struct {
	Data string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type XMLObject struct {
	Type      string     `xml:"activity:object-type"`
	ID        string     `xml:"id"`
	Title     string     `xml:"title"`
	Verb      string     `xml:"activity:verb"`
	Published time.Time  `xml:"published"`
	Updated   time.Time  `xml:"updated"`
	Content   XMLContent `xml:"content"`
	Links     []Link     `xml:"link"`
	Object    *XMLObject `xml:"activity:object"`
}

type apObject struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Object struct {
	Type    string
	ID      string
	Content string
	Name    string
}

func (o *Object) UnmarshalJSON(data []byte) error {
	var ap apObject
	err := json.Unmarshal(data, &ap)
	if err == nil {
		o.Type = ap.Type
		o.ID = ap.ID
		o.Content = ap.Content
		o.Name = ap.Name
	}
	return err
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(apObject{
		Context: ActivityStreams,
		Type:    o.Type,
		ID:      o.ID,
		Content: o.Content,
		Name:    o.Name,
	})
}
