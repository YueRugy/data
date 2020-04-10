package main

import (
	"errors"
	"fmt"
)

const (
	defaultSize = 16
)

type List struct {
	array  []int
	length int
	//cap    int
}

func (l *List) Add(element int) {
	l.expansion(cap(l.array) + 1)
	l.array[l.length] = element
	l.length++
}

func NewList() *List {
	l := &List{
		length: 0,
	}
	l.array = make([]int, 16)
	return l
}

func (l *List) Set(element, index int) error {
	if l.out(index) {
		return errors.New("index out range")
	}
	l.array[index] = element
	//是否需要扩容
	//l.expansion(index)
	return nil
}
func (l *List) expansion(index int) {
	if index >= l.length {
		//l.cap = index * 2
		//创建新数组指向
		//var tCap = l.cap
		l.resize(index * 2)
		//l.cap = index * 2
	}
}
func (l *List) Remove() bool {
	if l.Empty() {
		return false
	}
	l.length--
	return true
}
func (l *List) resize(newCap int) {
	tempArray := make([]int, newCap)
	for i := 0; i < l.length; i++ {
		tempArray[i] = l.array[i]
	}
	l.array = tempArray
}
func (l *List) IndexOf(ele int) int {
	if l.Empty() {
		return -1
	} else {
		for v, i := range l.array {
			if v == ele {
				return i
			}
		}
	}
	return -1
}
func (l *List) RemoveIndex(index int) int {
	if l.out(index) {
		fmt.Println("index range out")
		return -1
	}
	//平移
	for i := index; i < l.length-2; i++ {
		l.array[i] = l.array[i+1]
	}
	l.length--
	return index
}
func (l *List) Get(index int) (int, error) {
	if l.out(index) {
		return -1, errors.New("index out range")
	} else {
		for v, i := range l.array {
			if i == index {
				return v, nil
			}
		}
	}
	return -1, errors.New("index out range")
}
func (l *List) out(index int) bool {
	return index < 0 || index >= l.length
}
func (l *List) Size() int {
	return l.length
}

func (l *List) Empty() bool {
	return l.length == 0
}
func (l *List) Contains(element int) bool {
	if l.Size() == 0 {
		return false
	}
	for ele, _ := range l.array {
		if ele == element {
			return true
		}
	}
	return false
}

func main() {
	l := NewList()
	fmt.Println(l.Empty())
	for i := 0; i < 20; i++ {
		l.Add(i)
	}
	//show(l)
	//fmt.Println(l.Get(5))
	l.Remove()
	fmt.Println(l.Contains(12))
	fmt.Println(l.IndexOf(3))
	//show(l)
	fmt.Println(l.Empty())
}
func show(l *List) {
	for i := 0; i < l.length; i++ {
		fmt.Println(l.array[i])
	}
}
