package tabbar

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// tabs
// output, logs, stats, opts, etc

type Model struct {
	width, height int
}

func New() Model {
	return Model{}
}

func (m *Model) SetSize(w, h int) { m.width, m.height = w, h }
func (m Model) Init() tea.Cmd     { return nil }

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height)

	return style.Render("tabs")
}
