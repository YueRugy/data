package trie

import "unicode/utf8"

const (
	isWord = true
	noWord = false
)

type Trie struct {
	size int
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		size: 0,
		root: &node{
			word:  noWord,
			value: nil,
			child: make(map[rune]*node),
		},
	}
}

func (t *Trie) Get(s string) *int {
	n := t.node(s)
	if n == nil {
		return nil
	}
	return n.value
}

type node struct {
	word  bool
	value *int
	child map[rune]*node
}

func (t *Trie) Empty() bool {
	return t.Size() == 0
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Clear() {
	t.size = 0
	t.root = nil
}

func (t *Trie) Add(str string, value int) *int {
	if str == "" {
		return nil
	}
	cm := t.root
	for _, v := range str {
		temp := cm.child[v]
		if temp == nil {
			temp = &node{
				word:  noWord,
				value: nil,
				child: make(map[rune]*node),
			}
		}
		cm.child[v] = temp
		cm = temp
	}
	if !cm.word {
		cm.value = &value
		cm.word = isWord
		t.size++
		return nil
	}
	oldValue := cm.value
	cm.value = &value
	return oldValue
}
func (t *Trie) Remove(str string) *int {
	if str == "" {
		return nil
	}
	n := t.node(str)
	if n == nil {
		return nil
	}
	t.size--
	oldValue := n.value
	if len(n.child) > 0 {
		n.value = nil
		n.word = noWord
	}
	length := utf8.RuneCountInString(str)
	count := -1
	cm := t.root
	for index := 0; index < length-1; index++ {
		v, width := utf8.DecodeRuneInString(str)
		str = str[width:]
		cm = cm.child[v]
		if len(cm.child) > 1 || cm.word {
			count = index
		}
	}
	cn := t.root
	for index := 0; index <= count; index++ {
		v, width := utf8.DecodeRuneInString(str)
		str = str[width:]
		cn = cn.child[v]
		if index == count {
			v, _ = utf8.DecodeRuneInString(str)
			cn.child[v] = nil
		}

	}

	return oldValue
}

func (t *Trie) StartWith(str string) bool {
	if str == "" {
		return false
	}

	cm := t.root
	for _, v := range str {
		cm = cm.child[v]
		if cm == nil {
			return false
		}
	}
	return true
}

func (t *Trie) Contains(s string) bool {
	return t.node(s) != nil
}

func (t *Trie) node(str string) *node {
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
	if cm.word {
		return cm
	}
	return nil
}
