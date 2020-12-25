package main

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	property []interface{}
	length   int
}

type ArrayLister interface {
	AddToEnd(property interface{})
	AddToHead(property interface{})
	Set(n int, property interface{}) error
	Get(n int) (interface{}, error)
	Insert(n int, property interface{}) error
	AppendList(property ...interface{})
	GetAll() (p []interface{})
	Clear()
	Len() int
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
	_ = a.Insert(0, property)
}
func (a *ArrayList) Insert(n int, property interface{}) error {
	if n >= a.Len() {
		return errors.New("The n  out of range")
	}
	if n < 0 {
		return errors.New("The n must greate zero")
	}
	l := a.Len()
	al := make([]interface{}, l+1)
	copy(al[:n], a.property[:n])
	al[n] = property
	copy(al[n+1:], a.property[n:])
	a.property = al
	a.length++
	return nil
}

func (a *ArrayList) Set(n int, property interface{}) error {
	if n >= a.Len() {
		return errors.New("The n  out of range")
	}
	if n < 0 {
		return errors.New("The n must greate zero")
	}
	a.property[n] = property
	return nil
}

func (a *ArrayList) Get(n int) (property interface{}, err error) {
	if n >= a.Len() {
		return nil, errors.New("The n  out of range")
	}
	if n < 0 {
		return nil, errors.New("The n must greate zero")
	}
	property = a.property[n]
	return property, nil
}

//
func (a *ArrayList) AppendList(property ...interface{}) {
	a.property = append(a.property, property...)
	fmt.Println(a.property)
	a.length += len(property)
}

func (a *ArrayList) GetAll() (property []interface{}) {
	return a.property
}

func (a *ArrayList) Clear() {
	//a = NewArray()

	a.property = make([]interface{}, 0)
	a.length = 0

}

func (a *ArrayList) Len() (len int) {
	return a.length
}

func main() {
	//目的是调用接口
	var b ArrayLister
	b = NewArray()
	fmt.Println(b)

	l := NewArray()
	l.AddToEnd(1)
	l.AddToEnd(2)
	l.AddToEnd(3)
	l.AddToEnd(5)
	err := l.Insert(3, 7)
	err = l.Insert(4, 8)
	l.AddToHead(11)
	fmt.Println(err)
	fmt.Println(l.property)
	fmt.Println(l.length)

	//
	l.Clear()
	fmt.Println("The clear : ", l.length)
	l.AddToEnd(1)
	l.AddToEnd(2)
	l.AddToEnd(3)
	l.AddToEnd(5)

	l.AppendList(6, 7, 8)
	fmt.Println(l.Len())
	fmt.Println(l.GetAll())

}
