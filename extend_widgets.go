package nicegoi

import (
	"fmt"
	"github.com/yaoguangduan/nicegoi/icons"
	"github.com/yaoguangduan/nicegoi/option"
	"github.com/yaoguangduan/nicegoi/option/menu"
	"reflect"
	"strings"
)

//===================box==============================

type Box struct {
	*valuedWidget
}

func createBox(elements ...IWidget) *Box {
	row := Box{valuedWidget: newValuedWidget("box", "")}
	row.e.Set("align", "center")
	row.e.AddChildren(elements...)
	return &row
}
func (w *Box) Align(align option.Align) *Box {
	w.e.Set("align", string(align))
	return w
}

func (w *Box) Horizontal() *Box {
	w.e.Set("direction", "horizontal")
	return w
}
func (w *Box) Vertical() *Box {
	w.e.Set("direction", "vertical")
	return w
}

func (w *Box) Remove(elements ...IWidget) *Box {
	w.e.RemoveChildren(elements...)
	return w
}
func (w *Box) RemoveByIdx(elements ...uint32) *Box {
	w.e.RemoveChildrenByIndex(elements...)
	return w
}
func (w *Box) AddItems(elements ...IWidget) *Box {
	w.e.AddChildren(elements...)
	return w
}

//===================card==============================

type Card struct {
	*valuedWidget
}

func NewCard(content string) *Card {
	return NewCardWithTitle("", content)
}
func NewCardWithTitle(title string, content string) *Card {
	card := Card{valuedWidget: newReadonlyWidget("card", content)}
	card.e.Set("title", title)
	card.e.AddChildren(newEmptyWidget(card.e.Eid()+"_A", "card_action"))
	card.e.AddChildren(newEmptyWidget(card.e.Eid()+"_F", "card_footer"))
	return &card
}
func (w *Card) SetTitle(title string) *Card {
	w.e.Set("title", title)
	return w
}

func (w *Card) SetContent(content string) *Card {
	w.e.Set("content", content)
	return w
}

func (w *Card) SetDesc(desc string) *Card {
	w.e.Set("desc", desc)
	return w
}
func (w *Card) SetWidth(width int) *Card {
	w.e.Set("width", width)
	return w
}
func (w *Card) AddActions(items ...IWidget) {
	for _, item := range items {
		w.e.Children()[0].AddChildren(item)
	}
}

func (w *Card) AddFooters(items ...IWidget) {
	for _, item := range items {
		w.e.Children()[1].AddChildren(item)
	}
}

//===================drawer==============================

type Drawer struct {
	*valuedWidget
}

func NewDrawer(header string) *Drawer {
	d := &Drawer{newValuedWidget("drawer", false)}
	d.e.Set("header", header)
	return d
}
func (d *Drawer) AddWidgets(widgets ...IWidget) *Drawer {
	d.e.AddChildren(widgets...)
	return d
}
func (d *Drawer) Open() *Drawer {
	d.e.Set("value", true)
	return d
}
func (d *Drawer) Close() *Drawer {
	d.e.Set("value", false)
	return d
}
func (d *Drawer) State() bool {
	return d.e.Get("value").(bool)
}
func (d *Drawer) SetPlace(p option.Placement) *Drawer {
	d.e.Set("placement", p)
	return d
}

//===================list==============================

type List struct {
	*valuedWidget
}

func NewList() *List {
	list := List{valuedWidget: newValuedWidget("list", "")}
	return &list
}

func (w *List) AddItems(items ...*ListItem) {
	for _, item := range items {
		w.e.AddChildren(item)
	}
}
func (w *List) RemoveItem(item *ListItem) {
	w.e.RemoveChildren(item)
}

func (w *List) RemoveItemByIdx(idx int) {
	w.e.RemoveChildrenByIndex(uint32(idx))
}

type ListItem struct {
	*valuedWidget
}

func NewListItem(text string) *ListItem {
	list := ListItem{valuedWidget: newReadonlyWidget("list_item", text)}
	return &list
}

func (w *ListItem) AddAction(action IWidget) *ListItem {
	w.e.AddChildren(action)
	return w
}

//===================menu==============================

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
	m.addMsgHandler(func(message *Message) {
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

//===================row==============================

type Row struct {
	*valuedWidget
}

func NewRow(items ...IWidget) *Row {
	w := &Row{newValuedWidget("row", "")}
	w.e.AddChildren(items...)
	return w
}
func (w *Row) SetGutter(h, v int) *Row {
	w.e.Set("gutter", []int{h, v})
	return w
}
func (w *Row) SetSpan(spans ...int) *Row {
	w.e.Set("span", spans)
	return w
}

func (w *Row) SetOffset(index, value int) *Row {
	m := w.e.Get("offset")
	if m == nil {
		m = make(map[int]int)
	}
	m.(map[int]int)[index] = value
	w.e.Set("offset", m)
	return w
}
func (w *Row) Justify(justify option.Justify) *Row {
	w.e.Set("justify", justify)
	return w
}

//=================table======================

type Table struct {
	*valuedWidget
}

func NewTable(data interface{}) *Table {
	cols, realData := parseColsAndData(data)
	if cols == nil || realData == nil {
		panic("invalid data format")
	}
	t := &Table{newReadonlyWidget("table", realData)}
	t.e.Set("columns", cols)
	t.e.Set("rowKey", cols[0]["colKey"])
	return t
}

func parseColsAndData(data interface{}) ([]map[string]string, interface{}) {
	var dt = reflect.ValueOf(data)
	if dt.Kind() == reflect.Interface {
		dt = dt.Elem()
	}
	if dt.Kind() == reflect.Ptr {
		dt = dt.Elem()
	}
	if dt.Kind() == reflect.Struct {
		var typ = reflect.TypeOf(data)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		return getColFromStruct(typ), []interface{}{data}
	} else if dt.Kind() == reflect.Slice {
		first := dt.Index(0)
		if first.Kind() == reflect.Ptr {
			first = first.Elem()
		}
		if first.Kind() == reflect.Interface {
			first = first.Elem()
		}
		if first.Kind() == reflect.Ptr {
			first = first.Elem()
		}
		if first.Kind() == reflect.Struct {
			return getColFromStruct(first.Type()), data
		} else if first.Kind() == reflect.Slice {
			cols := make([]map[string]string, 0)
			for i := 0; i < first.Len(); i++ {
				cell := fmt.Sprintf("%s", first.Index(i))
				cols = append(cols, map[string]string{"colKey": cell, "title": cell})
			}
			realData := make([]map[string]string, 0)
			for i := 1; i < dt.Len(); i++ {
				row := dt.Index(i)
				rowData := make(map[string]string)
				for j := 0; j < row.Len(); j++ {
					if first.Len() > j {
						rowData[fmt.Sprintf("%s", first.Index(j))] = fmt.Sprintf("%s", row.Index(j))
					}
				}
				realData = append(realData, rowData)
			}
			return cols, realData
		} else if first.Kind() == reflect.Map {
			cols := make([]map[string]string, 0)
			for _, key := range first.MapKeys() {
				cell := fmt.Sprintf("%s", key.Interface())
				cols = append(cols, map[string]string{"colKey": cell, "title": cell})
			}
			return cols, data
		}
	} else if dt.Kind() == reflect.Map {
		mv := dt.MapIndex(dt.MapKeys()[0])
		if mv.Kind() == reflect.Ptr {
			mv = mv.Elem()
		}
		if mv.Kind() == reflect.Slice {
			cols := make([]map[string]string, 0)
			realData := make([]map[string]string, 0)
			var maxLen = -1
			for _, k := range dt.MapKeys() {
				cell := fmt.Sprintf("%s", k)
				cols = append(cols, map[string]string{"colKey": cell, "title": cell})
				maxLen = dt.MapIndex(k).Len()
			}
			for i := 0; i < maxLen; i++ {
				row := make(map[string]string)
				for _, k := range dt.MapKeys() {
					valSlice := dt.MapIndex(k)
					cell := fmt.Sprintf("%s", k)
					if i < valSlice.Len() {
						v := fmt.Sprintf("%s", valSlice.Index(i))
						row[cell] = v
					} else {
						row[cell] = ""
					}
				}
				realData = append(realData, row)
			}
			return cols, realData

		}
	}
	return nil, nil
}

func getColFromStruct(typ reflect.Type) []map[string]string {
	cols := make([]map[string]string, 0)
	for i := range typ.NumField() {
		f := typ.Field(i)
		col := f.Tag.Get("json")
		if strings.Contains(col, ",") {
			col = strings.Split(col, ",")[0]
		}
		cols = append(cols, map[string]string{"colKey": col, "title": col})
	}
	return cols
}

//===============tab===============

type Tab struct {
	*valuedWidget
	onChange func(key string, widget IWidget)
	widgets  map[string]IWidget
}

func NewTab() *Tab {
	ret := &Tab{valuedWidget: newReadonlyWidget("tab", ""), widgets: make(map[string]IWidget)}
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
	ew.element().AddChildren(widget)
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

func (t *Tab) SetPlace(place option.Placement) *Tab {
	t.e.Set("place", place)
	return t
}
