package shared

import (
	"net/http"
	"time"
)

type Header struct {
	Key string
	Val []string
}
type JsonField struct {
	Key string
	Val interface{}
}

type Request struct {
	Query    string
	Method   string
	Duration time.Duration
	Status   uint

	ReqHeaders http.Header
	ReqBody    []byte

	ResHeaders http.Header
	ResBody    []byte
}

func (r *Request) FilterValue() string {
	return r.Query
}
