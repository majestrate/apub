package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func StorePrivateKey(keyfile string, k *rsa.PrivateKey) error {
	d := x509.MarshalPKCS1PrivateKey(k)
	block := &pem.Block{
		Bytes: d,
		Headers: map[string]string{
			"Comment": "I've turned myself into an RSA Private key!!! I'm crypto RIIIIICK!!!!!",
		},
		Type: "RSA PRIVATE KEY",
	}
	return ioutil.WriteFile(keyfile, pem.EncodeToMemory(block), 0600)
}

func LoadPrivateKey(keyfile string) (k *rsa.PrivateKey, err error) {
	var data []byte
	data, err = ioutil.ReadFile(keyfile)
	if err == nil {
		var block *pem.Block
		block, _ = pem.Decode(data)
		k, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	return
}
