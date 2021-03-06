package rbt

import (
	"fmt"
	"testing"
)

func visitorKV(k, v int) {
	fmt.Printf("k:%d  v:%d   ", k, v)
}

//func visitor(node *Node) {
//	fmt.Print(strconv.Itoa(node.K) + "\t")
//}
//func visitor1(node *Node) {
//	if node.Color == red {
//		fmt.Print(strconv.Itoa(node.K) + "\t")
//	}
//}
func TestHash_Put(t *testing.T) {
	hash := NewHash()
	for i := 0; i < 8; i++ {
		hash.Put(i<<4, i)
	}
	if len(hash.bucket) == defaultCap {
		t.Log("success")
	} else {
		t.Error("err")
	}
	for i := 8; i < 16; i++ {
		hash.Put(i<<4, i)
	}
	if len(hash.bucket) == (defaultCap << 1) {
		t.Log("success")
	} else {
		t.Error("err")
	}
	//rbt := hash.bucket[0]
	//rbt.LevelRange(visitor)
	//fmt.Println()
	//fmt.Println(rbt.Height())
	//rbt.LevelRange(visitor1)
	//fmt.Println()
	//fmt.Println(hash.Size())
	//hash.Put(0, 123)
	//fmt.Println(hash.Size())
}
func TestHash_Empty(t *testing.T) {

	//fmt.Println(1 << 1)
}
func TestHash_Size(t *testing.T) {

}

func TestHash_ContainsKey(t *testing.T) {
	hash := NewHash()
	for i := 0; i < 16; i++ {
		hash.Put(i<<4, i)
	}
	if hash.ContainsKey(16) == false {
		t.Log("success")
	}
	if hash.ContainsKey(0) == true {
		t.Log("success")
	}
}

func TestHash_Remove(t *testing.T) {
	hash := NewHash()
	for i := 0; i < 16; i++ {
		hash.Put(i<<4, i)
	}
	hash.Remove(0)
	rbt := hash.bucket[0]
	rbt.LevelRange(visitor)
	fmt.Println()
	fmt.Println(rbt.Height())
	rbt.LevelRange(visitor1)
	fmt.Println()
}

func TestHash_Get(t *testing.T) {

	hash := NewHash()
	for i := 0; i < 16; i++ {
		hash.Put(i<<4, i)
	}
	v1, err := hash.Get(16)
	if err == nil && v1 == 1 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestHash_ContainsValue(t *testing.T) {
	hash := NewHash()
	for i := 0; i < 16; i++ {
		hash.Put(i<<4, i)
	}
	b1 := hash.ContainsValue(1)
	b2 := hash.ContainsValue(178)
	if b1 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	if !b2 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestHash_Traversal(t *testing.T) {
	hash := NewHash()
	for i := 0; i < 16; i++ {
		hash.Put(i<<4, i)
	}
	hash.Traversal(visitorKV)
}
