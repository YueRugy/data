package trie

import "unicode/utf8"

type Node struct {
	word   bool
	k      rune
	v      *int
	father *Node
	child  map[rune]*Node
}

type Tr struct {
	size int
	root *Node
}

func (t *Tr) Get(s string) *int {
	node := t.node(s)
	if node == nil {
		return nil
	}
	return node.v
}

func (t *Tr) Empty() bool {
	return t.size == 0
}

func (t *Tr) Size() int {
	return t.size
}

func (t *Tr) Clear() {
	t.size = 0
	t.root = &Node{
		child: make(map[rune]*Node),
	}
}

func (t *Tr) Add(s string, i int) *int {
	if s == "" {
		return nil
	}
	cm := t.root
	for _, v := range s {
		temp := cm.child[v]
		if temp == nil {
			temp = &Node{
				word:   false,
				k:      v,
				v:      nil,
				father: cm,
				child:  make(map[rune]*Node),
			}
			cm.child[v] = temp
		}
		cm = temp
	}

	if cm.word {
		oldValue := cm.v
		cm.v = &i
		return oldValue
	}
	t.size++
	cm.v = &i
	cm.word = isWord

	return &i
}

func (t *Tr) Remove(s string) *int {
	if s == "" {
		return nil
	}
	node := t.node(s)
	if node == nil {
		return nil
	}

	if len(node.child) > 0 {
		oldValue := node.v
		node.v = nil
		node.word = noWord
		return oldValue
	}
	length := utf8.RuneCountInString(s)
	n := node
	for index, size := length, len(s); index >= 1; index-- {
		v, width := utf8.DecodeLastRuneInString(s)
		s = s[:size-width]
		size -= width
		if index == length {
			if len(n.child) > 1 {
				break
			}
		} else {
			if n.word || len(n.child) > 1 {
				break
			}
		}
		n.father.child[v] = nil
		n = n.father
	}
	return n.v
}

func (t *Tr) StartWith(s string) bool {
	if s == "" {
		return false
	}
	cm := t.root
	for _, v := range s {
		cm = cm.child[v]
		if cm == nil {
			return false
		}
	}
	return true
}

func (t *Tr) Contains(s string) bool {
	return t.node(s) != nil
}

func (t *Tr) node(str string) *Node {
	if str == "" {
		return nil
	}
	cm := t.root
	for _, v := range str {
		cm = cm.child[v]
		if cm == nil {
			return nil
		}
	}
	if !cm.word {
		return nil
	}

	return cm
}

func NewTr() *Tr {
	return &Tr{
		size: 0,
		root: &Node{
			word:   false,
			k:      0,
			v:      nil,
			father: nil,
			child:  make(map[rune]*Node),
		},
	}
}
