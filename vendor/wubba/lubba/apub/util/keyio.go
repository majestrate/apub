package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func DumpPubkey(k *rsa.PublicKey) (data []byte) {
	d := x509.MarshalPKCS1PublicKey(k)
	block := &pem.Block{
		Bytes: d,
		Type:  "RSA PUBLIC KEY",
	}
	return pem.EncodeToMemory(block)
}

func LoadPubkey(data []byte) (k *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	k, err = x509.ParsePKCS1PublicKey(block.Bytes)
	return
}

func DumpPrivkey(k *rsa.PrivateKey) (data []byte) {
	d := x509.MarshalPKCS1PrivateKey(k)
	block := &pem.Block{
		Bytes: d,
		Type:  "RSA PRIVATE KEY",
	}
	return pem.EncodeToMemory(block)
}

func LoadPrivateKey(data []byte) (k *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	k, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	return
}

func StorePrivateKeyFile(keyfile string, k *rsa.PrivateKey) error {
	return ioutil.WriteFile(keyfile, DumpPrivkey(k), 0600)
}

func LoadPrivateKeyFile(keyfile string) (k *rsa.PrivateKey, err error) {
	var data []byte
	data, err = ioutil.ReadFile(keyfile)
	if err == nil {
		k, err = LoadPrivateKey(data)
	}
	return
}
