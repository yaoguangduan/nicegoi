package ui

import (
	"nicegoi/internal/msgs"
	"nicegoi/internal/ui/icons"
	"nicegoi/internal/ws"
	"time"
)

//========================label=========================

type GoiLabel struct {
	*valueWidget
}

func (gb *GoiLabel) SetText(text string) {
	gb.set(text)
}

func NewGoiLabel(level int, text string) *GoiLabel {
	vw := newValueWidget("label", text)
	vw.opt.Set("level", level)
	l := &GoiLabel{valueWidget: vw}
	return l
}

//========================link=========================

type Link struct {
	*valueWidget
	onClick func(link *Link)
}

func (gb *Link) SetText(text string) {
	gb.opt.Set("value", text)
}
func (gb *Link) SetHref(href string) *Link {
	gb.opt.Set("href", href)
	return gb
}
func (gb *Link) SetOnClick(onClick func(self *Link)) *Link {
	gb.onClick = onClick
	return gb
}
func (gb *Link) SetPrefixIcon(icon icons.Icon) *Link {
	gb.opt.Set("prefix_icon", icon)
	return gb
}
func (gb *Link) SetSuffixIcon(icon icons.Icon) *Link {
	gb.opt.Set("suffix_icon", icon)
	return gb
}

func NewLink(text string) *Link {
	btn := &Link{valueWidget: newValueWidget("link", text)}
	ws.RegMsgHandle(btn.opt.Eid(), func(msg *msgs.Message) {
		if btn.onClick != nil {
			btn.onClick(btn)
		}
	})
	return btn
}

//========================button=========================

type GoiButton struct {
	*valueWidget
	onClick func(self *GoiButton)
}

func (gb *GoiButton) SetText(text string) {
	gb.opt.Set("value", text)
}
func (gb *GoiButton) SetIcon(icon icons.Icon) *GoiButton {
	gb.opt.Set("icon", icon)
	return gb
}
func (gb *GoiButton) SetOnClick(f func(self *GoiButton)) {
	gb.onClick = f
}

func NewGoiButton(text string, onClick func(self *GoiButton)) *GoiButton {
	btn := &GoiButton{valueWidget: newValueWidget("button", text), onClick: onClick}
	ws.RegMsgHandle(btn.opt.Eid(), func(msg *msgs.Message) {
		if btn.onClick != nil {
			btn.onClick(btn)
		}
	})
	return btn
}

//========================checkbox=========================

type GoiCheckbox struct {
	*valueWidget
	onChange func(self *GoiCheckbox, v bool)
}

func NewCheckbox(state bool, text string) *GoiCheckbox {
	c := &GoiCheckbox{valueWidget: newValueWidget("checkbox", state)}
	c.opt.Set("text", text)
	return c
}

func (c *GoiCheckbox) OnChange(f func(self *GoiCheckbox, checked bool)) *GoiCheckbox {
	c.onValChange(func(v any) {
		if f != nil {
			f(c, v.(bool))
		}
	})
	return c
}

//========================radio=========================

type Radio struct {
	*valueWidget
	onChange func(self *Radio, v string)
}

func NewRadio(selected string, items ...string) *Radio {
	c := &Radio{valueWidget: newValueWidget("radio", selected)}
	c.opt.Set("items", items)
	return c
}

func (c *Radio) OnChange(f func(self *Radio, selected string)) *Radio {
	c.onValChange(func(v any) {
		if f != nil {
			f(c, v.(string))
		}
	})
	return c
}

func (c *Radio) Select(val string) {
	c.set(val)
}

//========================Select=========================

type Select struct {
	*valueWidget
	onChange func(self *Select, v string)
}

func NewSelect(selected string, items ...string) *Select {
	c := &Select{valueWidget: newValueWidget("select", selected)}
	c.opt.Set("items", items)
	return c
}

func (c *Select) OnChange(f func(self *Select, selected string)) *Select {
	c.onValChange(func(v any) {
		if v == nil {
			v = ""
		}
		if f != nil {
			f(c, v.(string))
		}
	})
	return c
}

func (c *Select) Select(val string) {
	c.set(val)
}

//========================switch=========================

type Switch struct {
	*valueWidget
	onChange func(self *Switch, v string)
}

func NewSwitch(on bool) *Switch {
	c := &Switch{valueWidget: newValueWidget("switch", on)}
	return c
}

func (c *Switch) OnChange(f func(self *Switch, on bool)) *Switch {
	c.onValChange(func(v any) {
		if f != nil {
			f(c, v.(bool))
		}
	})
	return c
}
func (c *Switch) SetState(val bool) {
	c.set(val)
}

//========================input=========================

const (
	inputPlaceholder = "placeholder"
)

type Input struct {
	*valueWidget
}

func NewInput(onChange func(self *Input, val string)) *Input {
	w := &Input{valueWidget: newValueWidget("input", "")}
	w.onValChange(func(v any) {
		if onChange != nil {
			onChange(w, v.(string))
		}
	})
	return w
}
func (w *Input) SetPrepend(prepend string) *Input {
	w.opt.Set("prepend", prepend)
	return w
}
func (w *Input) SetAppend(append string) *Input {
	w.opt.Set("append", append)
	return w
}
func (w *Input) SetIcon(icon icons.Icon) *Input {
	w.opt.Set("icon", icon)
	return w
}
func (w *Input) EnablePassword() *Input {
	w.opt.Set("password", true)
	return w
}
func (w *Input) SetValue(value string) *Input {
	w.set(value)
	return w
}
func (w *Input) PlaceHolder(pl string) *Input {
	w.opt.Set(inputPlaceholder, pl)
	return w
}

// ========================loading=========================

type Loading struct {
	*valueWidget
}

func NewLoading(text string) *Loading {
	w := &Loading{valueWidget: newValueWidget("loading", false)}
	w.opt.Set("text", text)
	return w
}
func (w *Loading) Start() *Loading {
	w.set(true)
	return w
}
func (w *Loading) Stop() *Loading {
	w.set(false)
	return w
}
func (w *Loading) AddItems(items ...IWidget) *Loading {
	w.opt.AddChildren(items...)
	return w
}
func (w *Loading) GetState() bool {
	return w.get().(bool)
}

func (w *Loading) FullScreen() *Loading {
	w.opt.Set("fullscreen", true)
	return w
}

//========================progress=========================

type state string

const (
	ProgressActive  state = "active"
	ProgressError   state = "error"
	ProgressWarning state = "warning"
	ProgressSuccess state = "success"
)

type Progress struct {
	*valueWidget
}

func NewProgress(percent float32) *Progress {
	w := &Progress{valueWidget: newValueWidget("progress", percent)}
	w.opt.Set("theme", "line")
	return w
}
func (w *Progress) Update(percent float32) *Progress {
	w.set(percent)
	return w
}
func (w *Progress) Current() float32 {
	return w.get().(float32)
}
func (w *Progress) MarkState(s state) *Progress {
	w.opt.Set("state", s)
	return w
}
func (w *Progress) CircleStyle() *Progress {
	w.opt.Set("theme", "circle")
	return w
}

// ========================badge=========================

type Badge struct {
	*valueWidget
}

func NewBadge(count int) *Badge {
	w := &Badge{valueWidget: newValueWidget("badge", count)}
	return w
}
func (w *Badge) Count() int {
	return w.get().(int)
}
func (w *Badge) Incr(i int) *Badge {
	w.set(w.Count() + i)
	return w
}
func (w *Badge) Decr(i int) *Badge {
	w.set(w.Count() - i)
	return w
}
func (w *Badge) SetChild(c IWidget) *Badge {
	w.opt.AddChildren(c)
	return w
}

//========================description=========================

type Description struct {
	*valueWidget
}

// NewDescription map[string]string/struct
func NewDescription(cols int, daa interface{}) *Description {
	w := &Description{valueWidget: newValueWidget("description", cols)}
	w.opt.Set("data", daa)
	return w
}

//========================datetime=========================

const datetimeFormat = "2006-01-02 15:04:05"

type DateTime struct {
	*valueWidget
}

func NewDateTime(datetime time.Time) *DateTime {
	w := &DateTime{valueWidget: newValueWidget("datetime", datetime.Format(datetimeFormat))}
	return w
}
func (w *DateTime) OnChange(f func(self *DateTime, datetime time.Time, err error)) *DateTime {
	w.onValChange(func(v any) {
		if f != nil {
			t, err := time.Parse(datetimeFormat, v.(string))
			f(w, t, err)
		}
	})
	return w
}
func (w *DateTime) Set(value time.Time) {
	w.opt.Set("value", value.Format(datetimeFormat))
}
func (w *DateTime) Get() (time.Time, error) {
	dts := w.get().(string)
	parse, err := time.Parse(datetimeFormat, dts)
	return parse, err
}

//=======================================================================================

type valueWidget struct {
	opt IElement
	f   func(v any)
}

func (vw *valueWidget) Element() IElement {
	return vw.opt
}

func newValueWidget(kind string, value any) *valueWidget {
	w := &valueWidget{opt: NewElement(kind).Set("value", value)}
	ws.RegMsgHandle(w.opt.Eid(), func(message *msgs.Message) {
		w.opt.Modify("value", message.Data)
		if w.f != nil {
			w.f(message.Data)
		}
	})
	return w
}
func (vw *valueWidget) set(val any) {
	vw.opt.Set("value", val)
}
func (vw *valueWidget) get() any {
	return vw.opt.Get("value")
}

func (vw *valueWidget) Run() {
	vw.opt.Run()
}

func (vw *valueWidget) onValChange(f func(v any)) {
	vw.f = f
}

func (vw *valueWidget) Visible(visible bool) {
	vw.opt.Set("hide", !visible)
}
