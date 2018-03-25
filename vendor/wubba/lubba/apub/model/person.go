package model

type Person struct {
	Context      []interface{} `json:"@context"`
	Type         string        `json:"type"`
	Name         string        `json:"name"`
	ID           string        `json:"id"`
	URL          string        `json:"url"`
	Outbox       string        `json:"outbox"`
	Inbox        string        `json:"inbox"`
	PublicKey    PublicKey     `json:"publicKey"`
	PreferedName string        `json:"preferredUsername"`
	ManualFollow bool          `json:"manuallyApprovesFollowers"`
	Image        Image         `json:"image"`
	Icon         Image         `json:"icon"`
	Followers    string        `json:"followers"`
	Endpoints    EndpointMap   `json:"endpoints"`
}
