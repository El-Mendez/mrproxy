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
	Headers  http.Header
	Duration time.Duration
	Status   uint
	Body     []byte
}

func (r *Request) FilterValue() string {
	return r.Query
}
