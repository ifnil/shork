package conf

import "charm.land/bubbles/v2/key"

type HostKeyMap struct {
	Info    key.Binding
	Exec    key.Binding
	Session key.Binding
	Ping    key.Binding
}

type Capturing struct {
	Enter    key.Binding
	HistUp   key.Binding
	HistDown key.Binding
}

type KeyMap struct {
	NextPane  key.Binding
	PrevPane  key.Binding
	Quit      key.Binding
	ForceQuit key.Binding
	Confirm   key.Binding
	Select    key.Binding
	Release   key.Binding
	Insert    key.Binding

	Capturing Capturing
}

var DefaultKeyMap = KeyMap{
	NextPane:  key.NewBinding(key.WithKeys("tab", "l"), key.WithHelp("<tab>", "next pane")),
	PrevPane:  key.NewBinding(key.WithKeys("shift+tab", "h"), key.WithHelp("<s+tab>", "prev pane")),
	Quit:      key.NewBinding(key.WithKeys("q"), key.WithHelp("<q/esc>", "quit")),
	ForceQuit: key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("<q/esc>", "quit")),
	Confirm:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("<enter>", "confirm")),
	Select:    key.NewBinding(key.WithKeys("space"), key.WithHelp("<space>", "select")),
	Release:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("<esc>", "release")),
	Insert:    key.NewBinding(key.WithKeys("i"), key.WithHelp("<i>", "insert")),

	Capturing: Capturing{
		Enter:    key.NewBinding(key.WithKeys("enter"), key.WithHelp("<enter>", "confirm")),
		HistUp:   key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("<up>", "history up")),
		HistDown: key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("<down>", "history down")),
	},
}
