package menu

import "github.com/yaoguangduan/nicegoi/internal/ui/icons"

type Option struct {
	Collapse  bool                                 `json:"collapse"`
	Value     string                               `json:"value"`
	MenuItems []*ItemOption                        `json:"items,omitempty"`
	OnChange  func(menu *Option, item *ItemOption) `json:"-"`
}
type ItemOption struct {
	Value    string        `json:"value"`
	Label    string        `json:"label"`
	Icon     icons.Icon    `json:"icon,omitempty"`
	Children []*ItemOption `json:"children,omitempty"`
}

func New() *Option {
	return &Option{}
}
func NewWithSelect(val string) *Option {
	return &Option{Value: val}
}
func NewItem(label, value string) *ItemOption {
	return &ItemOption{Label: label, Value: value}
}
func NewItemWithIcon(label, value string, icon icons.Icon) *ItemOption {
	return &ItemOption{Label: label, Value: value, Icon: icon}
}
func (m *Option) SetOnChange(f func(menu *Option, item *ItemOption)) *Option {
	m.OnChange = f
	return m
}
func (m *Option) SetCollapse(c bool) *Option {
	m.Collapse = c
	return m
}
func (m *Option) AddItems(menuItem ...*ItemOption) *Option {
	m.MenuItems = append(m.MenuItems, menuItem...)
	return m
}
func (mi *ItemOption) SetIcon(icon icons.Icon) *ItemOption {
	mi.Icon = icon
	return mi
}
func (mi *ItemOption) AddItems(menuItem ...*ItemOption) *ItemOption {
	mi.Children = append(mi.Children, menuItem...)
	return mi
}
