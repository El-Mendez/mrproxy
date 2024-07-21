package requestTabs

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	request2 "mrproxy/shared"
	"strings"
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	tabStyle  = lipgloss.NewStyle().BorderForeground(request2.HighlightColor).Padding(0, 1)
	restColor = lipgloss.NewStyle().Foreground(request2.HighlightColor)

	activeRequestStyle = tabStyle.
				Copy().
				Border(tabBorderWithBottom("│", " ", "└")).
				Bold(true)
	inactiveRequestStyle = tabStyle.
				Copy().
				Border(tabBorderWithBottom("├", "─", "┴"))

	activeResponseStyle = tabStyle.
				Copy().
				Border(tabBorderWithBottom("┘", " ", "└")).
				Bold(true)
	inactiveResponseStyle = tabStyle.
				Copy().
				Border(tabBorderWithBottom("┴", "─", "┴"))

	windowStyle = lipgloss.NewStyle().
			BorderForeground(request2.HighlightColor).
			Align(lipgloss.Top).
			Border(lipgloss.NormalBorder()).
			UnsetBorderTop()
)

type Model struct {
	requestFocused bool
	viewport       viewport.Model
	request        *request2.Request
}

func New(request *request2.Request) Model {
	v := viewport.New(0, 0)
	v.HighPerformanceRendering = false
	return Model{
		true,
		v,
		request,
	}
}
func (m *Model) SetWidth(width int) {
	// account for padding
	windowStyle.Width(width - 2)
	m.viewport.Width = width - 2
}

func (m *Model) SetHeight(height int) {
	// account for padding and the tabs height themselves
	windowStyle.Height(height - 4)
	m.viewport.Height = height - 4
}

func (m *Model) SetRequest(request *request2.Request) {
	m.request = request
	m.viewport.SetContent(renderRequest(request))
	m.viewport.SetYOffset(0)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.requestFocused = !m.requestFocused
			break
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	doc := strings.Builder{}

	var row string
	if m.requestFocused {
		request := activeRequestStyle.Render("Request")
		response := inactiveResponseStyle.Render("Response")
		row = lipgloss.JoinHorizontal(lipgloss.Top, request, response)
	} else {
		request := inactiveRequestStyle.Render("Request")
		response := activeResponseStyle.Render("Response")
		row = lipgloss.JoinHorizontal(lipgloss.Top, request, response)
	}
	doc.WriteString(row)
	doc.WriteString(restColor.Render(strings.Repeat("─", windowStyle.GetWidth()-lipgloss.Width(row)+1) + "┐"))
	doc.WriteString("\n")

	if m.requestFocused {
		doc.WriteString(
			windowStyle.Render(m.viewport.View()),
		)
	} else {
		doc.WriteString(
			windowStyle.Render(m.viewport.View()),
		)
	}

	return doc.String()
}
