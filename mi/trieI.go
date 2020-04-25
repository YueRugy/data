package mi

type trieI interface {
	Empty() bool
	Size() int
	Clear()
	Add(string, int)
	Remove(string) int
	StartWith(string) bool
	Contains(string) bool
}
