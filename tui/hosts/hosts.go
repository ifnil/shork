package hosts

import (
	"sort"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ifnil/shork/internal/host"
	"github.com/ifnil/shork/internal/runner"
	"github.com/ifnil/shork/tui/conf"
	"github.com/ifnil/shork/tui/helpers"
	"github.com/ifnil/shork/tui/msgs"
)

type Model struct {
	width    int
	height   int
	focused  bool
	styles   styles
	list     list.Model
	hm       *host.HostMap
	selected map[string]struct{}
}

// TODO: make more robust
func NewModel(rr *runner.Runner) (Model, error) {
	// hm := host.NewHostMap()
	// if err := hm.LoadSSHConfig(viper.GetString("ssh_config_path")); err != nil {
	// 	return Model{}, err
	// }

	hm := rr.HostMap()

	items := []list.Item{}
	sortedItems := []string{}
	for _, name := range hm.Hosts() {
		sortedItems = append(sortedItems, name)
	}

	sort.Strings(sortedItems)
	for _, name := range sortedItems {
		items = append(items, item{
			title: name,
		})
	}

	l := list.New(items, itemDelegate{}, 15, 20)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)

	m := Model{
		hm:       hm,
		list:     l,
		selected: make(map[string]struct{}),
	}

	m.updateStyles(true)
	return m, nil
}

func (m *Model) Blur()  { m.focused = false }
func (m *Model) Focus() { m.focused = true }
func (m *Model) SetSize(w, h int) {
	m.width, m.height = w, h
	m.list.SetSize(w-2, h-3)
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Selected() []string {
	out := make([]string, 0, len(m.selected))
	for _, li := range m.list.Items() {
		if i, ok := li.(item); ok {
			if _, on := m.selected[i.title]; on {
				out = append(out, i.title)
			}
		}
	}
	return out
}

func (m Model) Current() host.Host {
	i := m.list.SelectedItem().(item)
	return m.hm.Get(i.title)
}

func (m Model) hostInfoCmd() tea.Cmd {
	return func() tea.Msg {
		i := m.list.SelectedItem().(item)
		h := m.hm.Get(i.title)

		return msgs.HostInfo{Addr: h.HostName}
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, conf.DefaultKeyMap.Select):
			if i, ok := m.list.SelectedItem().(item); ok {
				if _, on := m.selected[i.title]; on {
					delete(m.selected, i.title)
				} else {
					m.selected[i.title] = struct{}{}
				}
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, tea.Batch(cmd, m.hostInfoCmd())
}

func (m Model) View() string {
	border := lipgloss.Color("#ffffff")
	if m.focused {
		border = lipgloss.Color("63")
	}

	style := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Border(lipgloss.NormalBorder()).
		BorderForeground(border)

	return helpers.TitledBox("hosts", m.list.View(), style)
}
