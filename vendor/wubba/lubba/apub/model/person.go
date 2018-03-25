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
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	PublicKey PublicKey `json:"publicKey"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(apPerson{
		Context: ActivityStreams,
		Type:    PersonType,
		Name:    p.Name,
		URL:     p.URL,
		ID:      p.URL,
		PublicKey: PublicKey{
			ID:    p.URL + "#main-key",
			Owner: p.URL,
			Key:   p.PublicKey,
		},
	})
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var ap apPerson
	err := json.Unmarshal(data, &ap)
	if err == nil {
		p.URL = ap.ID
		p.Name = ap.Name
		p.PublicKey = ap.PublicKey.Key
	}
	return err
}
