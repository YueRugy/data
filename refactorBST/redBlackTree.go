package refactorBST

type RedBlackTree struct {
	*BinarySearchTree
}

const (
	black = iota
	red
)

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		&BinarySearchTree{
			&BinaryTree{
				size: 0,
				root: nil,
			},
		},
	}
}

func (rb *RedBlackTree) Add(ele int) {
	node := rb.BinarySearchTree.Add(ele)
	rb.AddAfter(node)
}

func (rb *RedBlackTree) color(node *Node, color int) *Node {
	if node == nil {
		return node
	}
	node.Color = color
	return node
}
func (rb *RedBlackTree) red(node *Node) {
	rb.color(node, red)
}
func (rb *RedBlackTree) black(node *Node) {
	rb.color(node, black)
}
func (rb *RedBlackTree) colorOf(node *Node) int {
	if node == nil {
		return black
	}
	return node.Color
}
func (rb *RedBlackTree) isBlack(node *Node) bool {
	return rb.colorOf(node) == black
}
func (rb *RedBlackTree) isRed(node *Node) bool {
	return rb.colorOf(node) == red
}

func (rb *RedBlackTree) AddAfter(node *Node) {
	if node == nil { //相同元素不添加
		return
	}
	if node.parent == nil { //root 节点 染黑即可
		rb.black(node)
		return
	}
	//默认是红色
	rb.red(node)
	//如果父节点是黑色直接添加即可
	if rb.isBlack(node.parent) {
		return
	} else if rb.isRed(node) { //double red 需要处理
		if rb.isBlack(node.parent.sibling()) { //如果uncle节点==black uncle节点不存在,在红黑树里颜色是black 。这一步祖父节点必然不为nil
			rb.process(node)
		} else { //叔父节点是红色存在叔父节点
			//uncle := node.sibling()
			//父节点叔父节点染黑
			rb.black(node.parent.sibling())
			rb.black(node.parent)
			rb.red(node.parent.parent)
			rb.AddAfter(node.parent.parent)
		}
	}
}

func (rb *RedBlackTree) process(node *Node) {
	direHandle := manner(node.parent.dire(), node.dire())
	switch direHandle {
	case ll:
		rb.ll(node.parent, node.parent.parent)
	case rr:
		rb.rr(node.parent, node.parent.parent)
	case lr:
		rb.lr(node, node.parent, node.parent.parent)
	case rl:
		rb.rl(node, node.parent, node.parent.parent)

	}
}

func (rb *RedBlackTree) ll(parent *Node, grand *Node) {
	rb.black(parent)
	rb.red(grand)
	parent.right = grand
	grand.left = nil
	if grand.parent == nil {
		rb.root = parent
		parent.parent = nil
	} else {
		if grand.dire() == left {
			grand.parent.left = parent
		} else {
			grand.parent.right = parent
		}
		parent.parent = grand.parent
	}
	grand.parent = parent
	/*parent.right = grand
	if parent.parent == nil {
		rb.root = parent
	} else {
		parent.parent = grand.parent
		if grand.dire() == left {
			parent.parent.left = parent
		} else {
			parent.parent.right = parent
		}
	}
	grand.parent = parent*/
}

func (rb *RedBlackTree) rr(parent *Node, grand *Node) {
	rb.red(grand)
	rb.black(parent)
	parent.left = grand
	grand.right = nil
	if grand.parent == nil {
		rb.root = parent
		parent.parent = nil
	} else {
		parent.parent = grand.parent
		if grand.dire() == left {
			grand.parent.left = parent
		} else {
			grand.parent.right = parent
		}
	}
	grand.parent = parent
}

func (rb *RedBlackTree) lr(node *Node, parent *Node, grand *Node) {

	parent.right = nil
	node.left = parent
	grand.left = node
	node.parent = grand
	parent.parent = node
	//rb.rr(node, parent)
	rb.ll(node, grand)
}

func (rb *RedBlackTree) rl(node *Node, parent *Node, grand *Node) {
	//rb.ll(node, parent)
	parent.left = nil
	node.right = parent
	grand.right = node
	node.parent = grand
	parent.parent = node
	rb.rr(node, grand)
}
