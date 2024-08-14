package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type IPage interface {
	Name() string
	Route() string
	OnInit()
	FullData() ([]byte, error)
	OnNewWsMsg(msg *msgs.Message)
	RouteTo(name string)
}

type pageManager struct {
	pagesRes map[string]IPage
}

func RegPageRes(page IPage) {
	pageMgr.pagesRes[page.Route()] = page
}
func (pm *pageManager) genPage(route string) IPage {
	p, ok := pm.pagesRes[route]
	if !ok {
		return nil
	}
	return p
}

var pageMgr = &pageManager{pagesRes: make(map[string]IPage)}

type Session struct {
	sync.Mutex
	ID      string
	padding []msgs.Message
	rcv     chan *msgs.Message
	page    IPage
}

//
//func (s *Session) HandleFetch(w http.ResponseWriter, r *http.Request) {
//	var data []byte
//	s.Lock()
//	data, err := json.Marshal(s.page.PaddingMsg())
//	if err != nil {
//		data = []byte("[]")
//	}
//	s.Unlock()
//	w.WriteHeader(200)
//	w.Header().Set("Content-Type", "application/json")
//	if string(data) != "[]" {
//		log.Println("write padding message", s.ID, string(data))
//	}
//	_, err = w.Write(data)
//	if err != nil {
//		log.Println("write err:", err)
//	}
//}

func (s *Session) HandleDelivery(w http.ResponseWriter, r *http.Request) {
	bys, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	msg := &msgs.Message{}
	err = json.Unmarshal(bys, msg)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	s.rcv <- msg
}

func (s *Session) handlePageReq(w http.ResponseWriter, r *http.Request) {
	if s.page == nil {
		route := r.URL.Query().Get("route")
		p := pageMgr.genPage(route)
		if p == nil {
			w.WriteHeader(404)
			_, _ = w.Write([]byte("404 page not found"))
		}
		s.page = p
	}
	if s.page != nil {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data, err := s.page.FullData()
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
		}
		_, _ = w.Write(data)
	}
}

func (s *Session) svrReceiver() {
	for message := range s.rcv {
		if s.page != nil {
			s.page.OnNewWsMsg(message)
		}
	}
}

func NewSession(uuid string) *Session {
	s := &Session{
		ID:      uuid,
		padding: []msgs.Message{},
		rcv:     make(chan *msgs.Message, 1024),
	}
	go s.svrReceiver()
	return s
}

type SessionManager struct {
	sync.Mutex
	Sessions map[string]*Session
}

func (sm *SessionManager) GetSession(id string) *Session {
	session, ok := sm.Sessions[id]
	if !ok {
		return nil
	}
	return session
}

func (sm *SessionManager) AddSession(s *Session) {
	sm.Lock()
	defer sm.Unlock()
	sm.Sessions[s.ID] = s
}

var sessionMgr = &SessionManager{sync.Mutex{}, make(map[string]*Session)}

var mux = http.NewServeMux()

func Run() {
	for _, p := range pageMgr.pagesRes {
		p.OnInit()
		RegMsgHandle(p.Route(), func(message *msgs.Message) {
			p.OnNewWsMsg(message)
		})
	}

	// 注册 pprof 的路由
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/api/ws", handleWebSocket)
	//mux.HandleFunc("/api/fetch", func(w http.ResponseWriter, r *http.Request) {
	//	uuid := r.URL.Query().Get("uuid")
	//	var session = sessionMgr.GetSession(uuid)
	//	if session == nil {
	//		w.WriteHeader(500)
	//		_, _ = w.Write([]byte("session not found"))
	//		return
	//	}
	//	session.HandleFetch(w, r)
	//})
	//mux.HandleFunc("/api/delivery", func(w http.ResponseWriter, r *http.Request) {
	//	uuid := r.URL.Query().Get("uuid")
	//	var session = sessionMgr.GetSession(uuid)
	//	if session == nil {
	//		w.WriteHeader(500)
	//		_, _ = w.Write([]byte("session not found"))
	//		return
	//	}
	//	session.HandleDelivery(w, r)
	//})
	mux.HandleFunc("/api/page", func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("uuid")
		var session = sessionMgr.GetSession(uuid)
		if session == nil {
			session = NewSession(uuid)
			sessionMgr.AddSession(session)
		}
		session.handlePageReq(w, r)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sendWebContent(w, r)
	})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen on port 0: %v", err)
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	log.Println("open http://localhost:" + strconv.Itoa(port))
	go func() {
		time.Sleep(time.Millisecond * 500)
		err := openBrowser("http://127.0.0.1:" + strconv.Itoa(port))
		if err != nil {
			panic(err)
		}
	}()
	log.Fatal(http.Serve(listener, mux))
}

// openBrowser 打开系统默认浏览器并访问指定的 URL
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}
	return exec.Command(cmd, args...).Start()
}

//go:embed index.html
var webFile embed.FS

var webContent []byte
var webLoad sync.Once

func loadContent() {
	file, err := webFile.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	webContent = file
}

func sendWebContent(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	webLoad.Do(loadContent)
	_, err := w.Write(webContent)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
}
