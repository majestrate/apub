package model

import "wubba/lubba/apub/json"

type APBase interface {
	json.Marshaler
	json.Unmarshaler
	APType() string
}
