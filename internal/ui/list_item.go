package ui

type Item struct {
	opt IElement
}

func NewListItem(text string) *Item {
	list := Item{opt: NewElement("list_item").Set("text", text)}
	return &list
}

func NewComplexListItem(title, desc string) *Item {
	list := Item{opt: NewElement("list_item").Set("title", title).Set("desc", desc)}
	return &list
}

func (w *Item) AddAction(action IWidget) *Item {
	w.opt.AddChildren(action)
	return w
}

func (w *Item) Element() IElement {
	return w.opt
}
