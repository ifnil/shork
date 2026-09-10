package conf

import "charm.land/bubbles/v2/key"

type KeyMap struct {
	NextPane  key.Binding
	PrevPane  key.Binding
	Quit      key.Binding
	ForceQuit key.Binding
	Confirm   key.Binding
	Select    key.Binding
	Release   key.Binding
	Insert    key.Binding

	Debug key.Binding

	InsertMode  InsertMode
	HostKeyMap  HostKeyMap
	ModalKeyMap ModalKeyMap
}

type ModalKeyMap struct {
	Close key.Binding
}

type HostKeyMap struct {
	Info    key.Binding
	Exec    key.Binding
	Session key.Binding
	Ping    key.Binding
}

type InsertMode struct {
	Enter    key.Binding
	HistUp   key.Binding
	HistDown key.Binding
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

	Debug: key.NewBinding(key.WithKeys("o")),

	InsertMode: InsertMode{
		Enter:    key.NewBinding(key.WithKeys("enter"), key.WithHelp("<enter>", "confirm")),
		HistUp:   key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("<up>", "history up")),
		HistDown: key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("<down>", "history down")),
	},

	HostKeyMap: HostKeyMap{
		Info: key.NewBinding(key.WithKeys("i")),
		Ping: key.NewBinding(key.WithKeys("p")),
	},

	ModalKeyMap: ModalKeyMap{
		Close: key.NewBinding(key.WithKeys("q")),
	},
}
