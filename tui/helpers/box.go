package helpers

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func TitledBox(title, body string, style lipgloss.Style) string {
	b := style.GetBorderStyle()

	inner := style.BorderTop(false)
	if h := style.GetHeight(); h > 0 {
		inner = inner.Height(h - 1)
	}
	rendered := inner.Render(body)

	boxW := lipgloss.Width(rendered)
	fill := boxW - lipgloss.Width(b.TopLeft) - lipgloss.Width(b.TopRight)

	label := " " + title + " "
	labelW := min(lipgloss.Width(label), fill)

	top := b.TopLeft + label + strings.Repeat(b.Top, fill-labelW) + b.TopRight
	top = lipgloss.NewStyle().
		Foreground(style.GetBorderTopForeground()).
		Render(top)

	return top + "\n" + rendered
}
