package run

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/tui/helpers"
)

type SetValueMsg string

type Model struct {
	width, height int
	focused       bool
	capturing     bool
	content       string
	style         lipgloss.Style
	t             textinput.Model
}

func New() Model {
	t := textinput.New()
	t.Placeholder = "cmd..."
	t.SetVirtualCursor(true)

	return Model{
		t:       t,
		height:  1,
		focused: false,
		style:   lipgloss.NewStyle(),
	}
}

func (m *Model) Blur()    { m.focused = false }
func (m *Model) Focus()   { m.focused = true }
func (m *Model) Clear()   { m.t.Reset() }
func (m *Model) Capture() { m.capturing = true; m.t.Focus() }
func (m *Model) Release() { m.capturing = false; m.content = m.t.Value(); m.t.Blur() }

func (m *Model) SetSize(w, h int) {
	m.width, m.height = w, h
	m.t.SetWidth(w)
}

func (m Model) Capturing() bool { return m.capturing }
func (m Model) Value() string   { return m.t.Value() }
func (m Model) Init() tea.Cmd   { return textinput.Blink }

func (m Model) Set(v string) tea.Cmd {
	return func() tea.Msg {
		return SetValueMsg(v)
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SetValueMsg:
		m.t.SetValue(string(msg))
		return m, nil
	}

	var cmd tea.Cmd
	m.t, cmd = m.t.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	border := lipgloss.Color("#fff")
	if m.focused {
		border = lipgloss.Color("63")
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(border).
		Width(m.width).Height(m.height)

	body := m.t.View()

	return helpers.TitledBox("run", body, style)
}
