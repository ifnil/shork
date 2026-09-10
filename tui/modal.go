package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	border      = lipgloss.NormalBorder()
	borderColor = lipgloss.Color("#772233")
)

type Modal struct {
	width, height int
	content       string
}

func NewModal(content string) Modal {
	return Modal{content: content}
}

func (m Modal) Init() tea.Cmd {
	return nil
}

func (m Modal) Update(msg tea.Msg) (Modal, tea.Cmd) {
	return m, nil
}

func (m Modal) View() string {
	box := lipgloss.NewStyle().
		Border(border).
		BorderForeground(borderColor).
		Render(m.content)
	return box
}
