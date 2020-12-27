package main

import "fmt"

type Set struct {
	intergerMap map[int]bool
}

func (set *Set) New() {
	set.intergerMap = make(map[int]bool)
}

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.intergerMap[element]
	return exists

}

func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.intergerMap[element] = true
	}
}
func (set *Set) DeleteElement(element int) {
	delete(set.intergerMap, element)

}
func (set *Set) InterSet(anotherSet *Set) *Set {
	for value, _ := range anotherSet.intergerMap {
		if !set.ContainsElement(value) {
			set.AddElement(value)
		}
	}
}
func main() {
	var set *Set
	set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
}
