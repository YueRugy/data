package trie

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

var trie *Trie

func init() {
	trie = NewTrie()
	trie.Add("doggy", 1)
	trie.Add("dog", 2)
	trie.Add("deh", 6)
	trie.Add("cat", 3)
	trie.Add("catalog", 4)
	trie.Add("cast", 7)
	trie.Add("岳伟超", 8)
}

func TestTrie_Remove(t *testing.T) {
	fmt.Println(trie.Remove("catalog"))
}

func TestTrie_Add(t *testing.T) {
	tr := NewTrie()
	tr.Add("doggy", 1)
	tr.Add("dog", 2)
	tr.Add("cat", 3)
}

func TestTrie_StartWith(t *testing.T) {
	if trie.StartWith("dog") {
		t.Log("success")
	} else {
		t.Error("failed")
	}

	if !trie.StartWith("doge") {
		t.Log("success")
	} else {
		t.Error("failed")
	}

}

func TestTrie_Contains(t *testing.T) {
	if trie.Contains("岳伟超") {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	if !trie.Contains("hello") {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestTrie_Get(t *testing.T) {
	if *trie.Get("dog") == 2 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestTrie_Size(t *testing.T) {
	if trie.Size() == 4 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func TestTrie_Empty(t *testing.T) {
	if !trie.Empty() {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}
func Test_Str(t *testing.T) {
	str := "hello岳伟超"

	fmt.Println(utf8.RuneCountInString(str))
	for _, v := range str {
		fmt.Printf("%c\t", v)
	}
}
