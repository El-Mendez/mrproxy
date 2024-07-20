package main

import (
	"time"
)

const PORT string = ":3333"

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

func main() {
	SetupProxy(PORT, nil)

	//p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	//go SetupProxy(PORT, p)
	//if _, err := p.Run(); err != nil {
	//	fmt.Printf("Alas, there's been an error: %v", err)
	//	os.Exit(1)
	//}
}
