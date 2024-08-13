package ui

type GoiCard struct {
	opt IElement
}

func NewCard(content string) *GoiCard {
	return NewCardWithTitle("", content)
}
func NewCardWithTitle(title string, content string) *GoiCard {
	card := GoiCard{opt: NewElement("card").Set("content", content).Set("title", title)}
	card.opt.AddChildren(NewEmptyWidget(card.opt.Eid()+"_A", "card_action"))
	card.opt.AddChildren(NewEmptyWidget(card.opt.Eid()+"_F", "card_footer"))
	return &card
}
func (w *GoiCard) SetTitle(title string) *GoiCard {
	w.opt.Set("title", title)
	return w
}

func (w *GoiCard) SetContent(content string) *GoiCard {
	w.opt.Set("content", content)
	return w
}

func (w *GoiCard) SetDesc(desc string) *GoiCard {
	w.opt.Set("desc", desc)
	return w
}
func (w *GoiCard) SetWidth(width int) *GoiCard {
	w.opt.Set("width", width)
	return w
}
func (w *GoiCard) AddActions(items ...IWidget) {
	for _, item := range items {
		w.opt.Children()[0].AddChildren(item)
	}
}

func (w *GoiCard) AddFooters(items ...IWidget) {
	for _, item := range items {
		w.opt.Children()[1].AddChildren(item)
	}
}

func (w *GoiCard) Visible(v bool) {
	w.opt.Set("hide", !v)
}

func (w *GoiCard) RemoveItemByIdx(idx int) {
	w.opt.RemoveChildrenByIndex(uint32(idx))
}

func (w *GoiCard) Element() IElement {
	return w.opt
}
