package main

import "fmt"

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

func main() {
	arr := [...]int{7, 4, 9, 2, 5, 8, 11, 3, 12, 1}
	bst := NewBST()
	for _, val := range arr {
		bst.Add(val)
	}
	//fmt.Println(bst)
	PreRange(bst.root)
	fmt.Println("--------------")
	MidRange(bst.root)
	fmt.Println("--------------")
	PostRange(bst.root)
	fmt.Println("--------------")
	LevelRange(bst.root)
}
