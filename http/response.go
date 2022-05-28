package http

import "encoding/json"

type Respose struct {
	IsValid bool `json:"isValid"`
}

func (r Respose) Encode() []byte {
	payload, _ := json.Marshal(r)

	return payload
}
