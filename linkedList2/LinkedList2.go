package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	ele  interface{}
	next *Node
}

type LinkedList2 struct {
	size int
	head *Node
}

func NewLinkedList2() *LinkedList2 {
	return &LinkedList2{
		size: 0,
		head: &Node{
			ele:  nil,
			next: nil,
		},
	}
}
func (l *LinkedList2) Add(ele interface{}) {
	temp := l.node(l.size)
	if temp != nil {
		temp.next = &Node{
			ele:  ele,
			next: nil,
		}
		l.size++
	}
}
func (l *LinkedList2) AddIndex(index int, ele interface{}) {
	temp := l.node(index)
	if temp != nil {
		n := &Node{
			ele:  ele,
			next: temp.next,
		}
		temp.next = n
		l.size++
	}
}
func (l *LinkedList2) Remove() {
	temp := l.node(l.size - 1)
	temp.next = nil
	l.size--
}
func (l *LinkedList2) Contains(ele interface{}) bool {
	temp := l.head
	for i := 1; i <= l.size; i++ {
		temp = temp.next
		if reflect.DeepEqual(ele, temp.ele) {
			return true
		}
	}
	return false
}

func (l *LinkedList2) Size() int {
	return l.size
}

func (l *LinkedList2) Empty() bool {
	return l.size == 0
}

func (l *LinkedList2) node(index int) *Node {
	if !l.RangeIndex(index) {
		temp := l.head
		for i := 0; i < index; i++ {
			temp = temp.next
		}
		return temp
	}
	return nil
}

func (l *LinkedList2) RangeIndex(index int) bool {
	return index < 0 || index > l.size
}

func main() {
	l := NewLinkedList2()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Remove()
	l.Add(5)
	l.AddIndex(4, 10)
	show(l)
}

func show(l *LinkedList2) {
	temp := l.head
	for i := 1; i <= l.size; i++ {
		temp = temp.next
		fmt.Println(temp.ele)
	}
}
