package ui

import "github.com/yaoguangduan/nicegoi/internal/ui/align"

type Box struct {
	opt IElement
}

func (w *Box) Run() {
	w.opt.Run()
}

func NewBox(elements ...IWidget) *Box {
	row := Box{opt: NewElement("box").Set("align", "center")}
	row.opt.AddChildren(elements...)
	return &row
}
func (w *Box) SetAlign(align align.Align) *Box {
	w.opt.Set("align", string(align))
	return w
}

func (w *Box) Horizontal() *Box {
	w.opt.Set("direction", "horizontal")
	return w
}
func (w *Box) Vertical() *Box {
	w.opt.Set("direction", "vertical")
	return w
}
func (w *Box) WithSeparator() *Box {
	w.opt.Set("separator", true)
	return w
}
func (w *Box) Remove(elements ...IWidget) {
	w.opt.RemoveChildren(elements...)
}
func (w *Box) RemoveByIdx(elements ...uint32) {
	w.opt.RemoveChildrenByIndex(elements...)
}
func (w *Box) AddItems(elements ...IWidget) {
	w.opt.AddChildren(elements...)
}
func (w *Box) Element() IElement {
	return w.opt
}
