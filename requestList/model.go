package requestList

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type Model struct {
	list    list.Model
	focused bool
}

func New(items []list.Item, width int, height int) Model {
	return Model{
		focused: true,
		list: list.New(
			items,
			requestRenderer{},
			width,
			height,
		),
	}
}

func (m *Model) SetFocus(focused bool) {
	m.focused = focused
}

func (m *Model) SetWidth(width int) {
	m.list.SetWidth(width)
	itemStyle = itemStyle.MaxWidth(width)
	selectedItemStyle = selectedItemStyle.MaxWidth(width)
}

func (m *Model) SetHeight(height int) {
	m.list.SetHeight(height - 1)
}

func (m Model) SelectedItem() list.Item {
	return m.list.SelectedItem()
}

func (m Model) Items() []list.Item {
	return m.list.Items()
}

func (m *Model) InsertItem(index int, item list.Item) tea.Cmd {
	cmd := m.list.InsertItem(index, item)
	if index != 0 {
		m.list.CursorDown()
	}
	return cmd
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}
