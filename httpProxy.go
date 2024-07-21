package main

import (
	"encoding/json"
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
		Headers:  make([]shared.Header, len(r.Header)),
	}

	i := 0
	for key, val := range r.Header {
		request.Headers[i] = shared.Header{
			Key: key,
			Val: val,
		}
		i++
	}
	if program != nil {
		program.Send(incomingMsg{request: &request})
	}

	if r.Header.Get("Content-Type") == "application/json" {
		var body interface{}
		d := json.NewDecoder(r.Body)
		err := d.Decode(&body)
		if err == nil {
			request.Body = parseJsonValue(body)
		}
	} else if r.Header.Get("Content-Type") == "text/plain" {
		s, _ := io.ReadAll(r.Body)
		request.Body = string(s)
	}

	request.Status = 200
	request.Duration = time.Since(start)
	io.WriteString(w, "This is my website!\n")
}

func parseJsonValue(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		return v
	case float64:
		return value
	case bool:
		return v
	case []interface{}:
		for i := 0; i < len(v); i++ {
			v[i] = parseJsonValue(v[i])
		}
		return value
	case map[string]interface{}:
		i := 0
		result := make([]shared.JsonField, len(v))
		for k, field := range v {
			result[i] = shared.JsonField{
				Key: k,
				Val: parseJsonValue(field),
			}
			i++
		}
		return result
	}

	fmt.Printf("unknown type %T \n", value)
	return nil
}
