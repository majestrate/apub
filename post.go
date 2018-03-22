package apub

import (
	"encoding/xml"
	"html"
	"time"
	"wubba/lubba/apub"
	"wubba/lubba/apub/model"
)

type RecipInfo struct {
	Users []string
}

type Post struct {
	Message string
	To      RecipInfo
	From    string
	Posted  time.Time
	Updated time.Time
	Feed    string
	Self    string
}

func (p *Post) FeedURL() string {
	return p.Feed
}

func (p *Post) PostedAt() time.Time {
	return p.Posted
}

func (p *Post) Text() string {
	return html.EscapeString(p.Message)
}

func (p *Post) HTML() string {
	return p.Message
}

func (p *Post) Mentions() []string {
	return p.To.Users
}

func (p *Post) PostURL() string {
	return p.Self
}

func (p *Post) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return nil
}

type entry struct {
	model.Object
}

func (p *Post) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(&entry{
		model.Object{
			Type: model.TypeNote,
			Verb: model.VerbPost,
			ID:   p.PostURL(),
			Content: model.Content{
				Data: p.HTML(),
				Type: "html",
			},
			Updated:   p.Updated,
			Published: p.Posted,
			Title:     "New note by " + p.From,
			Links: []model.Link{
				model.Link{
					Href: p.FeedURL(),
					Rel:  "self",
					Type: apub.AtomMime,
				},
				model.Link{
					Href: p.PostURL(),
					Rel:  "alternate",
					Type: apub.HTMLMime,
				},
			},
		},
	}, start)
}
