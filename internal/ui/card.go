package ui

type GoiCard struct {
	*valuedWidget
}

func NewCard(content string) *GoiCard {
	return NewCardWithTitle("", content)
}
func NewCardWithTitle(title string, content string) *GoiCard {
	card := GoiCard{valuedWidget: newReadonlyWidget("card", content)}
	card.e.Set("title", title)
	card.e.AddChildren(newEmptyWidget(card.e.Eid()+"_A", "card_action"))
	card.e.AddChildren(newEmptyWidget(card.e.Eid()+"_F", "card_footer"))
	return &card
}
func (w *GoiCard) SetTitle(title string) *GoiCard {
	w.e.Set("title", title)
	return w
}

func (w *GoiCard) SetContent(content string) *GoiCard {
	w.e.Set("content", content)
	return w
}

func (w *GoiCard) SetDesc(desc string) *GoiCard {
	w.e.Set("desc", desc)
	return w
}
func (w *GoiCard) SetWidth(width int) *GoiCard {
	w.e.Set("width", width)
	return w
}
func (w *GoiCard) AddActions(items ...IWidget) {
	for _, item := range items {
		w.e.Children()[0].AddChildren(item)
	}
}

func (w *GoiCard) AddFooters(items ...IWidget) {
	for _, item := range items {
		w.e.Children()[1].AddChildren(item)
	}
}
