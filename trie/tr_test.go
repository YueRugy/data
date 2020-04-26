package trie

import "testing"

var tr *Tr

func init() {
	tr = NewTr()
	tr.Add("doggy", 1)
	tr.Add("dog", 2)
	tr.Add("deh", 6)
	tr.Add("cat", 3)
	tr.Add("catalog", 4)
	tr.Add("cast", 7)
	tr.Add("岳伟超", 8)
}
func TestTr_Add(t *testing.T) {

}
func TestTr_Size(t *testing.T) {
	size := tr.Size()
	if size == 7 {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestTr_Empty(t *testing.T) {

	if !tr.Empty() {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestTr_Get(t *testing.T) {
	t.Log(*(tr.Get("catalog")))
	t.Log(*(tr.Get("cat")))
	t.Log(*(tr.Get("cast")))
}

func TestTr_Contains(t *testing.T) {

	if tr.Contains("cat") {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestTr_StartWith(t *testing.T) {
	if tr.StartWith("ca") {
		t.Log("success")
	} else {
		t.Error("failed")
	}

}
func TestTr_Remove(t *testing.T) {
	tr.Remove("catalog")
	t.Log(tr.Contains("cat"))
	tr.Remove("cast")
	t.Log(tr.Contains("cast"))
	t.Log(tr.Contains("cat"))
	tr.Remove("岳伟超")
}
