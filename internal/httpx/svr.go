package httpx

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"nicegoi/internal/ws"
	"os/exec"
	"runtime"
	"time"
)

//go:embed asset/*
var staticFiles embed.FS

func Run() {

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", ws.HandleWebSocket)

	sub, err := fs.Sub(staticFiles, "asset")
	if err != nil {
		log.Fatal(err)
	}
	fileServer := http.FileServer(http.FS(sub))
	mux.Handle("/", fileServer)
	log.Println("Serving at localhost:8848...")
	go func() {
		time.Sleep(time.Second)
		err = openBrowser("http://127.0.0.1:8848")
		if err != nil {
			panic(err)
		}
	}()
	log.Fatal(http.ListenAndServe(":8848", mux))
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
