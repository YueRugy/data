package main

import (
	"fmt"
)

type BinarySearchTree struct {
	size int
	root *Node
}

type Node struct {
	ele    int
	left   *Node
	right  *Node
	parent *Node
}

func (b *BinarySearchTree) Size() int {
	return b.size
}

func (b *BinarySearchTree) Empty() bool {
	return b.Size() == 0
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (b *BinarySearchTree) Add(ele int) {
	node := &Node{
		ele: ele,
	}
	if b.root == nil {
		b.root = node
		b.size++
		return
	}
	// 找到合适的节点
	temp := b.root
	var parent *Node
	for ; temp != nil; {
		parent = temp
		if compare(ele, temp.ele) > 0 {
			temp = temp.right
		} else if compare(ele, temp.ele) < 0 {
			temp = temp.left
		} else {
			break
		}
	}
	if parent == nil {
		return
	}
	if compare(ele, parent.ele) > 0 {
		parent.right = node
		node.parent = parent
		b.size++
	} else if compare(ele, parent.ele) < 0 {
		parent.left = node
		node.parent = parent
		b.size++
	} else {
		node.parent = parent
		node.left = parent.left
		node.right = parent.right
	}

}

func PreRange(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.ele)
	fmt.Print("\t")
	PreRange(n.left)
	PreRange(n.right)
}
func MidRange(n *Node) {
	if n == nil {
		return
	}
	MidRange(n.left)
	fmt.Print(n.ele)
	fmt.Print("\t")
	MidRange(n.right)
}

func PostRange(n *Node) {
	if n == nil {
		return
	}
	PostRange(n.left)
	PostRange(n.right)
	fmt.Print(n.ele)
	fmt.Print("\t")
}

func LevelRange(root *Node) {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for ; len(queue) != 0; {
		node := queue[0]
		fmt.Print(node.ele)
		fmt.Print("\t")
		queue = queue[1:]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func compare(c1, c2 int) int {
	return c1 - c2
}

func PreRangeFunc(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	visitor(node.ele)
	PreRangeFunc(node.left, visitor)
	PreRangeFunc(node.right, visitor)
}

func MidRangeFunc(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	MidRangeFunc(node.left, visitor)
	visitor(node.ele)
	MidRangeFunc(node.right, visitor)
}

func PostRangeFunc(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	PostRangeFunc(node.left, visitor)
	PostRangeFunc(node.right, visitor)
	visitor(node.ele)
}

func LevelRangeFunc(root *Node, visitor func(int)) {
	if root == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for ; len(queue) > 0; {
		node := queue[0]
		visitor(node.ele)
		queue = queue[1:]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func main() {
	arr := [...]int{7, 4, 9, 2, 5, 8, 11, 3, 12, 1}
	bst := NewBST()
	for _, val := range arr {
		bst.Add(val)
	}
	//fmt.Println(bst)
	visitor := func(ele int) {
		fmt.Printf("_%d_\t", ele)
	}
	PreRange(bst.root)
	fmt.Println()
	fmt.Println("---------------------")
	PreRangeFunc(bst.root, visitor)
	fmt.Println()
	fmt.Println("--------------")
	MidRange(bst.root)
	fmt.Println()
	fmt.Println("--------------")
	MidRangeFunc(bst.root, visitor)
	fmt.Println()
	fmt.Println("--------------")
	PostRange(bst.root)
	fmt.Println()
	fmt.Println("--------------")
	PostRangeFunc(bst.root, visitor)
	fmt.Println()
	fmt.Println("--------------")
	LevelRange(bst.root)
	fmt.Println()
	fmt.Println("--------------")
	LevelRangeFunc(bst.root, visitor)
}
