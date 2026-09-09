package tui

type pane int

const (
	paneHosts pane = iota
	paneRun
	paneCount // sentinel
)

func (p pane) next() pane { return (p + 1) % paneCount }
func (p pane) prev() pane { return (p + paneCount - 1) % paneCount }
