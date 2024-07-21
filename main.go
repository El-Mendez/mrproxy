package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

const thePORT string = ":3333"

func main() {
	//setupProxy(thePORT, nil)

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	go setupProxy(thePORT, p)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
