package main

import "fmt"

type ArrayList struct {
	property []interface{}
	length   int
}

func NewArray() *ArrayList {
	al := &ArrayList{
		property: make([]interface{}, 0),
		length:   0,
	}
	return al
}
func (a *ArrayList) AddToEnd(property interface{}) {
	a.property = append(a.property, property)
	a.length += 1
}
func (a *ArrayList) AddToHead(property interface{}) {
	l := a.Len()
	al := make([]interface{}, l+1)
	al[0] = property
	copy(al[1:], a.property)
	a.length += 1
}
func (a *ArrayList) Len() (len int) {
	return a.length
}
func (a *ArrayList) Clear() {
	//a = NewArray()
	//a = &ArrayList{
	//	property: make([]interface{}, 0),
	//	length:   0,
	//}
	a.property = make([]interface{}, 0)
	a.length = 0
	fmt.Println(a.length)
	fmt.Println(a.property)
}
func main() {
	l := NewArray()
	l.AddToEnd(1)
	l.AddToEnd(2)
	l.AddToEnd(3)
	l.AddToEnd(4)

	fmt.Println(l.property)
	fmt.Println(l)
	l.Clear()
	fmt.Println(l)
	fmt.Println(l.property)

}
