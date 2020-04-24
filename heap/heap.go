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

func (h *Heap) Remove() {
	panic("implement me")
}

func (h *Heap) Replace(num int) int {
	return num
}

func NewHeap() *Heap {
	return &Heap{
		size:  0,
		array: make([]int, defaultCap),
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
