package ui

type List struct {
	*valuedWidget
}

func NewList() *List {
	list := List{valuedWidget: newValuedWidget("list", "")}
	return &list
}

func (w *List) AddItems(items ...*ListItem) {
	for _, item := range items {
		w.e.AddChildren(item)
	}
}
func (w *List) RemoveItem(item *ListItem) {
	w.e.RemoveChildren(item)
}

func (w *List) RemoveItemByIdx(idx int) {
	w.e.RemoveChildrenByIndex(uint32(idx))
}

type ListItem struct {
	*valuedWidget
}

func NewListItem(text string) *ListItem {
	list := ListItem{valuedWidget: newValuedWidget("list_item", text)}
	return &list
}

func (w *ListItem) AddAction(action IWidget) *ListItem {
	w.e.AddChildren(action)
	return w
}
