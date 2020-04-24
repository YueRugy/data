package heap

import (
	"testing"
)

func TestHeap_Size(t *testing.T) {
	var heap *Heap
	num := heap.Size()
	if num == 0 {
		t.Log("success")
	} else {
		t.Log("error")
	}
}
func TestHeap_Add(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	heap := NewHeap()
	for _, v := range array {
		heap.Add(v)
	}
	t.Log(heap.array)
}
