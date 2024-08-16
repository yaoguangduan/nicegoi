package ui

import (
	"github.com/yaoguangduan/nicegoi/internal/server"
)

type LayoutFunc func() []IWidget

type XPage struct {
	name   string
	layout LayoutFunc
}

func (p *XPage) Name() string {
	return p.name
}
func (p *XPage) GenPage() server.IPage {
	root := createElement("-")
	root.AddChildren(p.layout()...)
	page := &Page{
		root:  root,
		route: "/" + p.Name(),
		name:  p.name,
		title: p.name,
	}
	pw := PageWidget{
		vw: newReadonlyWidget("PW", ""),
		p:  page,
	}
	page.delegate = &pw
	return page
}

func PageN(name string, fn LayoutFunc) {
	xp := XPage{name, fn}
	server.RegisterLayout(&xp)
}
