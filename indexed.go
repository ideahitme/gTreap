package treap

import (
	"math/rand"
	"time"
)

// Indexed is
type Indexed struct {
	root *Node
	size int
}

// NewIndexed returns new Indexed Tree
func NewIndexed() *Indexed {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Indexed{
		size: 0,
	}
}

// Insert adds new element to the treap
// index - index of the element - same as key in normal Tree
// value - value of the element - normally something to be updated
// if element with identical index is found, it is overwritten
func (t *Indexed) Insert(index, value int) {
	newNode := NewNodeWithValue(index, value)
	l, r := t.split(t.root, index-1)
	xl, xr := t.split(r, index)
	if xl == nil {
		t.size++
	}
	l = t.merge(l, newNode)
	t.root = t.merge(l, xr)
}

// Slice returns slice by InOrder traversing
// each element of the slice is a slice of form [index, value]
// slice is sorted by index in increasing order
func (t *Indexed) Slice() [][]int {
	result := make([][]int, 0, t.size)
	t.inOrder(t.root, func(cur *Node) {
		result = append(result, []int{cur.key, cur.value})
	})
	return result
}

// Remove deletes node with key == x from the Indexed Treap
func (t *Indexed) Remove(x int) {
	l, r := t.split(t.root, x-1)
	_, xr := t.split(r, x)
	t.root = t.merge(l, xr)
}

// Find returns value of element in the treap with index == x, second parameter
// indicates if element exists in the Treap
func (t *Indexed) Find(x int) (value int, exists bool) {
	n := t.find(t.root, x)
	if n == nil {
		return 0, false
	}
	return n.value, true
}

// Add adds "x" to all elements with index i <= t <= j
func (t *Indexed) Add(x int, i, j int) {

}

// Max returns maximum element (by value) with index i <= t <= j
func (t *Indexed) Max(i, j int) int {
	return -1
}

// Min returns minimum element (by value) with index i <= t <= j
func (t *Indexed) Min(i, j int) int {
	return -1
}

/** Private methods */

// split recursively splits the treap with a given root into two Treaps l, r based on key
// l - has all node keys less or equal to key
// r - has all node keys strictly greater than key
func (t *Indexed) split(root *Node, key int) (l *Node, r *Node) {
	if root == nil {
		return
	}
	if root.key > key {
		subL, subR := t.split(root.left, key)
		subR.SetParent(root)
		root.left = subR
		return subL, root
	}
	subL, subR := t.split(root.right, key)
	subL.SetParent(root)
	root.right = subL
	return root, subR
}

// merge recursively merges two Treaps into one, maintaining Treap property
// produces new root
// it is required that every key in one of the Treaps is less than any key in another Treap
func (t *Indexed) merge(l, r *Node) (root *Node) {
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
		subR := t.merge(l.right, r)
		subR.SetParent(l)
		l.right = subR
		return l
	}
	subL := t.merge(l, r.left)
	subL.SetParent(r)
	r.left = subL
	return r
}

func (t *Indexed) inOrder(cur *Node, f func(cur *Node)) {
	if cur == nil {
		return
	}
	t.inOrder(cur.left, f)
	f(cur)
	t.inOrder(cur.right, f)
}

func (t *Indexed) find(cur *Node, x int) *Node {
	if cur == nil {
		return nil
	}
	if cur.key == x {
		return cur
	}
	if cur.key > x {
		return t.find(cur.left, x)
	}
	return t.find(cur.right, x)
}
