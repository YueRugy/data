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

func (rb *RedBlackTree)Add(ele int)  {
	rb.BinarySearchTree.Add(ele)
}

func (rb *RedBlackTree) color(node *Node, color int) *Node {
	if node == nil {
		return node
	}
	node.color = color
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
	return node.color
}
func (rb *RedBlackTree) isBlack(node *Node) bool {
	return rb.colorOf(node) == black
}
func (rb *RedBlackTree) isRed(node *Node) bool {
	return rb.colorOf(node) == red
}
