package main

import (
	"fmt"
	"github.com/data/refactorBST"
)

func main() {

	arr := [...]int{8, 6, 2, 7, 10, 9, 11}

	avl := refactorBST.NewAVL()
	for _, v := range arr {
		avl.Add(v)
	}
	fmt.Println()
}
