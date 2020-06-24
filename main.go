package main

import (
	"fmt"

	bst "code.golang.com/Datastruct/05-Tree/01-BST"
)

func main() {
	/*
		                      7
		   			   ------|------
		   			 /				\
		   			4				9
		   		  --|--			  --|--
		   		 /	   \		 /     \
		   		2		5		8	   11
		   	 --/--					 --|--
		   	/	  \			        /	  \
		   1	   3			   10	  12
											\
											13
	*/
	arr := []int{7, 4, 9, 2, 5, 8, 11, 1, 3, 10, 12}
	nBst := bst.New(nil)
	for _, v := range arr {
		err := nBst.AddByRecur(v)
		if err != nil {
			panic("Add element to BST occure error")
		}
	}

	fmt.Println(nBst.Rank(11))
	// for _, v := range arr {
	// 	fmt.Println(nBst.Contains(v))
	// }
	
}
