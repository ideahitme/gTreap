package treap

import (
	"math/rand"
	"time"
)

var _ Treap = &SimpleTreap{}

// SimpleTreap is the most simple implementation of treap
type SimpleTreap struct {
	root *Node
	cfg  *Config
	size int
}

/** Public methods */

// NewSimpleTreap returns SimpleTreap object with nil root
func NewSimpleTreap(cfg *Config) *SimpleTreap {
	rand.Seed(time.Now().UTC().UnixNano())
	return &SimpleTreap{
		cfg:  cfg,
		size: 0,
	}
}

// Add adds values to the Treap
func (st *SimpleTreap) Add(x ...int) {
	for _, el := range x {
		if !st.cfg.AllowDuplicates && st.find(st.root, el) {
			continue
		}
		newNode := NewNode(el)
		l, r := st.split(st.root, &el)
		st.root = st.merge(l, newNode)
		st.root = st.merge(st.root, r)
		st.size++
	}
}

// Find finds if value is present in the Treap
func (st *SimpleTreap) Find(x int) bool {
	return st.find(st.root, x)
}

// Slice returns sorted slice by InOrder traversing
func (st *SimpleTreap) Slice() []int {
	result := make([]int, 0, st.size)
	st.inOrderTrv(st.root, func(cur *Node) {
		result = append(result, cur.key)
	})
	return result
}

/** Private methods */

// split recursively splits the treap with a given root into two Treaps l, r based on key
// l - has all node keys less or equal to key
// r - has all node keys strictly greater than key
func (st *SimpleTreap) split(root *Node, key *int) (l *Node, r *Node) {
	if root == nil {
		return
	}
	if root.key > *key {
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

func (st *SimpleTreap) inOrderTrv(cur *Node, f func(cur *Node)) {
	if cur == nil {
		return
	}
	st.inOrderTrv(cur.left, f)
	f(cur)
	st.inOrderTrv(cur.right, f)
}

func (st *SimpleTreap) find(cur *Node, x int) bool {
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
