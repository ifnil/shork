package tui

import (
	"container/list"
	"context"
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/internal/runner"
	"github.com/ifnil/shork/tui/conf"
	"github.com/ifnil/shork/tui/hosts"
	"github.com/ifnil/shork/tui/msgs"
	"github.com/ifnil/shork/tui/output"
	"github.com/ifnil/shork/tui/run"
	"github.com/ifnil/shork/tui/status"
	"github.com/ifnil/shork/tui/tabbar"
)

type Model struct {
	ctx           context.Context
	width, height int
	focus         pane
	content       string
	modalVisible  bool

	hosts  hosts.Model
	tabs   tabbar.Model
	output output.Model
	run    run.Model
	status status.Model

	modal Modal

	layout Layout
	hist   *list.List

	rr *runner.Runner
}

func NewModel(ctx context.Context, rr *runner.Runner) (Model, error) {
	h, err := hosts.NewModel(rr)
	if err != nil {
		return Model{
			status: status.New(),
		}, err
	}

	h.Focus()
	return Model{
		ctx: ctx,

		hosts:  h,
		status: status.New(),
		output: output.New(),
		tabs:   tabbar.New(),
		run:    run.New(),

		modal: NewModal("empty"),
		hist:  list.New(),
		rr:    rr,
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

func (m *Model) toggleModal() {
	m.modalVisible = !m.modalVisible
}

// broadcast is not currently  used, but we'll keep it around
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

	m.tabs, cmd = m.tabs.Update(msg)
	cmds = append(cmds, cmd)

	m.modal, cmd = m.modal.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (m Model) capturing() bool {
	switch m.focus {
	case paneRun:
		return m.run.Capturing()
	}
	return false
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.hosts.Init(),
		m.status.Init(),
		m.run.Init(),
		m.output.Init(),
		m.tabs.Init(),
		m.modal.Init(),
	)
}

// runs the command using the internal runner
// and returns the result wrapped in a tea.Msg
// NOTE: move to msgs
type RunResult struct {
	Host   string
	Result string
}

func (m Model) runCmd(host, cmd string) tea.Cmd {
	return func() tea.Msg {
		r := m.rr.RunCmd(m.ctx, host, cmd)
		if r.Err != nil {
			return nil
		}

		return RunResult{
			Host:   host,
			Result: strings.Trim(r.Output, "\n"),
		}
	}
}

//---

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// request to spawn a modal
	case msgs.SpawnModal:
		// i'm aware this is empty
		return m, nil

	// host info
	case msgs.HostInfo:
		var cmd tea.Cmd
		m.status, cmd = m.status.Update(msg)
		return m, cmd

	// ssh command result
	case RunResult:
		m.output.AddLine(msg.Host, msg.Result)
		return m, nil

	// compute layout
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.layout = Compute(msg.Width, msg.Height)

		m.run.SetSize(m.layout.Run.W, m.layout.Run.H)
		m.output.SetSize(m.layout.Output.W, m.layout.Output.H)
		m.hosts.SetSize(m.layout.Hosts.W, m.layout.Hosts.H)
		m.status.SetSize(m.layout.Status.W, m.layout.Status.H)
		m.tabs.SetSize(m.layout.Tabs.W, m.layout.Tabs.H)
		return m, nil

	// handle keys
	case tea.KeyPressMsg:
		if key.Matches(msg, conf.DefaultKeyMap.ForceQuit) {
			return m, tea.Quit
		}

		// insert mode keybinds
		if m.capturing() {
			// TODO: get selected hosts
			switch {
			case key.Matches(msg, conf.DefaultKeyMap.InsertMode.Enter):
				if m.run.Value() == "" {
					return m, nil
				}
				m.run.Release()
				c := m.run.Value()
				m.hist.PushBack(c)
				m.run.Clear()

				// TODO: confirm modal
				return m, m.runCmd(m.hosts.Selected()[0], c)

			case key.Matches(msg, conf.DefaultKeyMap.Release):
				m.run.Release()
				return m, nil
			}
			return m, m.updateFocused(msg)
		}

		// normal mode keybinds
		switch {
		case key.Matches(msg, conf.DefaultKeyMap.Debug):
			m.toggleModal()
			return m, nil

		case key.Matches(msg, conf.DefaultKeyMap.Quit):
			// close modal instead of quitting
			if m.modalVisible {
				m.toggleModal()
				return m, nil
			}
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
	return m, m.broadcast(msg)
}

func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("")
	}

	// arg order doesn't determine stacking
	l := m.layout
	c := lipgloss.NewCompositor(
		l.Hosts.Layer(m.hosts.View()),
		l.Tabs.Layer(m.tabs.View()),
		l.Output.Layer(m.output.View()),
		l.Run.Layer(m.run.View()),
		l.Status.Layer(m.status.View()),
	)

	// modals will need an explicit Z value set
	if m.modalVisible {
		p := l.Body.Center(lipgloss.Width(m.modal.View()), lipgloss.Height(m.modal.View()))
		c.AddLayers(p.CenterLayer(m.modal.View()).Z(1))
	}

	view := tea.NewView(
		lipgloss.NewCanvas(m.width, m.height).Compose(c).Render(),
	)
	view.AltScreen = true
	view.MouseMode = tea.MouseModeCellMotion
	return view
}

func Run(ctx context.Context, rr *runner.Runner) error {
	m, err := NewModel(ctx, rr)
	if err != nil {
		return err
	}

	p := tea.NewProgram(m, tea.WithContext(ctx))
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
