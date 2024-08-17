package main

import (
	"github.com/yaoguangduan/nicegoi"
)

type Home struct {
}

func (t *Home) Name() string {
	return ""
}

func (t *Home) Layout(ctx nicegoi.PageContext) {
	input := ctx.Input(nil)
	btn := ctx.Button("Click Me!", func(self *nicegoi.Button) {
		is := input.GetValue()
		if is == "" {
			self.Ctx().MsgInfo("hello world")
		} else {
			self.Ctx().MsgSuccess("hello " + is)
		}
	})
	ctx.Box(input, btn).Vertical()
}

func main() {
	nicegoi.Run(new(Home))
}
