package ui

import "github.com/yaoguangduan/nicegoi/internal/ui/align"

type Row struct {
	*valueWidget
}

func NewRow(items ...IWidget) *Row {
	w := &Row{newValueWidget("row", "")}
	w.opt.AddChildren(items...)
	return w
}
func (w *Row) SetGutter(h, v int) *Row {
	w.opt.Set("gutter", []int{h, v})
	return w
}
func (w *Row) SetSpan(spans ...int) *Row {
	w.opt.Set("span", spans)
	return w
}

func (w *Row) SetOffset(index, value int) *Row {
	m := w.opt.Get("offset")
	if m == nil {
		m = make(map[int]int)
	}
	m.(map[int]int)[index] = value
	w.opt.Set("offset", m)
	return w
}
func (w *Row) Justify(justify align.Justify) *Row {
	w.opt.Set("justify", justify)
	return w
}
