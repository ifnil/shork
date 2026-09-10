package modal

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/tui/conf"
	"github.com/ifnil/shork/tui/msgs"
)

var keys = conf.DefaultKeyMap.ModalKeyMap

type Confirm struct {
	prompt string
	yes    bool
}

func NewConfirm(prompt string) Confirm {
	return Confirm{prompt: prompt}
}

func answer(ok bool) tea.Cmd {
	return func() tea.Msg {
		return msgs.ModalResult{OK: ok}
	}
}

func (c Confirm) Init() tea.Cmd {
	return nil
}

func (c Confirm) Update(msg tea.Msg) (Content, tea.Cmd) {
	k, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return c, nil
	}

	switch {
	case key.Matches(k, keys.Yes):
		return c, answer(true)
	case key.Matches(k, keys.Cancel, keys.No):
		return c, answer(false)
	case key.Matches(k, keys.Accept):
		return c, answer(c.yes)
	case key.Matches(k, keys.Toggle):
		c.yes = !c.yes
	}
	return c, nil
}

func (c Confirm) View() string {
	on := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("170")).
		Padding(0, 2)

	off := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Padding(0, 2)

	yes, no := off.Render("yes"), on.Render("no")
	if c.yes {
		yes, no = on.Render("yes"), off.Render("no")
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		c.prompt,
		"",
		lipgloss.JoinHorizontal(lipgloss.Top, yes, no),
	)
}
