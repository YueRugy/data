package mi

type HeapI interface {
	Size() int
	Empty() bool
	Add(int)
	Get() int
	Clear()
	Remove() int
	Replace(int) int
}
