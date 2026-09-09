package tui

import (
	"container/list"
	"context"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/tui/conf"
	"github.com/ifnil/shork/tui/hosts"
	"github.com/ifnil/shork/tui/msgs"
	"github.com/ifnil/shork/tui/output"
	"github.com/ifnil/shork/tui/run"
	"github.com/ifnil/shork/tui/status"
)

type Model struct {
	width   int
	height  int
	focus   pane
	content string
	hist    *list.List

	status status.Model
	run    run.Model
	hosts  hosts.Model
	output output.Model
}

func NewModel() (Model, error) {
	h, err := hosts.NewModel()
	if err != nil {
		return Model{
			status: status.New(),
		}, err
	}

	h.Focus()
	return Model{
		status: status.New(),
		output: output.New(),
		run:    run.New(),
		hist:   list.New(),
		hosts:  h,
	}, nil
}

func (m *Model) setFocus(p pane) {
	m.hosts.Blur()
	m.run.Blur()

	m.focus = p
	switch p {
	case paneHosts:
		m.hosts.Focus()
	case paneRun:
		m.run.Focus()
	}
}

func (m *Model) updateFocused(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch m.focus {
	case paneHosts:
		m.hosts, cmd = m.hosts.Update(msg)
	case paneRun:
		m.run, cmd = m.run.Update(msg)
	}

	return cmd
}

func (m *Model) broadcast(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.output, cmd = m.output.Update(msg)
	cmds = append(cmds, cmd)

	m.hosts, cmd = m.hosts.Update(msg)
	cmds = append(cmds, cmd)

	m.status, cmd = m.status.Update(msg)
	cmds = append(cmds, cmd)

	m.run, cmd = m.run.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.hosts.Init(), m.status.Init(), m.run.Init(), m.output.Init())
}

func (m Model) capturing() bool {
	switch m.focus {
	case paneRun:
		return m.run.Capturing()
	}
	return false
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.HostInfo:
		var cmd tea.Cmd
		m.status, cmd = m.status.Update(msg)
		return m, cmd

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, m.broadcast(msg)

	case tea.KeyPressMsg:
		if key.Matches(msg, conf.DefaultKeyMap.ForceQuit) {
			return m, tea.Quit
		}

		if m.capturing() {
			switch {
			case key.Matches(msg, conf.DefaultKeyMap.Capturing.Enter):
				if m.run.Value() == "" {
					return m, nil
				}

				m.run.Release()

				// send command
				// push to hist
				m.hist.PushBack(m.run.Value())
				m.run.Clear()

				return m, nil

			case key.Matches(msg, conf.DefaultKeyMap.Release):
				m.run.Release()
				return m, nil
			}
			return m, m.updateFocused(msg)
		}

		switch {
		case key.Matches(msg, conf.DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, conf.DefaultKeyMap.NextPane):
			m.setFocus(m.focus.next())
			return m, nil
		case key.Matches(msg, conf.DefaultKeyMap.PrevPane):
			m.setFocus(m.focus.prev())
			return m, nil
		case key.Matches(msg, conf.DefaultKeyMap.Insert):
			m.run.Capture()
			return m, nil
		}
	}

	return m, m.updateFocused(msg)
}

func (m Model) View() tea.View {
	// style := lipgloss.NewStyle().
	// 	Width(m.width - lipgloss.Width(m.hosts.View())).
	// 	Height(m.height - lipgloss.Height(m.run.View()) - lipgloss.Height(m.status.View())).
	// 	Border(lipgloss.NormalBorder())

	// TODO: fix widths and heights

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Bottom,
			m.hosts.View(),
			lipgloss.JoinVertical(
				lipgloss.Left,
				m.output.View(),
				m.run.View(),
			),
		),
	)

	view := tea.NewView(body)
	view.AltScreen = true
	return view
}

func Run(ctx context.Context) error {
	m, err := NewModel()
	if err != nil {
		return err
	}

	p := tea.NewProgram(m, tea.WithContext(ctx))
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
