package shared

import "time"

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
	Headers  []Header
	Duration time.Duration
	Status   uint
	Body     interface{}
}

func (r *Request) FilterValue() string {
	return r.Query
}
