package hosts

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type item struct{ title string }

func (i item) FilterValue() string { return "" }

type itemDelegate struct {
	styles   *styles
	selected map[string]struct{}
}

func (d itemDelegate) Height() int  { return 1 }
func (d itemDelegate) Spacing() int { return 0 }

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	ts := lipgloss.NewStyle()
	if _, on := d.selected[i.title]; on {
		ts = ts.Foreground(lipgloss.Color("63"))
	}

	str := ts.Render(i.title)
	fn := d.styles.item.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.styles.selectedItem.PaddingLeft(-1).Render("• " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
