package main

import (
	"fmt"
	"github.com/data/refactorBST"
	"strconv"
)

func main() {

	arr := [...]int{69, 60, 96, 49, 54, 38, 10, 43, 78, 21, 82, 34, 56, 3, 81, 40}
	//bst := refactorBST.NewBst()
	/*	for _, v := range arr {
			bst.Add(v)
		}
	*/
	//54, 38, 69, 21, 43, 60, 82, 10, 34, 49, 56, 78, 96, 3, 81
	avl := refactorBST.NewAVL1()
	for _, v := range arr {
		avl.Add(v)
	}
	fmt.Println(avl.HeightByLevel())
	visitor := func(ele int) {
		fmt.Print(strconv.Itoa(ele) + "\t")
	}
	avl.LevelRange(visitor)
	fmt.Println()
	avl.Remove(40)
	avl.LevelRange(visitor)
	fmt.Println()
	avl.Remove(49)
	avl.LevelRange(visitor)

	/*bst.MidRange(func(i int) {
		//fmt.Print(strconv.Itoa(i) + "\t")
	})*/
	//fmt.Println()
	//avl.LevelRange(func(i int) {
	//	fmt.Print(strconv.Itoa(i) + "\t")
	//})
	//fmt.Println()
	//fmt.Println(avl.HeightByLevel())
}
