package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"strings"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

type requestRenderer struct{}

func (d requestRenderer) Height() int {
	return 1
}
func (d requestRenderer) Spacing() int {
	return 0
}
func (d requestRenderer) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d requestRenderer) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	r, ok := listItem.(*Request)
	if !ok {
		return
	}

	var str string
	if r.status == 0 {
		str = fmt.Sprintf("--- %s %s", r.method, r.query)
	} else {
		str = fmt.Sprintf("%d %s %s", r.status, r.method, r.query)
	}

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, ", "))
		}
	}

	fmt.Fprint(w, fn(str))
}
