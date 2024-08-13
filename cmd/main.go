package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/yaoguangduan/nicegoi/goi"
	"github.com/yaoguangduan/nicegoi/internal/option"
	"github.com/yaoguangduan/nicegoi/internal/option/menu"
	"github.com/yaoguangduan/nicegoi/internal/option/timeline"
	"github.com/yaoguangduan/nicegoi/internal/ui"
	"github.com/yaoguangduan/nicegoi/internal/ui/icons"
	"time"
)

func main() {

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	card := goi.Card("i am card content")

	card.AddActions(goi.Link("card Action"))
	card.AddFooters(goi.Link("add"), goi.Link("Setting"))
	card.SetTitle("card title")

	tab := ui.NewTab().SetPlace(option.Left)

	loading := ui.NewLoading("loading...")
	fullLoading := ui.NewLoading("loadingAll...").FullScreen()

	goi.Divider().SetText("button")
	goi.Box(
		goi.H6("Button:"),
		goi.Button("button", func(self *ui.Button) {
			goi.MsgSuccess("been clicked")
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
	)

	goi.Divider().SetText("input:")
	goi.Box(
		goi.H6("Input:"),
		goi.Input(func(ctx *ui.Input, val string) {
			goi.MsgWarn(fmt.Sprintf("you input:%s", val))
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
	)

	goi.Divider().SetText("link")
	goi.Box(
		goi.H6("link:"),
		goi.Link("点击").SetOnClick(func(self *ui.Link) {
			goi.NotifySuccess("I am Title", "you clicked link!")
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
	)
	goi.Divider().SetText("dropdown")
	goi.Box(
		goi.Dropdown("more...", "option 1", "option 2", "option 3").OnClick(func(self *ui.Dropdown, value string) {
			goi.MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
		goi.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Dashed).OnClick(func(self *ui.Dropdown, value string) {
			goi.MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
		goi.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Base).OnClick(func(self *ui.Dropdown, value string) {
			goi.MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
		goi.Dropdown("more...", "option 1", "option 2", "option 3").SetVariant(option.Outline).OnClick(func(self *ui.Dropdown, value string) {
			goi.MsgSuccess(fmt.Sprintf("you clicked:%s", value))
		}),
	)
	goi.Divider().SetText("tag")
	tags := make([]ui.IWidget, 0)
	variants := []option.Variant{option.Outline, option.Text, option.Base, option.Dashed}
	themes := []option.Theme{option.Default, option.Primary, option.Danger, option.Success, option.Warning}
	for _, variant := range variants {
		for _, theme := range themes {
			tags = append(tags, goi.Tag(fmt.Sprintf("%s/%s", variant, theme)).SetVariant(variant).SetTheme(theme))
		}
	}
	goi.Box(tags...)

	goi.Divider().SetText("tag input")
	goi.Row(
		goi.TagInput(func(self *ui.TagInput, values []string) {
			goi.NotifySuccess("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Success).PlaceHolder("success theme"),
		goi.TagInput(func(self *ui.TagInput, values []string) {
			goi.NotifyInfo("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Primary).PlaceHolder("primary theme"),
		goi.TagInput(func(self *ui.TagInput, values []string) {
			goi.NotifyInfo("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Default).PlaceHolder("default theme"),
		goi.TagInput(func(self *ui.TagInput, values []string) {
			goi.NotifyError("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Danger).PlaceHolder("danger theme"),
		goi.TagInput(func(self *ui.TagInput, values []string) {
			goi.NotifyWarn("tag input", fmt.Sprintf("%v", values))
		}).SetTheme(option.Warning).PlaceHolder("warning theme"),
	).SetSpan(2, 2, 2, 2, 2).SetGutter(10, 0).Justify(option.RowSpaceAround)

	goi.Divider().SetText("card")
	goi.Box(
		goi.H6("Card:"),
		card)

	goi.Divider().SetText("checkbox")
	goi.Box(
		goi.H6("checkbox:"),
		goi.Checkbox(true, "check me").OnChange(func(self *ui.Checkbox, checked bool) {
			goi.MsgInfo(fmt.Sprintf("checkbox change:%v", checked))
		}))

	goi.Divider().SetText("radio")
	radio := goi.Radio("value 1", "value 1", "value 2", "value 3").OnChange(func(self *ui.Radio, selected string) {
		goi.MsgSuccess(fmt.Sprintf("you has selected %s", selected))
	})
	goi.Box(
		goi.H6("radio:"),
		radio)

	goi.Divider().SetText("select")
	goi.Box(
		goi.H6("select:"),
		goi.Select("value 1", "value 1", "value 2", "value 3").OnChange(func(self *ui.Select, selected string) {
			goi.MsgWarn("radio will change!")
			radio.Select(selected)
		}))

	goi.Divider().SetText("switch")
	goi.Box(
		goi.H6("switch:"),
		goi.Switch(false).OnChange(func(self *ui.Switch, on bool) {
			goi.MsgWarn(fmt.Sprintf("switch on:%v", on))
		}))

	goi.Divider().SetText("datetime")
	goi.Box(
		goi.H6("datetime:"),
		goi.DateTime(time.Now()).OnChange(func(self *ui.DateTime, datetime time.Time, err error) {
			if err != nil {
				goi.MsgError("datetime err:" + err.Error())
			} else {
				goi.MsgSuccess("datetime changed:" + datetime.Format("2006-01-02 15:04:05"))
			}
		}))

	goi.Divider().SetText("list")
	l := goi.List()
	item1 := ui.NewListItem("list item 1").AddAction(goi.Link("operate 1").SetOnClick(func(self *ui.Link) {
		goi.MsgInfo("item 1 operated")
	}))
	item2 := ui.NewListItem("list item 2").AddAction(goi.Link("operate 2").SetOnClick(func(self *ui.Link) {
		goi.MsgInfo("item 2 operated")
	}))
	l.AddItems(item1, item2)
	goi.Box(
		goi.H6("list:"),
		l)

	goi.Divider().SetText("menu")
	m := menu.New().AddItems(
		menu.NewItem("MenuOption 1", "m1").SetIcon(icons.Home).AddItems(menu.NewItem("MenuOption 1-1", "m11"), menu.NewItem("MenuOption 1-2", "m12")),
		menu.NewItem("MenuOption 2", "m2").SetIcon(icons.Edit),
	).SetOnChange(func(root *menu.Option, item *menu.ItemOption) {
		goi.MsgWarn("menu selected:" + item.Value)
		v4, err := uuid.NewGen().NewV4()
		if err == nil {
			item.Label = v4.String()[0:5]
		}
	})
	goi.Box(goi.H6("menu:"), ui.NewMenu(m))

	goi.Divider().SetText("tab")
	goi.Box(
		goi.H6("Tabs:"),
		goi.Radio("left", "top", "left", "right", "bottom").OnChange(func(self *ui.Radio, selected string) {
			tab.SetPlace(option.Placement(selected))
		}),
		tab.Add("button", goi.Button("button", nil)).
			AddWithIcon("link", icons.Link, goi.Link("Remove Cur Tab").SetOnClick(func(self *ui.Link) {
				tab.Remove("link")
			})).
			Add("input", goi.Input(nil)).
			SetOnChange(func(key string, widget ui.IWidget) {
				goi.MsgSuccess("tag change " + key)
				if key == "button" {
					b := widget.(*ui.Button)
					v4, err := uuid.NewGen().NewV4()
					if err == nil {
						b.SetText(v4.String()[0:5])
					}
				}
			}),
	)

	goi.Divider().SetText("table")
	goi.Box(
		goi.H6("table:"),
		ui.NewTable([]interface{}{menu.NewItem("label1", "value1"), menu.NewItem("label2", "value2")}),
	)

	goi.Divider().SetText("loading")
	goi.Box(
		goi.H6("Loading:"),
		goi.Button("start/stop", func(self *ui.Button) {
			if loading.GetState() {
				loading.Stop()
			} else {
				loading.Start()
			}
		}),
		loading.AddItems(goi.Card("this is a long content for loading show").SetWidth(300)),
		goi.Divider().Vertical(),
		goi.H6("full screen:"),
		goi.Button("full screen loading", func(self *ui.Button) {
			fullLoading.Start()
			go func() {
				time.Sleep(time.Second * 2)
				fullLoading.Stop()
			}()
		}),
		fullLoading,
	)

	goi.Divider().SetText("progress")

	p := goi.Progress(20).CircleStyle()
	goi.H6("progress:")
	p0 := goi.Progress(20)
	goi.Box(
		goi.Button("add", func(self *ui.Button) {
			p0.Update(p0.Current() + 2)
			p.Update(p.Current() + 2)
		}),
		p,
		goi.Divider().Vertical(),
		goi.H6("error:"),
		goi.Progress(45).CircleStyle().MarkState(ui.ProgressError),
		goi.Divider().Vertical(),
		goi.H6("warning:"),
		goi.Progress(45).CircleStyle().MarkState(ui.ProgressWarning),
		goi.Divider().Vertical(),
		goi.H6("success:"),
		goi.Progress(100).CircleStyle().MarkState(ui.ProgressSuccess),
	)
	goi.Divider().SetText("description")
	goi.Description(2, map[string]string{"Name": "NceGoi", "Tel": "12288884444", "Area": "China SHangHai", "Address": "XUJIAHU DASHIJIE"})

	goi.Divider().SetText("badge")
	badge := goi.Badge(29)
	goi.Box(
		goi.H6("Badge:"),
		badge.SetChild(goi.Button("add", func(self *ui.Button) {
			badge.Incr(1)
		})),
	)

	goi.Divider().SetText("timeline")
	goi.Timeline().Add(timeline.Primary("2024-10-01", "something common")).
		Add(timeline.Success("2024-10-21", "something success")).
		Add(timeline.Warning("2024-11-01", "something warning").WithDetail("a detail description")).
		Add(timeline.Error("2024-10-21", "something error"))

	goi.Divider().SetText("drawer")
	dr := goi.Drawer("Header").AddWidgets(goi.Label("drawer content"))
	dl := goi.Drawer("Header").AddWidgets(goi.Label("drawer content")).SetPlace(option.Left)
	dt := goi.Drawer("Header").AddWidgets(goi.Label("drawer content")).SetPlace(option.Top)
	db := goi.Drawer("Header").AddWidgets(goi.Label("drawer content")).SetPlace(option.Bottom)
	goi.Box(
		goi.Button("right", func(self *ui.Button) {
			dr.Open()
		}),
		goi.Button("left", func(self *ui.Button) {
			dl.Open()
		}),
		goi.Button("top", func(self *ui.Button) {
			dt.Open()
		}),
		goi.Button("bottom", func(self *ui.Button) {
			db.Open()
		}),
	)

	goi.Divider().SetText("row")
	goi.Label("default:")
	goi.Row(goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4"))
	goi.Label("with gutter:")
	goi.Row(goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4"), goi.Input(nil).PlaceHolder("span 4")).
		SetGutter(10, 0)
	goi.Label("with span:")
	goi.Row(goi.Input(nil).PlaceHolder("span 2"), goi.Input(nil).PlaceHolder("span 5"), goi.Input(nil).PlaceHolder("span 5")).
		SetSpan(2, 5, 5)
	goi.Label("with offset:")
	goi.Row(goi.Input(nil).PlaceHolder("span 4 offset 0"), goi.Input(nil).PlaceHolder("span 4 offset 4")).SetSpan(4, 4).
		SetOffset(1, 4)

	goi.H4("align:")
	goi.Label("left:")
	goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowLeft)
	goi.Label("center:")
	goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowCenter)
	goi.Label("right:")
	goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowRight)
	goi.Label("space around:")
	goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowSpaceAround)
	goi.Label("space between:")
	goi.Row(goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3"), goi.Input(nil).PlaceHolder("span 3")).SetSpan(3, 3, 3).Justify(option.RowSpaceBetween)
	goi.Run()

}
