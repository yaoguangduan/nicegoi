package main

import (
	"github.com/yaoguangduan/nicegoi/goi"
	"github.com/yaoguangduan/nicegoi/internal/ui"
)

func main() {
	goi.H4("Hello NiceGOI!")
	goi.Button("Button", func(self *ui.Button) {
		goi.MsgSuccess("You Clicked!")
	})
	goi.Run()
}
