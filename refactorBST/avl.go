package refactorBST

import (
	"math"
	"reflect"
)

const (
	maxAbsBalanceFactor = 1
)

type AVL struct {
	*BinarySearchTree
}

func NewAVL() *AVL {
	return &AVL{
		&BinarySearchTree{
			&BinaryTree{
				size: 0,
				root: nil,
			},
		},
	}
}

func (avl *AVL) Add(ele int) {
	node := &Node{
		Ele:    ele,
		height: 1,
	}
	if avl.Empty() {
		avl.root = node
		avl.size++
		avl.AfterAdd(node)
		return
	}

	//寻找合适的节点添加
	resultNode := avl.root //用于保存找到的节点
	for temp := avl.root; temp != nil; {
		resultNode = temp
		if compare(ele, temp.Ele) > 0 {
			temp = temp.right
		} else if compare(ele, temp.Ele) < 0 {
			temp = temp.left
		} else {
			return //相同元素策略 不添加
		}
	}
	if compare(ele, resultNode.Ele) > 0 {
		resultNode.right = node
	} else if compare(ele, resultNode.Ele) < 0 {
		resultNode.left = node
	} else {
		return
	}
	node.parent = resultNode
	avl.size++
	avl.AfterAdd(node)
}

func (avl *AVL) AfterAdd(node *Node) {
	temp := node.parent
	for temp != nil {
		if avl.isBalance(temp) {
			avl.updateHeight(temp)
		} else { //恢复平衡
			avl.reBalance(temp)
			break
		}
		temp = temp.parent
	}
	//avl.updateHeight(node)
}

func (avl *AVL) updateHeight(node *Node) {
	lh, rh := avl.getLhRH(node)
	height := 1 + int(math.Max(float64(lh), float64(rh)))
	node.height = height
}

func (avl *AVL) BalanceFactor(node *Node) int {
	lh, rh := avl.getLhRH(node)
	return lh - rh
}

func (avl *AVL) getLhRH(node *Node) (int, int) {
	lh, rh := 0, 0
	if node.left != nil {
		lh = node.left.height
	}
	if node.right != nil {
		rh = node.right.height
	}
	return lh, rh
}

func (avl *AVL) isBalance(node *Node) bool {
	return int(math.Abs(float64(avl.BalanceFactor(node)))) <= maxAbsBalanceFactor
}

func (avl *AVL) reBalance(grand *Node) {
	parent := avl.taliChild(grand)
	node := avl.taliChild(parent)
	if reflect.DeepEqual(parent, grand.left) {
		if reflect.DeepEqual(node, parent.left) { //LL 右旋 单旋
			avl.LL(grand, parent)
		} else { //LR 先右旋在左旋
			avl.LR(grand, parent, node)
		}
	} else {
		if reflect.DeepEqual(node, parent.right) { //RR 左旋 单旋
			avl.RR(grand, parent)
		} else { //RL
			avl.RL(grand, parent, node)
		}
	}
}

func (avl *AVL) taliChild(p *Node) *Node {
	if p.left == nil {
		return p.right
	} else if p.right == nil {
		return p.left
	}
	if p.left.height > p.right.height {
		return p.left
	} else if p.left.height < p.right.height {
		return p.right
	} else {
		//如果左右高度相等返回和p父节点相同方向的节点
		if reflect.DeepEqual(p, p.parent.left) {
			return p.left
		} else {
			return p.right
		}
	}
}

func (avl *AVL) LL(grand *Node, parent *Node) { //LL 右旋 G 右旋
	grand.left = parent.right
	if parent.right != nil {
		parent.right.parent = grand
	}
	parent.right = grand
	if grand.parent == nil {
		avl.root = parent
		parent.parent = nil
	} else if reflect.DeepEqual(grand, grand.parent.left) {
		parent.parent = grand.parent
		grand.parent.left = parent

	} else if reflect.DeepEqual(grand, grand.parent.right) {
		parent.parent = grand.parent
		grand.parent.right = parent
	}
	grand.parent = parent
	avl.updateHeight(grand)
	avl.updateHeight(parent)
}

func (avl *AVL) RR(grand *Node, parent *Node) { // RR 左旋 G 左旋
	grand.right = parent.left
	if parent.left != nil {
		parent.left.parent = grand
		parent.parent = nil
	}
	parent.left = grand
	//parent.parent = grand.parent
	if grand.parent == nil {
		avl.root = parent
	} else if reflect.DeepEqual(grand, grand.parent.left) {
		parent.parent = grand.parent
		grand.parent.left = parent
	} else if reflect.DeepEqual(grand, grand.parent.right) {
		parent.parent = grand.parent
		grand.parent.right = parent
	}
	grand.parent = parent
	avl.updateHeight(grand)
	avl.updateHeight(parent)
}

func (avl *AVL) LR(grand *Node, parent *Node, node *Node) { //LR 先右旋在左旋
	avl.RR(parent, node)
	avl.LL(grand, node)
}

func (avl *AVL) RL(grand *Node, parent *Node, node *Node) {
	avl.LL(parent, node)
	avl.RR(grand, node)
}
