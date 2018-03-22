package util

import (
	"encoding/base64"
	"math/big"
)

func BigIntToBase64(i *big.Int) string {
	return base64.URLEncoding.EncodeToString(i.Bytes())
}
