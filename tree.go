package treap

import (
	"math/rand"
	"time"
)

// Tree is the most simple implementation of treap
type Tree struct {
	root *Node
	cfg  *Config
	size int
}

/** Public methods */

// NewTree returns Tree object with nil root
func NewTree(cfg *Config) *Tree {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Tree{
		cfg:  cfg,
		size: 0,
	}
}

// Insert adds values to the Treap
func (st *Tree) Insert(x ...int) {
	for _, el := range x {
		if !st.cfg.AllowDuplicates && st.find(st.root, el) {
			continue
		}
		newNode := NewNode(el)
		l, r := st.split(st.root, el)
		st.root = st.merge(l, newNode)
		st.root = st.merge(st.root, r)
		st.size++
	}
}

// Find finds if value is present in the Treap
func (st *Tree) Find(x int) bool {
	return st.find(st.root, x)
}

// Slice returns sorted slice by InOrder traversing
func (st *Tree) Slice() []int {
	result := make([]int, 0, st.size)
	st.inOrder(st.root, func(cur *Node) {
		result = append(result, cur.key)
	})
	return result
}

// Remove deletes node with key == x from the Treap
// if multiple such elements exist only one is deleted
func (st *Tree) Remove(x int) {
	l, r := st.split(st.root, x-1)
	xl, xr := st.split(r, x)
	if xl != nil {
		l = st.merge(l, xl.left)
		xr = st.merge(xl.right, xr)
	}
	st.root = st.merge(l, xr)
}

// RemoveAll deletes all node with key == x from the Treap
func (st *Tree) RemoveAll(x int) {
	l, r := st.split(st.root, x-1)
	_, xr := st.split(r, x)
	st.root = st.merge(l, xr)
}

/** Private methods */

// split recursively splits the treap with a given root into two Treaps l, r based on key
// l - has all node keys less or equal to key
// r - has all node keys strictly greater than key
func (st *Tree) split(root *Node, key int) (l *Node, r *Node) {
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
func (st *Tree) merge(l, r *Node) (root *Node) {
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

func (st *Tree) inOrder(cur *Node, f func(cur *Node)) {
	if cur == nil {
		return
	}
	st.inOrder(cur.left, f)
	f(cur)
	st.inOrder(cur.right, f)
}

func (st *Tree) find(cur *Node, x int) bool {
	if cur == nil {
		return false
	}
	if cur.key == x {
		return true
	}
	if cur.key > x {
		return st.find(cur.left, x)
	}
	return st.find(cur.right, x)
}
