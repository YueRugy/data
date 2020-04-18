package refactorBST

import (
	"fmt"
	"math"
	"reflect"
)

type BinaryTree struct {
	size int
	root *Node
}
type Node struct {
	ele    int
	height int
	color  int
	left   *Node
	right  *Node
	parent *Node
}

//返回兄弟节点
func (node *Node) sibling() *Node {
	if node == nil || node.parent == nil {
		return nil
	}
	if reflect.DeepEqual(node, node.parent.left) {
		return node.parent.right
	} else {
		return node.parent.left
	}
}

func (bt *BinaryTree) Size() int {
	return bt.size
}

func (bt *BinaryTree) Empty() bool {
	return bt.Size() == 0
}

func (bt *BinaryTree) PreRange(visitor func(int)) {
	PreRange(bt.root, visitor)
}
func (bt BinaryTree) PostRange(visitor func(int)) {
	PostRange(bt.root, visitor)
}
func (bt *BinaryTree) MidRange(visitor func(int)) {
	MidRange(bt.root, visitor)
}
func (bt *BinaryTree) LevelRange(visitor func(int)) {
	LevelRange(bt.root, visitor)
}

func (bt *BinaryTree) Height() int {
	return Height(bt.root)
}
func (bt *BinaryTree) HeightByLevel() int {
	return HeightByLevel(bt.root)
}

//使用中序遍历 寻找前序节点
func (bt *BinaryTree) Predecessor(node *Node) *Node {
	if node == nil {
		return node //节点==nil return nil
	}
	// node.left.right.right..... 最后一个right 就是前序节点
	if node.left != nil {
		temp := node.left
		for temp.right != nil {
			temp = temp.right
		}
		return temp
	}
	//node.parent.parent.... 直到某个节点是父节点的右节点
	if node.left == nil && node.parent != nil {
		temp := node
		for temp.parent != nil {
			if reflect.DeepEqual(temp, temp.parent.right) {
				return temp.parent
			}
			temp = temp.parent
		}
	}
	return nil
}

//使用中序遍历寻找后续节点
func (bt *BinaryTree) Successor(node *Node) *Node {
	if node == nil {
		return node
	}
	// node.right.left.left....
	if node.right != nil {
		temp := node.right
		for temp.left != nil {
			temp = temp.left
		}
		return temp
	}

	if node.right == nil && node.parent != nil {
		temp := node
		for temp.parent != nil {
			if reflect.DeepEqual(temp, temp.parent.left) {
				return temp.parent
			}
			temp = temp.parent
		}
	}

	return nil
}

//使用层序遍历判断是不是完全二叉树
func (bt *BinaryTree) Complete() bool {
	if bt.root == nil {
		return false
	}
	queue := make([]*Node, 0)
	queue = append(queue, bt.root)
	leaf := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if leaf {
			if node.left != nil || node.right != nil {
				return false
			}
		}
		if node.left == nil && node.right == nil { //遍历到叶子节点后面的也是叶子节点
			leaf = true
		} else if node.left != nil && node.right != nil { //加入队列
			queue = append(queue, node.left, node.right)
		} else if node.left == nil && node.right != nil {
			return false
		} else {
			leaf = true
			queue = append(queue, node.left)
		}

	}
	return true
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

func Height(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(Height(node.left)), float64(Height(node.right))))
}

//前序遍历
func PreRange(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	visitor(node.ele)
	PreRange(node.left, visitor)
	PreRange(node.right, visitor)
}

//后序遍历
func PostRange(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	PostRange(node.left, visitor)
	PostRange(node.right, visitor)
	visitor(node.ele)
}

//中序遍历
func MidRange(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	MidRange(node.left, visitor)
	visitor(node.ele)
	MidRange(node.right, visitor)
}

//层序遍历
func LevelRange(node *Node, visitor func(int)) {
	if node == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		n := queue[0]
		visitor(n.ele)
		queue = queue[1:]
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

type BinarySearchTree struct {
	*BinaryTree
}

func NewBst() *BinarySearchTree {
	return &BinarySearchTree{
		&BinaryTree{
			size: 0,
			root: nil,
		},
	}
}

func (bst *BinarySearchTree) Add(ele int) *Node {
	node := &Node{
		ele: ele,
	}

	if bst.Empty() {
		bst.root = node
		bst.size++
		return node
	}
	//寻找合适的节点添加
	resultNode := bst.root //用于保存找到的节点
	for temp := bst.root; temp != nil; {
		resultNode = temp
		if compare(ele, temp.ele) > 0 {
			temp = temp.right
		} else if compare(ele, temp.ele) < 0 {
			temp = temp.left
		} else {
			return nil //相同元素策略 不添加
		}
	}
	if compare(ele, resultNode.ele) > 0 {
		resultNode.right = node
	} else if compare(ele, resultNode.ele) < 0 {
		resultNode.left = node
	} else {
		return nil
	}
	node.parent = resultNode
	bst.size++
	return node
}

func (bst *BinarySearchTree) Contains(ele int) bool {
	//node := bst.root
	//for node != nil {
	//	if compare(ele, node.ele) > 0 {
	//		node = node.right
	//	} else if compare(ele, node.ele) < 0 {
	//		node = node.left
	//	} else {
	//		return true
	//	}
	//}
	//return false
	return bst.node(ele) != nil
}

func (bst *BinarySearchTree) Remove(ele int) *Node {
	if bst.root == nil {
		return nil
	}
	resNode := bst.node(ele)
	if resNode == nil {
		return nil
	}
	//度为2的情况 前继节点度==1||==0 left.right.right if left==nil parent.parent
	var replaceNode *Node
	if resNode.left != nil && resNode.right != nil {
		replaceNode = bst.Predecessor(resNode)
		resNode.ele = replaceNode.ele //复值 删除前继节点
		resNode = replaceNode
	}
	//度==1
	if resNode.left != nil && resNode.right == nil {
		parent := resNode.parent
		if parent == nil {
			bst.root = resNode.left
		} else {
			if reflect.DeepEqual(parent.left, resNode) {
				parent.left = resNode.left
			} else {
				parent.right = resNode.left
			}
			resNode.parent = parent
		}
	}

	if resNode.left == nil && resNode.right != nil {
		parent := resNode.parent
		if parent == nil {
			bst.root = resNode.right
		} else {
			if reflect.DeepEqual(resNode, parent.left) {
				parent.left = resNode.right
			} else {
				parent.right = resNode.right
			}
		}
	}

	if resNode.right == nil && resNode.left == nil {
		parent := resNode.parent
		if parent == nil {
			bst.root = nil
		} else {
			if reflect.DeepEqual(resNode, parent.left) {
				parent.left = nil
			} else {
				parent.right = nil
			}
		}
	}
	bst.size--
	return resNode
}

func (bst *BinarySearchTree) node(ele int) *Node {
	if bst.root == nil {
		return nil
	}
	node := bst.root
	for node != nil {
		if compare(ele, node.ele) > 0 {
			node = node.right
		} else if compare(ele, node.ele) < 0 {
			node = node.left
		} else {
			break
		}
	}
	return node
}

func compare(c1, c2 int) int {
	return c1 - c2
}
func main() {

	//bst.Remove(10)
	fmt.Println()

	//fmt.Println(bst.Contains(8))
	//for _, v := range arr {
	//	fmt.Print(bst.Contains(v))
	//	fmt.Print("\t")
	//}
	//fmt.Println()
	//fmt.Println(bst.Contains(16)) //false
	//fmt.Println(bst.Complete())
	//fmt.Println(bst.Successor(bst.root).ele)            //8
	//fmt.Println(bst.Successor(bst.root.left.right).ele) //7
	//fmt.Println(bst.Height())
	//fmt.Println(bst.HeightByLevel())
	//fmt.Println(bst.Predecessor(bst.root).ele)            //5
	//fmt.Println(bst.Predecessor(bst.root.right.left).ele) //7
	//	visitor := func(ele int) {
	//		fmt.Print("_" + strconv.Itoa(ele) + "_\t")
	//	}
	//	bst.PreRange(visitor)
	//	fmt.Println("------------------")
	//	bst.PostRange(visitor)
	//	fmt.Println("------------------")
	//	bst.MidRange(visitor)
	//	fmt.Println("------------------")
	//	bst.LevelRange(visitor)

}
