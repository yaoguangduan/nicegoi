package nice

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"sync"
)

type PageDefFn func(goiCtx GoiContext) []IWidget

func Page(name string, fn PageDefFn) {
	if fn == nil {
		return
	}
	layoutFuncList = append(layoutFuncList, pageDef{name: name, fn: fn})
}

var pageQueryData = make(map[string]map[string]any)

var layoutFuncList = make([]pageDef, 0)

type pageDef struct {
	name string
	fn   PageDefFn
}

func findPageDefFn(name string) PageDefFn {
	for _, d := range layoutFuncList {
		if d.name == name {
			return d.fn
		}
	}
	return nil
}

type pageManager struct {
	sync.Mutex
	pages map[string]*page
}

func (pm *pageManager) getOrCreate(name string, uuid string) *page {
	p, ok := pm.pages[uuid]
	if ok {
		return p
	}
	fn := findPageDefFn(name)
	if fn == nil {
		return nil
	}
	pm.Lock()
	defer pm.Unlock()
	p = &page{layout: fn, route: "/" + name, name: name, title: name, uuid: uuid}
	pw := &pageWidget{owner: p}
	p.delegate = pw
	pm.pages[uuid] = p
	return p
}

var pageMgr = &pageManager{pages: make(map[string]*page)}

type page struct {
	uuid     string
	layout   PageDefFn
	root     *element
	route    string
	name     string
	delegate *pageWidget
	title    string
	conn     *websocket.Conn
	onNewMsg func(id string, cmd string, data any)
}

func (p *page) serverNewCon(conn *websocket.Conn) {
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

func (p *page) handleRcv(conn *websocket.Conn) {
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

func (p *page) handlePageReq(w http.ResponseWriter, r *http.Request) {
	if p.root == nil {
		p.root = createElement("-")
		var data map[string]any
		if d, ok := pageQueryData[p.uuid]; ok {
			fmt.Println("owner receive query data:", d)
			data = d
		}
		cx := &pageCtx{q: make(map[string]any)}
		cx.q = data
		p.root.AddChildren(p.layout(cx)...)
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

func setElementPage(root *element, p *page) {
	root.Page = p
	for _, e := range root.Elements {
		setElementPage(e, p)
	}
}

func (p *page) FullData() ([]byte, error) {
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
func (p *page) SendMessage(id, cmd string, data any) {
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

func (p *page) OnNewWsMsg(msg *Message) {
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
func (p *page) SetNewMsgHandler(f func(id string, cmd string, data any)) {
	p.onNewMsg = f
}

func (p *page) RouteTo(name string, data map[string]any) {
	if data == nil {
		p.SendMessage("EID0", "route", map[string]any{"name": name})
	} else {
		v4, err := uuid.NewGen().NewV4()
		if err != nil {
			log.Println("new v4 err:", err)
			return
		}
		v4s := v4.String()
		uid := v4s[0:strings.Index(v4s, "-")]
		pageQueryData[uid] = data
		p.SendMessage("EID0", "route", map[string]any{"name": name, "uuid": uid})
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
func Run(fn PageDefFn) {
	Page("", fn)
	if len(layoutFuncList) <= 0 {
		panic("at least one page is needed")
	}
	var p = ""
	root := findPageDefFn("")
	if root == nil {
		p = layoutFuncList[0].name
	}
	start(p)
}

type pageWidget struct {
	vw    *valuedWidget
	owner *page
}

func (pw *pageWidget) MsgSuccess(msg string) {
	pw.owner.SendMessage("EID0", "message", map[string]interface{}{"level": 0, "msg": msg})
}
func (pw *pageWidget) MsgInfo(msg string) {
	pw.owner.SendMessage("EID0", "message", map[string]interface{}{"level": 1, "msg": msg})
}
func (pw *pageWidget) MsgWarn(msg string) {
	pw.owner.SendMessage("EID0", "message", map[string]interface{}{"level": 2, "msg": msg})
}
func (pw *pageWidget) MsgError(msg string) {
	pw.owner.SendMessage("EID0", "message", map[string]interface{}{"level": 3, "msg": msg})
}

func (pw *pageWidget) NotifySuccess(title, text string) {
	pw.owner.SendMessage("EID0", "notify", map[string]interface{}{"level": 0, "title": title, "text": text})
}
func (pw *pageWidget) NotifyInfo(title, text string) {
	pw.owner.SendMessage("EID0", "notify", map[string]interface{}{"level": 1, "title": title, "text": text})
}
func (pw *pageWidget) NotifyWarn(title, text string) {
	pw.owner.SendMessage("EID0", "notify", map[string]interface{}{"level": 2, "title": title, "text": text})
}
func (pw *pageWidget) NotifyError(title, text string) {
	pw.owner.SendMessage("EID0", "notify", map[string]interface{}{"level": 3, "title": title, "text": text})
}

func (pw *pageWidget) RouteTo(name string, data map[string]any) {
	pw.owner.RouteTo(name, data)
}
func (pw *pageWidget) SetTitle(title string) {
	pw.owner.title = title
}
