package main

import (
	"fmt"
	"github.com/data/refactorBST"
	"strconv"
)

func main() {

	arr := []int{94, 28, 70, 86, 89, 72, 24, 7, 75, 33, 23, 9, 55, 22, 80, 30, 18}

	//test1(arr)
	fmt.Println()
	test2(arr)

	//bst := refactorBST.NewBst()
	/*	for _, v := range arr {
			bst.Add(v)
		}
	*/
	//54, 38, 69, 21, 43, 60, 82, 10, 34, 49, 56, 78, 96, 3, 81
	/*avl := refactorBST.NewAVL1()
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
	avl.LevelRange(visitor)*/

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

func test2(arr []int) {
	avl := refactorBST.NewAVL1()
	for _, v := range arr {
		avl.Add(v)
	}
	visitor := func(node *refactorBST.Node) {
		fmt.Print(strconv.Itoa(node.Ele) + "\t")
	}
	avl.LevelRange(visitor)
	fmt.Println()
	avl.Remove(70)
	avl.LevelRange(visitor)
	fmt.Println()
	avl.Remove(22)
	avl.LevelRange(visitor)
	fmt.Println()

}

func test1(arr []int) {
	visitor := func(node *refactorBST.Node) {
		fmt.Print(strconv.Itoa(node.Ele) + "\t")
	}
	rb := refactorBST.NewRedBlackTree()
	for _, v := range arr {
		rb.Add(v)
	}
	rb.LevelRange(visitor)
	fmt.Println()
	fmt.Println(rb.HeightByLevel())
	//rb.MidRange(visitor)
	visitor1 := func(node *refactorBST.Node) {
		if node.Color == 1 {
			fmt.Print(strconv.Itoa(node.Ele) + "\t")
		}
	}
	rb.LevelRange(visitor1)
}
