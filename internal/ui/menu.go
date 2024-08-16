package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/msgs"
	"github.com/yaoguangduan/nicegoi/internal/option/menu"
)

type Menu struct {
	*valuedWidget
}

func NewMenu(m *menu.Option) *Menu {
	w := &valuedWidget{e: createElement("menu").Set("root", m)}
	if m.Value != "" {
		w.set(m.Value)
	}
	w.e.Set("collapse", m.Collapse)
	mu := &Menu{w}

	return mu
}
func (m *Menu) SetOnChange(onChange func(self *Menu, menu *menu.Option, item *menu.ItemOption)) *Menu {
	m.addMsgHandler(func(message *msgs.Message) {
		selected := message.Data.(string)
		m.e.Modify("value", selected)
		root := m.e.Get("root").(*menu.Option)
		if onChange != nil {
			onChange(m, root, findItem(root.MenuItems, selected))
			m.e.OnModify("root")
		}
	})
	return m
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
