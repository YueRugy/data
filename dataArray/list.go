package main

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	defaultSize = 16
)

type List struct {
	array  []interface{}
	length int
	//cap    int
}

func (l *List) Add(element interface{}) {
	l.expansion(l.length + 1)
	l.array[l.length] = element
	l.length++
}

func NewList() *List {
	l := &List{
		length: 0,
	}
	l.array = make([]interface{}, defaultSize)
	return l
}

func (l *List) Set(element interface{}, index int) error {
	if l.out(index) {
		return errors.New("index out range")
	}
	l.array[index] = element
	//是否需要扩容
	//l.expansion(index)
	return nil
}
func (l *List) expansion(index int) {
	if index >= cap(l.array) {
		//l.cap = index * 2
		//创建新数组指向
		//var tCap = l.cap
		l.resize(index << 1)
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
	tempArray := make([]interface{}, newCap)
	for i := 0; i < l.length; i++ {
		tempArray[i] = l.array[i]
	}
	l.array = tempArray
}
func (l *List) IndexOf(ele interface{}) int {
	if l.Empty() {
		return -1
	} else {
		for i, v := range l.array {
			if reflect.DeepEqual(v, ele) {
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

	l.length--
	//平移
	for i := index; i < l.length; i++ {
		l.array[i] = l.array[i+1]
	}
	return index
}
func (l *List) Get(index int) (interface{}, error) {
	if l.out(index) {
		return nil, errors.New("index out range")
	} else {
		for i, v := range l.array {
			if i == index {
				return v, nil
			}
		}
	}
	return nil, errors.New("index out range")
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
func (l *List) Contains(element interface{}) bool {
	if l.Size() == 0 {
		return false
	}
	for ele := range l.array {
		if element == ele {
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
	l.RemoveIndex(12)
	fmt.Println(l.Contains(12))
	fmt.Println(l.IndexOf(3))
	show(l)
	fmt.Println(l.Empty())
	fmt.Println(l.Get(4))
}
func show(l *List) {
	for i := 0; i < l.length; i++ {
		fmt.Println(l.array[i])
	}
}
