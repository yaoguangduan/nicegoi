package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/server"
)

type IWidget interface {
	Element() *Element
	Page() server.IPage
}

type emptyWidget struct {
	opt *Element
}

func (w *emptyWidget) Element() *Element {
	return w.opt
}
func (w *emptyWidget) Page() server.IPage {
	return w.opt.Page
}

func newEmptyWidget(eid string, kind string) IWidget {
	element := NewElement(kind)
	element.Id = eid
	return &emptyWidget{opt: element}
}

//=======================================================================================

type valuedWidget struct {
	e *Element
	f func(v any)
}

func (vw *valuedWidget) Element() *Element {
	return vw.e
}
func (vw *valuedWidget) Page() server.IPage {
	return vw.e.Page
}
func (vw *valuedWidget) AddMsgHandler(f func(message *msgs.Message)) *valuedWidget {
	vw.e.Handlers = append(vw.e.Handlers, f)
	return vw
}
func newValuedWidget(kind string, value any) *valuedWidget {
	w := &valuedWidget{e: NewElement(kind).Set("value", value)}
	w.AddMsgHandler(func(message *msgs.Message) {
		w.e.Modify("value", message.Data)
		if w.f != nil {
			w.f(message.Data)
		}
	})
	w.e.AttachWidget(w)
	return w
}
func newReadonlyWidget(kind string, value any) *valuedWidget {
	w := &valuedWidget{e: NewElement(kind).Set("value", value)}
	w.AddMsgHandler(func(message *msgs.Message) {
		if w.f != nil {
			w.f(message.Data)
		}
	})
	w.e.AttachWidget(w)
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
