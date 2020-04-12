package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	ele  interface{}
	next *Node
	prev *Node
}
type DoubleLinkedList struct {
	size int
	head *Node
	tail *Node
}

func NewDouble() *DoubleLinkedList {
	return &DoubleLinkedList{
		size: 0,
	}
}
func (l *DoubleLinkedList) Add(ele interface{}) {
	node := &Node{
		ele: ele,
	}
	if l.tail == nil {
		l.head, l.tail = node, node
	} else {
		/*node.prev = l.tail
		l.tail.next = node
		l.tail = node*/
		node.prev, l.tail.next, l.tail = l.tail, node, node
	}
	l.size++
}
func (l *DoubleLinkedList) AddIndex(index int, ele interface{}) {
	if !l.rangeIndex(index) {
		node := l.node(index)
		newNode := &Node{
			ele: ele,
		}
		if node == nil || l.Size() == index {
			l.Add(ele)
		} else {
			prev := node.prev
			if prev == nil {
				l.head = newNode
			} else {
				newNode.prev = prev
				prev.next = newNode
			}
			node.prev = newNode
			newNode.next = node
			l.size++
		}
	}
}
func (l *DoubleLinkedList) Remove() {
	if l.size > 0 {
		if l.size == 1 {
			l.head, l.tail = nil, nil
		} else {
			l.tail.prev.next = nil
			l.tail = l.tail.prev
		}
		l.size--
	}
}

func (l *DoubleLinkedList) IndexOf(ele interface{}) int {
	index := -1
	temp := l.head
	flag := false
	for i := 0; i < l.Size(); i++ {
		index++
		if reflect.DeepEqual(temp.ele, ele) {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag == false {
		index = -1
	}
	return index
}
func (l *DoubleLinkedList) Contains(ele interface{}) bool {
	return l.IndexOf(ele) > -1
}
func (l *DoubleLinkedList) Empty() bool {
	return l.size == 0
}
func (l *DoubleLinkedList) Size() int {
	return l.size
}

func (l *DoubleLinkedList) node(index int) *Node {
	if !l.rangeIndex(index) {
		if index < l.size>>1 {
			temp := l.head
			for i := 0; i < index; i++ {
				temp = temp.next
			}
			return temp
		} else {
			temp := l.tail
			for i := index - 1; i > index; i-- {
				temp = temp.next
			}
			return temp
		}
	}
	return nil
}
func (l *DoubleLinkedList) rangeIndex(index int) bool {
	return index < 0 || index > l.Size()
}

func main() {
	l := NewDouble()
	l.AddIndex(0, 0)
	l.Add(1)
	l.AddIndex(1, 999)
	show(l)
	fmt.Println("-----------")
	l.Remove()
	show(l)
}
func show(l *DoubleLinkedList) {
	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Println(temp.ele)
	}
}
