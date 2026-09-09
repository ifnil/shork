package hosts

import (
	"charm.land/bubbles/v2/list"
	"charm.land/lipgloss/v2"
)

const padding int = 1

type styles struct {
	title        lipgloss.Style
	item         lipgloss.Style
	selectedItem lipgloss.Style
	pagination   lipgloss.Style
	help         lipgloss.Style
	quitText     lipgloss.Style
}

func newStyles(darkBG bool) styles {
	var s styles
	s.title = lipgloss.NewStyle()
	s.item = lipgloss.NewStyle().PaddingLeft(padding).PaddingRight(padding)
	s.selectedItem = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).
		PaddingLeft(padding).PaddingRight(padding)

	s.pagination = list.DefaultStyles(darkBG).PaginationStyle
	s.quitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	return s
}

func (m *Model) updateStyles(isDark bool) {
	m.styles = newStyles(isDark)
	m.list.Styles.Title = m.styles.title
	m.list.Styles.PaginationStyle = m.styles.pagination
	m.list.Styles.HelpStyle = m.styles.help
	m.list.SetDelegate(itemDelegate{styles: &m.styles, selected: m.selected})
}
