package executorhttp

import (
	"encoding/json"
)

type Request struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   []byte            `json:"body"`
}

func NewRequest(payload []byte) (*Request, error) {
	r := &Request{}

	if err := json.Unmarshal(payload, r); err != nil {
		return nil, err
	}

	return r, nil
}

func RequestPayload(r Request) []byte {
	out, _ := json.Marshal(r)
	return out
}
