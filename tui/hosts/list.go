package hosts

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var padding = 1

type styles struct {
	title        lipgloss.Style
	item         lipgloss.Style
	selectedItem lipgloss.Style
	pagination   lipgloss.Style
	help         lipgloss.Style
	quitText     lipgloss.Style
}

func newStyles(darkBG bool) styles {
	// padding := 1
	// var s styles
	// s.title = lipgloss.NewStyle()
	// s.item = lipgloss.NewStyle().PaddingLeft(padding).PaddingRight(3)
	// s.selectedItem = lipgloss.NewStyle().PaddingLeft(padding).PaddingRight(3).Foreground(lipgloss.Color("170"))
	// s.pagination = list.DefaultStyles(darkBG).PaginationStyle.PaddingLeft(4)
	// s.quitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	var s styles
	s.title = lipgloss.NewStyle()
	s.item = lipgloss.NewStyle().PaddingLeft(padding).PaddingRight(padding)
	s.selectedItem = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).
		PaddingLeft(padding).PaddingRight(padding)

	s.pagination = list.DefaultStyles(darkBG).PaginationStyle
	s.quitText = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	return s
}

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
			return d.styles.selectedItem.PaddingLeft(-1).Render("•" + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
