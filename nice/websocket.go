package nice

import (
	"github.com/gorilla/websocket"
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
	conn        *websocket.Conn
	rcvHandlers []func(route, uuid string, msg *Message)
	sendLock    sync.Mutex
	once        sync.Once
	uuid        string
}

var connCtxMap = make(map[string]*WsConnContext)

func createWsConnContext(uuid string, path string) *WsConnContext {
	wsc := &WsConnContext{
		uuid:        uuid,
		Route:       path,
		rcvHandlers: make([]func(route, uuid string, msg *Message), 0),
	}
	return wsc
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
	uuid := r.URL.Query().Get("uuid")
	pg := pageMgr.getOrCreate("           ", uuid)
	if pg == nil {
		log.Printf("ERROR:can not find page of %s,no src websocket\n", uuid)
		_ = conn.Close()
		return
	}
	pg.serverNewCon(conn)
}
