package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/ws"
)

type IWidget interface {
	Element() IElement
}

type emptyWidget struct {
	opt IElement
}

func (w *emptyWidget) Element() IElement {
	return w.opt
}

func newEmptyWidget(eid string, kind string) IWidget {
	element := NewElement(kind).(*Element)
	element.Id = eid
	return &emptyWidget{opt: element}
}

//=======================================================================================

type valuedWidget struct {
	e IElement
	f func(v any)
}

func (vw *valuedWidget) Element() IElement {
	return vw.e
}

func newValuedWidget(kind string, value any) *valuedWidget {
	w := &valuedWidget{e: NewElement(kind).Set("value", value)}
	ws.RegMsgHandle(w.e.Eid(), func(message *msgs.Message) {
		w.e.Modify("value", message.Data)
		if w.f != nil {
			w.f(message.Data)
		}
	})
	return w
}
func (vw *valuedWidget) set(val any) {
	vw.e.Set("value", val)
}
func (vw *valuedWidget) get() any {
	return vw.e.Get("value")
}

func (vw *valuedWidget) onValChange(f func(v any)) {
	vw.f = f
}

func (vw *valuedWidget) Visible(visible bool) {
	vw.e.Set("hide", !visible)
}
