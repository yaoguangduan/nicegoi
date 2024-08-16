package main

import (
	"github.com/yaoguangduan/nicegoi/goi"
)

func main() {
	goi.H4("Hello NiceGOI!")
	goi.Button("Button", func(self *server.Button) {
		goi.MsgSuccess("You Clicked!")
	})
	goi.Run()
}
