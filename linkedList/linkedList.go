package main

import (
	"errors"
	"fmt"
	"reflect"
)

type LinkedList struct {
	size  int
	first *Node
}

type Node struct {
	ele  interface{}
	next *Node
}

func NewLinked() *LinkedList {
	return &LinkedList{
		size:  0,
		first: nil,
	}
}
func (l *LinkedList) Add(ele interface{}) {
	if l.size == 0 {
		node := &Node{
			ele:  ele,
			next: l.first,
		}
		l.first = node
		l.size++
		return
	}
	temp, err := l.node(l.size - 1)
	if err == nil {
		node := &Node{
			ele:  ele,
			next: nil,
		}
		temp.next = node
		l.size++
	}
}
func (l *LinkedList) node(index int) (*Node, error) {
	if !l.rangeIndex(index) {
		temp := l.first
		for i := 0; i < index; i++ {
			temp = temp.next
		}
		return temp, nil
	}
	return nil, errors.New("index error")
}
func (l *LinkedList) rangeIndex(index int) bool {
	return index < 0 || index > l.size
}

func (l *LinkedList) AddIndex(index int, ele interface{}) {
	if index == 0 {
		node := &Node{
			ele:  ele,
			next: l.first,
		}
		l.first = node
		l.size++
		return
	}
	temp, err := l.node(index - 1)
	if err == nil {
		node := &Node{
			ele:  ele,
			next: temp.next,
		}
		temp.next = node
		l.size++
	}
}
func (l *LinkedList) Remove() {
	if l.size == 0 {
		return
	}
	if l.size == 1 {
		l.first = nil
		l.size--
	}
	temp, err := l.node(l.size - 2)
	if err == nil {
		temp.next = nil
		l.size--
	}
}
func (l *LinkedList) Size() int {
	return l.size
}
func (l *LinkedList) Empty() bool {
	return l.size == 0
}
func (l *LinkedList) Set(ele interface{}, index int) {
	temp, err := l.node(index)
	if err == nil {
		temp.ele = ele
	}
}
func (l *LinkedList) Get(index int) interface{} {

	temp, err := l.node(index)
	if err == nil {
		return temp.ele
	}
	return nil
}
func (l *LinkedList) Contains(ele interface{}) bool {
	for temp := l.first; temp != nil; temp = temp.next {
		if reflect.DeepEqual(temp.ele, ele) {
			return true
		}
	}
	return false
}
func (l *LinkedList) IndexOf(ele interface{}) int {
	index := -1
	esc := false
	for temp := l.first; temp != nil; temp = temp.next {
		if reflect.DeepEqual(temp.ele, ele) {
			esc = true
			break
		}
		index++
	}
	if !esc {
		index = -1
	}
	return index
}

/*func main() {

	l := NewLinked()
	l.Add(3)
//	l.Add(4)
//	l.Add(5)
//	l.Add(6)
//	l.AddIndex(0, 0)
//	l.AddIndex(2, 10)
//	l.AddIndex(5, 7)
	fmt.Println(l.Empty())
	fmt.Println(l.size)
	l.Set(4,0)
	//l.Remove()
	l.Show()
}
*/
func (l *LinkedList) Show() {
	for t := l.first; t != nil; t = t.next {
		fmt.Println(t.ele)
	}
}
