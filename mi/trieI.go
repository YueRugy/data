package mi

type TrieI interface {
	Get(string) int
	Empty() bool
	Size() int
	Clear()
	Add(string, int)
	Remove(string) int
	StartWith(string) bool
	Contains(string) bool
}
