package model

import "wubba/lubba/apub/json"

type Image struct {
	URL string
}

func (i Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"type": ImageType,
		"url":  i.URL,
	})
}

func (i *Image) UnmarshalJSON(data []byte) error {
	j := make(map[string]interface{})
	err := json.Unmarshal(data, &j)
	if err == nil {
		i.URL = j["url"].(string)
	}
	return err
}
