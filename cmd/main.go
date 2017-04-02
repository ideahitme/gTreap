package main

import (
	"fmt"

	"github.com/ideahitme/gTreap"
)

func main() {
	/**
	Simple Treap - randomized binary search tree
	Methods:
		Insert(...int) - add elements to the Treap
		Find(x int) bool - find if element exists in the Treap
		Slice() []int - returns the sorted elements as slice
		Remove(x int) - removes one occurence of Node with key == x
		RemoveAll(x int) removes all occurences of x
	*/
	simpleDupl := treap.NewTree(&treap.DefaultConfig) //allows duplicates
	simpleDupl.Insert(10, 2, 15, -9, 15, 203, 41)
	fmt.Println(simpleDupl.Slice()) // [-9 2 10 15 15 41 203]
	simpleDupl.Remove(15)
	fmt.Println(simpleDupl.Slice()) // [-9 2 10 15 41 203]
	simpleDupl.Insert(41)
	fmt.Println(simpleDupl.Slice()) // [-9 2 10 15 41 41 203]
	simpleDupl.RemoveAll(41)
	fmt.Println(simpleDupl.Slice()) // [-9 2 10 15 203]

	simpleNoDupl := treap.NewTree(&treap.Config{AllowDuplicates: false})
	simpleNoDupl.Insert(1, 1, 1, 2)
	fmt.Println(simpleNoDupl.Slice()) // [1 2]

	/**
	Indexed Treap - randomized binary search tree supporting efficient element rearranging
	Methods:
		Add(...int) - add elements to the Treap
		Find(x int) bool - find if element exists in the Treap

	*/
}
