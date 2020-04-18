package refactorBST

import (
	"math"
	"reflect"
)

const (
	ll = iota
	rr
	lr
	rl
	max = 1
)

const (
	left = iota
	right
)

type avl1 struct {
	*BinarySearchTree
}

func NewAVL1() *avl1 {
	return &avl1{
		&BinarySearchTree{
			&BinaryTree{
				size: 0,
				root: nil,
			},
		},
	}
}

func (avl *avl1) Add(ele int) {
	node := avl.BinarySearchTree.Add(ele)

	avl.AfterAdd(node)
	/*node := &Node{
		ele:    ele,
		height: 1,
	}

	if avl.Empty() {
		avl.root = node
		avl.size++
		return
	}
	//寻找合适的节点添加
	resultNode := avl.root //用于保存找到的节点
	for temp := avl.root; temp != nil; {
		resultNode = temp
		if compare(ele, temp.ele) > 0 {
			temp = temp.right
		} else if compare(ele, temp.ele) < 0 {
			temp = temp.left
		} else {
			return //相同元素策略 不添加
		}
	}
	if compare(ele, resultNode.ele) > 0 {
		resultNode.right = node
	} else if compare(ele, resultNode.ele) < 0 {
		resultNode.left = node
	} else {
		return
	}
	node.parent = resultNode
	avl.size++
	avl.AfterAdd(node)*/
}

func (avl *avl1) AfterAdd(node *Node) {
	if node == nil {
		return
	}
	temp := node.parent
	node.height = 1

	for temp != nil {
		if avl.isBalance(temp) { //是否平衡 平衡 更新高度
			avl.updateHeight(temp)
		} else {
			avl.reBalance(temp)
			break
		}
		temp = temp.parent
	}
}

func (avl *avl1) Remove(ele int) {
	node := avl.BinarySearchTree.Remove(ele)
	avl.removeAfter(node)
}

//左右子树的高度绝对值是否大于max
func (avl *avl1) isBalance(node *Node) bool {
	lh, rh := avl.getHeight(node)
	return int(math.Abs(float64(lh)-float64(rh))) <= max
}

func (avl *avl1) getHeight(node *Node) (int, int) {
	lh, rh := 0, 0
	if node.left != nil {
		lh = node.left.height
	}
	if node.right != nil {
		rh = node.right.height
	}
	return lh, rh
}

func (avl *avl1) updateHeight(node *Node) {
	lh, rh := avl.getHeight(node)
	node.height = 1 + int(math.Max(float64(lh), float64(rh)))
}

func (avl *avl1) reBalance(grand *Node) {
	parent, pDire := avl.tailChild(grand)
	node, nDire := avl.tailChild(parent)
	manner := avl.manner(pDire, nDire)
	switch manner {
	case ll:
		avl.ll(grand, parent)
	case rr:
		avl.rr(grand, parent)
	case lr:
		avl.lr(grand, parent, node)
	case rl:
		avl.rl(grand, parent, node)

	}
}

func (avl *avl1) tailChild(node *Node) (*Node, int) {
	if node.left == nil {
		return node.right, right
	}

	if node.right == nil {
		return node.left, left
	}

	if node.left.height > node.right.height {
		return node.left, left
	} else if node.right.height > node.left.height {
		return node.right, right
	} else { //如果左右子树高度相同去与父树同方向的子树节点
		if avl.dire(node) == left {
			return node.parent.left, left
		} else {
			return node.parent.right, right
		}
	}
}

func (avl *avl1) dire(node *Node) int {
	if node != nil && node.parent != nil {
		if reflect.DeepEqual(node, node.parent.left) {
			return left
		} else {
			return right
		}
	}
	return left
}

func (avl *avl1) manner(dire int, dire2 int) int {
	if dire == left {
		if dire2 == left {
			return ll
		} else {
			return lr
		}
	} else {
		if dire2 == right {
			return rr
		} else {
			return rl
		}
	}
}

func (avl *avl1) ll(grand *Node, parent *Node) {
	grand.left = parent.right
	if parent.right != nil {
		parent.right.parent = grand
	}

	parent.right = grand
	avl.execute(grand, parent)
}

/*func (avl *avl1) execute(grand *Node, parent *Node) {
	if grand.parent == nil {
		avl.root = parent
		parent.parent = nil
	} else {
		if avl.dire(grand) == left {
			grand.parent.left = parent
		} else {
			grand.parent.right = parent
		}
		parent.parent = grand.parent
	}
	grand.parent = parent
	avl.updateHeight(grand)
	avl.updateHeight(parent)
}*/

func (avl *avl1) rr(grand *Node, parent *Node) {
	grand.right = parent.left
	if parent.left != nil {
		parent.left = grand
	}
	parent.left = grand
	avl.execute(grand, parent)
}

func (avl *avl1) execute(grand *Node, parent *Node) {
	if grand.parent == nil {
		avl.root = parent
		parent.parent = nil
	} else {
		if avl.dire(grand) == left {
			grand.parent.left = parent
		} else {
			grand.parent.right = parent
		}
		parent.parent = grand.parent
	}
	grand.parent = parent
	avl.updateHeight(parent)
	avl.updateHeight(grand)
}

func (avl *avl1) lr(grand *Node, parent *Node, node *Node) {
	avl.rr(parent, node)
	avl.ll(grand, node)
}

func (avl *avl1) rl(grand *Node, parent *Node, node *Node) {
	avl.ll(parent, node)
	avl.rr(grand, node)
}

func (avl *avl1) removeAfter(node *Node) {
	temp := node.parent
	for temp != nil {
		if avl.isBalance(temp) {
			avl.updateHeight(temp)
		} else {
			avl.reBalance(temp)
		}
		temp = temp.parent
	}
}
