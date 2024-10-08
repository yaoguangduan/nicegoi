package menu

import (
	"github.com/yaoguangduan/nicegoi/icons"
)

type Option struct {
	Collapse  bool          `json:"collapse"`
	Value     string        `json:"value"`
	MenuItems []*ItemOption `json:"items,omitempty"`
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
