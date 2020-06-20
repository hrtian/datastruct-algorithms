package main

import (
	"code.golang.com/Datastruct/05-Tree/01-Bindary-Seaarch-Tree"
	"fmt"
)

func main() {
	arr := []int{7, 4, 9, 2, 5, 8, 11, 1, 3}

	bst := BST.New()

	for v, _ := range arr {
		bst.Add(v)
	}

	fmt.Println(bst)
}
