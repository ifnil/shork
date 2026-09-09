package status

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/tui/helpers"
	"github.com/ifnil/shork/tui/msgs"
)

type Model struct {
	width, height int
	content       string
}

func New() Model {
	return Model{
		content: "loading...",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.HostInfo:
		m.content = msg.Addr
	case tea.WindowSizeMsg:
		m.height = 1
		m.width = helpers.Percent(msg.Width, 0.7)
	}
	return m, nil
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		PaddingLeft(1).
		Width(m.width).
		Height(1).
		Render(m.content)

	return style
}
