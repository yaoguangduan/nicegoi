package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/yaoguangduan/nicegoi/goi"
	"github.com/yaoguangduan/nicegoi/nice"
	"github.com/yaoguangduan/nicegoi/nice/icons"
	"github.com/yaoguangduan/nicegoi/nice/option"
	"github.com/yaoguangduan/nicegoi/nice/option/menu"
	"github.com/yaoguangduan/nicegoi/nice/option/timeline"
	"time"
)

func hello(goiCtx nice.GoiContext) []nice.IWidget {
	label := goi.Label(fmt.Sprintf("hello %v", goiCtx.Query().GetOr("name", "world!")))
	return []nice.IWidget{
		label,
	}
}

func home(goiCtx nice.GoiContext) []nice.IWidget {

	card := goi.Card("i am card content")

	card.AddActions(goi.Link("card Action"))
	card.AddFooters(goi.Link("add"), goi.Link("Setting"))
	card.SetTitle("card title")
	tab := nice.NewTab().SetPlace(option.Left)
	loading := nice.NewLoading("loading...")
	fullLoading := nice.NewLoading("loadingAll...").FullScreen()
	disable := goi.Button("Disable", nil)
	disable.SetDisable(true)

	tags := make([]nice.IWidget, 0)
	variants := []option.TagVariant{option.TagVarDark, option.TagVarOutline, option.TagVarLight, option.TagVarLightOutline}
	themes := []option.Theme{option.Default, option.Primary, option.Danger, option.Success, option.Warning}
	for _, variant := range variants {
		for _, theme := range themes {
			tags = append(tags, goi.Tag(fmt.Sprintf("%s/%s", variant, theme)).SetVariant(variant).SetTheme(theme))
		}
	}

	radio := goi.Radio("value 1", "value 1", "value 2", "value 3").OnChange(func(self *nice.Radio, selected string) {
		self.Page().MsgSuccess(fmt.Sprintf("you has selected %s", selected))
	})

	st := goi.Select("banana", "banana", "apple", "orange")

	p0 := goi.Progress(20)

	l := goi.List()
	item1 := nice.NewListItem("list item 1").AddAction(goi.Link("operate 1").SetOnClick(func(self *nice.Link) {
		self.Page().MsgInfo("item 1 operated")
	}))
	item2 := nice.NewListItem("list item 2").AddAction(goi.Link("operate 2").SetOnClick(func(self *nice.Link) {
		self.Page().MsgInfo("item 2 operated")
	}))
	l.AddItems(item1, item2)

	m := menu.New().AddItems(
		menu.NewItem("MenuOption 1", "m1").SetIcon(icons.Home).AddItems(menu.NewItem("MenuOption 1-1", "m11"), menu.NewItem("MenuOption 1-2", "m12")),
		menu.NewItem("MenuOption 2", "m2").SetIcon(icons.Edit),
	)

	p := goi.Progress(20).CircleStyle()

	dr := goi.Drawer("Header").AddWidgets(goi.Label("drawer content"))
	dl := goi.Drawer("Header").AddWidgets(goi.Label("drawer content")).SetPlace(option.Left)
	dt := goi.Drawer("Header").AddWidgets(goi.Label("drawer content")).SetPlace(option.Top)
	db := goi.Drawer("Header").AddWidgets(goi.Label("drawer content")).SetPlace(option.Bottom)

	badge := goi.Badge(29)

	gotoInput := goi.Input(nil).PlaceHolder("input your name")
	return []nice.IWidget{

		goi.H1("Hello NiceGOI!"),

		goi.Divider().SetText("button"),
		goi.Box(
			goi.H6("Button:"),
			goi.Button("button", func(self *nice.Button) {
				self.Page().MsgWarn("button clicked!")
			}),
			goi.Divider().Vertical(),
			goi.H6("with Icon:"),
			goi.Button("Icon", nil).SetIcon(icons.Add),
			goi.Divider().Vertical(),
			goi.Button("Success", nil).SetTheme(option.Success),
			goi.Divider().Vertical(),
			goi.Button("Danger", nil).SetTheme(option.Danger),
			goi.Divider().Vertical(),
			goi.Button("Warning", nil).SetTheme(option.Warning),
		),
		goi.Box(
			goi.H6("Variant:"),
			goi.Button("Outline", nil).SetVariant(option.Outline),
			goi.Button("Dashed", nil).SetVariant(option.Dashed),
			goi.Button("Text", nil).SetVariant(option.Text),
		),
		goi.Box(
			goi.H6("Shape:"),
			goi.Button("Rectangle", nil).SetShape(option.Rectangle),
			goi.Button("", nil).SetIcon(icons.Home).SetShape(option.Square),
			goi.Button("Round", nil).SetShape(option.Round),
			goi.Button("", nil).SetIcon(icons.Edit).SetShape(option.Circle),
		),
		goi.Box(
			goi.H6("State/Size:"),
			disable,
			goi.Button("Loading", nil).Loading(true),
			goi.Button("Small", nil).SetSize(option.Small),
			goi.Button("Large", nil).SetSize(option.Large),
		),

		goi.Divider().SetText("input:"),
		goi.Box(
			goi.H6("Input:"),
			goi.Input(func(ctx *nice.Input, val string) {
				ctx.Page().MsgWarn(fmt.Sprintf("you input:%s", val))
			}),
			goi.Divider().Vertical(),
			goi.H6("with prefix/suffix:"),
			goi.Input(nil).SetPrepend("https://").SetAppend(".com"),
			goi.Divider().Vertical(),
			goi.H6("with icon:"),
			goi.Input(nil).SetIcon(icons.ArrowRight).PlaceHolder("example.com"),
			goi.Divider().Vertical(),
			goi.H6("password:"),
			goi.Input(nil).EnablePassword(),
		),

		goi.Divider().SetText("link"),
		goi.Box(
			goi.H6("link:"),
			goi.Link("点击").SetOnClick(func(self *nice.Link) {
				self.Page().NotifySuccess("I am Title", "you clicked link!")
			}),
			goi.Divider().Vertical(),
			goi.H6("with icon:"),
			goi.Link("跳转链接").SetPrefixIcon(icons.Link).SetHref("https://google.com"),
			goi.Divider().Vertical(),
			goi.Link("Success").SetTheme(option.Success),
			goi.Divider().Vertical(),
			goi.Link("Danger").SetTheme(option.Danger),
			goi.Divider().Vertical(),
			goi.Link("Warning").SetTheme(option.Warning),
			goi.Divider().Vertical(),
			goi.Link("Small").SetSize(option.Small),
			goi.Divider().Vertical(),
			goi.Link("Large").SetSize(option.Large),
			goi.Divider().Vertical(),
			goi.Link("underline").Underlined(),
		),
		goi.Divider().SetText("dropdown"),
		goi.Box(
			goi.Dropdown("more...", "option 1", "option 2", "option 3").OnClick(func(self *nice.Dropdown, value string) {
				self.Page().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
			}),
			goi.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Dashed).OnClick(func(self *nice.Dropdown, value string) {
				self.Page().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
			}),
			goi.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Base).OnClick(func(self *nice.Dropdown, value string) {
				self.Page().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
			}),
			goi.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Outline).OnClick(func(self *nice.Dropdown, value string) {
				self.Page().MsgSuccess(fmt.Sprintf("you clicked:%s", value))
			}),
		),
		goi.Divider().SetText("tag"),
		goi.Box(tags...),
		goi.Box(
			goi.H6("Size:"),
			goi.Tag("medium").SetSize(option.Medium),
			goi.Tag("small").SetSize(option.Small),
			goi.Tag("large").SetSize(option.Large),
		),
		goi.Box(
			goi.H6("Shape:"),
			goi.Tag("round").SetShape(option.Round),
			goi.Tag("square").SetShape(option.Square),
			goi.Tag("mark").SetShape(option.Mark),
		),

		goi.Divider().SetText("tag input"),
		goi.Row(
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifySuccess("tag input", fmt.Sprintf("%v", values))
			}).SetTheme(option.Success).SetPlaceHolder("success theme"),
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifyInfo("tag input", fmt.Sprintf("%v", values))
			}).SetTheme(option.Primary).SetPlaceHolder("primary theme"),
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifyInfo("tag input", fmt.Sprintf("%v", values))
			}).SetTheme(option.Default).SetPlaceHolder("default theme"),
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifyError("tag input", fmt.Sprintf("%v", values))
			}).SetTheme(option.Danger).SetPlaceHolder("danger theme"),
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifyWarn("tag input", fmt.Sprintf("%v", values))
			}).SetTheme(option.Warning).SetPlaceHolder("warning theme"),
		).SetSpan(2, 2, 2, 2, 2).SetGutter(10, 0).Justify(option.RowSpaceAround),
		goi.Box(),
		goi.Row(
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifySuccess("tag input", fmt.Sprintf("%v", values))
			}).SetLabel("With Label:").SetPlaceHolder("labeled input"),
			goi.TagInput(func(self *nice.TagInput, values []string) {
				self.Page().NotifyInfo("tag input", fmt.Sprintf("%v", values))
			}).SetMax(3).SetPlaceHolder("max 3 tags"),
		).SetSpan(4, 4).SetGutter(10, 0).Justify(option.RowStart),

		goi.Divider().SetText("card"),
		goi.Box(
			goi.H6("Card:"),
			card),

		goi.Divider().SetText("checkbox"),
		goi.Box(
			goi.H6("checkbox:"),
			goi.Checkbox(true, "check me").OnChange(func(self *nice.Checkbox, checked bool) {
				self.Page().MsgInfo(fmt.Sprintf("checkbox change:%v", checked))
			})),

		goi.Divider().SetText("radio"),
		goi.Box(
			goi.H6("radio:"),
			radio),

		goi.Divider().SetText("select"),
		goi.Box(
			goi.H6("select:"),
			goi.Select("value 1", "value 1", "value 2", "value 3").OnChange(func(self *nice.Select, selected string) {
				self.Page().MsgWarn("radio will change!")
				radio.Select(selected)
			}),
			goi.H6("labeled:"),
			goi.Select("value 1", "value 1", "value 2", "value 3").SetLabel("Opt:"),
			goi.H6("placeholder:"),
			goi.Select("", "value 1", "value 2", "value 3").SetPlaceholder("please select"),
		),
		goi.Box(
			st,
			goi.Divider().Vertical(),
			goi.Label("clearable:"),
			goi.Switch(false).OnChange(func(self *nice.Switch, on bool) {
				st.SetClearable(on)
			}),
			goi.Label("filterable:"),
			goi.Switch(false).OnChange(func(self *nice.Switch, on bool) {
				st.SetFilterable(on)
			}),
			goi.Label("loading:"),
			goi.Switch(false).OnChange(func(self *nice.Switch, on bool) {
				st.SetLoading(on)
			}),
		),

		goi.Divider().SetText("switch"),
		goi.Box(
			goi.H6("switch:"),
			goi.Switch(false).OnChange(func(self *nice.Switch, on bool) {
				self.Page().MsgWarn(fmt.Sprintf("switch on:%v", on))
			})),

		goi.Divider().SetText("route to new page"),
		goi.Box(
			gotoInput,
			goi.Link("goto /hello with query param").SetOnClick(func(self *nice.Link) {
				var name = gotoInput.GetValue()
				if name == "" {
					self.Page().RouteTo("hello", nil)
				} else {
					self.Page().RouteTo("hello", map[string]any{"name": gotoInput.GetValue()})
				}
			}),
		),

		goi.Divider().SetText("datetime"),
		goi.Box(
			goi.H6("datetime:"),
			goi.DateTime(time.Now()).OnChange(func(self *nice.DateTime, datetime time.Time, err error) {
				if err != nil {
					self.Page().MsgError("datetime err:" + err.Error())
				} else {
					self.Page().MsgSuccess("datetime changed:" + datetime.Format("2006-01-02 15:04:05"))
				}
			})),

		goi.Divider().SetText("list"),
		goi.Box(
			goi.H6("list:"),
			l),

		goi.Divider().SetText("menu"),
		goi.Box(goi.H6("menu:"), nice.NewMenu(m).SetOnChange(func(self *nice.Menu, root *menu.Option, item *menu.ItemOption) {
			self.Page().MsgWarn("menu selected:" + item.Value)
			v4, err := uuid.NewGen().NewV4()
			if err == nil {
				item.Label = v4.String()[0:5]
			}
		})),

		goi.Divider().SetText("tab"),
		goi.Box(
			goi.H6("Tabs:"),
			goi.Radio("left", "top", "left", "right", "bottom").OnChange(func(self *nice.Radio, selected string) {
				tab.SetPlace(option.Placement(selected))
			}),
			tab.Add("button", goi.Button("button", nil)).
				AddWithIcon("link", icons.Link, goi.Link("Remove Cur Tab").SetOnClick(func(self *nice.Link) {
					tab.Remove("link")
				})).
				Add("input", goi.Input(nil)).
				SetOnChange(func(key string, widget nice.IWidget) {
					tab.Page().MsgSuccess("tag change " + key)
					if key == "button" {
						b := widget.(*nice.Button)
						v4, err := uuid.NewGen().NewV4()
						if err == nil {
							b.SetText(v4.String()[0:5])
						}
					}
				}),
		),

		goi.Divider().SetText("table"),
		goi.Box(
			goi.H6("table:"),
			nice.NewTable([]interface{}{menu.NewItem("label1", "value1"), menu.NewItem("label2", "value2")}),
		),

		goi.Divider().SetText("loading"),
		goi.Box(
			goi.H6("Loading:"),
			goi.Button("start/stop", func(self *nice.Button) {
				if loading.GetState() {
					loading.Stop()
				} else {
					loading.Start()
				}
			}),
			loading.AddItems(goi.Card("this is a long content for loading show").SetWidth(300)),
			goi.Divider().Vertical(),
			goi.H6("full screen:"),
			goi.Button("full screen loading", func(self *nice.Button) {
				fullLoading.Start()
				go func() {
					time.Sleep(time.Second * 2)
					fullLoading.Stop()
				}()
			}),
			fullLoading,
		),

		goi.Divider().SetText("progress"),

		goi.H6("progress:"),
		goi.Box(
			goi.Button("add", func(self *nice.Button) {
				p0.Update(p0.Current() + 2)
				p.Update(p.Current() + 2)
			}),
			p,
			goi.Divider().Vertical(),
			goi.H6("error:"),
			goi.Progress(45).CircleStyle().MarkState(nice.ProgressError),
			goi.Divider().Vertical(),
			goi.H6("warning:"),
			goi.Progress(45).CircleStyle().MarkState(nice.ProgressWarning),
			goi.Divider().Vertical(),
			goi.H6("success:"),
			goi.Progress(100).CircleStyle().MarkState(nice.ProgressSuccess),
		),
		goi.Divider().SetText("description"),
		goi.Description(2, map[string]string{"Name": "NceGoi", "Tel": "12288884444", "Area": "China SHangHai", "Address": "XUJIAHU DASHIJIE"}),

		goi.Divider().SetText("badge"),
		goi.Box(
			goi.H6("Badge:"),
			badge.SetChild(goi.Button("add", func(self *nice.Button) {
				badge.Incr(1)
			})),
		),

		goi.Divider().SetText("timeline"),
		goi.Timeline().Add(timeline.Primary("2024-10-01", "something common")).
			Add(timeline.Success("2024-10-21", "something success")).
			Add(timeline.Warning("2024-11-01", "something warning").WithDetail("a detail description")).
			Add(timeline.Error("2024-10-21", "something error")),

		goi.Divider().SetText("drawer"),
		dr, dl, dt, db,
		goi.Box(
			goi.Button("right", func(self *nice.Button) {
				dr.Open()
			}),
			goi.Button("left", func(self *nice.Button) {
				dl.Open()
			}),
			goi.Button("top", func(self *nice.Button) {
				dt.Open()
			}),
			goi.Button("bottom", func(self *nice.Button) {
				db.Open()
			}),
		),

		goi.Divider().SetText("row"),
		goi.Label("default:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4")),
		goi.Label("with gutter:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4")).
			SetGutter(10, 0),
		goi.Label("with span:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 2"), goi.Input(nil).PlaceHolder("span 5"), goi.Input(nil).PlaceHolder("span 5")).
			SetSpan(2, 5, 5),
		goi.Label("with offset:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 4 offset 0"), goi.Input(nil).PlaceHolder("span 4 offset 4")).SetSpan(4, 4).
			SetOffset(1, 4),

		goi.H4("align:"),
		goi.Label("start:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowStart),
		goi.Label("center:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowCenter),
		goi.Label("end:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowEnd),
		goi.Label("space around:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowSpaceAround),
		goi.Label("space between:"),
		goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowSpaceBetween),
	}
}

func main() {

	nice.Page("hello", hello)
	nice.Page("", home)
	goi.Run()

}
