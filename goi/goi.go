package goi

import (
	"github.com/yaoguangduan/nicegoi/internal"
	"github.com/yaoguangduan/nicegoi/internal/option/menu"
	"github.com/yaoguangduan/nicegoi/internal/option/timeline"
	"time"
)

func Button(name string, onClick func(self *nice.Button)) *nice.Button {
	return nice.NewButton(name, onClick)
}

func Link(text string) *nice.Link {
	return nice.NewLink(text)
}
func Box(elements ...nice.IWidget) *nice.Box {
	return nice.NewBox(elements...)
}
func List() *nice.List {
	return nice.NewList()
}
func Input(onChange func(ctx *nice.Input, val string)) *nice.Input {
	return nice.NewInput(onChange)
}
func Card(content string) *nice.Card {
	return nice.NewCard(content)
}
func Label(text string) *nice.Label {
	return nice.NewLabel(0, text)
}
func H1(text string) *nice.Label {
	return nice.NewLabel(1, text)
}
func H2(text string) *nice.Label {
	return nice.NewLabel(2, text)
}
func H3(text string) *nice.Label {
	return nice.NewLabel(3, text)
}
func H4(text string) *nice.Label {
	return nice.NewLabel(4, text)
}
func H5(text string) *nice.Label {
	return nice.NewLabel(5, text)
}
func H6(text string) *nice.Label {
	return nice.NewLabel(6, text)
}
func Checkbox(state bool, text string) *nice.Checkbox {
	return nice.NewCheckbox(state, text)
}
func Radio(selected string, items ...string) *nice.Radio {
	return nice.NewRadio(selected, items...)
}

func Select(selected string, items ...string) *nice.Select {
	return nice.NewSelect(selected, items...)
}

func Switch(on bool) *nice.Switch {
	return nice.NewSwitch(on)
}

func DateTime(t time.Time) *nice.DateTime {
	return nice.NewDateTime(t)
}

func Menu(opt *menu.Option) *nice.Menu {
	return nice.NewMenu(opt)
}

func Tab() *nice.Tab {
	return nice.NewTab()
}
func Table(data interface{}) *nice.Table {
	return nice.NewTable(data)
}

func Loading(text string) *nice.Loading {
	return nice.NewLoading(text)
}
func Progress(percent float32) *nice.Progress {
	return nice.NewProgress(percent)
}
func Description(cols int, data interface{}) *nice.Description {
	return nice.NewDescription(cols, data)
}
func Badge(count int) *nice.Badge {
	return nice.NewBadge(count)
}
func Divider() *nice.Divider {
	return nice.NewDivider()
}
func Row(items ...nice.IWidget) *nice.Row {
	return nice.NewRow(items...)
}

func Timeline(opts ...*timeline.Option) *nice.Timeline {
	return nice.NewTimeline(opts...)
}
func Drawer(header string) *nice.Drawer {
	return nice.NewDrawer(header)
}
func Dropdown(text string, options ...string) *nice.Dropdown {
	return nice.NewDropdown(text, options...)
}
func Tag(text string) *nice.Tag {
	return nice.NewTag(text)
}
func TagInput(f func(self *nice.TagInput, values []string)) *nice.TagInput {
	return nice.NewTagInput(f)
}
func Run() {
	nice.Run()
}
