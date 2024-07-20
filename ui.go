package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type incomingMsg struct{ request *Request }
type model struct {
	requests []*Request
}

func initialModel() model {
	return model{
		requests: make([]*Request, 0),
	}
}
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case incomingMsg:
		m.requests = append(m.requests, msg.request)
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("program is running! We've gotten %d request so far!", len(m.requests))
}
