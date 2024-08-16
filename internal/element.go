package nice

import (
	"github.com/yaoguangduan/nicegoi/internal/util"
	"maps"
	"slices"
)

var elements = make([]*element, 0)

type element struct {
	Data     map[string]any       `json:"data"`
	Attr     map[string]any       `json:"attr"`
	Par      *element             `json:"-"`
	Id       string               `json:"eid"`
	Kind     string               `json:"type"`
	Elements []*element           `json:"elements"`
	Page     *page                `json:"-"`
	W        IWidget              `json:"-"`
	Handlers []func(msg *Message) `json:"-"`
}

func createElement(kind string) *element {
	e := &element{
		Kind:     kind,
		Elements: make([]*element, 0),
		Id:       util.AllocEID(),
		Data:     make(map[string]any),
		Attr:     make(map[string]any),
	}
	elements = append(elements, e)
	return e
}

func (e *element) Duplicate() *element {
	ed := createElement(e.Kind)
	for _, c := range e.Elements {
		ed.Elements = append(ed.Elements, c.Duplicate())
	}
	maps.Copy(ed.Data, e.Data)
	ed.Page = e.Page
	ed.W = e.W
	for _, f := range e.Handlers {
		ed.Handlers = append(ed.Handlers, f)
	}
	return ed
}

func (e *element) GetWidget() IWidget {
	return e.W
}
func (e *element) AttachWidget(w IWidget) {
	e.W = w
}
func (e *element) Parent() *element {
	return e.Par
}
func (e *element) SetParent(p *element) {
	e.Par = p
}
func (e *element) Type() string {
	return e.Kind
}
func (e *element) SetVisible(visible bool) {
	e.Set("hide", !visible)
}
func (e *element) Get(key string) any {
	return e.Data[key]
}
func (e *element) Set(key string, value any) *element {
	//old, exist := e.Data[key]
	e.Data[key] = value
	e.OnModify(key)
	return e
}

func (e *element) GetAttr(key string) any {
	return e.Data[key]
}
func (e *element) SetAttr(key string, value any) *element {
	//old, exist := e.Data[key]
	e.Attr[key] = value
	e.OnModify(key)
	return e
}

func (e *element) Modify(key string, value any) *element {
	e.Data[key] = value
	return e
}

func (e *element) get(key string) any {
	return e.Data[key]
}
func (e *element) OnModify(fields ...string) {
	res := make(map[string]any)
	res["data"] = getUpdated(e.Data, fields...)
	res["attr"] = getUpdated(e.Attr, fields...)
	if e.Page != nil {
		e.Page.SendMessage(e.Id, "diff", res)
	}
}

func getUpdated(data map[string]any, fields ...string) map[string]any {
	ret := make(map[string]any)
	for _, field := range fields {
		val, exist := data[field]
		if exist {
			ret[field] = val
		}
	}
	return ret
}

func (e *element) Eid() string {
	return e.Id
}
func (e *element) AddChildren(cc ...IWidget) {
	if e.Elements == nil {
		e.Elements = make([]*element, 0)
	}
	added := make([]any, 0)
	for _, c := range cc {
		ce := c.element()
		ce.Par = e
		added = append(added, ce)
		e.Elements = append(e.Elements, ce)
	}
	if e.Page != nil {
		e.Page.SendMessage(e.Id, "add", added)
	}
}

func (e *element) RemoveChildrenByIndex(ii ...uint32) {
	if e.Elements == nil {
		return
	}
	slices.Sort(ii)
	slices.Reverse(ii)
	removed := make([]any, 0)
	for _, idx := range ii {
		if len(e.Elements) <= int(idx) {
			continue
		}
		eid := e.Elements[idx].Eid()
		e.Elements[idx].SetParent(nil)
		slices.Delete(elements, int(idx), int(idx))
		removed = append(removed, eid)
		e.Elements = append(e.Elements[:idx], e.Elements[idx+1:]...)
	}
	e.Page.SendMessage(e.Id, "remove", removed)
}
func (e *element) RemoveChildren(cc ...IWidget) {
	if e.Elements == nil {
		return
	}
	removed := make([]uint32, 0)
	for _, c := range cc {
		idx := slices.IndexFunc(e.Elements, func(element *element) bool {
			return c.element().Eid() == element.Eid()
		})
		if idx != -1 {
			removed = append(removed, uint32(idx))
		}
	}
	e.RemoveChildrenByIndex(removed...)
}

func (e *element) Children() []*element {
	if e.Elements == nil {
		e.Elements = make([]*element, 0)
	}
	return e.Elements
}
