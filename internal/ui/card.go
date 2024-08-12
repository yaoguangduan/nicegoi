package ui

type GoiCard struct {
	opt IElement
}

func (w *GoiCard) Run() {
	w.opt.Run()
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
func (w *GoiCard) SetTitle(title string) {
	w.opt.Set("title", title)
}

func (w *GoiCard) SetContent(content string) {
	w.opt.Set("content", content)
}

func (w *GoiCard) SetDesc(desc string) {
	w.opt.Set("desc", desc)
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
