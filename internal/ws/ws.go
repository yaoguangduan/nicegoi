package ws

import (
	"encoding/json"
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
var (
	connMap           = sync.Map{}
	rcvHandlers       = make(map[string]func(msg *msgs.Message), 0)
	onConnectHandlers = make([]func(c *websocket.Conn), 0)
	receiver          = make(chan *msgs.Message, 1024)
	sender            = make(chan *msgs.Message, 1024)
	sendLock          sync.Mutex
)

func Active() bool {
	var a = false
	connMap.Range(func(key, value interface{}) bool {
		a = true
		return false
	})
	return a
}

func Send(eid string, kind string, msg any) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("json marshal err %v", err)
		return
	}
	sender <- &msgs.Message{Eid: eid, Kind: kind, Data: string(data)}
}

func RegMsgHandle(eid string, handler func(*msgs.Message)) {
	rcvHandlers[eid] = handler
}

func OnNewConn(f func(conn *websocket.Conn)) {
	onConnectHandlers = append(onConnectHandlers, f)
}

var once sync.Once

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	once.Do(func() {
		go handleSend()
		go handleReceive()
	})
	go handleConn(conn)
}

func handleReceive() {
	for msg := range receiver {
		eid := msg.Eid
		h, exist := rcvHandlers[eid]
		if exist {
			h(msg)
		}
	}
}

func handleSend() {
	for msg := range sender {
		data, err := json.Marshal(msg)
		if err != nil {
			log.Println("send marshal error:", err)
			continue
		}
		log.Printf("send new message:%s", string(data))
		sendLock.Lock()
		var toDelete []*websocket.Conn
		connMap.Range(func(key, value interface{}) bool {
			conn := key.(*websocket.Conn)
			err = conn.WriteMessage(1, data)
			if err != nil {
				log.Println("send error:", err)
				toDelete = append(toDelete, conn)
			}
			return true
		})
		for _, c := range toDelete {
			connMap.Delete(c)
		}
		sendLock.Unlock()
	}
}

func handleConn(conn *websocket.Conn) {
	sendLock.Lock()
	for _, h := range onConnectHandlers {
		h(conn)
	}
	connMap.Store(conn, nil)
	sendLock.Unlock()
	for {
		up := &msgs.Message{}
		err := conn.ReadJSON(up)
		if err != nil {
			log.Println("Read error:", err)
			err = conn.Close()
			if err != nil {
				log.Println("close conn error:", err)
			}
			connMap.Delete(conn)
			break
		}
		log.Printf("receive new msg:%+v", up)
		receiver <- up
	}
}
