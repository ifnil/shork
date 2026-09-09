package output

import (
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	width, height int
	content       string
	t             table.Model
}

func New() Model {
	t := table.New()
	return Model{t: t}
}

func (m *Model) SetSize(w, h int) { m.width, m.height = w, h }

func (m Model) Init() tea.Cmd { return nil }
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	// switch msg := msg.(type) {
	// }
	return m, nil
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(m.width).
		Height(m.height).
		Render(m.content)

	return style
}
