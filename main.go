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
	avl := refactorBST.NewAVL()
	for _, v := range arr {
		avl.Add(v)
	}
	/*bst.MidRange(func(i int) {
		//fmt.Print(strconv.Itoa(i) + "\t")
	})*/
	//fmt.Println()
	avl.LevelRange(func(i int) {
		fmt.Print(strconv.Itoa(i) + "\t")
	})
	fmt.Println()
	fmt.Println(avl.HeightByLevel())
}
