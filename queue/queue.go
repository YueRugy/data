package main

import "fmt"

type Queue struct {
	size int
	head *Node
	tail *Node
}
type Node struct {
	ele  interface{}
	prev *Node
	next *Node
}

func NewQueue() *Queue {
	return &Queue{
		size: 0,
	}
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Clear() {
	q.head, q.tail = nil, nil
}

func (q *Queue) EnQueue(ele interface{}) {
	n := &Node{
		ele: ele,
	}
	if q.tail == nil {
		q.tail, q.head = n, n
	} else {
		q.tail.next, n.prev, q.tail = n, q.tail, n
	}
	q.size++
}

func (q *Queue) DeQueue() interface{} {
	if q.size == 0 {
		return nil
	}
	ele := q.head.ele
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return ele
}

func (q *Queue) front() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.head.ele
}

func main() {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	fmt.Println(q.Empty())
	q.EnQueue(5)
	show(q)
	fmt.Println("-----------")
	for i := q.size; i > 0; i-- {
		fmt.Println(q.front())
		fmt.Println("------------")
		q.DeQueue()
	}
	fmt.Println(q.size)
}
func show(q *Queue) {
	for temp := q.head; temp != nil; temp = temp.next {
		fmt.Println(temp.ele)
	}
}
