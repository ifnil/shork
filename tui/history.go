package tui

import "container/list"

const MAX_HIST = 1000

type History struct {
	h []string
	c int

	l list.List
}

func (h *History) Previous() string {
	if h.c <= 0 {
		h.c = 0
		return h.h[0]
	}
	h.c--
	return h.h[h.c]
}

func (h *History) Next() string {
	if h.c >= MAX_HIST {
		h.c = MAX_HIST
		return h.h[MAX_HIST]
	}
	h.c++
	return h.h[h.c]
}

func (h *History) Append(v string) { h.h = append(h.h, v) }
func (h History) Current() string  { return h.h[h.c] }
func (h History) Dump() []string   { return h.h }
