package main

import (
	"github.com/yaoguangduan/nicegoi/goi"
	"github.com/yaoguangduan/nicegoi/nice"
)

func home(ctx nice.GoiContext) []nice.IWidget {
	input := goi.Input(nil)
	button := goi.Button("Click Me", func(self *nice.Button) {
		is := input.GetValue()
		if is == "" {
			self.Page().MsgInfo("hello world")
		} else {
			self.Page().MsgSuccess("hello " + is)
		}
	})
	v := goi.Box().Vertical().AddItems(input, button)
	return []nice.IWidget{v}
}

func main() {
	goi.Run(home)
}
