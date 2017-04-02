# gTreap
Treap data structure implementation in Go 

## Simple Treap

```
package main

import (
	"github.com/ideahitme/gTreap"
)

func main(){
	/**
	Tree Treap - randomized binary search tree
	Methods:
		Insert(...int) - add elements to the Treap
		Find(x int) bool - find if element exists in the Treap
		Slice() []int - returns the sorted elements as slice
		Remove(x int) - removes one occurence of Node with key == x
		RemoveAll(x int) removes all occurences of x

	All operations are O(log(n)) in terms of time complexity
	*/
	treeDupl := treap.NewTree(&treap.DefaultConfig) //allows duplicates
	treeDupl.Insert(10, 2, 15, -9, 15, 203, 41)
	fmt.Println(treeDupl.Slice()) // [-9 2 10 15 15 41 203]
	treeDupl.Remove(15)
	fmt.Println(treeDupl.Slice()) // [-9 2 10 15 41 203]
	treeDupl.Insert(41)
	fmt.Println(treeDupl.Slice()) // [-9 2 10 15 41 41 203]
	treeDupl.RemoveAll(41)
	fmt.Println(treeDupl.Slice()) // [-9 2 10 15 203]

	treeNoDupl := treap.NewTree(&treap.Config{AllowDuplicates: false})
	treeNoDupl.Insert(1, 1, 1, 2)
	fmt.Println(treeNoDupl.Slice()) // [1 2]
	fmt.Printf("%d exists: %t\n", 1, treeNoDupl.Find(1))
	fmt.Printf("%d exists: %t\n", 3, treeNoDupl.Find(3))

	/**
	Indexed Treap - randomized binary search tree supporting efficient range updates
	Methods:
		Insert(index, value int)
		Slice() [][]int
		Find(x int) (value int, exists bool)
		Remove(x int)
	*/
	indexed := treap.NewIndexed()
	indexed.Insert(0, 10)
	indexed.Insert(10, 0)
	indexed.Insert(5, 5)
	indexed.Insert(5, 7)         //overwrites previous index
	fmt.Println(indexed.Slice()) // [[0 10] [5 7] [10 0]]
	fmt.Println(indexed.Find(5)) // 7 true
	fmt.Println(indexed.Find(1)) // 0 false
	indexed.Remove(2)            // does not exist
	indexed.Remove(5)            // removes [5 7]
	fmt.Println(indexed.Slice()) // [[0 10] [10 0]]
}
```

