package tree

import (
	"saod/common"
	"strconv"
)

// red-black tree in essense is a binary tree with balancing capabilities from b-tree
// insertion and deletion are weird because it is actually an abstration over b-tree

// tree can be unbalanced due to:
// - insertion of elements in sorted order
// - insertion some large or small value which will funnel most of remaining values to one side of tree
// 'balanced' means that each of tree's node needs to have roughly the same number descendents on its left side as it has on its right.

// rules of red-black tree:
// - root and nil nodes are always black
// - every path from root to leaf, including nil node, must contain the same number of black nodes (black height)
// - two consecutive red nodes are prohibited (THE MOST IMPORTANT RULE)

// black height violation is about correct balance (violated at REMOVE) (recall b-tree: black node is its central element)
// red consecutive violation is about correct representation of red-black tree as b-tree (violated at INSERT) (recall b-tree: red nodes are edges of its one node)

// insertion -----------------------------------------------------------------------------------------------------------
// trigger for colors flip:
// 1. two consecutive red nodes
// 2. red uncle
// actually does not do anything, it propagates "problem" to upper level

// trigger for rotation (balancing)
// 1. two consecutive red nodes
// 2. black uncle
// really changes structure of tree
// ---------------------------------------------------------------------------------------------------------------------

// NOTE: it is assumed that there are NO DUPLICATE keys (can be handled by randomizing process in insertion algorithm)

// rotation types:
// - single
//
// *
//  *    ->  *
//   *      * *
//
// - double (straightening of bent tree)
//
//  *          *
// *     ->   *   ->   *
//  *        *        * *
//

// rules: [will be used in IsValid()]
//
//
//

// for self check https://www.cs.usfca.edu/~galles/visualization/RedBlack.html

// wiki:
// For every 2–3–4 tree, there are corresponding red–black trees with data elements in the same order.
// The insertion and deletion operations on 2–3–4 trees are also equivalent to color-flipping and rotations in red–black trees.
// This makes 2–3–4 trees an important tool for understanding the logic behind red–black trees,
// and this is why many introductory algorithm texts introduce 2–3–4 trees just before red–black trees,
// even though 2–3–4 trees are not often used in practice.

// implementation example: https://github.com/niedong/stl_rbtree/blob/63034ac2a1aac3fe855235031f174a4198732c4f/src/rbtree.c#L383

const (
	RB_TREE_COLOR_RED = iota
	RB_TREE_COLOR_BLACK

	RB_TREE_TRAVERSE_PRE_ORDER
	RB_TREE_TRAVERSE_IN_ORDER
	RB_TREE_TRAVERSE_POST_ORDER
)

type RBTreeNode struct {
	d int
	l *RBTreeNode
	r *RBTreeNode
	p *RBTreeNode
	c int
}

type RBTree struct {
	root *RBTreeNode
}

func NewRBTree() *RBTree {

	return &RBTree{}
}

func (t *RBTree) Insert(data int) bool {

	if t.Contains(data) {
		return false
	}

	if t.root == nil {
		t.root = &RBTreeNode{d: data, c: RB_TREE_COLOR_BLACK}
		return true
	}

	currNode := t.root
	for currNode != nil {
		if data > currNode.d { // insert to right side
			if currNode.r != nil {
				currNode = currNode.r
				continue
			} else {
				currNode.r = &RBTreeNode{d: data, c: RB_TREE_COLOR_RED}
				currNode.r.p = currNode
				t.fixInsert(currNode.r) // fix insert
				break
			}
		} else { // insert to left side
			if currNode.l != nil {
				currNode = currNode.l
				continue
			} else {
				currNode.l = &RBTreeNode{d: data, c: RB_TREE_COLOR_RED}
				currNode.l.p = currNode
				t.fixInsert(currNode.l) // fix insert
				break
			}
		}
	}

	return true
}

func (t *RBTree) fixInsert(node *RBTreeNode) {

	if node.p == nil { // fixing of root is just making it black
		node.c = RB_TREE_COLOR_BLACK
		return
	}

	if node.p.c == RB_TREE_COLOR_BLACK { // otherwise 'two consecutive red nodes' violation
		return
	}

	var uncleNode *RBTreeNode
	if node.p.l == node && node.p.p != nil {
		uncleNode = node.p.p.r
	}
	if node.p.r == node && node.p.p != nil {
		uncleNode = node.p.p.l
	}

	if uncleNode != nil && uncleNode.c == RB_TREE_COLOR_RED { // RED uncle: colors flip - split in terms of b-tree
		node.p.c = RB_TREE_COLOR_BLACK
		uncleNode.c = RB_TREE_COLOR_BLACK
		node.p.p.c = RB_TREE_COLOR_RED

		t.fixInsert(node.p.p)
	} else { // BLACK uncle: rotation - compaction of existing node

		// rotation (normalization of triad of nodes, i.e. make it as a 'triangle' - red-black-red, which represents one node of 2-3-4 b-tree)
		// fix colors (inverse parent's color and respective color for childs)
	}
}

func (t *RBTree) Remove(value int) bool {

	// basic idea: deletion always from leaf (as deletion from regular binary tree)

	// P.S. or just mark value as 'removed' :)

	// TODO: it remains to implement this :D

	t.fixRemove()

	return false
}

func (t *RBTree) fixRemove() {

	// scheme:
	// if red leaf or parent of leaf (preserving its color) - simply delete
	// if all is black - sibling and its childs - move DB to parent (sibling becomes red, parent is either black if was red or DB if was black)
	// if at least one red in sibling and its childs - make colors flip
	// - if red is in sibling or in its near child - move that red color to parent, then make rotation
	// - if red is in far child of sibling - just make it black, then make rotation

	// TODO: do
}

func (t *RBTree) Contains(data int) bool {

	node := t.findNode(data)
	return node != nil
}

func (t *RBTree) findNode(value int) *RBTreeNode {

	if t.root == nil {
		return nil
	}

	currNode := t.root
	for currNode != nil {
		if currNode.d == value {
			return currNode
		} else if value > currNode.d {
			currNode = currNode.r
		} else {
			currNode = currNode.l
		}
	}

	return nil
}

func (t *RBTree) AcceptVisitor(v common.INodeVisitor, policy int) {

	t.traverse(t.root, v, policy)
}

func (t *RBTree) traverse(node *RBTreeNode, v common.INodeVisitor, policy int) {

	if node == nil {
		return
	}

	// collect data
	dataStr := strconv.Itoa(node.d)
	if node.c == RB_TREE_COLOR_RED {
		dataStr += "r"
	}

	var childsData []string
	if node.l != nil {
		childStr := strconv.Itoa(node.l.d)
		if node.l.c == RB_TREE_COLOR_RED {
			childStr += "r"
		}

		childsData = append(childsData, childStr)
	}

	if node.r != nil {
		childStr := strconv.Itoa(node.r.d)
		if node.r.c == RB_TREE_COLOR_RED {
			childStr += "r"
		}

		childsData = append(childsData, childStr)
	}

	// visit childs by policy
	if RB_TREE_TRAVERSE_PRE_ORDER == policy {
		v.VisitNode(dataStr, childsData)
	}
	t.traverse(node.l, v, policy)
	if RB_TREE_TRAVERSE_IN_ORDER == policy {
		v.VisitNode(dataStr, childsData)
	}
	t.traverse(node.r, v, policy)
	if RB_TREE_TRAVERSE_POST_ORDER == policy {
		v.VisitNode(dataStr, childsData)
	}
}

func (t *RBTree) IsValid() bool {

	if t.root == nil {
		return true
	}

	if t.root.c != RB_TREE_COLOR_BLACK { // root must be black
		return false
	}

	leafs := t.CollectLeafs(t.root)
	prevBlacksCount := -1

	// ascend from each leaf to root counting the number of blacks and checking for two consecutive reds
	for _, leaf := range leafs {

		blacksCount := 0
		currNode := leaf
		prevWasRed := false

		for currNode != nil {
			if currNode.c == RB_TREE_COLOR_BLACK {
				prevWasRed = false
				blacksCount++
			} else if currNode.c == RB_TREE_COLOR_RED {
				if prevWasRed { // reds check
					return false
				}

				prevWasRed = true
			} else {
				return false // node must be either red or black
			}

			currNode = currNode.p
		}

		if prevBlacksCount != -1 && blacksCount != prevBlacksCount { // blacks check
			return false
		}

		prevBlacksCount = blacksCount
	}

	return true
}

func (t *RBTree) CollectLeafs(node *RBTreeNode) []*RBTreeNode {

	leafs := []*RBTreeNode{}

	if node == nil {
		return leafs
	}

	// pre-order traverse
	if node.l == nil && node.r == nil {
		return append(leafs, node)
	}

	leafs = append(leafs, t.CollectLeafs(node.l)...)
	leafs = append(leafs, t.CollectLeafs(node.r)...)
	return leafs
}

func cloneRBTreeNode(node *RBTreeNode) *RBTreeNode {
	if node == nil {
		return nil
	}

	clone := &RBTreeNode{}
	clone.d = node.d
	clone.c = node.c
	clone.p = node.p
	clone.l = cloneRBTreeNode(node.l)
	clone.r = cloneRBTreeNode(node.r)

	return clone
}

func (t *RBTree) Clone() *RBTree {

	clone := &RBTree{}
	clone.root = cloneRBTreeNode(t.root)

	return clone
}

type NodeCounter struct {
	Size int
}

func (nc *NodeCounter) VisitNode(name string, childs []string) {

	nc.Size++
}

func (t *RBTree) GetSize() int {

	nodeCounter := NodeCounter{}
	t.AcceptVisitor(&nodeCounter, RB_TREE_TRAVERSE_PRE_ORDER)
	return nodeCounter.Size
}

// https://stackoverflow.com/questions/5813639/how-does-a-red-black-tree-work

// For searches and traversals, it's the same as any binary tree.

// For inserts and deletes, more sophisticated algorithms are applied which aim to ensure that the tree cannot be too unbalanced. These guarantee that all single-item operations will always run in at worst O(log n) time, whereas in a simple binary tree the binary tree can become so unbalanced that it's effectively a linked list, giving O(n) worst case performance for each single-item operation.

// The basic idea of the red-black tree is to imitate a B-tree with up to 3 keys and 4 children per node. B-trees (or variations such as B+ trees) are mainly used for database indexes and for data stored on hard disk.

// Each binary tree node has a "colour" - red or black. Each black node is, in the B-tree analogy, the subtree root for the subtree that fits within that B-tree node. If this node has red children, they are also considered part of the same B-tree node. So it is possible (though not done in practice) to convert a red-black tree to a B-tree and back, with (most) structure preserved. The only possible anomoly is that when a B-tree node has two keys and three children, you have a choice of which key to goes in the black node in the equivalent red-black tree.

// For example, with red-black trees, every line from root to leaf has the same number of black nodes. This rule is derived from the B-tree rule that all leaf nodes are at the same depth.

// Although this is the basic idea from which red-black trees are derived, the algorithms used in practice for inserts and deletes are modified to enforce all the B-tree rules (there might be a minor exception - I forget) during updates, but are tailored for the binary tree form. This means that doing a red-black tree insert or delete may give a different structure for the result than that you'd expect comparing with doing the B-tree insert or delete.

// For more detail, follow the Wikipedia link that MigDus already supplied.
