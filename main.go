package main

import (
	"fmt"
	bst "code.golang.com/Datastruct/05-Tree/01-Bindary-Seaarch-Tree"
)

func main() {
	arr := []int{7, 4, 9, 2, 5, 8, 11, 3}

	nBst := bst.New(nil)

	for _, v := range arr {
		err := nBst.Add(v)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(nBst)
}
