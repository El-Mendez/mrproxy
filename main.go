package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

const PORT string = ":3333"

func main() {
	//SetupProxy(PORT, nil)

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	go SetupProxy(PORT, p)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
