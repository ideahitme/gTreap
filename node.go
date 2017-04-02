package treap

import (
	"math"
	"math/rand"
)

// Node storage for treap node
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	key    int
	p      int //priority - randomized value
	size   int
	value  int
}

// NewNode returns new node with randomly generated `p` value
func NewNode(key int) *Node {
	p := int(rand.Int31n(math.MaxInt32))
	return &Node{key: key, p: p, size: 1}
}

// NewNodeWithValue returns new node with additional field for value. Used for `Slice` Treap with element index acting as node key
func NewNodeWithValue(key, value int) *Node {
	n := NewNode(key)
	n.value = value
	return n
}

// SetParent sets parent for a given node, if the node is non-nil
func (n *Node) SetParent(p *Node) {
	if n != nil {
		n.parent = p
	}
}

// Size returns size of the subtree including current node
func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return n.size
}
