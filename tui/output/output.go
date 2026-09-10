package output

import (
	"fmt"
	"strings"
	"time"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Line struct {
	HostName string
	Output   string
}

func NewLine(host, output string) Line {
	return Line{
		HostName: host,
		Output:   output,
	}
}

type Model struct {
	width, height int
	content       []string
	vp            viewport.Model
}

func New() Model {
	vp := viewport.New()
	vp.SoftWrap = true
	vp.MouseWheelEnabled = true
	vp.LeftGutterFunc = func(gc viewport.GutterContext) string {
		ts := time.Now().Format(time.TimeOnly)
		gutter := fmt.Sprintf("%s ", ts)
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#505050")).Render(gutter)
	}

	return Model{
		content: []string{},
		vp:      vp,
	}
}

func (m *Model) SetSize(w, h int) {
	m.width, m.height = w, h

	m.vp.SetWidth(m.width - 2)
	m.vp.SetHeight(m.height - 3)
}

// TODO: line styling
func (m *Model) AddLine(host, s string) {
	h := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#757575")).
		Render(host)

	line := lipgloss.NewStyle().
		Foreground(lipgloss.White).
		Render(s)

	r := lipgloss.NewStyle().Render(h, line)
	m.content = append(m.content, strings.Trim(r, "\n"))
	m.vp.SetContentLines(m.content)
}

func (m Model) Init() tea.Cmd {
	return m.vp.Init()
}
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.vp, cmd = m.vp.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(m.width).
		Height(m.height).
		Render(m.vp.View())

	return style
}
