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
func (t *Tree) Insert(x ...int) {
	for _, el := range x {
		if !t.cfg.AllowDuplicates && t.find(t.root, el) {
			continue
		}
		newNode := NewNode(el)
		l, r := t.split(t.root, el)
		t.root = t.merge(l, newNode)
		t.root = t.merge(t.root, r)
		t.size++
	}
}

// Find finds if value is present in the Treap
func (t *Tree) Find(x int) bool {
	return t.find(t.root, x)
}

// Slice returns sorted slice by InOrder traversing
func (t *Tree) Slice() []int {
	result := make([]int, 0, t.size)
	t.inOrder(t.root, func(cur *Node) {
		result = append(result, cur.key)
	})
	return result
}

// Remove deletes node with key == x from the Treap
// if multiple such elements exist only one is deleted
func (t *Tree) Remove(x int) {
	l, r := t.split(t.root, x-1)
	xl, xr := t.split(r, x)
	if xl != nil {
		l = t.merge(l, xl.left)
		xr = t.merge(xl.right, xr)
	}
	t.root = t.merge(l, xr)
}

// RemoveAll deletes all node with key == x from the Treap
func (t *Tree) RemoveAll(x int) {
	l, r := t.split(t.root, x-1)
	_, xr := t.split(r, x)
	t.root = t.merge(l, xr)
}

/** Private methods */

// split recursively splits the treap with a given root into two Treaps l, r based on key
// l - has all node keys less or equal to key
// r - has all node keys strictly greater than key
func (t *Tree) split(root *Node, key int) (l *Node, r *Node) {
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
func (t *Tree) merge(l, r *Node) (root *Node) {
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

func (t *Tree) inOrder(cur *Node, f func(cur *Node)) {
	if cur == nil {
		return
	}
	t.inOrder(cur.left, f)
	f(cur)
	t.inOrder(cur.right, f)
}

func (t *Tree) find(cur *Node, x int) bool {
	if cur == nil {
		return false
	}
	if cur.key == x {
		return true
	}
	if cur.key > x {
		return t.find(cur.left, x)
	}
	return t.find(cur.right, x)
}
