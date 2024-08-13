package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/ui/icons"
	"github.com/yaoguangduan/nicegoi/internal/ui/place"
)

type Tab struct {
	*valuedWidget
	onChange func(key string, widget IWidget)
	widgets  map[string]IWidget
}

func NewTab() *Tab {
	ret := &Tab{valuedWidget: newValuedWidget("tab", ""), widgets: make(map[string]IWidget)}
	ret.e.Set("place", "top")
	ret.onValChange(func(v any) {
		k := v.(string)
		if ret.onChange != nil {
			ret.onChange(k, ret.widgets[k])
		}
	})
	return ret
}

func (t *Tab) Add(val string, widget IWidget) *Tab {
	return t.AddWithIcon(val, "", widget)
}
func (t *Tab) AddWithIcon(val string, icon icons.Icon, widget IWidget) *Tab {
	ew := newValuedWidget(t.e.Eid()+"_S", val)
	ew.e.Set("icon", icon)
	ew.Element().AddChildren(widget)
	t.e.AddChildren(ew)
	if len(t.e.Children()) == 1 {
		t.set(t.e.Children()[0].Type())
	}
	t.widgets[val] = widget
	return t
}

func (t *Tab) Remove(val string) *Tab {
	delete(t.widgets, val)
	var idx = -1
	for i, w := range t.e.Children() {
		if w.Get("value").(string) == val {
			idx = i
			break
		}
	}
	if idx < 0 {
		return t
	}
	t.e.RemoveChildrenByIndex(uint32(idx))
	return t
}

func (t *Tab) SetOnChange(onChange func(key string, widget IWidget)) *Tab {
	t.onChange = onChange
	return t
}

func (t *Tab) SetPlace(place place.Placement) *Tab {
	t.e.Set("place", place)
	return t
}
