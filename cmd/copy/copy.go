package main

import "fmt"

type ele struct {
}
type page interface {
	ELe() []*ele
}
type Layout func()

type my struct {
}

func (m my) ELe() []*ele {
	e1 := &ele{}
	e2 := &ele{}
	return []*ele{e1, e2}
}
func main() {
	m := new(my)
	fmt.Printf("%p\n", m.ELe())
	fmt.Printf("%p\n", m.ELe())
}
