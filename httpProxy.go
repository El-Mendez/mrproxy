package main

import (
	"bytes"
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"mrproxy/shared"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func setupProxy(port string, program *tea.Program, proxyUrl *url.URL) {
	http.HandleFunc("/", handleIncomingGenerator(program, proxyUrl))

	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func handleIncomingGenerator(program *tea.Program, proxyUrl *url.URL) func(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

	return func(w http.ResponseWriter, r *http.Request) {
		handleIncoming(program, proxy, w, r)
	}
}

func handleIncoming(program *tea.Program, proxy *httputil.ReverseProxy, w http.ResponseWriter, r *http.Request) {
	request := shared.Request{
		Query:      r.RequestURI,
		Method:     r.Method,
		ReqHeaders: r.Header,
	}

	if program != nil {
		program.Send(incomingMsg{request: &request})
	}

	contentType := r.Header.Get("Content-Type")
	// avoid cacheing responses (304 are not visible)
	r.Header.Del("If-Modified-Since")
	r.Header.Del("If-None-Match")
	if strings.Contains(contentType, "text/html") || strings.Contains(contentType, "application/json") || strings.Contains(contentType, "text/plain") {
		s, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(s))
		request.ReqBody = s
	}

	wp := NewProxyResponse(w, &request)
	proxy.ServeHTTP(wp, r)
	wp.Done(program)
}
