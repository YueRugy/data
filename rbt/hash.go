package rbt

import (
	"errors"
	"fmt"
)

type Hash struct {
	size   int
	bucket []*RedBlackTree
}

const (
	defaultCap = 2 << 4
)

func NewHash() *Hash {
	return &Hash{
		size:   0,
		bucket: make([]*RedBlackTree, defaultCap),
	}
}

func (hash *Hash) ContainsKey(k int) bool {
	if hash == nil {
		return false
	}
	code := hashcode(k)
	index := hash.index(code)
	rbt := hash.bucket[index]
	if rbt == nil {
		return false
	}
	node := rbt.node(k)
	fmt.Println(node)
	return rbt.node(k) != nil
}

func (hash *Hash) Get(k int) (int, error) {
	if hash == nil {
		return 0, errors.New("hash is nil")
	}
	code := hashcode(k)
	index := hash.index(code)
	rbt := hash.bucket[index]
	if rbt == nil {
		return 0, nil //暂时用-1代替
	}
	node := rbt.node(k)
	if node == nil {
		return 0, nil
	}
	return rbt.node(k).v, nil
}

func (hash *Hash) Put(k, v int) {
	code := hashcode(k)
	index := hash.index(code)
	node := &Node{
		K:    k,
		v:    v,
		code: code,
	}
	rbt := hash.bucket[index]
	if rbt == nil {
		rbt = NewRedBlackTree()
		hash.bucket[index] = rbt
	}
	success := rbt.Add(node)
	if success {
		hash.size++
	}
}
func (hash *Hash) Remove(k int) {
	if hash == nil {
		return
	}
	code := hashcode(k)
	index := hash.index(code)
	rbt := hash.bucket[index]
	if rbt == nil {
		return
	}
	success := rbt.Remove(k)
	if success {
		hash.size--
	}
}

func (hash *Hash) Size() int {
	return hash.size
}

func (hash *Hash) index(code int) int {
	return code & (hash.size - 1)
}

func (hash *Hash) Empty() bool {
	return hash.Size() == 0
}

func hashcode(num int) int {
	return num
	//todo
}
