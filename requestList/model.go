package requestList

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

	windowStyle = lipgloss.NewStyle().Align(lipgloss.Top)
)

type Model struct {
	list   list.Model
	follow bool
}

func New(items []list.Item, width int, height int) Model {
	list := list.New(
		items,
		requestRenderer{},
		width,
		height,
	)
	list.SetShowHelp(false)

	return Model{
		follow: true,
		list:   list,
	}
}

func (m *Model) SetFollow(follow bool) {
	m.follow = follow
}

func (m *Model) SetWidth(width int) {
	windowStyle.Width(width)
	m.list.SetWidth(width)
	itemStyle = itemStyle.MaxWidth(width)
	selectedItemStyle = selectedItemStyle.MaxWidth(width)
}

func (m *Model) SetHeight(height int) {
	windowStyle.Height(height)
	m.list.SetHeight(height - 1)
}

func (m *Model) Clear() {
	m.list.SetItems(make([]list.Item, 0))
}

func (m Model) SelectedItem() list.Item {
	return m.list.SelectedItem()
}

func (m Model) Items() []list.Item {
	return m.list.Items()
}

func (m *Model) InsertItem(index int, item list.Item) tea.Cmd {
	cmd := m.list.InsertItem(index, item)
	if index != 0 || m.follow {
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
	return windowStyle.Render(m.list.View())
}
