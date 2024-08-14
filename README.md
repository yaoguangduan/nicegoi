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
	"github.com/yaoguangduan/nicegoi/internal/ui"
)

func main() {
	goi.H4("Hello NiceGOI!")
	goi.Button("Button", func(self *ui.Button) {
		goi.MsgSuccess("You Clicked!")
	})
	goi.Run()
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