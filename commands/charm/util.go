package charm

import "github.com/charmbracelet/lipgloss"

func Margin(i ...int) Style {
	return NewStyle(lipgloss.NewStyle().Margin(i...))
}

func Indent(level int) Style {
	return Margin(0, 0, 0, level)
}