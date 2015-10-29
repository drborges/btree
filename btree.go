package btree

type items []Item
type children []*node

type Item interface {
	Less(Item) bool
}

// node is an internal node in a tree.
//
// It must at all times maintain the invariant that either
//   * len(children) == 0, len(items) unconstrained
//   * len(children) == len(items) + 1
type node struct {
	tree     *BTree
	parent   *node
	items    items
	children children
}

type BTree struct {
	degree int
	length int
	root   *node
}

// A B-tree of order m is an m-ary search tree with the following properties:
// The root is either a leaf or has at least two children
// Each node, except for the root and the leaves, has between (m/2) and (m) children
// Each path from the root to a leaf has the same length.
// The root, each internal node, and each leaf is typically a disk block.
// Each internal node has up to (m - 1) key values and up to m pointers (to children)
// The records are typically stored in leaves (in some organizations, they are also stored in internal nodes)
//
// See http://lcm.csa.iisc.ernet.in/dsa/node122.html
func New(degree int) *BTree {
	if degree <= 1 {
		panic("bad degree")
	}
	return &BTree{
		degree:   degree,
	}
}

// Len returns the number of items currently in the tree
func (b *BTree) Len() int {
	return b.length
}