package webfinger

import "wubba/lubba/apub"

type Link = apub.Link

type XRD struct {
	NS      string `xml:"xmlns,attr"`
	Subject string `xml:"subject"`
	Alias   string `xml:"alias"`
	Links   []Link `xml:"link"`
}
