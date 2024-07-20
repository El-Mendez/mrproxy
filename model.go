package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const defaultWidth = 20
const defaultHeight = 14

type incomingMsg struct{ request *Request }
type model struct {
	list     list.Model
	width    int
	height   int
	selected *Request
}

func initialModel() model {
	return model{
		list: list.New(
			make([]list.Item, 0),
			requestRenderer{},
			defaultWidth,
			defaultHeight,
		),
		width:    defaultWidth,
		height:   defaultHeight,
		selected: nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.selected == nil {
			m.list.SetSize(msg.Width, msg.Height)
		} else {
			m.list.SetSize(msg.Width/2, msg.Height)
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if selected, ok := m.list.SelectedItem().(*Request); ok {
				m.selected = selected
				m.list.SetWidth(m.width / 2)
				return m, nil
			}
		}

	case incomingMsg:
		m.list.InsertItem(0, msg.request)
		if m.list.Index() != 0 {
			m.list.CursorDown()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.selected != nil {
		return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), fmt.Sprintf("%+v", m.selected))
	}
	if len(m.list.Items()) == 0 {
		return "No items yet! :)"
	}
	return m.list.View()
}
