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
}

// NewNode returns new node with randomly generated `p` value
func NewNode(key int) *Node {
	p := int(rand.Int31n(math.MaxInt32))
	return &Node{key: key, p: p}
}

// SetParent sets parent for a given node, if the node is non-nil
func (n *Node) SetParent(p *Node) {
	if n != nil {
		n.parent = p
	}
}
