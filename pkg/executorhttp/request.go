package executorhttp

import (
	"encoding/base64"
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

func RequestPayload(r *Request) []byte {
	out, _ := json.Marshal(r)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(out)))
	base64.StdEncoding.Encode(dst, out)
	return dst
}
