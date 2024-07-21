package shared

import "github.com/charmbracelet/lipgloss"

var (
	HighlightColor  = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	HighlightedText = lipgloss.NewStyle().Foreground(HighlightColor).Bold(true)
)
