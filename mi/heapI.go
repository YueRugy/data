package main

type HeapI interface {
	Size() int
	Empty() bool
	Add(int)
	Get() int
	Clear()
	Remove()
	Replace(int) int
}
