package model

import (
	"crypto/rsa"
	"wubba/lubba/apub/json"
	"wubba/lubba/apub/util"
)

type PublicKey struct {
	Owner string
	ID    string
	Key   *rsa.PublicKey
}

func (k PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"owner":        k.Owner,
		"id":           k.ID,
		"publicKeyPem": string(util.DumpPubkey(k.Key)),
	})
}

func (k *PublicKey) UnmarshalJSON(data []byte) error {
	j := make(map[string]interface{})
	err := json.Unmarshal(data, &j)
	if err == nil {
		k.Owner = j["owner"].(string)
		k.ID = j["id"].(string)
		keypem := j["publicKeyPem"].(string)
		k.Key, err = util.LoadPubkey([]byte(keypem))
	}
	return err
}
