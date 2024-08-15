package server

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
	rcvHandlers []func(route, uuid string, msg *msgs.Message)
	sender      chan *msgs.Message
	sendLock    sync.Mutex
	once        sync.Once
}

func (wsc *WsConnContext) serverNewCon(conn *websocket.Conn, uuid string) {
	wsc.once.Do(func() {
		go wsc.handleSend()
	})
	conn.SetCloseHandler(func(code int, text string) error {
		log.Println("warn:websocket closed ", code, text, uuid)
		wsc.connMap.Delete(uuid)
		return nil
	})
	wsc.connMap.Store(uuid, conn)
	go wsc.handleConn(conn, uuid)
}

var routeCtxMap = make(map[string]*WsConnContext)

func createWsConnContext(route string) *WsConnContext {
	wsc := &WsConnContext{
		Route:       route,
		rcvHandlers: make([]func(route, uuid string, msg *msgs.Message), 0),
		sender:      make(chan *msgs.Message, 1024),
	}
	return wsc
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

func RegMsgHandle(route string, handler func(string, string, *msgs.Message)) {
	ctx, ok := routeCtxMap[route]
	if !ok {
		ctx = createWsConnContext(route)
		routeCtxMap[route] = ctx
	}
	ctx.rcvHandlers = append(ctx.rcvHandlers, handler)
}

var handlerLock sync.Mutex

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
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
	uuid := r.URL.Query().Get("uuid")
	fmt.Println("new route conn:", path, uuid)
	context, ok := routeCtxMap[path]
	if !ok {
		context = createWsConnContext(path)
		routeCtxMap[path] = context
	}
	context.serverNewCon(conn, uuid)
}

func (wsc *WsConnContext) handleSend() {
	for msg := range wsc.sender {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Println("send marshal error:", err)
			continue
		}
		log.Printf("send new message:%s", string(data))
		wsc.connMap.Range(func(key, value interface{}) bool {
			conn := value.(*websocket.Conn)
			err = conn.WriteMessage(1, data)
			if err != nil {
				log.Println("send error to conn:", conn.RemoteAddr(), err)
			}
			return true
		})
	}
}

func (wsc *WsConnContext) handleConn(conn *websocket.Conn, uuid string) {
	var errCnt = 0
	for {
		up := &msgs.Message{}
		err := conn.ReadJSON(up)
		if err != nil {
			log.Println("Read error:", uuid, err)
			if errCnt > 3 {
				err = conn.Close()
				if err != nil {
					log.Println("close conn error:", uuid, err)
				}
				wsc.connMap.Delete(uuid)
				break
			} else {
				errCnt++
			}
			continue
		}
		errCnt = 0
		log.Printf("receive new msg:%+v", up)
		for _, f := range wsc.rcvHandlers {
			f(wsc.Route, uuid, up)
		}
	}
}
