package requestTabs

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	shared "mrproxy/shared"
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
	tabStyle  = lipgloss.NewStyle().BorderForeground(shared.HighlightColor).Padding(0, 1)
	restColor = lipgloss.NewStyle().Foreground(shared.HighlightColor)

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

	insideStyle = lipgloss.NewStyle().
			Align(lipgloss.Top)

	windowStyle = lipgloss.NewStyle().
			BorderForeground(shared.HighlightColor).
			Align(lipgloss.Top).
			Border(lipgloss.NormalBorder()).
			UnsetBorderTop()
)

type Model struct {
	requestFocused bool
	reqViewport    viewport.Model
	resViewport    viewport.Model
	request        *shared.Request
	reqContent     string
	resContent     string
}

func New(request *shared.Request) Model {
	reqV := viewport.New(0, 0)
	reqV.HighPerformanceRendering = false

	resV := viewport.New(0, 0)
	resV.HighPerformanceRendering = false
	return Model{
		true,
		reqV,
		resV,
		request,
		"",
		"",
	}
}
func (m *Model) SetWidth(width int) {
	// account for padding
	windowStyle.Width(width - 2)
	m.reqViewport.Width = width - 2
	m.resViewport.Width = width - 2
	insideStyle.Width(width - 2)

	m.resViewport.SetContent(insideStyle.Render(m.resContent))
	m.reqViewport.SetContent(insideStyle.Render(m.reqContent))
}

func (m *Model) SetHeight(height int) {
	// account for padding and the tabs height themselves
	windowStyle.Height(height - 4)
	m.reqViewport.Height = height - 4
	m.resViewport.Height = height - 4
	insideStyle.Height(height - 6)

	m.resViewport.SetContent(insideStyle.Render(m.resContent))
	m.reqViewport.SetContent(insideStyle.Render(m.reqContent))
}

func (m *Model) SetRequest(request *shared.Request) {
	m.request = request
	m.reqContent = renderRequest(request.Method, request.Query, request.ReqHeaders, request.ReqBody)
	m.reqViewport.SetContent(m.reqContent)
	m.reqViewport.SetYOffset(0)

	m.resContent = renderRequest(request.Method, request.Query, request.ResHeaders, request.ResBody)
	m.resViewport.SetContent(m.resContent)
	m.resViewport.SetYOffset(0)

	m.resViewport.SetContent(insideStyle.Render(m.resContent))
	m.reqViewport.SetContent(insideStyle.Render(m.reqContent))
}

func (m *Model) SetResponse(request *shared.Request) {
	m.resContent = renderRequest(request.Method, request.Query, request.ResHeaders, request.ResBody)
	m.resViewport.SetContent(m.resContent)
	m.resViewport.SetYOffset(0)

	m.resViewport.SetContent(insideStyle.Render(m.resContent))
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
	if m.requestFocused {
		m.reqViewport, cmd = m.reqViewport.Update(msg)
	} else {
		m.resViewport, cmd = m.resViewport.Update(msg)
	}
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
			windowStyle.Render(m.reqViewport.View()),
		)
	} else {
		doc.WriteString(
			windowStyle.Render(m.resViewport.View()),
		)
	}

	return doc.String()
}
