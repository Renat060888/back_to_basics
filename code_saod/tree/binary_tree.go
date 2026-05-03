package tree

type TreeNode struct {
	d int
	l *TreeNode
	r *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &TreeNode{d: data}
	}

	// TODO: create child
}

func (t *Tree) Delete(data int) bool {

	// TODO: do
	return false
}

func (t *Tree) Search(data int) bool {

	// TODO: do
	return false
}

func (t *Tree) GetMin() *int {

	// TODO: do
	return nil
}

func (t *Tree) GetMax() *int {

	// TODO: do
	return &t.root.d
}

func (t *Tree) GetSize() int {

	// TODO: do
	return 0
}

func (t *Tree) Clear() {

	// TODO: do
}

func (t *Tree) IsBalanced() bool {

	return false
}
