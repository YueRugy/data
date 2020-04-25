package trie

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
func (t *Trie) Remove(s string) *int {
	panic("implement me")
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
