package ui

import "github.com/yaoguangduan/nicegoi/internal/option"

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
