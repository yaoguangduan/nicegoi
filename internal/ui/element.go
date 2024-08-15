package ui

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/server"
	"github.com/yaoguangduan/nicegoi/internal/util"
	"log"
	"maps"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
)

var elements = make([]*element, 0)

type element struct {
	Data     map[string]any            `json:"data"`
	Attr     map[string]any            `json:"attr"`
	Par      *element                  `json:"-"`
	Id       string                    `json:"eid"`
	Kind     string                    `json:"type"`
	Elements []*element                `json:"elements"`
	Page     *PageWidget               `json:"-"`
	W        IWidget                   `json:"-"`
	Handlers []func(msg *msgs.Message) `json:"-"`
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
		e.Page.p.SendMessage(e.Id, "diff", res)
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
		e.Page.p.SendMessage(e.Id, "add", added)
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
	e.Page.p.SendMessage(e.Id, "remove", removed)
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

type Page struct {
	sync.Mutex
	root  *element
	route string
	name  string
	self  *PageWidget
	title string
}

func (p *Page) Name() string {
	return p.name
}
func (p *Page) Route() string {
	return p.route
}
func (p *Page) OnInit() {
	if p == RootPage {
		pw := &PageWidget{}
		p.self = pw
		pw.p = p
	}
	for _, ele := range elements {
		if ele.Parent() == nil {
			ele.SetParent(RootPage.root)
			RootPage.root.Elements = append(RootPage.root.Elements, ele)
		}
	}
	setElementPage(p.root, p)
}

func setElementPage(root *element, p *Page) {
	root.Page = p.self
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
	data := make(map[string]any)
	data["elements"] = p.root.Elements
	data["title"] = p.title
	bys, err := json.Marshal(data)
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

var QID = atomic.Uint64{}

func (p *Page) RouteTo(name string, data ...any) {
	if data == nil {
		p.SendMessage("EID0", "route", map[string]any{"name": name})
	} else {
		id := fmt.Sprintf("QID%d", QID.Add(1))
		server.AppendQueryData(id, data)
		p.SendMessage("EID0", "route", map[string]any{"name": name, "qid": id})
	}
}
func findElement(element *element, id string) *element {
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
	e := &element{Id: "EID0"}
	var route = name
	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}
	page := &Page{name: name, root: e, route: route}
	page.title = name
	server.RegPageRes(page)
	return page
}

var RootPage = createPage("/")

type PageWidget struct {
	vw *valuedWidget
	p  *Page
}

func NewPage(name string) *PageWidget {
	pw := &PageWidget{p: createPage(name), vw: &valuedWidget{e: createElement("PAGE-" + name)}}
	pw.p.self = pw
	return pw
}

func (p *PageWidget) AddItems(widgets ...IWidget) {
	p.p.root.AddChildren(widgets...)
}

func (p *PageWidget) RouteTo(name string, data ...any) {
	p.p.RouteTo(name, data...)
}
func (p *PageWidget) SetTitle(title string) {
	p.p.title = title
}

func SetTitle(title string) {
	RootPage.title = title
}
