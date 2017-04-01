package main

import (
	"fmt"

	"github.com/ideahitme/gTreap"
)

func main() {
	simpleDupl := treap.NewSimpleTreap(&treap.DefaultConfig) //allows duplicates
	simpleDupl.Add(10, 2, 15, -9, 15, 203, 41)
	fmt.Println(simpleDupl.Slice()) //prints sorted entries

	simpleNoDupl := treap.NewSimpleTreap(&treap.Config{AllowDuplicates: false})
	simpleNoDupl.Add(1, 1, 1, 2)
	fmt.Println(simpleNoDupl.Slice())

}
