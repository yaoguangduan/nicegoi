package goi

import (
	"github.com/yaoguangduan/nicegoi/internal/option/menu"
	"github.com/yaoguangduan/nicegoi/internal/option/timeline"
	"github.com/yaoguangduan/nicegoi/internal/server"
	"github.com/yaoguangduan/nicegoi/internal/ui"
	"time"
)

func Button(name string, onClick func(self *ui.Button)) *ui.Button {
	return ui.NewButton(name, onClick)
}

func Link(text string) *ui.Link {
	return ui.NewLink(text)
}
func Box(elements ...ui.IWidget) *ui.Box {
	return ui.NewBox(elements...)
}
func List() *ui.List {
	return ui.NewList()
}
func Input(onChange func(ctx *ui.Input, val string)) *ui.Input {
	return ui.NewInput(onChange)
}
func Card(content string) *ui.GoiCard {
	return ui.NewCard(content)
}
func MsgSuccess(msg string) {
	ui.Message(0, msg)
}
func MsgInfo(msg string) {
	ui.Message(1, msg)
}
func MsgWarn(msg string) {
	ui.Message(2, msg)
}
func MsgError(msg string) {
	ui.Message(3, msg)
}

func NotifySuccess(title, text string) {
	ui.Notify(0, title, text)
}
func NotifyInfo(title, text string) {
	ui.Notify(1, title, text)
}
func NotifyWarn(title, text string) {
	ui.Notify(2, title, text)
}
func NotifyError(title, text string) {
	ui.Notify(3, title, text)
}
func Label(text string) *ui.Label {
	return ui.NewLabel(0, text)
}
func H1(text string) *ui.Label {
	return ui.NewLabel(1, text)
}
func H2(text string) *ui.Label {
	return ui.NewLabel(2, text)
}
func H3(text string) *ui.Label {
	return ui.NewLabel(3, text)
}
func H4(text string) *ui.Label {
	return ui.NewLabel(4, text)
}
func H5(text string) *ui.Label {
	return ui.NewLabel(5, text)
}
func H6(text string) *ui.Label {
	return ui.NewLabel(6, text)
}
func Checkbox(state bool, text string) *ui.Checkbox {
	return ui.NewCheckbox(state, text)
}
func Radio(selected string, items ...string) *ui.Radio {
	return ui.NewRadio(selected, items...)
}

func Select(selected string, items ...string) *ui.Select {
	return ui.NewSelect(selected, items...)
}

func Switch(on bool) *ui.Switch {
	return ui.NewSwitch(on)
}

func DateTime(t time.Time) *ui.DateTime {
	return ui.NewDateTime(t)
}

func Menu(opt *menu.Option) *ui.Menu {
	return ui.NewMenu(opt)
}

func Tab() *ui.Tab {
	return ui.NewTab()
}
func Table(data interface{}) *ui.Table {
	return ui.NewTable(data)
}

func Loading(text string) *ui.Loading {
	return ui.NewLoading(text)
}
func Progress(percent float32) *ui.Progress {
	return ui.NewProgress(percent)
}
func Description(cols int, data interface{}) *ui.Description {
	return ui.NewDescription(cols, data)
}
func Badge(count int) *ui.Badge {
	return ui.NewBadge(count)
}
func Divider() *ui.Divider {
	return ui.NewDivider()
}
func Row(items ...ui.IWidget) *ui.Row {
	return ui.NewRow(items...)
}

func Timeline(opts ...*timeline.Option) *ui.Timeline {
	return ui.NewTimeline(opts...)
}
func Drawer(header string) *ui.Drawer {
	return ui.NewDrawer(header)
}
func Dropdown(text string, options ...string) *ui.Dropdown {
	return ui.NewDropdown(text, options...)
}
func Tag(text string) *ui.Tag {
	return ui.NewTag(text)
}
func TagInput(f func(self *ui.TagInput, values []string)) *ui.TagInput {
	return ui.NewTagInput(f)
}
func Page(name string) *ui.PageWidget {
	return ui.NewPage(name)
}
func RouteTo(name string) {
}

func Run() {
	server.Run()
}
