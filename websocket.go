package nicegoi

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
	pg := pageInstMgr.getOrCreate("           ", uuid)
	if pg == nil {
		log.Printf("ERROR:can not find pageInstance of %s,no src websocket\n", uuid)
		_ = conn.Close()
		return
	}
	pg.serverNewCon(conn)
}
