package model

import (
	"crypto/rsa"
	"wubba/lubba/apub/json"
)

type Person struct {
	Name      string
	URL       string
	PublicKey *rsa.PublicKey
}

type apPerson struct {
	Context   string    `json:"@context"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	PublicKey PublicKey `json:"publicKey"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(apPerson{
		Context: ActivityStreams,
		Type:    PersonType,
		Name:    p.Name,
		URL:     p.URL,
		PublicKey: PublicKey{
			ID:    p.URL + "#main-key",
			Owner: p.URL,
			Key:   p.PublicKey,
		},
	})
}
