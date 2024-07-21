package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"mrproxy/requestList"
	"mrproxy/requestTabs"
	request2 "mrproxy/shared"
)

const defaultWidth = 20
const defaultHeight = 14
const minWidth = 80

type incomingMsg struct{ request *request2.Request }
type updatedMsg struct{ request *request2.Request }
type model struct {
	list     requestList.Model
	tabs     requestTabs.Model
	selected *request2.Request
	width    int
	height   int
}

func initialModel() model {
	return model{
		list: requestList.New(
			make([]list.Item, 0),
			defaultWidth,
			defaultHeight,
		),
		tabs:     requestTabs.New(nil),
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
		m.list.SetHeight(msg.Height)
		m.tabs.SetHeight(msg.Height)

		if m.selected == nil {
			m.list.SetWidth(msg.Width)
		} else {
			m.list.SetWidth(msg.Width / 2)
		}

		if m.width <= minWidth {
			m.tabs.SetWidth(msg.Width)
		} else {
			m.tabs.SetWidth(msg.Width - msg.Width/2)
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if m.selected != nil {
				m.selected = nil
				m.list.SetFollow(false)
				m.list.SetWidth(m.width)
				return m, nil
			} else {
				return m, tea.Quit
			}
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if selected, ok := m.list.SelectedItem().(*request2.Request); ok {
				m.list.SetFollow(true)
				m.tabs.SetRequest(selected)
				m.selected = selected
				m.list.SetWidth(m.width / 2)
				return m, nil
			}
		}

	case incomingMsg:
		return m, m.list.InsertItem(0, msg.request)

	case updatedMsg:
		if m.selected == msg.request {
			m.tabs.SetResponse(msg.request)
		}
	}

	var cmd tea.Cmd
	if m.selected == nil {
		m.list, cmd = m.list.Update(msg)
	} else {
		m.tabs, cmd = m.tabs.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.selected != nil {
		if m.width <= minWidth {
			return m.tabs.View()
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), m.tabs.View())
	}
	if len(m.list.Items()) == 0 {
		return "No items yet! :)"
	}
	return m.list.View()
}
