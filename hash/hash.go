package main

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	defaultCap = 1 << 4
	black      = 0
	red        = 1
	left       = 0
	right      = 1
)

type Hash struct {
	size   int
	bucket []*Node
}

func (hash *Hash) Put(k, v int) {
	node := &Node{
		k:    k,
		v:    v,
		code: hashCode(k),
	}
	index := hash.index(k)
	root := hash.bucket[index]
	fmt.Printf("root  %p,bucket %p", root, hash.bucket[index])
	if root == nil {
		add(&root, node)
		hash.bucket[index] = root
		hash.size++
	} else {

		oldNode := add(&root, node)
		if oldNode != nil {
			hash.size++
			hash.bucket[index] = root
		}
		fmt.Println()
	}

}

func (hash *Hash) index(k int) int {
	return hashCode(k) & (len(hash.bucket) - 1)
}

type Node struct {
	k      int
	v      int
	color  int
	code   int
	parent *Node
	left   *Node
	right  *Node
}

func add(root **Node, node *Node) *Node {
	if node == nil {
		return nil
	}
	resNode := node
	if *root == nil {
		*root = node
	} else { //寻找合适的节点
		for temp := *root; temp != nil; {
			resNode = temp
			if compare(node.code, temp.code, node, temp) > 0 {
				temp = temp.right
			} else if compare(node.code, temp.code, node, temp) < 0 {
				temp = temp.left
			} else {
				return nil
			}
		}
		if compare(node.code, resNode.code, node, resNode) > 0 {
			resNode.right = node
		} else if compare(node.code, resNode.code, node, resNode) < 0 {
			resNode.left = node
		}
		node.parent = resNode
	}
	afterAdd(root, node)
	return resNode
}

func (node *Node) colorOf() int {
	return node.color
}

func (node *Node) black() {
	node.color = black
}

func (node *Node) red() {
	node.color = red
}

func (node *Node) isBlack() bool {
	if node == nil {
		return true
	}
	return node.colorOf() == black
}

func (node *Node) isRed() bool {
	if node == nil {
		return false
	}
	return node.colorOf() == red
}
func (node *Node) dire() int {
	if node != nil && node.parent != nil {
		if reflect.DeepEqual(node, node.parent.left) {
			return left
		} else {
			return right
		}
	}
	return left
}

func afterAdd(root **Node, node *Node) {
	parent := node.parent
	if parent == nil { //根节点
		(*root).black()
		return
	}
	node.red()            //default red
	if parent.isBlack() { //父亲是黑色节点
		return
	}
	uncle := sibling(parent)
	grand := parent.parent
	flag := grand.parent == nil
	if uncle.isRed() { //叔父节点是红色
		uncle.black()
		parent.black()
		grand.red()
		afterAdd(root, grand)
	} else {
		if parent.dire() == left {
			if node.dire() == right {
				rotateRR(parent, node)
				parent = node
			}
			rotateLL(grand, parent)
		} else {
			if node.dire() == left {
				rotateLL(parent, node)
				parent = node
			}
			rotateRR(grand, parent)
		}
		parent.black()
		grand.red()
		if flag {
			*root = parent
			parent.parent = nil
		}
		fmt.Println()
	}
}

func compare(h1, h2 int, n1, n2 *Node) int {
	res := h1 - h2
	if res != 0 {
		return res
	}
	return n1.v - n2.v
}

func sibling(node *Node) *Node {
	if node == nil || node.parent == nil {
		return nil
	}
	if reflect.DeepEqual(node, node.parent.left) {
		return node.parent.right
	} else {
		return node.parent.left
	}
}
func NewHash() *Hash {
	return &Hash{
		size:   0,
		bucket: make([]*Node, defaultCap),
	}
}

func rotateLL(p, c *Node) {
	p.left = c.right
	if c.right != nil {
		c.right.parent = p
	}
	c.right = p
	if p.parent != nil {
		c.parent = p.parent
		flag := p.dire() == left
		if flag {
			p.parent.left = c
		} else {
			p.parent.right = c
		}
	}
	p.parent = c
}

func rotateRR(p, c *Node) {
	p.right = c.left
	if c.left != nil {
		c.left.parent = p
	}
	c.left = p
	if p.parent != nil {
		c.parent = p.parent
		flag := p.dire() == left
		if flag {
			p.parent.left = c
		} else {
			p.parent.right = c
		}
	} else {
		c.parent = nil
	}
	p.parent = c
}

func main() {
	arr := []int{94, 28, 70, 86, 89, 72, 24, 7, 75, 33, 23, 9, 55, 22, 80, 30, 18}
	//test1(arr)
	hash := NewHash()
	for _, v := range arr {
		hash.Put(0, v)
	}
	f1(hash.bucket[0])
	//fmt.Println(hash)
}

func test2(x, y *int) {
	x, y = y, x
}

func test1(arr []int) {
	var root *Node
	for _, v := range arr {
		node := &Node{
			k:      0,
			v:      v,
			code:   hashCode(0),
			color:  0,
			parent: nil,
			left:   nil,
			right:  nil,
		}
		add(&root, node)
	}
	f1(root)
}

func f1(root *Node) {
	visitor := func(node *Node) {
		fmt.Print(strconv.Itoa(node.v) + "\t")
	}
	visitor1 := func(node *Node) {
		if node.color == red {
			fmt.Print(strconv.Itoa(node.v) + "\t")
		}
	}
	LevelRange(root, visitor)
	fmt.Println()
	fmt.Println(HeightByLevel(root))
	LevelRange(root, visitor1)
}

func HeightByLevel(root *Node) int {
	if root == nil {
		return 0
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	levelSize := 1 //用于记录每层有多少节点 第一次levelSize=root节点=1 没遍历一次levelSize-- 维护一个队列
	//遍历队列中的节点元素元素然后讲该节点的左右节点加入到队列中
	//levelSize=0 时意味 某层已经遍历完 此时队列里的元素数量就是下一层节点的数量
	count := 0 //记录层数
	for len(queue) > 0 {
		levelSize--
		node := queue[0]
		queue = queue[1:]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		if levelSize == 0 {
			count++
			levelSize = len(queue)
		}

	}
	return count
}

//层序遍历
func LevelRange(node *Node, visitor func(*Node)) {
	if node == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		n := queue[0]
		visitor(n)
		queue = queue[1:]
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}
func hashCode(num int) int {
	return num
	//TODO
}
