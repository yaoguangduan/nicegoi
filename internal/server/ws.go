package httpx

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"log"
	"net/http"
	"sync"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

type WsConnContext struct {
	Route       string
	connMap     sync.Map
	rcvHandlers []func(msg *msgs.Message)
	sender      chan *msgs.Message
	receiver    chan *msgs.Message
	sendLock    sync.Mutex
	once        sync.Once
}

func (wsc *WsConnContext) serverNewCon(conn *websocket.Conn) {
	wsc.once.Do(func() {
		go wsc.handleSend()
		go wsc.handleReceive()
	})
	go wsc.handleConn(conn)
}

var routeCtxMap = make(map[string]*WsConnContext)

func createWsConnContext(route string) *WsConnContext {
	wsc := &WsConnContext{
		Route:       route,
		rcvHandlers: make([]func(msg *msgs.Message), 0),
		sender:      make(chan *msgs.Message),
		receiver:    make(chan *msgs.Message),
	}
	return wsc
}

var (
	onConnectHandlers = make([]func(route string, c *websocket.Conn), 0)
)

func Active(route string) bool {
	ctx, ok := routeCtxMap[route]
	if !ok {
		return false
	}
	var a = false
	ctx.connMap.Range(func(key, value interface{}) bool {
		a = true
		return false
	})
	return a
}

func Send(route string, eid string, kind string, msg any) {
	ctx, ok := routeCtxMap[route]
	if !ok {
		return
	}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("json marshal err %v", err)
		return
	}
	ctx.sender <- &msgs.Message{Eid: eid, Kind: kind, Data: string(data)}
}

func RegMsgHandle(route string, handler func(*msgs.Message)) {
	ctx, ok := routeCtxMap[route]
	if !ok {
		ctx = createWsConnContext(route)
		routeCtxMap[route] = ctx
	}
	ctx.rcvHandlers = append(ctx.rcvHandlers, handler)
}

func OnNewWsConn(f func(route string, conn *websocket.Conn)) {
	onConnectHandlers = append(onConnectHandlers, f)
}

var handlerLock sync.Mutex

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	handlerLock.Lock()
	defer handlerLock.Unlock()
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	query := r.URL.Query()
	path := query.Get("route")
	if path == "" {
		path = "/"
	}
	fmt.Println("new route conn:", path)
	context, ok := routeCtxMap[path]
	if !ok {
		context = createWsConnContext(path)
		routeCtxMap[path] = context
	}
	context.serverNewCon(conn)
}

func (wsc *WsConnContext) handleReceive() {
	for msg := range wsc.receiver {
		for _, f := range wsc.rcvHandlers {
			f(msg)
		}
	}
}

func (wsc *WsConnContext) handleSend() {
	for msg := range wsc.sender {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Println("send marshal error:", err)
			continue
		}
		log.Printf("send new message:%s", string(data))
		wsc.sendLock.Lock()
		var toDelete []*websocket.Conn
		wsc.connMap.Range(func(key, value interface{}) bool {
			conn := key.(*websocket.Conn)
			err = conn.WriteMessage(1, data)
			if err != nil {
				log.Println("send error:", err)
				toDelete = append(toDelete, conn)
			}
			return true
		})
		for _, c := range toDelete {
			wsc.connMap.Delete(c)
		}
		wsc.sendLock.Unlock()
	}
}

func (wsc *WsConnContext) handleConn(conn *websocket.Conn) {
	wsc.sendLock.Lock()
	for _, h := range onConnectHandlers {
		h(wsc.Route, conn)
	}
	wsc.connMap.Store(conn, nil)
	wsc.sendLock.Unlock()
	for {
		up := &msgs.Message{}
		err := conn.ReadJSON(up)
		if err != nil {
			log.Println("Read error:", err)
			err = conn.Close()
			if err != nil {
				log.Println("close conn error:", err)
			}
			wsc.connMap.Delete(conn)
			break
		}
		log.Printf("receive new msg:%+v", up)
		wsc.receiver <- up
	}
}
