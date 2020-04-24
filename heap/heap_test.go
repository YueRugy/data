package heap

import (
	"fmt"
	"reflect"
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
	array1 := []int{6, 29, 26, 51, 34, 54, 28, 86, 80, 79, 65, 58, 85, 93, 39, 95}
	fmt.Println(heap.array)
	if reflect.DeepEqual(heap.array, array1) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	//flag := true
	//for index := 0; index < len(array); index++ {
	//	if heap.array[index] != array1[index] {
	//		flag = false
	//	}
	//}
	//if flag {
	//	t.Log("success")
	//} else {
	//	t.Error("failed")
	//}
}
