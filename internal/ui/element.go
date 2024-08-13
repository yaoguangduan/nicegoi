package ui

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/yaoguangduan/nicegoi/internal/httpx"
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/util"
	"github.com/yaoguangduan/nicegoi/internal/ws"
	"log"
	"slices"
)

var elements = make([]IElement, 0)

type IElement interface {
	Eid() string
	Type() string
	Parent() IElement
	SetParent(IElement)
	Children() []IElement
	AddChildren(...IWidget)
	RemoveChildren(...IWidget)
	RemoveChildrenByIndex(...uint32)
	Modify(key string, value any) IElement
	Set(key string, value any) IElement
	Get(key string) any
}
type Data map[string]any
type Element struct {
	Data     `json:"data"`
	Par      IElement   `json:"-"`
	Id       string     `json:"eid"`
	Kind     string     `json:"type"`
	Elements []IElement `json:"elements"`
}

func NewElement(kind string) IElement {
	e := &Element{
		Kind:     kind,
		Elements: make([]IElement, 0),
		Id:       util.AllocEID(),
		Data:     make(map[string]any),
	}
	elements = append(elements, e)
	return e
}
func (e *Element) Parent() IElement {
	return e.Par
}
func (e *Element) SetParent(p IElement) {
	e.Par = p
}
func (e *Element) Type() string {
	return e.Kind
}
func (e *Element) SetVisible(visible bool) {
	e.Set("hide", !visible)
}
func (e *Element) Get(key string) any {
	return e.Data[key]
}
func (e *Element) Set(key string, value any) IElement {
	//old, exist := e.Data[key]
	e.Data[key] = value
	e.OnModify(key)
	return e
}

func (e *Element) Modify(key string, value any) IElement {
	e.Data[key] = value
	return e
}

func (e *Element) get(key string) any {
	return e.Data[key]
}
func (e *Element) OnModify(fields ...string) {
	if !ws.Active() {
		return
	}
	res := make(map[string]any)
	res["data"] = getUpdated(*e, fields...)
	ws.Send(e.Id, "diff", res)
}

func getUpdated(e Element, fields ...string) map[string]any {
	data := make(map[string]any)
	for _, field := range fields {
		val, exist := e.Data[field]
		if exist {
			data[field] = val
		}
	}
	return data
}

func (e *Element) Eid() string {
	return e.Id
}
func (e *Element) AddChildren(cc ...IWidget) {
	if e.Elements == nil {
		e.Elements = make([]IElement, 0)
	}
	added := make([]any, 0)
	for _, c := range cc {
		ce := c.Element()
		ce.SetParent(e)
		added = append(added, ce)
		e.Elements = append(e.Elements, ce)
	}
	if ws.Active() {
		ws.Send(e.Id, "add", added)
	}
}

func (e *Element) RemoveChildrenByIndex(ii ...uint32) {
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

	if ws.Active() {
		ws.Send(e.Id, "remove", removed)
	}
}
func (e *Element) RemoveChildren(cc ...IWidget) {
	if e.Elements == nil {
		return
	}
	removed := make([]uint32, 0)
	for _, c := range cc {
		idx := slices.IndexFunc(e.Elements, func(element IElement) bool {
			return c.Element().Eid() == element.Eid()
		})
		if idx != -1 {
			removed = append(removed, uint32(idx))
		}
	}
	e.RemoveChildrenByIndex(removed...)
}

func (e *Element) Children() []IElement {
	if e.Elements == nil {
		e.Elements = make([]IElement, 0)
	}
	return e.Elements
}

type Root struct {
	Element
}

var RootElement = &Root{Element{
	Id:       "EID0",
	Elements: make([]IElement, 0),
}}

func (re Root) SendRootMessage(cmd string, data any) {
	ws.Send(re.Id, cmd, data)
}
func Run() {
	for _, ele := range elements {
		if ele.Parent() == nil {
			ele.SetParent(RootElement)
			RootElement.Elements = append(RootElement.Elements, ele)
		}
	}
	ws.OnNewConn(func(conn *websocket.Conn) {
		bys, err := json.Marshal(RootElement.Elements)
		if err != nil {
			log.Printf("conn full msg error ;%s", err)
		}
		log.Printf("send conn msg:%v", msgs.Message{Eid: RootElement.Eid(), Kind: "data", Data: string(bys)})
		conn.WriteJSON(msgs.Message{Eid: RootElement.Eid(), Kind: "add", Data: string(bys)})
	})
	httpx.Run()
}
