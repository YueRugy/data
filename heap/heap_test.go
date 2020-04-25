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
func TestHeap_Remove(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	heap := NewHeap()
	for _, v := range array {
		heap.Add(v)
	}
	heap.Remove()
	fmt.Println(heap.array)
}
func TestSlice(t *testing.T) {
	array := []int{54, 80, 29, 79}
	a2 := make([]int, len(array))
	a2 = append(a2, array...)
	t.Log(&array[0], &a2[0])
	a1 := array[1:3]
	a1 = append(a1, 6)
	a1 = append(a1, 16)
	//a1 = append(a1, array...)
	//a1 = append(a1, 4)
	t.Log(&array[1], &a1[0])
	t.Log(array)

}
func TestNewHeapSlice(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	//array1 := []int{6, 29, 26, 51, 34, 54, 28, 86, 80, 79, 65, 58, 85, 93, 39, 95}
	heap := NewHeapSlice(array)
	t.Log(heap.array)
	s := make([]int, len(array))
	l := heap.size
	for index := 0; index < l; index++ {
		s[index] = heap.Remove()
	}
	t.Log(s)
	t.Log(heap.array)
}
func TestTopK(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	sli := TopK(3, array)
	fmt.Println(sli)
}

func TestHeapSort(t *testing.T) {
	array := []int{54, 80, 29, 79, 6, 58, 93, 86, 51, 65, 34, 39, 85, 26, 28, 95}
	array = HeapSort(array)
	t.Log(array)
}
