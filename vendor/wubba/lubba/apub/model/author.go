package model

type Author struct {
	ID           string `xml:"id"`
	ObjectType   string `xml:"activity:object"`
	URI          string `xml:"uri"`
	Summary      string `xml:"summary"`
	Name         string `xml:"name"`
	Links        []Link `xml:"link"`
	AP           bool   `xml:"ap_enabled"`
	PreferedName string `xml:"poco:preferredUsername"`
	DisplayName  string `xml:"poco:displayName"`
	Note         string `xml:"poco:note"`
}
