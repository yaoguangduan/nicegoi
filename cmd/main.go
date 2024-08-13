package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/yaoguangduan/nicegoi/goi"
	"github.com/yaoguangduan/nicegoi/internal/ui"
	"github.com/yaoguangduan/nicegoi/internal/ui/icons"
	"github.com/yaoguangduan/nicegoi/internal/ui/menu"
	"github.com/yaoguangduan/nicegoi/internal/ui/place"
	"time"
)

func main() {

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	card := goi.Card("i am card content")

	card.AddActions(goi.Link("card Action"))
	card.AddFooters(goi.Link("add"), goi.Link("Setting"))
	card.SetTitle("card title")

	l := goi.List()
	item1 := ui.NewListItem("list item 1").AddAction(goi.Link("operate 1").SetOnClick(func(self *ui.Link) {
		goi.MsgInfo("item 1 operated")
	}))
	item2 := ui.NewListItem("list item 2").AddAction(goi.Link("operate 2").SetOnClick(func(self *ui.Link) {
		goi.MsgInfo("item 2 operated")
	}))
	l.AddItems(item1, item2)
	radio := goi.Radio("value 1", "value 1", "value 2", "value 3").OnChange(func(self *ui.Radio, selected string) {
		goi.MsgSuccess(fmt.Sprintf("you has selected %s", selected))
	})

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

	tab := ui.NewTab().SetPlace(place.Left)

	loading := ui.NewLoading("loading...")
	fullLoading := ui.NewLoading("loadingAll...").FullScreen()
	p0 := goi.Progress(20)
	p := goi.Progress(20).CircleStyle()

	badge := goi.Badge(29)
	goi.Box(
		goi.H6("Button:"),
		goi.Button("button", func(self *ui.GoiButton) {
			goi.MsgSuccess("been clicked")
		}),
		goi.Button("Icon", nil).SetIcon(icons.Add),
	)
	goi.Box(
		goi.H6("Input:"),
		goi.Input(func(ctx *ui.Input, val string) {
			goi.MsgWarn(fmt.Sprintf("you input:%s", val))
		}),
		goi.Input(nil).SetPrepend("https://").SetAppend(".com"),
		goi.Input(nil).SetIcon(icons.ArrowRight).PlaceHolder("example.com"),
		goi.Input(nil).EnablePassword(),
	).WithSeparator()
	goi.Box(
		goi.H6("Link:"),
		goi.Link("点击").SetOnClick(func(self *ui.Link) {
			goi.NotifySuccess("I am Title", "you clicked link!")
		}),
		goi.Link("跳转链接").SetPrefixIcon(icons.Link).SetHref("https://google.com"),
	).WithSeparator()
	goi.Box(
		goi.H5("Card:"),
		card)
	goi.Box(
		goi.H5("checkbox:"),
		goi.Checkbox(true, "check me").OnChange(func(self *ui.GoiCheckbox, checked bool) {
			goi.MsgInfo(fmt.Sprintf("checkbox change:%v", checked))
		}))

	goi.Box(
		goi.H5("radio:"),
		radio)

	goi.Box(
		goi.H5("select:"),
		goi.Select("value 1", "value 1", "value 2", "value 3").OnChange(func(self *ui.Select, selected string) {
			goi.MsgWarn("radio will change!")
			radio.Select(selected)
		}))

	goi.Box(
		goi.H5("switch:"),
		goi.Switch(false).OnChange(func(self *ui.Switch, on bool) {
			goi.MsgWarn(fmt.Sprintf("switch on:%v", on))
		}))

	goi.Box(
		goi.H5("datetime:"),
		goi.DateTime(time.Now()).OnChange(func(self *ui.DateTime, datetime time.Time, err error) {
			if err != nil {
				goi.MsgError("datetime err:" + err.Error())
			} else {
				goi.MsgSuccess("datetime changed:" + datetime.Format("2006-01-02 15:04:05"))
			}
		}))

	goi.Box(
		goi.H5("list:"),
		l)
	goi.Box(goi.H6("menu:"), ui.NewMenu(m))
	goi.Box(
		goi.H5("Tabs:"),
		goi.Radio("left", "top", "left", "right", "bottom").OnChange(func(self *ui.Radio, selected string) {
			tab.SetPlace(place.Placement(selected))
		}),
		tab.Add("button", goi.Button("button", nil)).
			AddWithIcon("link", icons.Link, goi.Link("Remove Cur Tab").SetOnClick(func(self *ui.Link) {
				tab.Remove("link")
			})).
			Add("input", goi.Input(nil)).
			SetOnChange(func(key string, widget ui.IWidget) {
				goi.MsgSuccess("tag change " + key)
				if key == "button" {
					b := widget.(*ui.GoiButton)
					v4, err := uuid.NewGen().NewV4()
					if err == nil {
						b.SetText(v4.String()[0:5])
					}
				}
			}),
	)
	goi.Box(
		goi.H5("table:"),
		ui.NewTable([]interface{}{menu.NewItem("label1", "value1"), menu.NewItem("label2", "value2")}),
	)
	goi.Box(
		goi.H5("Loading:"),
		goi.Button("start/stop", func(self *ui.GoiButton) {
			if loading.GetState() {
				loading.Stop()
			} else {
				loading.Start()
			}
		}),
		loading.AddItems(goi.Card("this is a long content for loading show")),
		goi.Button("full screen loading", func(self *ui.GoiButton) {
			fullLoading.Start()
			go func() {
				time.Sleep(time.Second * 2)
				fullLoading.Stop()
			}()
		}),
		fullLoading,
	)
	goi.Box(
		goi.H5("progress:"),
		p0,
		goi.Button("add", func(self *ui.GoiButton) {
			p0.Update(p0.Current() - 2)
			p.Update(p.Current() + 2)
		}),
		p,
		goi.Progress(45).CircleStyle().MarkState(ui.ProgressError),
		goi.Progress(45).CircleStyle().MarkState(ui.ProgressWarning),
	)
	goi.Box(
		goi.H5("description:"),
		goi.Description(2, map[string]string{"Name": "NceGoi", "Tel": "12288884444", "Area": "China SHangHai", "Address": "XUJIAHU DASHIJIE"}),
	)
	goi.Box(
		goi.H5("Badge:"),
		badge.SetChild(goi.Button("add", func(self *ui.GoiButton) {
			badge.Incr(1)
		})),
	)

	goi.Run()

}
