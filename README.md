### NiceGoi
NiceGOI is an easy-to-use, Golang-based UI framework, which shows up in your web browser. You can create buttons,inputs,card and much more.

It is great for micro web apps,tools app.

##### run belong command to see demo:
```shell
go run github.com/yaoguangduan/nicegoi/cmd/nicegoi_demo@latest
```

##### Features
- write like native gui,use like web server
- browser-based graphical user interface
- standard GUI elements like label, button, checkbox, switch, slider, input

##### Usage

```go
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

```

##### Showcase

###### buttons
![Button](./docs/buttons.png)
###### inputs
![Input](./docs/inputs.png)
###### links
![Link](./docs/links.png)
###### tags
![tag](./docs/tags.png)
###### tag-input
![ti](./docs/tag-inputs.png)
###### progress
![progress](./docs/progress.png)

###### add more...