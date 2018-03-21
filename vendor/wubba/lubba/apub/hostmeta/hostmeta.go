package hostmeta

import (
	"encoding/xml"
)

type MetaInfo struct {
	Template string
}

type miXRDLink struct {
	Rel      string `xml:"rel,attr"`
	Template string `xml:"template,attr"`
	Type     string `xml:"type,attr"`
}

type XRD struct {
	Link miXRDLink `xml:"Link"`
	NS   string    `xml:"xmlns,attr"`
}

func (info *MetaInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var x XRD
	err := d.Decode(&x)
	if err == nil {
		info.Template = x.Link.Template
	}
	return err
}

func (info *MetaInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(&XRD{
		NS: XMLNS,
		Link: miXRDLink{
			Rel:      XMLRel,
			Template: info.Template,
			Type:     MimeType,
		},
	})
}
