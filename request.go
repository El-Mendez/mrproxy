package main

import "time"

type RequestHeader struct {
	key string
	val []string
}
type JsonField struct {
	key string
	val interface{}
}

type Request struct {
	query    string
	method   string
	headers  []RequestHeader
	duration time.Duration
	status   uint
	body     interface{}
}

func (r *Request) FilterValue() string {
	return r.query
}
