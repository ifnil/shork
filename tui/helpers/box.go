package helpers

import (
	"math"
	"strings"

	"charm.land/lipgloss/v2"
)

func TitledBox(title, body string, style lipgloss.Style) string {
	b := style.GetBorderStyle()

	inner := style.BorderTop(false).Render(body)

	boxW := lipgloss.Width(inner)
	fill := boxW - lipgloss.Width(b.TopLeft) - lipgloss.Width(b.TopRight)

	label := " " + title + " "
	labelW := min(lipgloss.Width(label), fill)

	top := b.TopLeft + label + strings.Repeat(b.Top, fill-labelW) + b.TopRight
	top = lipgloss.NewStyle().
		Foreground(style.GetBorderTopForeground()).
		Render(top)

	return top + "\n" + inner
}

func Percent(w int, p float64) int {
	return int(math.Round(float64(w) * p))
}
