package main

import (
	bst "code.golang.com/Datastruct/05-Tree/01-Bindary-Seaarch-Tree"
	"fmt"
)

func main() {
	arr := []int{7, 4, 9, 2, 5, 8, 11, 1, 3, 10, 12}

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
	nBst := bst.New(nil)

	for _, v := range arr {
		err := nBst.Add(v)
		if err != nil {
			panic(err)
		}
	}

	/*visitor := func(e interface{}) bool {
		fmt.Printf("%v\t", e)
		if utils.Comparator(e, 2) == 0 {
			return true
		}
		return false
	}

	fmt.Println("PreRecur:")
	nBst.PreRecur(&bst.Visitor{Visit: visitor})

	fmt.Println("\nIneRecur:")
	nBst.InRecur(&bst.Visitor{Visit: visitor})

	fmt.Println("\nPostRecur:")
	nBst.PostRecur(&bst.Visitor{Visit: visitor})

	fmt.Println("\nLevelTravel:")
	nBst.LevelTravel(&bst.Visitor{Visit: visitor})*/

	/*for i := 0; i < 150; i++ {
		rand.Seed(time.Now().UnixNano())
		if err := nBst.Add(rand.Intn(100)); err != nil {
			panic(err)
		}
	}
	fmt.Println("-------- test --------")
	fmt.Println(nBst.GetHByIota())
	fmt.Println(nBst.GetHByRecur())*/
	fmt.Println(nBst)
	fmt.Println(nBst.IsComplete())
}
