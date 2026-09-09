package tui

type Rect struct{ W, H int }
type Layout struct {
	Hosts  Rect
	Run    Rect
	Status Rect
	Output Rect
}

func Compute(w, h int) Layout {
	const (
		hostsW  = 24 // fixed sidebar
		runH    = 3  // one text row + border
		statusH = 1  // no border
	)

	bodyH := h - statusH
	return Layout{
		Hosts:  Rect{hostsW, bodyH},
		Output: Rect{w - hostsW, bodyH - runH},
		Run:    Rect{w - hostsW, runH},
		Status: Rect{w, statusH},
	}
}
