package main

import (
	"bytes"
	tea "github.com/charmbracelet/bubbletea"
	"mrproxy/shared"
	"net/http"
	"strings"
	"time"
)

type ProxyResponse struct {
	statusCode int
	w          http.ResponseWriter
	request    *shared.Request
	start      time.Time
	buffer     *bytes.Buffer
}

func NewProxyResponse(w http.ResponseWriter, request *shared.Request) *ProxyResponse {
	return &ProxyResponse{
		0,
		w,
		request,
		time.Now(),
		new(bytes.Buffer),
	}
}

func (p *ProxyResponse) Header() http.Header {
	return p.w.Header()
}
func (p *ProxyResponse) Write(bytes []byte) (int, error) {
	p.buffer.Write(bytes)
	return p.w.Write(bytes)
}
func (p *ProxyResponse) WriteHeader(statusCode int) {
	p.statusCode = statusCode
	p.w.WriteHeader(statusCode)
}

func (p *ProxyResponse) Done(program *tea.Program) {
	p.request.Duration = time.Since(p.start)
	if p.statusCode == 0 {
		p.request.Status = 200
	} else {
		p.request.Status = uint(p.statusCode)
	}

	p.request.ResHeaders = p.w.Header()
	contentType := p.w.Header().Get("Content-Type")
	if strings.Contains(contentType, "text/html") || strings.Contains(contentType, "application/json") || strings.Contains(contentType, "text/plain") {
		p.request.ResBody = p.buffer.Bytes()
	}
	program.Send(updatedMsg{request: p.request})
}
