package util

import (
	"crypto/rsa"
	"math/big"
)

func EncodeRSAPublicKey(pub rsa.PublicKey) string {
	return "RSA." + BigIntToBase64(pub.N) + "." + BigIntToBase64(big.NewInt(int64(pub.E)))
}
