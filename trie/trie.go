package trie

type Trie struct {
	size int
	root *node
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

func (t *Trie) Add(s string, i int) {
	panic("implement me")
}

func (t *Trie) Remove(s string) *int {
	panic("implement me")
}

func (t *Trie) StartWith(s string) bool {
	panic("implement me")
}

func (t *Trie) Contains(s string) bool {
	return t.node(s) != nil
}

func (t *Trie) node(str string) *node {
	if t.root == nil || str == "" {
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
