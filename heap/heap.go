package heap

const (
	defaultCap = 16
)

type Heap struct {
	size  int
	array []int
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Empty() bool {
	return h.Size() == 0
}

func (h *Heap) Add(ele int) {
	h.ensureCap(h.size + 1)
	//selfIndex := h.size
	h.array[h.size] = ele
	h.siftUp(h.size)
	h.size++
}

func (h *Heap) siftUp(selfIndex int) {
	node := h.array[selfIndex]
	for {
		if selfIndex <= 0 {
			break
		}
		pin := (selfIndex - 1) >> 1
		if node >= h.array[pin] {
			break
		} else {
			h.array[selfIndex] = h.array[pin]
			selfIndex = pin
		}
	}
	h.array[selfIndex] = node
}

//func (h *Heap) siftUp(selfIndex int) {
//	for {
//		if selfIndex <= 0 {
//			break
//		}
//		fatherIndex := (selfIndex - 1) >> 1
//		if h.array[selfIndex] >= h.array[fatherIndex] {
//			break
//		} else {
//			h.array[fatherIndex], h.array[selfIndex] = h.array[selfIndex], h.array[fatherIndex]
//			selfIndex = fatherIndex
//		}
//	}
//}

func (h *Heap) Get() int {
	if h.Empty() {
		return 0
	}
	return h.array[0]
}

func (h *Heap) Clear() {
	if h.Empty() {
		return
	}
	h.size = 0
	for i := 0; i < len(h.array); i++ {
		h.array[i] = 0
	}
}

func (h *Heap) Remove() int {
	if h.size == 0 {
		return 0
	}
	res := h.array[0]
	if h.size == 1 {
		h.size--
		h.array[h.size] = 0
		return res
	} else {
		h.size--
		h.array[0] = h.array[h.size]
		node := h.array[0]
		h.array[h.size] = 0
		index := 0
		h.siftDown(node, index)
	}
	return res
}

func (h *Heap) findMinIndex(index, compare int) int {
	left := index<<1 + 1
	if left >= h.size {
		return -1
	}
	if left == h.size-1 {
		if h.array[left] > compare {
			return -1
		}
		return left
	} else {
		minIndex := -1
		if h.array[left] < h.array[left+1] {
			minIndex = left
		} else {
			minIndex = left + 1
		}

		if compare < h.array[minIndex] {
			return -1
		}

		return minIndex

	}

}

func (h *Heap) Replace(num int) int {
	if h.size == 0 {
		h.array[h.size] = num
		h.size++
		return 0
	}
	ele := h.array[0]
	h.array[0] = num
	if h.size == 1 {
		return ele
	}
	index := 0
	h.siftDown(num, index)
	return ele
}

func (h *Heap) siftDown(num int, index int) {
	for {
		res := h.findMinIndex(index, num)
		if res < 0 {
			break
		}
		h.array[index] = h.array[res]
		index = res
	}
	h.array[index] = num
}

func NewHeap() *Heap {
	return &Heap{
		size:  0,
		array: make([]int, defaultCap),
	}
}
func NewHeapSlice(sli []int) *Heap {
	if sli == nil {
		return NewHeap()
	}
	l := len(sli)
	if l < defaultCap {
		l = defaultCap
	}
	heap := &Heap{
		size:  len(sli),
		array: make([]int, l),
	}
	for index := 0; index < heap.size; index++ {
		heap.array[index] = sli[index]
	}
	//copy(heap.array, sli)
	heap.heapify()
	return heap
}

func (h *Heap) heapify() {
	for index := h.size>>1 - 1; index >= 0; index-- {
		h.siftDown(h.array[index], index)
	}
}

func (h *Heap) ensureCap(num int) {
	length := len(h.array)
	if num <= length {
		return
	}
	newArr := make([]int, length+length>>1)
	newArr = append(newArr, h.array...)
	h.array = newArr
}
func TopK(k int, sli []int) []int {
	heap := NewHeap()
	for index := 0; index < len(sli); index++ {
		if index < k {
			heap.Add(sli[index])
		} else if sli[index] > heap.Get() {
			heap.Replace(sli[index])
		}
	}
	return heap.array
}

func HeapSort(arr []int) []int {
	heap := NewHeapSlice(arr)
	for index := heap.size - 1; index >= 0; index-- {
		heap.array[index] = heap.Remove()
	}
	return heap.array
}
