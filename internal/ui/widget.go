package ui

type IWidget interface {
	Element() IElement
}

type EmptyWidget struct {
	opt IElement
}

func (w *EmptyWidget) Element() IElement {
	return w.opt
}

func NewEmptyWidget(eid string, kind string) IWidget {
	element := NewElement(kind).(*Element)
	element.Id = eid
	return &EmptyWidget{opt: element}
}
