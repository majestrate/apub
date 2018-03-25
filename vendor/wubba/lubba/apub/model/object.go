package model

import "time"
import "wubba/lubba/apub/json"

type Content struct {
	Data string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type Object struct {
	Type      string    `xml:"activity:object-type"`
	ID        string    `xml:"id"`
	Title     string    `xml:"title"`
	Verb      string    `xml:"activity:verb"`
	Published time.Time `xml:"published"`
	Updated   time.Time `xml:"updated"`
	Content   Content   `xml:"content"`
	Links     []Link    `xml:"link"`
	Object    *Object   `xml:"activity:object"`
}

func (o *Object) APType() string {
	return ObjectType
}

type apObject struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (o *Object) UnmarshalJSON(data []byte) error {
	var ap apObject
	err := json.Unmarshal(data, &ap)
	if err == nil {
		o.Type = ap.Type
		o.Title = ap.Name
		o.Content.Data = ap.Content
	}
	return err
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(apObject{
		Context: ActivityStreams,
		Type:    ObjectType,
		ID:      o.ID,
		Name:    o.Title,
		Content: o.Content.Data,
	})
}
