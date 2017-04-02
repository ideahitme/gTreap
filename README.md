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
		Simple Treap - randomized binary search tree
		Methods: 
			Add(...int) - add elements to the Treap
			Find(x int) bool - find if element exists in the Treap
	*/
	simpleDupl := treap.NewSimpleTreap(&treap.DefaultConfig) //allows duplicates
	simpleDupl.Add(10, 2, 15, -9, 15, 203, 41)
	fmt.Println(simpleDupl.Slice()) // [-9 2 10 15 15 41 203]

	simpleNoDupl := treap.NewSimpleTreap(&treap.Config{AllowDuplicates: false})
	simpleNoDupl.Add(1, 1, 1, 2)
	fmt.Println(simpleNoDupl.Slice()) // [1 2]

	/** 
		Indexed Treap - randomized binary searhc tree supporting efficient element rearranging 
		Methods: 
			Add(...int) - add elements to the Treap
			Find(x int) bool - find if element exists in the Treap
		
	*/
}
```

