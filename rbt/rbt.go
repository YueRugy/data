package rbt

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	left, black = iota, iota
	right, red
	unknown, _
)

type Node struct {
	K, v, code, Color   int
	left, right, parent *Node
}

type RedBlackTree struct {
	size int
	root *Node
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		size: 0,
		root: nil,
	}
}

func (node *Node) Compare(c1, c2 int, n *Node) int {
	if c1 == c2 {
		return node.K - n.K
	}
	return c1 - c2
}

func (node *Node) Predecessor() *Node {
	if node == nil {
		return nil
	}
	var resNode *Node
	if node.left != nil {
		for temp := node.left; temp != nil; {
			resNode = temp
			temp = temp.right
		}
		return resNode
	}

	if node.left == nil && node.parent != nil {
		for temp := node; temp != nil; {
			if temp.dire() == right {
				return temp
			}
			temp = temp.parent
		}
	}
	return nil
}

func (node *Node) dire() int {
	if node == nil || node.parent == nil {
		return unknown
	}
	if reflect.DeepEqual(node, node.parent.left) {
		return left
	}
	return right
}

func (node *Node) red() {
	node.Color = red
}
func (node *Node) black() {
	node.Color = black
}

func (node *Node) isBlack() bool {
	return node == nil || node.colorOf() == black
}
func (node *Node) isRed() bool {
	if node == nil {
		return false
	}
	return node.colorOf() == red
}

func (node *Node) colorOf() int {
	if node == nil {
		return black
	}
	return node.Color
}

func (node *Node) coloring(color int) {
	if node == nil {
		return
	}
	node.Color = color
}

func (node *Node) sibling() *Node {
	if node.dire() == unknown {
		return nil
	}
	if node.dire() == left {
		return node.parent.right
	}
	return node.parent.left
}

//-------------------------------------------------

func (rbt *RedBlackTree) Remove(k int) bool {
	node := rbt.node(k)
	if node == nil {
		return false
	}

	var childNode *Node = nil
	//dire := unknown
	if node.left != nil && node.right != nil {
		preNode := node.Predecessor()
		node.code = preNode.code
		node.K = preNode.K
		node.v = preNode.v
		node = preNode
	}
	flag := node.dire()
	dire := unknown
	if node.left != nil {
		if flag == left {
			node.parent.left = node.left
			dire = left
		} else if flag == right {
			node.parent.right = node.left
			dire = right
		} else {
			rbt.root = node.left
		}
		node.left.parent = node.parent
		childNode = node.left
	} else if node.right != nil {
		if flag == left {
			node.parent.left = node.right
			dire = left
		} else if flag == right {
			node.parent.right = node.right
			dire = right
		} else {
			rbt.root = node.right
		}
		node.right.parent = node.parent
		childNode = node.right
	} else {
		if flag == left {
			node.parent.left = nil
			dire = left
		} else if flag == right {
			node.parent.right = nil
			dire = right
		} else {
			rbt.root = nil
		}
	}
	rbt.size--
	rbt.afterRemove(node, childNode, dire)
	return true
}

func (rbt *RedBlackTree) afterRemove(node, childNode *Node, dire int) {
	if node == nil || node.isRed() {
		return
	}

	if childNode.isRed() {
		childNode.red() //node 是根节点的情况也包括在内
		return
	}

	parent := node.parent
	var sibling *Node
	if dire == left {
		sibling = parent.right
	} else {
		sibling = parent.left
	}

	deleteLeft := dire == left
	if deleteLeft {
		if sibling.isRed() {
			rotateRR(parent, sibling)
			sibling.black()
			parent.red()
			if sibling.parent == nil {
				rbt.root = sibling
			}
			sibling = parent.right
		}

		if sibling.left.isBlack() && sibling.right.isBlack() {
			pb := parent.isBlack()
			sibling.red()
			parent.black()
			if pb {
				rbt.afterRemove(parent, nil, parent.dire())
			}
		} else {
			if sibling.right.isBlack() {
				rotateLL(sibling, sibling.left)
				sibling = parent.right
			}
			rotateRR(parent, sibling)
			if sibling.parent == nil {
				rbt.root = sibling
			}
			sibling.coloring(parent.Color)
			parent.black()
			sibling.right.black()
		}

	} else {
		if sibling.isRed() {
			rotateLL(parent, sibling)
			sibling.black()
			parent.red()
			if sibling.parent == nil {
				rbt.root = sibling
			}
			sibling = parent.left
		}

		if sibling.left.isBlack() && sibling.right.isBlack() {
			pb := parent.isBlack()
			parent.black()
			sibling.red()
			if pb {
				rbt.afterRemove(parent.parent, nil, parent.dire())
			}
		} else {
			if sibling.left.isBlack() {
				rotateRR(sibling, sibling.right)
				sibling = parent.left
			}
			rotateLL(parent, sibling)
			if sibling.parent == nil {
				rbt.root = sibling
			}
			sibling.coloring(parent.Color)
			parent.black()
			sibling.left.black()
		}
	}
}

func (rbt *RedBlackTree) Add(node *Node) bool {
	if rbt.root == nil {
		rbt.root = node
		rbt.size++
		return true
	}

	//寻找合适的节点添加
	resNode := rbt.root
	for temp := rbt.root; temp != nil; {
		resNode = temp
		flag := node.Compare(node.code, temp.code, temp)
		if flag > 0 {
			temp = temp.right
		} else if flag < 0 {
			temp = temp.left
		} else {
			break
		}
	}

	flag := node.Compare(node.code, resNode.code, resNode)
	if flag > 0 {
		resNode.right = node
		node.parent = resNode
	} else if flag < 0 {
		resNode.left = node
		node.parent = resNode
	} else {
		resNode.v = node.v
		resNode.K = node.K
		resNode.code = node.code
		return false
	}
	rbt.size++
	rbt.addAfter(node)
	return true
	/*for temp := rbt.root; temp != nil; {
		flag := node.Compare(node.code, temp.code, temp)
		if flag > 0 {
			if temp.right == nil {
				temp.right = node
				node.parent = temp
				return
			}
			temp = temp.right
		} else if flag < 0 {
			if temp.left == nil {
				temp.left = node
				node.parent = temp
				return
			}
			temp = temp.left
		} else {
			temp.code = node.code
			temp.k = node.k
			temp.v = node.v
			return
		}
	}*/

}

func (rbt *RedBlackTree) addAfter(node *Node) {
	if node.parent == nil { //node ==root
		rbt.root.black()
		return
	}
	node.red()
	parent := node.parent
	if parent.isBlack() {
		return
	}
	uncle := parent.sibling()
	grand := parent.parent //一定存在
	if uncle.isRed() {
		uncle.black()
		parent.black()
		grand.red()
		rbt.addAfter(grand)
	} else {
		flag := parent.dire()
		if flag == left {
			if node.dire() == right {
				rotateRR(parent, node)
				parent = node
			}
			rotateLL(grand, parent)
			if parent.parent == nil {
				rbt.root = parent
			}
		} else {
			if node.dire() == left {
				rotateLL(parent, node)
				parent = node
			}
			rotateRR(grand, parent)
			if parent.parent == nil {
				rbt.root = parent
			}
		}
		parent.black()
		grand.red()
	}
}

func (rbt *RedBlackTree) ContainK(k int) bool {
	return rbt.node(k) != nil
}

func (rbt *RedBlackTree) node(k int) *Node {
	if rbt.root != nil {
		for temp := rbt.root; temp != nil; {
			flag := rbt.compare(k, temp.K)
			if flag > 0 {
				temp = temp.right
			} else if flag < 0 {
				temp = temp.left
			} else {
				return temp
			}
		}
	}
	return nil
}

func (rbt *RedBlackTree) compare(n1, n2 int) int {
	return n1 - n2
}

func (rbt *RedBlackTree) MidRange(root *Node, visitor func(node *Node)) {
	rbt.MidRange(rbt.root.left, visitor)
	visitor(root)
	rbt.MidRange(root.right, visitor)
}

func (rbt *RedBlackTree) Height() int {
	return rbt.LevelRange(nil)
}

func (rbt *RedBlackTree) LevelRange(visitor func(n *Node)) int {
	queue := make([]*Node, 0)
	queue = append(queue, rbt.root)
	count, level := 0, 1
	for len(queue) > 0 {
		node := queue[0]
		if visitor != nil {
			visitor(node)
		}
		queue = queue[1:]
		level--
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		if level == 0 {
			count++
			level = len(queue)
		}
	}
	return count
}

func (rbt *RedBlackTree) Size() int {
	return rbt.size
}

func (rbt *RedBlackTree) Empty() bool {
	return rbt.Size() == 0
}

//-------------------------------------------------

func rotateLL(g, p *Node) {
	g.left = p.right
	if p.right != nil {
		p.right.parent = g
	}
	p.right = g
	if g.parent != nil {
		if g.dire() == left {
			g.parent.left = p
		} else {
			g.parent.right = p
		}
	}
	p.parent = g.parent
	g.parent = p
}
func rotateRR(g, p *Node) {
	g.right = p.left
	if p.left != nil {
		p.left.parent = g
	}
	p.left = g
	p.parent = g.parent
	if g.parent != nil {
		if g.dire() == left {
			g.parent.left = p
		} else {
			g.parent.right = p
		}
	}
	g.parent = p
}

//--------------------------------------------
//func main() {
//	//fmt.Println(left, black, right, red, unknown)
//
//	//arr := []int{94, 28, 70, 86, 89, 72, 24, 7, 75, 33, 23, 9, 55, 22, 80, 30, 18}
//	//testAdd(arr)
//	testRemove()
//}

func visitor(node *Node) {
	fmt.Print(strconv.Itoa(node.K) + "\t")
}
func visitor1(node *Node) {
	if node.Color == red {
		fmt.Print(strconv.Itoa(node.K) + "\t")
	}
}

func testAdd(arr []int) {
	rbt := NewRedBlackTree()
	for _, v := range arr {
		node := &Node{
			K: v,
			v: v,
		}
		rbt.Add(node)
	}
	rbt.LevelRange(visitor)
	fmt.Println()
	fmt.Println(rbt.Height())
	rbt.LevelRange(visitor1)
}

func testRemove() {
	rbt := NewRedBlackTree()
	for i := 0; i < 16; i++ {
		node := &Node{
			K: i << 4,
			v: i,
		}
		rbt.Add(node)
	}
	rbt.Remove(0)
	rbt.LevelRange(visitor)
	fmt.Println()
	fmt.Println(rbt.Height())
	rbt.LevelRange(visitor1)
}
