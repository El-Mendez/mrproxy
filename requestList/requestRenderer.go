package requestList

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"mrproxy/shared"
	"strings"
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
	r, ok := listItem.(*shared.Request)
	if !ok {
		return
	}

	var str string
	if r.Status == 0 {
		str = fmt.Sprintf("--- %s %s", r.Method, r.Query)
	} else {
		str = fmt.Sprintf("%d %s %s", r.Status, r.Method, r.Query)
	}

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, ", "))
		}
	}

	fmt.Fprint(w, fn(str))
}
