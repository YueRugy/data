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
	} else if rb.isRed(node.parent) { //double red 需要处理
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

func (rb *RedBlackTree) Remove(ele int) {
	node, chNode := rb.BinarySearchTree.Remove(ele)
	rb.afterRemove(node, chNode)
}

func (rb *RedBlackTree) afterRemove(node *Node, chNode *Node) {

	if node == nil || rb.isRed(node) { //没有找到这个元素或者root==nil ||node 根节点||自身是红色节点
		return
	}
	if rb.isRed(chNode) {
		rb.black(chNode) //如果替代的节点的子节点是red 染黑即可
		return
	}
	parent := node.parent
	var sibling *Node
	deleteLeft := parent.left == nil
	if deleteLeft {
		sibling = parent.right
	} else {
		sibling = parent.left
	}

	if deleteLeft { //被删除的节点在左边
		if rb.isRed(sibling) {
			rotateRR(parent, sibling)
			rb.black(sibling)
			rb.red(parent)
			if sibling.parent == nil {
				rb.root = sibling
			}
			sibling = parent.right
		}

		if rb.isBlack(sibling.right) && rb.isBlack(sibling.left) {
			flag := rb.isBlack(parent)
			rb.black(parent)
			rb.red(sibling)
			if flag {
				rb.afterRemove(parent, nil)
			}
		} else {
			if rb.isBlack(sibling.right) {
				rotateLL(sibling, sibling.left)
				sibling = parent.left
			}

			rotateRR(parent, sibling)
			if sibling.parent == nil {
				rb.root = sibling
			}
			rb.color(sibling, rb.colorOf(parent))
			rb.black(parent)
			rb.black(sibling.right)
		}

	} else {                   //被删除的节点在右边
		if rb.isRed(sibling) { //兄弟节点是红色
			rotateLL(parent, sibling)
			rb.black(sibling)
			rb.red(parent)
			if sibling.parent == nil {
				rb.root = sibling
			}
			sibling = parent.left
		}

		//兄弟节点必然是黑色
		//兄弟节点不可以借
		if rb.isBlack(sibling.left) && rb.isBlack(sibling.right) { //下溢 合并
			pBlock := rb.isBlack(parent)
			rb.black(parent)
			rb.red(sibling)
			if pBlock {
				rb.afterRemove(parent, nil)
			}
		} else { //兄弟节点必然有一个红色节点
			if rb.isBlack(sibling.left) {
				rotateRR(sibling, sibling.right)
				sibling = parent.left
			}

			rotateLL(parent, sibling)
			if sibling.parent == nil {
				rb.root = sibling
			}
			rb.color(sibling, rb.colorOf(parent)) //旋转后 兄弟节点染成同原来父节点一样颜色
			rb.black(parent)
			rb.black(sibling.left)
		}

		/*if rb.isBlack(sibling.left) {
			rotateRR(sibling, sibling.right)
		}*/

	}

}

//node 度只可能是1或0 红黑树的删除只可能在度=1或者0 参考二叉树的中序遍历 if left!=nil left.right.right..... else parent.parent 知道parent.parent.left=parent..left
func (rb *RedBlackTree) afterRemove1(node *Node, chNode *Node) {
	if node == nil || node.parent == nil || rb.isRed(node) { //没有找到这个元素或者root==nil ||node 根节点||自身是红色节点
		return
	}

	if rb.isRed(chNode) {
		rb.black(chNode) //如果替代的节点的子节点是red 染黑即可
	} else { //下溢删除的是黑色节点所以一定有个兄弟节点（rbTree）是黑色
		//兄弟节点(b tree)能够借给我元素 红黑树逻辑上转换成b树 将黑色节点和它的红色节点合成一个超级节点
		//所以兄弟节点(二叉树）的颜色必然是黑色
		parent := node.parent
		if sibling := parent.left; sibling != nil {
			if rb.isBlack(sibling) {
				if nephew := sibling.left; nephew != nil { //ll 先判断同方向的一遍可能值需要一次旋转就可以
					parent.left = sibling.right
					if sibling.right != nil {
						sibling.right.parent = parent
					}
					sibling.right = parent
					sibling.parent = parent.parent
					parent.parent = sibling
					if sibling.parent == nil { //旋转可能导致根节点变化
						rb.root = sibling
					}
					sibling.Color = parent.Color
					rb.black(parent)
					rb.black(nephew)
				} else if nephew := sibling.right; nephew != nil && sibling.left == nil { //lr
					nephew.parent = parent //rr
					parent.left = nephew
					nephew.left = sibling
					sibling.parent = nephew
					//ll
					nephew.parent = parent.parent
					if nephew.parent == nil {
						rb.root = nephew
					}
					nephew.right = parent
					parent.parent = nephew
					nephew.Color = parent.Color
					rb.black(parent)
					rb.black(sibling)
				} else {
					if rb.isRed(parent) {
						rb.black(parent)
						rb.red(sibling)
					} else {
						rb.black(parent)
						rb.red(sibling)
						rb.afterRemove(parent, sibling)
					}
				}
			} else {
				//parent ll
				parent.left = sibling.right
				sibling.right.parent = parent
				sibling.right = nil
				sibling.right = parent
			}
		} else if sibling := parent.right; sibling != nil {
			if rb.isBlack(sibling) {
				if nephew := sibling.right; nephew != nil { //rr先判断同方向的一遍可能值需要一次旋转就可以
					parent.right = sibling.left
					if sibling.left != nil {
						sibling.left.parent = parent
					}
					sibling.left = parent
					sibling.parent = parent.parent
					parent.parent = sibling
					if sibling.parent == nil { //旋转可能导致根节点变化
						rb.root = sibling
					}
					sibling.Color = parent.Color
					rb.black(parent)
					rb.black(nephew)
				} else if nephew := sibling.left; nephew != nil && sibling.right == nil { //rl
					nephew.parent = parent //rr
					parent.right = nephew
					nephew.right = sibling
					sibling.parent = nephew
					//ll
					nephew.parent = parent.parent
					if nephew.parent == nil {
						rb.root = nephew
					}
					nephew.left = parent
					parent.parent = nephew
					nephew.Color = parent.Color
					rb.black(parent)
					rb.black(sibling)
				} else {
					if rb.isRed(parent) {
						rb.black(parent)
						rb.red(sibling)
					} else {
						rb.black(parent)
						rb.red(sibling)
						rb.afterRemove(parent, sibling)
					}
				}
			} else {

			}
		}
	}

}

/*
func (rb *RedBlackTree) Remove(ele int) {
	node := rb.node(ele)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		replaceNode := rb.Predecessor(node)
		if replaceNode == nil {
			return
		}
		node.Ele = replaceNode.Ele
		node = replaceNode
	}

	if node.left != nil && node.right == nil {
		if node.parent == nil {
			rb.root = node.left
			rb.black(node.left)
			node.left.parent = nil
		} else {
			if node.dire() == left {
				node.parent.left = node.left
			} else {
				node.parent.right = node.left
			}
			node.left.parent = node.parent
			rb.black(node.left)
		}
	} else if node.right != nil && node.left == nil {
		if node.parent == nil {
			rb.root = node.right
			rb.black(node.right)
			node.right.parent = nil
		} else {
			if node.dire() == left {
				node.parent.left = node.right
			} else {
				node.parent.right = node.right
			}
			node.right.parent = node.parent
			rb.black(node.right)
		}
	} else if node.right == nil && node.left == nil {
		if node.parent == nil {
			rb.root = nil
			return
		} else {
			if rb.isRed(node) {
				if node.dire() == left {
					node.parent.left = nil
				} else {
					node.parent.right = nil
				}
			} else {

			}
		}
	}

}*/
func rotateLL(p, c *Node) {
	p.left = c.right
	if c.right != nil {
		c.right.parent = p
	}
	c.right = p
	c.parent = p.parent
	if p.parent != nil {
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
	c.parent = p.parent
	if p.parent != nil {
		flag := p.dire() == left
		if flag {
			p.parent.left = c
		} else {
			p.parent.right = c
		}
	}

	p.parent = c
}
