package model

import "time"

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
}
