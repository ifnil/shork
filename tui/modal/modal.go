package modal

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/tui/helpers"
)

var (
	border      = lipgloss.NormalBorder()
	borderColor = lipgloss.Color("#772233")
)

type Content interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Content, tea.Cmd)
	View() string
}

type Sizer interface {
	SetSize(w, h int) Content
}

type Modal struct {
	width, height int
	title         string
	body          Content
}

func New() Modal { return Modal{} }

func (m Modal) Active() bool  { return m.body != nil }
func (m Modal) Init() tea.Cmd { return nil }

func (m *Modal) Open(title string, body Content) tea.Cmd {
	m.title, m.body = title, body
	m.resizeBody()
	return body.Init()
}

func (m *Modal) Close() {
	m.title = ""
	m.body = nil
}

func (m *Modal) SetSize(w, h int) {
	m.width, m.height = w, h
	m.resizeBody()
}

func (m *Modal) resizeBody() {
	if s, ok := m.body.(Sizer); ok {
		m.body = s.SetSize(m.width-2, m.height-2)
	}
}

func (m Modal) Update(msg tea.Msg) (Modal, tea.Cmd) {
	if m.body == nil {
		return m, nil
	}

	var cmd tea.Cmd
	m.body, cmd = m.body.Update(msg)
	return m, cmd
}

func (m Modal) View() string {
	if m.body == nil {
		return ""
	}

	body := m.body.View()
	style := lipgloss.NewStyle().
		Border(border).
		BorderForeground(borderColor).
		Width(min(lipgloss.Width(body)+2, m.width)).
		Height(min(lipgloss.Height(body)+2, m.height)).
		Align(lipgloss.Center, lipgloss.Center)
	return helpers.TitledBox(m.title, body, style)
}
