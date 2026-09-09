package tui

// calculate grid
// arrange, track, and assign tiles

const (
	tileSize = 1
)

type tile struct {
	w, h, x, y int
}

type Layout struct {
	Width, Height int
	grid          []tile
}

func NewLayout(w, h int) Layout {
	g := make([]tile, w*h)

	return Layout{
		Width:  w,
		Height: h,

		grid: g,
	}
}
