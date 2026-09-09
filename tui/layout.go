package tui

type Rect struct{ X, Y, W, H int }

type Layout struct {
	Hosts  Rect
	Run    Rect
	Status Rect
	Tabs   Rect
	Output Rect
	Body   Rect
}

func (r Rect) Center(w, h int) Rect {
	w, h = min(w, r.W), min(h, r.H)
	return Rect{
		X: r.X + (r.W-w)/2,
		Y: r.Y + (r.H-h)/2,
		W: w, H: h,
	}
}

func Compute(w, h int) Layout {
	const (
		hostsW  = 24 // fixed sidebar
		runH    = 3  // one text row + border
		statusH = 1  // no border
		tabsH   = 1
	)

	bodyH := h - statusH
	rightX, rightW := hostsW, w-hostsW
	return Layout{
		Hosts:  Rect{X: 0, Y: 0, W: hostsW, H: bodyH},
		Tabs:   Rect{X: rightX, Y: 0, W: rightW, H: tabsH},
		Output: Rect{X: rightX, Y: tabsH, W: rightW, H: bodyH - tabsH - runH},
		Run:    Rect{X: rightX, Y: bodyH - runH, W: rightW, H: runH},
		Status: Rect{X: 0, Y: bodyH, W: w, H: statusH},
		Body:   Rect{X: 0, Y: 0, W: w, H: bodyH},
	}
}
