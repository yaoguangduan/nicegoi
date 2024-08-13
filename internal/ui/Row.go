package ui

import "github.com/yaoguangduan/nicegoi/internal/ui/align"

type Row struct {
	*valuedWidget
}

func NewRow(items ...IWidget) *Row {
	w := &Row{newValuedWidget("row", "")}
	w.e.AddChildren(items...)
	return w
}
func (w *Row) SetGutter(h, v int) *Row {
	w.e.Set("gutter", []int{h, v})
	return w
}
func (w *Row) SetSpan(spans ...int) *Row {
	w.e.Set("span", spans)
	return w
}

func (w *Row) SetOffset(index, value int) *Row {
	m := w.e.Get("offset")
	if m == nil {
		m = make(map[int]int)
	}
	m.(map[int]int)[index] = value
	w.e.Set("offset", m)
	return w
}
func (w *Row) Justify(justify align.Justify) *Row {
	w.e.Set("justify", justify)
	return w
}
