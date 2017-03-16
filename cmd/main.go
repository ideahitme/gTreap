package main

import (
	"github.com/ideahitme/gTreap"
)

func main() {
	st := treap.NewSimpleTreap()
	st.Add(10, 2, 15, -9, 203, 41)
	st.PrintSlice() //prints sorted entries
}
