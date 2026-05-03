package tree

// since AVL has strict balancing rules, its height is less than that of the Red-Black one, and therefore search is faster
// ...but on the other hand - the stricter balancing the more rotations are needed, and therefore more calculations...

type AVLTreeNode struct {
	d int
	l *AVLTreeNode
	r *AVLTreeNode
	h int
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {

	return &AVLTree{}
}

func (t *AVLTree) IsBalanced() bool {

	return false
}
