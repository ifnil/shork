package output

import (
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/tui/helpers"
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

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h := helpers.Percent(msg.Height, 0.9)
		w := helpers.Percent(msg.Width, 0.9)
		m.width = w
		m.height = int(h)
	}
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
