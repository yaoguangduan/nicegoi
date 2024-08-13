package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/option"
)

type Box struct {
	*valuedWidget
}

func NewBox(elements ...IWidget) *Box {
	row := Box{valuedWidget: newValuedWidget("box", "")}
	row.e.Set("align", "center")
	row.e.AddChildren(elements...)
	return &row
}
func (w *Box) Align(align option.Align) *Box {
	w.e.Set("align", string(align))
	return w
}

func (w *Box) Horizontal() *Box {
	w.e.Set("direction", "horizontal")
	return w
}
func (w *Box) Vertical() *Box {
	w.e.Set("direction", "vertical")
	return w
}

func (w *Box) Remove(elements ...IWidget) {
	w.e.RemoveChildren(elements...)
}
func (w *Box) RemoveByIdx(elements ...uint32) {
	w.e.RemoveChildrenByIndex(elements...)
}
func (w *Box) AddItems(elements ...IWidget) {
	w.e.AddChildren(elements...)
}
