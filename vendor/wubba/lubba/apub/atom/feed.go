package atom

import (
	"encoding/xml"
	"time"
	"wubba/lubba/apub"
	"wubba/lubba/apub/model"
)

type Feed struct {
	Author    model.Author
	ID        string
	Title     string
	Logo      string
	Published time.Time
	Updated   time.Time
	Links     []model.Link
	Entries   []apub.Post
}

type feed struct {
	Author     model.Author `xml:"author"`
	ID         string       `xml:"id"`
	Title      string       `xml:"title"`
	Logo       string       `xml:"logo"`
	Published  time.Time    `xml:"published"`
	Updated    time.Time    `xml:"updated"`
	Links      []model.Link `xml:"link"`
	Entries    []apub.Post  `xml:"entry"`
	NS         string       `xml:"xmlns,attr"`
	NSThr      string       `xml:"xmlns:thr,attr"`
	NSActivity string       `xml:"xmlns:activity,attr"`
	NSPoco     string       `xml:"xmlns:poco,attr"`
	NSOstatus  string       `xml:"xmlns:ostatus,attr"`
	NSMedia    string       `xml:"xmlns:media,attr"`
}

func (f Feed) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&feed{
		Author:     f.Author,
		ID:         f.ID,
		Title:      f.Title,
		Logo:       f.Logo,
		Published:  f.Published,
		Updated:    f.Updated,
		Links:      f.Links,
		Entries:    f.Entries,
		NS:         apub.AtomXMLNS,
		NSThr:      apub.ThrXMLNS,
		NSActivity: apub.ActivityXMLNS,
		NSPoco:     apub.PocoXMLNS,
		NSOstatus:  apub.OStatusXMLNS,
		NSMedia:    apub.MediaXMLNS,
	})
}

func (f *Feed) AppendPost(post apub.Post) {
	f.Entries = append(f.Entries, post)
}

func (f *Feed) Posts() []apub.Post {
	return f.Entries
}
