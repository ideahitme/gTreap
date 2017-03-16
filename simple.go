package treap

import "fmt"

// SimpleTreap is the most simple implementation of treap
type SimpleTreap struct {
	root *Node
}

/** Public methods */

// NewSimpleTreap returns SimpleTreap object with nil root
func NewSimpleTreap() *SimpleTreap {
	return &SimpleTreap{}
}

// Add adds values to the Treap
func (st *SimpleTreap) Add(x ...int) {
	for _, el := range x {
		newNode := NewNode(el)
		l, r := st.split(st.root, el)
		st.root = st.merge(l, newNode)
		st.root = st.merge(st.root, r)
	}
}

// PrintSlice prints the result as sorted array by InOrder traversing
func (st *SimpleTreap) PrintSlice() {
	st.inOrderFunc(st.root, func(cur *Node) {
		fmt.Printf("%d ", cur.key)
	})
	fmt.Printf("\n")
}

/** Private methods */

// split recursively splits the treap with a given root into two Treaps l, r based on key
// l - has all node keys less or equal to key
// r - has all node keys strictly greater than key
func (st *SimpleTreap) split(root *Node, key int) (l *Node, r *Node) {
	if root == nil {
		return
	}
	if root.key > key {
		subL, subR := st.split(root.left, key)
		subR.SetParent(root)
		root.left = subR
		return subL, root
	}
	subL, subR := st.split(root.right, key)
	subL.SetParent(root)
	root.right = subL
	return root, subR
}

// merge recursively merges two Treaps into one, maintaining Treap property
// produces new root
// it is required that every key in one of the Treaps is less than any key in another Treap
func (st *SimpleTreap) merge(l, r *Node) (root *Node) {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.key > r.key {
		l, r = r, l
	}
	if l.p > r.p {
		subR := st.merge(l.right, r)
		subR.SetParent(l)
		l.right = subR
		return l
	}
	subL := st.merge(l, r.left)
	subL.SetParent(r)
	r.left = subL
	return r
}

func (st *SimpleTreap) inOrderFunc(cur *Node, f func(cur *Node)) {
	if cur == nil {
		return
	}
	st.inOrderFunc(cur.left, f)
	f(cur)
	st.inOrderFunc(cur.right, f)
}
