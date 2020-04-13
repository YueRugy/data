package main

import "fmt"

const (
	defaultSize = 10
)

type CycleQueue struct {
	size     int
	front    int
	elements []interface{}
}

func NewCycleQueue() *CycleQueue {
	return &CycleQueue{
		size:     0,
		front:    0,
		elements: make([]interface{}, defaultSize),
	}
}

func (c *CycleQueue) Empty() bool {
	return c.Size() == 0
}
func (c *CycleQueue) Size() int {
	return c.size
}
func (c *CycleQueue) Front() interface{} {
	if c.Size() > 0 {
		return c.elements[c.front]
	}
	return nil
}

func (c *CycleQueue) DeQueue() interface{} {
	if c.Size() > 0 {
		ele := c.elements[c.front]
		c.elements[c.front] = nil
		c.front = c.index(1)
		c.size--
		return ele
	}
	return nil
}

func (c *CycleQueue) index(index int) int {
	return (c.front + index) % len(c.elements)
}
func (c *CycleQueue) EnQueue(ele interface{}) {
	c.ensureCap(c.Size() + 1)
	c.elements[c.index(c.Size())] = ele
	c.size++
}
func (c *CycleQueue) ensureCap(size int) {
	if size > len(c.elements) {
		newCap := len(c.elements) + len(c.elements)>>1
		newEle := make([]interface{}, newCap)
		for i := 0; i < len(c.elements); i++ {
			newEle[i] = c.elements[c.index(i)]
		}
		c.elements = newEle
		c.front = 0
	}
}
func main() {
	c := NewCycleQueue()
	for i := 0; i < 11; i++ {
		c.EnQueue(i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(c.DeQueue())
	}
	fmt.Println("--------------")
	show(c)
	for i := 0; i < 10; i++ {
		c.EnQueue(11 + i)
	}
	fmt.Println("---------")
	show(c)
}
func show(c *CycleQueue) {
	for i := 0; i < c.Size(); i++ {
		fmt.Println(c.elements[c.index(i)])
	}
	fmt.Println(c)
}
