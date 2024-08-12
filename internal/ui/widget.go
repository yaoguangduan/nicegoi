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
	return &EmptyWidget{opt: &Element{
		Data:     nil,
		Id:       eid,
		Kind:     kind,
		Elements: nil,
	}}
}
