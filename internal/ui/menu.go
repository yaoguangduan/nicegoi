package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/option/menu"
	"github.com/yaoguangduan/nicegoi/internal/ws"
)

type Menu struct {
	*valuedWidget
}

func NewMenu(m *menu.Option) *Menu {
	w := &valuedWidget{e: NewElement("menu").Set("root", m)}
	if m.Value != "" {
		w.set(m.Value)
	}
	w.e.Set("collapse", m.Collapse)
	mu := &Menu{w}
	ws.RegMsgHandle(w.e.Eid(), func(message *msgs.Message) {
		selected := message.Data.(string)
		w.e.Modify("value", selected)
		root := mu.e.Get("root").(*menu.Option)
		if m.OnChange != nil {
			m.OnChange(root, findItem(root.MenuItems, selected))
			w.e.(*Element).OnModify("root")
		}
	})
	return mu
}

func findItem(root []*menu.ItemOption, selected string) *menu.ItemOption {
	for _, item := range root {
		if item.Value == selected {
			return item
		}
		f := findItem(item.Children, selected)
		if f != nil {
			return f
		}
	}
	return nil
}
