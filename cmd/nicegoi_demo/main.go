package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/yaoguangduan/nicegoi"
	"github.com/yaoguangduan/nicegoi/icons"
	"github.com/yaoguangduan/nicegoi/option"
	"github.com/yaoguangduan/nicegoi/option/menu"
	"github.com/yaoguangduan/nicegoi/option/timeline"
	"time"
)

type Hello struct {
}

func (h *Hello) Name() string {
	return "hello"
}

func (h *Hello) Layout(ctx nicegoi.PageContext) {
	ctx.Label(fmt.Sprintf("hello %v", ctx.Query().GetOr("name", "world!")))
}

type Main struct {
}

func (mai *Main) Name() string {
	return ""
}

func (mai *Main) Layout(ctx nicegoi.PageContext) {

	card := ctx.Card("i am card content")

	card.AddActions(ctx.Link("card Action"))
	card.AddFooters(ctx.Link("add"), ctx.Link("Setting"))
	card.SetTitle("card title")
	tab := ctx.Tab().SetPlace(option.Left)
	loading := ctx.Loading("loading...")
	fullLoading := ctx.Loading("loadingAll...").FullScreen()
	disable := ctx.Button("Disable", nil)
	disable.SetDisable(true)

	tags := make([]nicegoi.IWidget, 0)
	variants := []option.TagVariant{option.TagVarDark, option.TagVarOutline, option.TagVarLight, option.TagVarLightOutline}
	themes := []option.Theme{option.Default, option.Primary, option.Danger, option.Success, option.Warning}
	for _, variant := range variants {
		for _, theme := range themes {
			tags = append(tags, ctx.Tag(fmt.Sprintf("%s/%s", variant, theme)).SetVariant(variant).SetTheme(theme))
		}
	}

	radio := ctx.Radio("value 1", "value 1", "value 2", "value 3").OnChange(func(self *nicegoi.Radio, selected string) {
		self.Ctx().MsgSuccess(fmt.Sprintf("you has selected %s", selected))
	})

	st := ctx.Select("banana", "banana", "apple", "orange")

	p0 := ctx.Progress(20)

	l := ctx.List()
	l.NewItem("list item 1").AddAction(ctx.Link("operate 1").SetOnClick(func(self *nicegoi.Link) {
		self.Ctx().MsgInfo("item 1 operated")
	}))
	l.NewItem("list item 2").AddAction(ctx.Link("operate 2").SetOnClick(func(self *nicegoi.Link) {
		self.Ctx().MsgInfo("item 2 operated")
	}))

	m := menu.New().AddItems(
		menu.NewItem("MenuOption 1", "m1").SetIcon(icons.Home).AddItems(menu.NewItem("MenuOption 1-1", "m11"), menu.NewItem("MenuOption 1-2", "m12")),
		menu.NewItem("MenuOption 2", "m2").SetIcon(icons.Edit),
	)

	p := ctx.Progress(20).CircleStyle()

	dr := ctx.Drawer("Header").AddWidgets(ctx.Label("drawer content"))
	dl := ctx.Drawer("Header").AddWidgets(ctx.Label("drawer content")).SetPlace(option.Left)
	dt := ctx.Drawer("Header").AddWidgets(ctx.Label("drawer content")).SetPlace(option.Top)
	db := ctx.Drawer("Header").AddWidgets(ctx.Label("drawer content")).SetPlace(option.Bottom)

	badge := ctx.Badge(29)

	gotoInput := ctx.Input(nil).PlaceHolder("input your name")

	ctx.H1("Hello NiceGOI!")

	ctx.Divider().SetText("button")
	ctx.Box(
		ctx.H6("Button:"),
		ctx.Button("button", func(self *nicegoi.Button) {
			self.Ctx().MsgWarn("button clicked!")
		}),
		ctx.Divider().Vertical(),
		ctx.H6("with Icon:"),
		ctx.Button("Icon", nil).SetIcon(icons.Add),
		ctx.Divider().Vertical(),
		ctx.Button("Success", nil).SetTheme(option.Success),
		ctx.Divider().Vertical(),
		ctx.Button("Danger", nil).SetTheme(option.Danger),
		ctx.Divider().Vertical(),
		ctx.Button("Warning", nil).SetTheme(option.Warning),
	)
	ctx.Box(
		ctx.H6("Variant:"),
		ctx.Button("Outline", nil).SetVariant(option.Outline),
		ctx.Button("Dashed", nil).SetVariant(option.Dashed),
		ctx.Button("Text", nil).SetVariant(option.Text),
	)
	ctx.Box(
		ctx.H6("Shape:"),
		ctx.Button("Rectangle", nil).SetShape(option.Rectangle),
		ctx.Button("", nil).SetIcon(icons.Home).SetShape(option.Square),
		ctx.Button("Round", nil).SetShape(option.Round),
		ctx.Button("", nil).SetIcon(icons.Edit).SetShape(option.Circle),
	)
	ctx.Box(
		ctx.H6("ProgressState/Size:"),
		disable,
		ctx.Button("Loading", nil).Loading(true),
		ctx.Button("Small", nil).SetSize(option.Small),
		ctx.Button("Large", nil).SetSize(option.Large),
	)

	ctx.Divider().SetText("input:")
	ctx.Box(
		ctx.H6("Input:"),
		ctx.Input(func(ctx *nicegoi.Input, val string) {
			ctx.Ctx().MsgWarn(fmt.Sprintf("you input:%s", val))
		}),
		ctx.Divider().Vertical(),
		ctx.H6("with prefix/suffix:"),
		ctx.Input(nil).SetPrepend("https://").SetAppend(".com"),
		ctx.Divider().Vertical(),
		ctx.H6("with icon:"),
		ctx.Input(nil).SetIcon(icons.ArrowRight).PlaceHolder("example.com"),
		ctx.Divider().Vertical(),
		ctx.H6("password:"),
		ctx.Input(nil).EnablePassword(),
	)

	ctx.Divider().SetText("link")
	ctx.Box(
		ctx.H6("link:"),
		ctx.Link("点击").SetOnClick(func(self *nicegoi.Link) {
			self.Ctx().NotifySuccess("I am Title", "you clicked link!")
		}),
		ctx.Divider().Vertical(),
		ctx.H6("with icon:"),
		ctx.Link("跳转链接").SetPrefixIcon(icons.Link).SetHref("https://google.com"),
		ctx.Divider().Vertical(),
		ctx.Link("Success").SetTheme(option.Success),
		ctx.Divider().Vertical(),
		ctx.Link("Danger").SetTheme(option.Danger),
		ctx.Divider().Vertical(),
		ctx.Link("Warning").SetTheme(option.Warning),
		ctx.Divider().Vertical(),
		ctx.Link("Small").SetSize(option.Small),
		ctx.Divider().Vertical(),
		ctx.Link("Large").SetSize(option.Large),
		ctx.Divider().Vertical(),
		ctx.Link("underline").Underlined(),
	)
	ctx.Divider().SetText("dropdown")
	ctx.Box(
		ctx.Dropdown("more...", "option 1", "option 2", "option 3").OnClick(func(self *nicegoi.Dropdown, value string) {
			self.Ctx().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
		ctx.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Dashed).OnClick(func(self *nicegoi.Dropdown, value string) {
			self.Ctx().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
		ctx.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Base).OnClick(func(self *nicegoi.Dropdown, value string) {
			self.Ctx().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
		ctx.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Outline).OnClick(func(self *nicegoi.Dropdown, value string) {
			self.Ctx().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
	)
	ctx.Divider().SetText("tag")
	ctx.Box(tags...)
	ctx.Box(
		ctx.H6("Size:"),
		ctx.Tag("medium").SetSize(option.Medium),
		ctx.Tag("small").SetSize(option.Small),
		ctx.Tag("large").SetSize(option.Large),
	)
	ctx.Box(
		ctx.H6("Shape:"),
		ctx.Tag("round").SetShape(option.Round),
		ctx.Tag("square").SetShape(option.Square),
		ctx.Tag("mark").SetShape(option.Mark),
	)

	ctx.Divider().SetText("tag input")
	ctx.Row(
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifySuccess("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Success).SetPlaceHolder("success theme"),
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifyInfo("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Primary).SetPlaceHolder("primary theme"),
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifyInfo("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Default).SetPlaceHolder("default theme"),
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifyError("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Danger).SetPlaceHolder("danger theme"),
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifyWarn("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Warning).SetPlaceHolder("warning theme"),
	).SetSpan(2, 2, 2, 2, 2).SetGutter(10, 0).Justify(option.RowSpaceAround)
	ctx.Box()
	ctx.Row(
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifySuccess("tag input", fmt.Sprintf("%v", values))
		}).SetLabel("With Label:").SetPlaceHolder("labeled input"),
		ctx.TagInput(func(self *nicegoi.TagInput, values []string) {
			self.Ctx().NotifyInfo("tag input", fmt.Sprintf("%v", values))
		}).SetMax(3).SetPlaceHolder("max 3 tags"),
	).SetSpan(4, 4).SetGutter(10, 0).Justify(option.RowStart)

	ctx.Divider().SetText("card")
	ctx.Box(
		ctx.H6("Card:"),
		card)

	ctx.Divider().SetText("checkbox")
	ctx.Box(
		ctx.H6("checkbox:"),
		ctx.Checkbox(true, "check me").OnChange(func(self *nicegoi.Checkbox, checked bool) {
			self.Ctx().MsgInfo(fmt.Sprintf("checkbox change:%v", checked))
		}))
	ctx.Divider().SetText("radio")
	ctx.Box(
		ctx.H6("radio:"),
		radio)

	ctx.Divider().SetText("select")
	ctx.Box(
		ctx.H6("select:"),
		ctx.Select("value 1", "value 1", "value 2", "value 3").OnChange(func(self *nicegoi.Select, selected string) {
			self.Ctx().MsgWarn("radio will change!")
			radio.Select(selected)
		}),
		ctx.H6("labeled:"),
		ctx.Select("value 1", "value 1", "value 2", "value 3").SetLabel("Opt:"),
		ctx.H6("placeholder:"),
		ctx.Select("", "value 1", "value 2", "value 3").SetPlaceholder("please select"),
	)
	ctx.Box(
		st,
		ctx.Divider().Vertical(),
		ctx.Label("clearable:"),
		ctx.Switch(false).OnChange(func(self *nicegoi.Switch, on bool) {
			st.SetClearable(on)
		}),
		ctx.Label("filterable:"),
		ctx.Switch(false).OnChange(func(self *nicegoi.Switch, on bool) {
			st.SetFilterable(on)
		}),
		ctx.Label("loading:"),
		ctx.Switch(false).OnChange(func(self *nicegoi.Switch, on bool) {
			st.SetLoading(on)
		}),
	)

	ctx.Divider().SetText("switch")
	ctx.Box(
		ctx.H6("switch:"),
		ctx.Switch(false).OnChange(func(self *nicegoi.Switch, on bool) {
			self.Ctx().MsgWarn(fmt.Sprintf("switch on:%v", on))
		}))

	ctx.Divider().SetText("route to new page")
	ctx.Box(
		gotoInput,
		ctx.Link("goto /hello with query param").SetOnClick(func(self *nicegoi.Link) {
			var name = gotoInput.GetValue()
			if name == "" {
				self.Ctx().RouteTo("hello", nil)
			} else {
				self.Ctx().RouteTo("hello", map[string]any{"name": gotoInput.GetValue()})
			}
		}),
		ctx.Divider().Vertical(),
		ctx.Link("goto no define page").SetOnClick(func(self *nicegoi.Link) {
			self.Ctx().RouteTo("something404", nil)
		}),
		ctx.Divider().Vertical(),
		ctx.Button("reload current page", func(self *nicegoi.Button) {
			self.Ctx().ReloadPage()
		}).SetVariant(option.Text),
	)

	ctx.Divider().SetText("datetime")
	ctx.Box(
		ctx.H6("datetime:"),
		ctx.DateTime(time.Now()).OnChange(func(self *nicegoi.DateTime, datetime time.Time, err error) {
			if err != nil {
				self.Ctx().MsgError("datetime err:" + err.Error())
			} else {
				self.Ctx().MsgSuccess("datetime changed:" + datetime.Format("2006-01-02 15:04:05"))
			}
		}))

	ctx.Divider().SetText("list")
	ctx.Box(
		ctx.H6("list:"),
		l)

	ctx.Divider().SetText("menu")
	ctx.Box(ctx.H6("menu:"), ctx.Menu(m).SetOnChange(func(self *nicegoi.Menu, root *menu.Option, item *menu.ItemOption) {
		self.Ctx().MsgWarn("menu selected:" + item.Value)
		v4, err := uuid.NewGen().NewV4()
		if err == nil {
			item.Label = v4.String()[0:5]
		}
	}))

	ctx.Divider().SetText("tab")
	ctx.Box(
		ctx.H6("Tabs:"),
		ctx.Radio("left", "top", "left", "right", "bottom").OnChange(func(self *nicegoi.Radio, selected string) {
			tab.SetPlace(option.Placement(selected))
		}),
		tab.Add("button", ctx.Button("button", nil)).
			AddWithIcon("link", icons.Link, ctx.Link("Remove Cur Tab").SetOnClick(func(self *nicegoi.Link) {
				tab.Remove("link")
			})).
			Add("input", ctx.Input(nil)).
			SetOnChange(func(key string, widget nicegoi.IWidget) {
				tab.Ctx().MsgSuccess("tag change " + key)
				if key == "button" {
					b := widget.(*nicegoi.Button)
					v4, err := uuid.NewGen().NewV4()
					if err == nil {
						b.SetText(v4.String()[0:5])
					}
				}
			}),
	)

	ctx.Divider().SetText("table")
	ctx.Box(
		ctx.H6("table:"),
		ctx.Table([]interface{}{menu.NewItem("label1", "value1"), menu.NewItem("label2", "value2")}),
	)

	ctx.Divider().SetText("loading")
	ctx.Box(
		ctx.H6("Loading:"),
		ctx.Button("start/stop", func(self *nicegoi.Button) {
			if loading.GetState() {
				loading.Stop()
			} else {
				loading.Start()
			}
		}),
		loading.AddItems(ctx.Card("this is a long content for loading show").SetWidth(300)),
		ctx.Divider().Vertical(),
		ctx.H6("full screen:"),
		ctx.Button("full screen loading", func(self *nicegoi.Button) {
			fullLoading.Start()
			go func() {
				time.Sleep(time.Second * 2)
				fullLoading.Stop()
			}()
		}),
		fullLoading,
	)

	ctx.Divider().SetText("progress")

	ctx.H6("progress:")
	ctx.Box(
		p0,
		ctx.Button("add", func(self *nicegoi.Button) {
			p0.Update(p0.Current() + 2)
			p.Update(p.Current() + 2)
		}),
		p,
		ctx.Divider().Vertical(),
		ctx.H6("error:"),
		ctx.Progress(45).CircleStyle().MarkState(option.ProgressError),
		ctx.Divider().Vertical(),
		ctx.H6("warning:"),
		ctx.Progress(45).CircleStyle().MarkState(option.ProgressWarning),
		ctx.Divider().Vertical(),
		ctx.H6("success:"),
		ctx.Progress(100).CircleStyle().MarkState(option.ProgressSuccess),
	)
	ctx.Divider().SetText("description")
	ctx.Description(2, map[string]string{"Name": "NceGoi", "Tel": "12288884444", "Area": "China SHangHai", "Address": "XUJIAHU DASHIJIE"})

	ctx.Divider().SetText("badge")
	ctx.Box(
		ctx.H6("Badge:"),
		badge.SetChild(ctx.Button("add", func(self *nicegoi.Button) {
			badge.Incr(1)
		})),
	)
	ctx.Divider().SetText("timeline")
	ctx.Timeline().Add(timeline.Primary("2024-10-01", "something common")).
		Add(timeline.Success("2024-10-21", "something success")).
		Add(timeline.Warning("2024-11-01", "something warning").WithDetail("a detail description")).
		Add(timeline.Error("2024-10-21", "something error"))

	ctx.Divider().SetText("drawer")

	ctx.Box(
		ctx.Button("right", func(self *nicegoi.Button) {
			dr.Open()
		}),
		ctx.Button("left", func(self *nicegoi.Button) {
			dl.Open()
		}),
		ctx.Button("top", func(self *nicegoi.Button) {
			dt.Open()
		}),
		ctx.Button("bottom", func(self *nicegoi.Button) {
			db.Open()
		}),
	)

	ctx.Divider().SetText("row")
	ctx.Label("default:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 4"), ctx.Input(nil).PlaceHolder("span 4"), ctx.Input(nil).PlaceHolder("span 4"))
	ctx.Label("with gutter:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 4"), ctx.Input(nil).PlaceHolder("span 4"), ctx.Input(nil).PlaceHolder("span 4")).
		SetGutter(10, 0)
	ctx.Label("with span:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 2"), ctx.Input(nil).PlaceHolder("span 5"), ctx.Input(nil).PlaceHolder("span 5")).
		SetSpan(2, 5, 5)
	ctx.Label("with offset:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 4 offset 0"), ctx.Input(nil).PlaceHolder("span 4 offset 4")).SetSpan(4, 4).
		SetOffset(1, 4)

	ctx.H4("align:")
	ctx.Label("start:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowStart)
	ctx.Label("center:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowCenter)
	ctx.Label("end:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowEnd)
	ctx.Label("space around:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowSpaceAround)
	ctx.Label("space between:")
	ctx.Row(ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3"), ctx.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowSpaceBetween)
}

func main() {
	nicegoi.Run(new(Hello), new(Main))

}
