package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"net/url"
	"os"
)

const thePORT string = ":3333"

func main() {
	p := tea.NewProgram(initialModel(thePORT), tea.WithAltScreen())

	var target string
	var port string
	if len(os.Args) > 1 {
		target = os.Args[1]
	} else {
		target = "http://localhost:3000"
	}

	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[2])
	} else {
		port = thePORT
	}

	proxyUrl, err := url.Parse(target)
	if err != nil {
		fmt.Println("URL Error:", err)
		os.Exit(1)
	}
	go setupProxy(port, p, proxyUrl)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
