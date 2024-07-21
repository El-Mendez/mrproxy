package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"net/url"
	"os"
)

const thePORT string = ":3333"

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	proxyUrl, err := url.Parse("http://localhost:3000")
	if err != nil {
		fmt.Println("URL Error:", err)
		os.Exit(1)
	}
	go setupProxy(thePORT, p, proxyUrl)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
