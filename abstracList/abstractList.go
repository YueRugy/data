package abstracList

type AbstractList interface {
	Add(element interface{})
	Set(element interface{}, index int) error
	Remove() bool
	IndexOf(ele interface{}) int
	RemoveIndex(index int) int
	Get(index int) (interface{}, error)
	Size() int
	Empty() bool
	Contains(element int) bool
}
