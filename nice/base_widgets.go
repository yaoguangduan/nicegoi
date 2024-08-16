package nice

import (
	"github.com/yaoguangduan/nicegoi/nice/icons"
	"github.com/yaoguangduan/nicegoi/nice/option"
	"github.com/yaoguangduan/nicegoi/nice/option/timeline"
	"time"
)

//========================label=========================

type Label struct {
	*valuedWidget
}

func (gb *Label) SetText(text string) *Label {
	gb.set(text)
	return gb
}

func NewLabel(level int, text string) *Label {
	vw := newReadonlyWidget("label", text)
	vw.e.Set("level", level)
	l := &Label{valuedWidget: vw}
	return l
}

//========================divider=========================

type Divider struct {
	*valuedWidget
}

func (d *Divider) SetText(text string) *Divider {
	d.set(text)
	return d
}

func NewDivider() *Divider {
	vw := newReadonlyWidget("divider", "")
	vw.e.Set("layout", "horizontal")
	vw.e.Set("align", "left")
	l := &Divider{valuedWidget: vw}
	return l
}

func (d *Divider) Vertical() *Divider {
	d.e.Set("layout", "vertical")
	return d
}
func (d *Divider) AlignCenter() *Divider {
	d.e.Set("align", "center")
	return d
}
func (d *Divider) AlignRight() *Divider {
	d.e.Set("align", "right")
	return d
}

//========================link=========================

type Link struct {
	*valuedWidget
	onClick func(link *Link)
}

func (gb *Link) SetText(text string) {
	gb.e.Set("value", text)
}
func (gb *Link) SetHref(href string) *Link {
	gb.e.SetAttr("href", href)
	return gb
}
func (gb *Link) SetOnClick(onClick func(self *Link)) *Link {
	gb.onClick = onClick
	return gb
}
func (gb *Link) SetPrefixIcon(icon icons.Icon) *Link {
	gb.e.Set("prefix_icon", icon)
	return gb
}
func (gb *Link) SetSuffixIcon(icon icons.Icon) *Link {
	gb.e.Set("suffix_icon", icon)
	return gb
}
func (gb *Link) SetTheme(theme option.Theme) *Link {
	gb.e.SetAttr("theme", theme)
	return gb
}

func (gb *Link) SetSize(size option.Size) *Link {
	gb.e.SetAttr("size", size)
	return gb
}

func (gb *Link) Underlined() *Link {
	gb.e.SetAttr("underline", true)
	return gb
}

func NewLink(text string) *Link {
	btn := &Link{valuedWidget: newReadonlyWidget("link", text)}
	btn.addMsgHandler(func(msg *Message) {
		if btn.onClick != nil {
			btn.onClick(btn)
		}
	})
	btn.e.SetAttr("theme", option.Primary)
	return btn
}

//========================button=========================

type Button struct {
	*valuedWidget
	onClick func(self *Button)
}

func (gb *Button) SetText(text string) {
	gb.e.Set("value", text)
}
func (gb *Button) SetIcon(icon icons.Icon) *Button {
	gb.e.Set("icon", icon)
	return gb
}
func (gb *Button) SetOnClick(f func(self *Button)) {
	gb.onClick = f
}
func (gb *Button) SetTheme(theme option.Theme) *Button {
	gb.e.SetAttr("theme", theme)
	return gb
}
func (gb *Button) SetVariant(variant option.Variant) *Button {
	gb.e.SetAttr("variant", variant)
	return gb
}

func (gb *Button) SetSize(size option.Size) *Button {
	gb.e.SetAttr("size", size)
	return gb
}

func (gb *Button) SetShape(shape option.Shape) *Button {
	gb.e.SetAttr("shape", shape)
	return gb
}
func (gb *Button) Loading(loading bool) *Button {
	gb.e.SetAttr("loading", loading)
	return gb
}
func NewButton(text string, onClick func(self *Button)) *Button {
	btn := &Button{valuedWidget: newReadonlyWidget("button", text), onClick: onClick}
	btn.addMsgHandler(func(msg *Message) {
		if btn.onClick != nil {
			btn.onClick(btn)
		}
	})
	btn.e.SetAttr("theme", option.Primary)
	return btn
}

//========================checkbox=========================

type Checkbox struct {
	*valuedWidget
	onChange func(self *Checkbox, v bool)
}

func NewCheckbox(state bool, text string) *Checkbox {
	c := &Checkbox{valuedWidget: newValuedWidget("checkbox", state)}
	c.e.Set("text", text)
	return c
}

func (c *Checkbox) OnChange(f func(self *Checkbox, checked bool)) *Checkbox {
	c.onValChange(func(v any) {
		if f != nil {
			f(c, v.(bool))
		}
	})
	return c
}

//========================radio=========================

type Radio struct {
	*valuedWidget
	onChange func(self *Radio, v string)
}

func NewRadio(selected string, items ...string) *Radio {
	c := &Radio{valuedWidget: newValuedWidget("radio", selected)}
	c.e.Set("items", items)
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
	*valuedWidget
	onChange func(self *Select, v string)
}

func NewSelect(selected string, items ...string) *Select {
	c := &Select{valuedWidget: newValuedWidget("select", selected)}
	c.e.Set("items", items)
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
func (c *Select) SetClearable(clearable bool) *Select {
	c.e.SetAttr("clearable", clearable)
	return c
}
func (c *Select) SetFilterable(filterable bool) *Select {
	c.e.SetAttr("filterable", filterable)
	return c
}

func (c *Select) SetLoading(loading bool) *Select {
	c.e.SetAttr("loading", loading)
	return c
}

func (c *Select) SetLabel(label string) *Select {
	c.e.SetAttr("label", label)
	return c
}

func (c *Select) SetPlaceholder(placeholder string) *Select {
	c.e.SetAttr("placeholder", placeholder)
	return c
}

func (c *Select) SetSize(size option.Size) *Select {
	c.e.SetAttr("size", size)
	return c
}

//========================switch=========================

type Switch struct {
	*valuedWidget
	onChange func(self *Switch, v string)
}

func NewSwitch(on bool) *Switch {
	c := &Switch{valuedWidget: newValuedWidget("switch", on)}
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
	*valuedWidget
}

func NewInput(onChange func(self *Input, val string)) *Input {
	w := &Input{valuedWidget: newValuedWidget("input", "")}
	w.onValChange(func(v any) {
		if onChange != nil {
			onChange(w, v.(string))
		}
	})
	return w
}
func (w *Input) SetPrepend(prepend string) *Input {
	w.e.Set("prepend", prepend)
	return w
}
func (w *Input) SetAppend(append string) *Input {
	w.e.Set("append", append)
	return w
}
func (w *Input) SetIcon(icon icons.Icon) *Input {
	w.e.Set("icon", icon)
	return w
}
func (w *Input) EnablePassword() *Input {
	w.e.Set("password", true)
	return w
}
func (w *Input) SetValue(value string) *Input {
	w.set(value)
	return w
}
func (w *Input) PlaceHolder(pl string) *Input {
	w.e.Set(inputPlaceholder, pl)
	return w
}

func (w *Input) GetValue() string {
	return w.get().(string)
}

// ========================loading=========================

type Loading struct {
	*valuedWidget
}

func NewLoading(text string) *Loading {
	w := &Loading{valuedWidget: newValuedWidget("loading", false)}
	w.e.Set("text", text)
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
	w.e.AddChildren(items...)
	return w
}
func (w *Loading) GetState() bool {
	return w.get().(bool)
}

func (w *Loading) FullScreen() *Loading {
	w.e.Set("fullscreen", true)
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
	*valuedWidget
}

func NewProgress(percent float32) *Progress {
	w := &Progress{valuedWidget: newValuedWidget("progress", percent)}
	w.e.Set("theme", "line")
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
	w.e.Set("state", s)
	return w
}
func (w *Progress) CircleStyle() *Progress {
	w.e.Set("theme", "circle")
	return w
}

// ========================badge=========================

type Badge struct {
	*valuedWidget
}

func NewBadge(count int) *Badge {
	w := &Badge{valuedWidget: newValuedWidget("badge", count)}
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
	w.e.AddChildren(c)
	return w
}

//========================description=========================

type Description struct {
	*valuedWidget
}

// NewDescription map[string]string/struct
func NewDescription(cols int, daa interface{}) *Description {
	w := &Description{valuedWidget: newReadonlyWidget("description", cols)}
	w.e.Set("data", daa)
	return w
}

//========================datetime=========================

const datetimeFormat = "2006-01-02 15:04:05"

type DateTime struct {
	*valuedWidget
}

func NewDateTime(datetime time.Time) *DateTime {
	w := &DateTime{valuedWidget: newValuedWidget("datetime", datetime.Format(datetimeFormat))}
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
	w.e.Set("value", value.Format(datetimeFormat))
}
func (w *DateTime) Get() (time.Time, error) {
	dts := w.get().(string)
	parse, err := time.Parse(datetimeFormat, dts)
	return parse, err
}

//========================timeline=========================

type Timeline struct {
	*valuedWidget
}

func NewTimeline(options ...*timeline.Option) *Timeline {
	w := &Timeline{newReadonlyWidget("timeline", options)}
	w.e.Set("layout", "vertical")
	return w
}
func (w *Timeline) Horizontal() *Timeline {
	w.e.Set("layout", "horizontal")
	return w
}
func (w *Timeline) Add(options ...*timeline.Option) *Timeline {
	opts := w.get().([]*timeline.Option)
	opts = append(opts, options...)
	w.set(opts)
	return w
}

//========================dropdown=========================

type Dropdown struct {
	*valuedWidget
	f func(dropdown *Dropdown, value string)
}

func NewDropdown(text string, opts ...string) *Dropdown {
	w := &Dropdown{valuedWidget: newReadonlyWidget("dropdown", "")}
	w.e.Set("options", opts)
	w.e.Set("text", text)
	w.e.Set("theme", "primary")
	w.e.Set("variant", "text")
	return w
}
func (w *Dropdown) OnClick(f func(self *Dropdown, value string)) *Dropdown {
	w.onValChange(func(v any) {
		sv := v.(string)
		if f != nil {
			f(w, sv)
		}
	})
	return w
}
func (gb *Dropdown) SetTheme(theme option.Theme) *Dropdown {
	gb.e.Set("theme", theme)
	return gb
}
func (gb *Dropdown) SetVariant(variant option.Variant) *Dropdown {
	gb.e.Set("variant", variant)
	return gb
}

//========================dropdown=========================

type Tag struct {
	*valuedWidget
}

func NewTag(text string) *Tag {
	w := &Tag{valuedWidget: newReadonlyWidget("tag", text)}
	w.e.SetAttr("theme", "primary")
	w.e.SetAttr("variant", option.TagVarOutline)
	return w
}
func (gb *Tag) SetTheme(theme option.Theme) *Tag {
	gb.e.SetAttr("theme", theme)
	return gb
}
func (gb *Tag) SetVariant(variant option.TagVariant) *Tag {
	gb.e.SetAttr("variant", variant)
	return gb
}
func (gb *Tag) SetSize(size option.Size) *Tag {
	gb.e.SetAttr("size", size)
	return gb
}
func (gb *Tag) SetShape(shape option.Shape) *Tag {
	gb.e.SetAttr("shape", shape)
	return gb
}

//========================input=========================

type TagInput struct {
	*valuedWidget
}

func NewTagInput(onChange func(self *TagInput, values []string)) *TagInput {
	w := &TagInput{valuedWidget: newValuedWidget("tag-input", make([]string, 0))}
	w.onValChange(func(v any) {
		if onChange != nil {
			onChange(w, w.GetValues())
		}
	})
	return w
}
func (w *TagInput) SetPlaceHolder(pl string) *TagInput {
	w.e.SetAttr(inputPlaceholder, pl)
	return w
}

func (w *TagInput) SetTheme(theme option.Theme) *TagInput {
	w.e.SetAttr("theme", theme)
	return w
}

func (w *TagInput) SetMax(max int) *TagInput {
	w.e.SetAttr("max", max)
	return w
}

func (w *TagInput) SetLabel(label string) *TagInput {
	w.e.SetAttr("label", label)
	return w
}

func (w *TagInput) GetValues() []string {
	vv := make([]string, 0)
	for _, vvv := range w.get().([]interface{}) {
		vv = append(vv, vvv.(string))
	}
	return vv
}
