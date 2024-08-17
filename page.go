package nicegoi

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/yaoguangduan/nicegoi/option/menu"
	"github.com/yaoguangduan/nicegoi/option/timeline"
	"github.com/yaoguangduan/nicegoi/util"
	"log"
	"net/http"
	"sync"
	"time"
)

type IPage interface {
	Name() string
	Layout(ctx PageContext)
}

func addPage(pages ...IPage) {
	for _, p := range pages {
		if p != nil {
			pageList = append(pageList, p)
		}
	}
}

var pageQueryData = make(map[string]map[string]any)

var pageList = make([]IPage, 0)

func findPage(name string) IPage {
	for _, d := range pageList {
		if d.Name() == name {
			return d
		}
	}
	return nil
}

type pageManager struct {
	sync.Mutex
	pages map[string]*pageInstance
}

func (pm *pageManager) getOrCreate(name string, uuid string) *pageInstance {
	p, ok := pm.pages[uuid]
	if ok {
		return p
	}
	pl := findPage(name)
	if pl == nil {
		return nil
	}
	pm.Lock()
	defer pm.Unlock()
	p = &pageInstance{layout: pl, route: "/" + name, name: name, title: name, uuid: uuid}
	pw := &pageInstCreateContext{owner: p, q: make(map[string]any)}
	p.delegate = pw
	pm.pages[uuid] = p
	return p
}

var pageInstMgr = &pageManager{pages: make(map[string]*pageInstance)}

type pageInstance struct {
	uuid     string
	layout   IPage
	root     *element
	route    string
	name     string
	delegate *pageInstCreateContext
	title    string
	conn     *websocket.Conn
	onNewMsg func(id string, cmd string, data any)
}

func (p *pageInstance) serverNewCon(conn *websocket.Conn) {
	if p.conn != nil {
		err := p.conn.Close()
		if err != nil {
			log.Println("owner.conn close err:", err)
		}
	}
	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("warn:websocket closed ", code, text, p.uuid)
		p.conn = nil
		return nil
	})
	p.conn = conn
	go p.handleRcv(conn)
}

func (p *pageInstance) handleRcv(conn *websocket.Conn) {
	var errCnt = 0
	for {
		up := &Message{}
		err := conn.ReadJSON(up)
		if err != nil {
			log.Println("Read error:", p.uuid, err)
			if errCnt > 3 {
				err = conn.Close()
				if err != nil {
					log.Println("close conn error:", p.uuid, err)
				}
				break
			} else {
				errCnt++
			}
			continue
		}
		errCnt = 0
		log.Printf("receive new msg:%+v", up)
		p.OnNewWsMsg(up)
	}
}

func (p *pageInstance) handlePageReq(w http.ResponseWriter, r *http.Request) {
	if p.root == nil {
		p.root = createElement("-")
		var data map[string]any
		if d, ok := pageQueryData[p.uuid]; ok {
			fmt.Println("owner receive query data:", d)
			data = d
		}
		p.delegate.q = data
		p.layout.Layout(p.delegate)
		setElementPage(p.root, p)
	}
	if p.root != nil {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data, err := p.FullData()
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(data)
	}
}

func setElementPage(root *element, p *pageInstance) {
	root.Page = p
	for _, e := range root.Elements {
		setElementPage(e, p)
	}
}

func (p *pageInstance) FullData() ([]byte, error) {
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
func (p *pageInstance) SendMessage(id, cmd string, data any) {
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Println("send ws msg marshal error", err)
		return
	}
	msg := &Message{Eid: id, Kind: cmd, Data: string(marshal)}
	dataBys, err := json.Marshal(msg)
	if err != nil {
		log.Println("send marshal error:", err)
		return
	}
	log.Printf("send new message:%s", string(dataBys))
	conn := p.conn
	err = conn.WriteMessage(1, dataBys)
	if err != nil {
		log.Println("send error to conn:", conn.RemoteAddr(), err)
	}
}

func (p *pageInstance) OnNewWsMsg(msg *Message) {
	id := msg.Eid
	ele := findElement(p.root, id)
	if ele == nil {
		log.Printf("ERROR find element fail ,id:%s,msg:%v", msg.Eid, msg)
		return
	}
	ele.ctx = &handlerContext{page: p}
	for _, f := range ele.Handlers {
		f(msg)
	}
}
func (p *pageInstance) SetNewMsgHandler(f func(id string, cmd string, data any)) {
	p.onNewMsg = f
}

func (p *pageInstance) RouteTo(name string, data map[string]any) {
	uid := util.GenUUID()
	if data != nil {
		pageQueryData[uid] = data
	}
	p.SendMessage("EID0", "route", map[string]any{"name": name, "uuid": uid})
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
func Run(pages ...IPage) {
	addPage(pages...)
	if len(pageList) <= 0 {
		panic("at least one pageInstance is needed")
	}
	var p = ""
	root := findPage("")
	if root == nil {
		p = pageList[0].Name()
	}
	start(p)
}

type pageInstCreateContext struct {
	vw    *valuedWidget
	owner *pageInstance
	q     map[string]any
}

func (pw *pageInstCreateContext) Query() Query {
	return pw.q
}
func (pw *pageInstCreateContext) Button(name string, onClick func(self *Button)) *Button {
	tmp := createButton(name, onClick)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Link(name string) *Link {
	tmp := createLink(name)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Box(elements ...IWidget) *Box {
	tmp := createBox(elements...)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) List() *List {
	tmp := NewList()
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Input(onChange func(ctx *Input, val string)) *Input {
	tmp := NewInput(onChange)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Card(content string) *Card {
	tmp := NewCard(content)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Label(text string) *Label {
	tmp := NewLabel(0, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) H1(text string) *Label {
	tmp := NewLabel(1, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) H2(text string) *Label {
	tmp := NewLabel(2, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) H3(text string) *Label {
	tmp := NewLabel(3, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) H4(text string) *Label {
	tmp := NewLabel(4, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) H5(text string) *Label {
	tmp := NewLabel(5, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) H6(text string) *Label {
	tmp := NewLabel(6, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Checkbox(state bool, text string) *Checkbox {
	tmp := NewCheckbox(state, text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Radio(selected string, items ...string) *Radio {
	tmp := NewRadio(selected, items...)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Select(selected string, items ...string) *Select {
	tmp := NewSelect(selected, items...)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Switch(on bool) *Switch {
	tmp := NewSwitch(on)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) DateTime(t time.Time) *DateTime {
	tmp := NewDateTime(t)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Menu(opt *menu.Option) *Menu {
	tmp := NewMenu(opt)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Tab() *Tab {
	tmp := NewTab()
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Table(data interface{}) *Table {
	tmp := NewTable(data)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Loading(text string) *Loading {
	tmp := NewLoading(text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Progress(percent float32) *Progress {
	tmp := NewProgress(percent)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Description(cols int, data interface{}) *Description {
	tmp := NewDescription(cols, data)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Badge(count int) *Badge {
	tmp := NewBadge(count)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Divider() *Divider {
	tmp := NewDivider()
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Row(items ...IWidget) *Row {
	tmp := NewRow(items...)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) Timeline(opts ...*timeline.Option) *Timeline {
	tmp := NewTimeline(opts...)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Drawer(header string) *Drawer {
	tmp := NewDrawer(header)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Dropdown(text string, options ...string) *Dropdown {
	tmp := NewDropdown(text, options...)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) Tag(text string) *Tag {
	tmp := NewTag(text)
	pw.owner.root.AddChildren(tmp)
	return tmp
}
func (pw *pageInstCreateContext) TagInput(f func(self *TagInput, values []string)) *TagInput {
	tmp := NewTagInput(f)
	pw.owner.root.AddChildren(tmp)
	return tmp
}

func (pw *pageInstCreateContext) RouteTo(name string, data map[string]any) {
	pw.owner.RouteTo(name, data)
}
func (pw *pageInstCreateContext) SetTitle(title string) {
	pw.owner.title = title
}
