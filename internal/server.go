package nice

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mux = http.NewServeMux()

func Run() {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/api/ws", handleWebSocket)

	mux.HandleFunc("/api/page", func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Query().Get("uuid")
		route := r.URL.Query().Get("route")
		var pg = pageMgr.getOrCreate(strings.TrimPrefix(route, "/"), uuid)
		if pg == nil {
			w.WriteHeader(200)
			write, err := w.Write([]byte("{\"error\":\"page not found\"}"))
			if err != nil {
				log.Println("page not found ret write err:", write, err)
			}
			return
		}
		pg.handlePageReq(w, r)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendWebContent(w, r)
	})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to listen on port 0: %v", err)
	}
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			log.Println("listener close err:", err)
		}
	}(listener)
	port := listener.Addr().(*net.TCPAddr).Port
	log.Println("open http://localhost:" + strconv.Itoa(port))
	go func() {
		time.Sleep(time.Millisecond * 500)
		err = openBrowser("http://127.0.0.1:" + strconv.Itoa(port))
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
