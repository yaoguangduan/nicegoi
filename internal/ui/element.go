package ui

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/server"
	"github.com/yaoguangduan/nicegoi/internal/util"
	"log"
	"maps"
	"slices"
	"strings"
	"sync"
)

var elements = make([]*Element, 0)

type Data map[string]any
type Element struct {
	Data     `json:"data"`
	Par      *Element                  `json:"-"`
	Id       string                    `json:"eid"`
	Kind     string                    `json:"type"`
	Elements []*Element                `json:"elements"`
	Page     *Page                     `json:"-"`
	W        IWidget                   `json:"-"`
	Handlers []func(msg *msgs.Message) `json:"-"`
}

func NewElement(kind string) *Element {
	e := &Element{
		Kind:     kind,
		Elements: make([]*Element, 0),
		Id:       util.AllocEID(),
		Data:     make(map[string]any),
	}
	elements = append(elements, e)
	return e
}

func (e *Element) Duplicate() *Element {
	ed := NewElement(e.Kind)
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

func (e *Element) GetWidget() IWidget {
	return e.W
}
func (e *Element) AttachWidget(w IWidget) {
	e.W = w
}
func (e *Element) Parent() *Element {
	return e.Par
}
func (e *Element) SetParent(p *Element) {
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
func (e *Element) Set(key string, value any) *Element {
	//old, exist := e.Data[key]
	e.Data[key] = value
	e.OnModify(key)
	return e
}

func (e *Element) Modify(key string, value any) *Element {
	e.Data[key] = value
	return e
}

func (e *Element) get(key string) any {
	return e.Data[key]
}
func (e *Element) OnModify(fields ...string) {
	res := make(map[string]any)
	res["data"] = getUpdated(*e, fields...)
	if e.Page != nil {
		e.Page.SendMessage(e.Id, "diff", res)
	}
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
		e.Elements = make([]*Element, 0)
	}
	added := make([]any, 0)
	for _, c := range cc {
		ce := c.Element()
		ce.Par = e
		added = append(added, ce)
		e.Elements = append(e.Elements, ce)
	}
	if e.Page != nil {
		e.Page.SendMessage(e.Id, "add", added)
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
	e.Page.SendMessage(e.Id, "remove", removed)
}
func (e *Element) RemoveChildren(cc ...IWidget) {
	if e.Elements == nil {
		return
	}
	removed := make([]uint32, 0)
	for _, c := range cc {
		idx := slices.IndexFunc(e.Elements, func(element *Element) bool {
			return c.Element().Eid() == element.Eid()
		})
		if idx != -1 {
			removed = append(removed, uint32(idx))
		}
	}
	e.RemoveChildrenByIndex(removed...)
}

func (e *Element) Children() []*Element {
	if e.Elements == nil {
		e.Elements = make([]*Element, 0)
	}
	return e.Elements
}

type Page struct {
	sync.Mutex
	root    *Element
	route   string
	name    string
	padding []*msgs.Message
}

func (p *Page) Name() string {
	return p.name
}
func (p *Page) Route() string {
	return p.route
}
func (p *Page) OnInit() {
	for _, ele := range elements {
		if ele.Parent() == nil {
			ele.SetParent(RootPage.root)
			RootPage.root.Elements = append(RootPage.root.Elements, ele)
		}
	}
	setElementPage(p.root, p)
}

func setElementPage(root *Element, p *Page) {
	root.Page = p
	for _, e := range root.Elements {
		setElementPage(e, p)
	}
}
func (p *Page) OnNewWsCon(conn *websocket.Conn) {
	bys, err := json.Marshal(p.root.Elements)
	if err != nil {
		log.Printf("conn full msg error ;%s", err)
	}
	log.Printf("send conn msg:%v", msgs.Message{Eid: p.root.Id, Kind: "add", Data: string(bys)})
	err = conn.WriteJSON(msgs.Message{Eid: p.root.Id, Kind: "add", Data: string(bys)})
	if err != nil {
		return
	}
}
func (p *Page) FullData() ([]byte, error) {
	bys, err := json.Marshal(p.root.Elements)
	if err != nil {
		log.Printf("conn full msg error ;%s", err)
		return nil, err
	}
	return bys, nil
}
func (p *Page) SendMessage(id, cmd string, data any) {
	server.Send(p.Route(), id, cmd, data)
}
func (p *Page) OnNewWsMsg(msg *msgs.Message) {
	id := msg.Eid
	ele := findElement(p.root, id)
	if ele == nil {
		log.Printf("ERROR find element fail ,id:%s,msg:%v", msg.Eid, msg)
		return
	}
	for _, f := range ele.Handlers {
		f(msg)
	}
}
func (p *Page) RouteTo(name string) {
	p.SendMessage("EID0", "route", name)
}

func findElement(element *Element, id string) *Element {
	if element.Id == id {
		return element
	} else {
		for _, ele := range element.Elements {
			find := findElement(ele, id)
			if find != nil {
				return find
			}
		}
	}
	return nil
}
func createPage(name string) *Page {
	e := &Element{Id: "EID0"}
	var route = name
	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}
	page := &Page{name: name, root: e, route: route}
	server.RegPageRes(page)
	return page
}

var RootPage = createPage("/")

type PageWidget struct {
	valuedWidget
	p *Page
}

func NewPage(name string) *PageWidget {
	return &PageWidget{p: createPage(name)}
}

func (p *PageWidget) AddItems(widgets ...IWidget) {
	p.p.root.AddChildren(widgets...)
}
