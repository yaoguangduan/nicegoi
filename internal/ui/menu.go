package ui

import (
	"nicegoi/internal/msgs"
	"nicegoi/internal/ui/menu"
	"nicegoi/internal/ws"
)

type Menu struct {
	*valueWidget
}

func NewMenu(m *menu.Option) *Menu {
	w := &valueWidget{opt: NewElement("menu").Set("root", m)}
	if m.Value != "" {
		w.set(m.Value)
	}
	w.opt.Set("collapse", m.Collapse)
	mu := &Menu{w}
	ws.RegMsgHandle(w.opt.Eid(), func(message *msgs.Message) {
		selected := message.Data.(string)
		w.opt.Modify("value", selected)
		root := mu.opt.Get("root").(*menu.Option)
		if m.OnChange != nil {
			m.OnChange(root, findItem(root.MenuItems, selected))
			w.opt.(*Element).OnModify("root")
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
