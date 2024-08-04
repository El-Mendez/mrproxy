package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"mrproxy/requestList"
	"mrproxy/requestTabs"
	shared "mrproxy/shared"
	"strings"
)

const defaultWidth = 20
const defaultHeight = 14
const minWidth = 80

type incomingMsg struct{ request *shared.Request }
type updatedMsg struct{ request *shared.Request }
type model struct {
	addresses []string
	fullWidth bool
	list      requestList.Model
	tabs      requestTabs.Model
	selected  *shared.Request
	width     int
	height    int
}

var windowStyle = lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Italic(true)

func initialModel(addresses []string) model {
	return model{
		addresses: addresses,
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

func (m *model) UpdatePanels(width int, height int, fullWidth bool, request *shared.Request) {
	m.width = width
	m.height = height
	windowStyle = windowStyle.Width(width).Height(height)
	m.fullWidth = fullWidth
	m.selected = request

	if request == nil {
		m.list.SetFollow(false)
		m.list.SetWidth(width)
		m.list.SetHeight(height)
	} else {
		m.list.SetFollow(true)
		if fullWidth || width <= minWidth {
			m.list.SetWidth(width)
			m.list.SetHeight(height)
			m.tabs.SetWidth(width)
			m.tabs.SetHeight(height)
		} else {
			m.list.SetWidth(width / 2)
			m.list.SetHeight(height)
			m.tabs.SetWidth(width - width/2)
			m.tabs.SetHeight(height)
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.UpdatePanels(msg.Width, msg.Height, m.fullWidth, m.selected)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "c":
			m.list.Clear()
			m.UpdatePanels(m.width, m.height, m.fullWidth, nil)
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.selected != nil {
				m.UpdatePanels(m.width, m.height, m.fullWidth, nil)
				return m, nil
			}
			return m, nil
		case "f":
			m.UpdatePanels(m.width, m.height, !m.fullWidth, m.selected)
			return m, nil
		case "enter":
			if m.selected == nil {
				if selected, ok := m.list.SelectedItem().(*shared.Request); ok {
					m.UpdatePanels(m.width, m.height, m.fullWidth, selected)
					m.tabs.SetRequest(selected)
					return m, nil
				}
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
		if m.width <= minWidth || m.fullWidth {
			return m.tabs.View()
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), m.tabs.View())
	}
	if len(m.list.Items()) == 0 {
		return windowStyle.Render(fmt.Sprintf("No hay solicitudes todavía.\n%s", strings.Join(m.addresses, "\n")))
	}
	return m.list.View()
}
