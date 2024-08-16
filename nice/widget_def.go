package nice

type IWidget interface {
	element() *element
}

type emptyWidget struct {
	opt *element
}

func (w *emptyWidget) element() *element {
	return w.opt
}
func (w *emptyWidget) Page() *pageWidget {
	return w.opt.Page.delegate
}

func newEmptyWidget(eid string, kind string) IWidget {
	e := createElement(kind)
	e.Id = eid
	return &emptyWidget{opt: e}
}

//=======================================================================================

type valuedWidget struct {
	e *element
	f func(v any)
}

func (vw *valuedWidget) element() *element {
	return vw.e
}
func (vw *valuedWidget) Page() *pageWidget {
	return vw.e.Page.delegate
}
func (vw *valuedWidget) addMsgHandler(f func(message *Message)) *valuedWidget {
	vw.e.Handlers = append(vw.e.Handlers, f)
	return vw
}
func newValuedWidget(kind string, value any) *valuedWidget {
	w := &valuedWidget{e: createElement(kind).Set("value", value)}
	w.addMsgHandler(func(message *Message) {
		w.e.Modify("value", message.Data)
		if w.f != nil {
			w.f(message.Data)
		}
	})
	return w
}
func newReadonlyWidget(kind string, value any) *valuedWidget {
	w := &valuedWidget{e: createElement(kind).Set("value", value)}
	w.addMsgHandler(func(message *Message) {
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

func (vw *valuedWidget) SetVisible(visible bool) {
	vw.e.Set("hide", !visible)
}
func (vw *valuedWidget) SetDisable(disable bool) {
	vw.e.SetAttr("disabled", disable)
}
