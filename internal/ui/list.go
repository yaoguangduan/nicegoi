package ui

type List struct {
	opt IElement
}

func NewList() *List {
	list := List{opt: NewElement("list")}
	return &list
}

func (w *List) AddItems(items ...*Item) {
	for _, item := range items {
		w.opt.AddChildren(item)
	}
}
func (w *List) Visible(v bool) {
	w.opt.Set("hide", !v)
}
func (w *List) RemoveItem(item *Item) {
	w.opt.RemoveChildren(item)
}

func (w *List) RemoveItemByIdx(idx int) {
	w.opt.RemoveChildrenByIndex(uint32(idx))
}

func (w *List) Element() IElement {
	return w.opt
}
