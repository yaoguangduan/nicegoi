package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/ui/icons"
	"github.com/yaoguangduan/nicegoi/internal/ui/place"
)

type Tab struct {
	*valueWidget
	onChange func(key string, widget IWidget)
	widgets  map[string]IWidget
}

func NewTab() *Tab {
	ret := &Tab{valueWidget: newValueWidget("tab", ""), widgets: make(map[string]IWidget)}
	ret.opt.Set("place", "top")
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
	ew := newValueWidget(t.opt.Eid()+"_S", val)
	ew.opt.Set("icon", icon)
	ew.Element().AddChildren(widget)
	t.opt.AddChildren(ew)
	if len(t.opt.Children()) == 1 {
		t.set(t.opt.Children()[0].Type())
	}
	t.widgets[val] = widget
	return t
}

func (t *Tab) Remove(val string) *Tab {
	delete(t.widgets, val)
	var idx = -1
	for i, w := range t.opt.Children() {
		if w.Get("value").(string) == val {
			idx = i
			break
		}
	}
	if idx < 0 {
		return t
	}
	t.opt.RemoveChildrenByIndex(uint32(idx))
	return t
}

func (t *Tab) SetOnChange(onChange func(key string, widget IWidget)) *Tab {
	t.onChange = onChange
	return t
}

func (t *Tab) SetPlace(place place.Placement) *Tab {
	t.opt.Set("place", place)
	return t
}
