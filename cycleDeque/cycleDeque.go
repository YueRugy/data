package main

import "fmt"

const (
	defaultCap = 10
)

type CycleDeque struct {
	size     int
	front    int
	elements []interface{}
}

func NewCycleDeque() *CycleDeque {
	return &CycleDeque{
		size:     0,
		front:    0,
		elements: make([]interface{}, defaultCap),
	}
}

func (c *CycleDeque) Empty() bool {
	return c.Size() == 0
}

func (c *CycleDeque) Size() int {
	return c.size
}

func (c *CycleDeque) index(num int) int {
	num += c.front
	if num < 0 {
		return len(c.elements) + num
	}
	if num >= len(c.elements) {
		return num - len(c.elements)
	}
	return num
}

func (c *CycleDeque) EnDequeFear(ele interface{}) {
	c.ensureCap(c.Size() + 1)
	c.elements[c.index(c.Size())] = ele
	c.size++
}

func (c *CycleDeque) EnDequeFront(ele interface{}) {
	c.ensureCap(c.Size() + 1)
	c.elements[c.index(-1)] = ele
	c.front = c.index(-1)
	c.size++
}

func (c *CycleDeque) Front() interface{} {
	if !c.Empty() {
		return c.elements[c.front]
	}
	return nil
}

func (c *CycleDeque) Fear() interface{} {
	if !c.Empty() {
		return c.elements[c.index(c.Size()-1)]
	}
	return nil
}
func (c *CycleDeque) DeDequeFront() interface{} {
	if !c.Empty() {
		res := c.elements[c.front]
		c.elements[c.front] = nil
		c.front = c.index(1)
		c.size--
		return res
	}
	return nil
}

func (c *CycleDeque) DeDequeFear() interface{} {
	if !c.Empty() {
		res := c.elements[c.index(c.Size())]
		c.elements[c.index(c.Size()-1)] = nil
		c.size--
		return res
	}
	return nil
}

func (c *CycleDeque) ensureCap(newSize int) {
	if newSize > len(c.elements) {
		newEle := make([]interface{}, len(c.elements)+len(c.elements)>>1)
		//copy old elements to new elements
		for i := 0; i < c.Size(); i++ {
			newEle[i] = c.elements[c.index(i)]
		}
		c.elements, c.front = newEle, 0
	}
}
func main() {
	c := NewCycleDeque()
	for i := 0; i < 10; i++ {
		c.EnDequeFront(i + 1)
		c.EnDequeFear(i + 100)
	}
	show(c)
	for i := 0; i < 3; i++ {
		c.DeDequeFront()
		c.DeDequeFear()
	}
	show(c)
	c.EnDequeFront(11)
	c.EnDequeFront(12)
	show(c)
}


func show(c *CycleDeque) {
	//	for i := 0; i < c.Size(); i++ {
	//		fmt.Println(c.elements[c.index(i)])                                                cycleDeque/cycleDeque.go:101
	//	}
	fmt.Println(c)
	fmt.Println("----------------")
}
