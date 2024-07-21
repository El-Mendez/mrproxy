package main

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"mrproxy/shared"
	"net/http"
	"os"
	"time"
)

func setupProxy(port string, program *tea.Program) {
	http.HandleFunc("/", handleIncomingGenerator(program))

	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func handleIncomingGenerator(program *tea.Program) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleIncoming(program, w, r)
	}
}

func handleIncoming(program *tea.Program, w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	request := shared.Request{
		Query:    r.RequestURI,
		Method:   r.Method,
		Duration: time.Since(start),
		Headers:  r.Header,
	}

	if program != nil {
		program.Send(incomingMsg{request: &request})
	}

	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Content-Type") == "text/plain" {
		s, _ := io.ReadAll(r.Body)
		request.Body = s
	}

	request.Status = 200
	request.Duration = time.Since(start)
	io.WriteString(w, "This is my website!\n")
}
