package model

type Attachment struct {
	FileName string
	Content  string
	Href     string
}

type apAttachment struct {
	URL     string `json:"url"`
	Content string `json:"content"`
}
