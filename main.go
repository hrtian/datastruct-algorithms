package main

import (
	"fmt"

	bst "code.golang.com/Datastruct/05-Tree/01-Bindary-Seaarch-Tree"
	utils "code.golang.com/Datastruct/Utils"
)

func main() {
	arr := []int{7, 4, 9, 2, 5, 8, 11, 1, 3, 10, 12}

	nBst := bst.New(nil)

	for _, v := range arr {
		err := nBst.Add(v)
		if err != nil {
			panic(err)
		}
	}

	nBst.LevelTravel(func(e interface{}) bool {
			fmt.Printf("%v\t", e)
			if utils.Comparator(e, 2) == 0 {
				return true
			}
			return false
		})
}
