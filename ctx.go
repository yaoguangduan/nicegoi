package nicegoi

import (
	"github.com/yaoguangduan/nicegoi/option/menu"
	"github.com/yaoguangduan/nicegoi/option/timeline"
	"time"
)

type Query map[string]any

func (q Query) Has(key string) bool {
	_, exist := q[key]
	return exist
}

func (q Query) Get(key string) any {
	return q[key]
}
func (q Query) GetOr(key string, def any) any {
	v, exist := q[key]
	if exist {
		return v
	}
	return def
}

type HandlerContext interface {
	RouteTo(name string, data map[string]any)
	ReloadPage()
	MsgSuccess(msg string)
	MsgInfo(msg string)
	MsgWarn(msg string)
	MsgError(msg string)
	NotifySuccess(title, text string)
	NotifyInfo(title, text string)
	NotifyWarn(title, text string)
	NotifyError(title, text string)
}

type handlerContext struct {
	page *pageInstance
}

func (pw *handlerContext) ReloadPage() {
	pw.page.SendMessage("EID0", "reload", "")
}

func (pw *handlerContext) RouteTo(name string, data map[string]any) {
	pw.page.RouteTo(name, data)
}
func (pw *handlerContext) MsgSuccess(msg string) {
	pw.page.SendMessage("EID0", "message", map[string]interface{}{"level": 0, "msg": msg})
}
func (pw *handlerContext) MsgInfo(msg string) {
	pw.page.SendMessage("EID0", "message", map[string]interface{}{"level": 1, "msg": msg})
}
func (pw *handlerContext) MsgWarn(msg string) {
	pw.page.SendMessage("EID0", "message", map[string]interface{}{"level": 2, "msg": msg})
}
func (pw *handlerContext) MsgError(msg string) {
	pw.page.SendMessage("EID0", "message", map[string]interface{}{"level": 3, "msg": msg})
}

func (pw *handlerContext) NotifySuccess(title, text string) {
	pw.page.SendMessage("EID0", "notify", map[string]interface{}{"level": 0, "title": title, "text": text})
}
func (pw *handlerContext) NotifyInfo(title, text string) {
	pw.page.SendMessage("EID0", "notify", map[string]interface{}{"level": 1, "title": title, "text": text})
}
func (pw *handlerContext) NotifyWarn(title, text string) {
	pw.page.SendMessage("EID0", "notify", map[string]interface{}{"level": 2, "title": title, "text": text})
}
func (pw *handlerContext) NotifyError(title, text string) {
	pw.page.SendMessage("EID0", "notify", map[string]interface{}{"level": 3, "title": title, "text": text})
}

type PageContext interface {
	Query() Query
	SetTitle(title string)
	Button(name string, onClick func(self *Button)) *Button
	Link(name string) *Link
	Box(elements ...IWidget) *Box
	List() *List
	Input(onChange func(ctx *Input, val string)) *Input
	Card(content string) *Card
	Label(text string) *Label
	H1(text string) *Label
	H2(text string) *Label
	H3(text string) *Label
	H4(text string) *Label
	H5(text string) *Label
	H6(text string) *Label
	Checkbox(state bool, text string) *Checkbox
	Radio(selected string, items ...string) *Radio
	Select(selected string, items ...string) *Select
	Switch(on bool) *Switch
	DateTime(t time.Time) *DateTime
	Menu(opt *menu.Option) *Menu
	Tab() *Tab
	Table(data interface{}) *Table
	Loading(text string) *Loading
	Progress(percent float32) *Progress
	Description(cols int, data interface{}) *Description
	Badge(count int) *Badge
	Divider() *Divider
	Row(items ...IWidget) *Row
	Timeline(opts ...*timeline.Option) *Timeline
	Drawer(header string) *Drawer
	Dropdown(text string, options ...string) *Dropdown
	Tag(text string) *Tag
	TagInput(f func(self *TagInput, values []string)) *TagInput
}
