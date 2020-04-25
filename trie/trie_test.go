package trie

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestTrie_Size(t *testing.T) {

}
func TestTrie_Empty(t *testing.T) {

}
func Test_Str(t *testing.T) {
	str := "hello岳伟超"
	fmt.Println(utf8.RuneCountInString(str))
	for _, v := range str {
		fmt.Printf("%c\t",v)
	}
}
