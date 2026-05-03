package tree

import (
	"fmt"
	"saod/common"
	"slices"
	"strconv"
)

// goal: reduce number of accesses to disk thanks to strong tree branching and therefore small height of b-tree
// number of accesses to disk is proportional to the height of b-tree
// truly self-balanced in the sense, that there is no any explicit code, which does it
// grows from the bottom up - each new split creates new root (?)

// 2-3-4 tree
//
//       10---20---30
//      /   |     \   \
// 1-2-3 11-12-13 .... 31-32-33
//

// rules: [will be used in IsValid()]
// - T is 'minimum degree' of tree (amount of child nodes)
// - each node can has at least T - 1 keys, and at most 2T - 1
// - e.g. 2-3-4 tree has T = 2 and may contain from 1 to 3 keys in each node

// for self check https://btree.app/

// '(!)' - key part of code

const (
	B_TREE_TRAVERSE_PRE_ORDER  = 0
	B_TREE_TRAVERSE_IN_ORDER   = 1
	B_TREE_TRAVERSE_POST_ORDER = 2

	// degree -> number of child nodes -> number of keys
	B_TREE_MINIMUM_DEGREE  = 2
	B_TREE_CHILD_NODES_MIN = B_TREE_MINIMUM_DEGREE
	B_TREE_CHILD_NODES_MAX = B_TREE_MINIMUM_DEGREE * 2
	B_TREE_KEYS_MIN        = B_TREE_CHILD_NODES_MIN - 1
	B_TREE_KEYS_MAX        = B_TREE_CHILD_NODES_MAX - 1
)

type BTreeNode struct {
	Data   []int
	Childs []*BTreeNode
	Parent *BTreeNode
}

type BTree struct {
	root *BTreeNode
}

// essense -------------------------------------------------------------------------------------------------------------

func (t *BTree) Insert(newVal int) bool {
	if t.Contains(newVal) {
		return false
	}

	if t.root == nil {
		t.root = &BTreeNode{Data: []int{newVal}}
		return true
	}

	// find target leaf node
	node := t.root
	for len(node.Childs) != 0 {

		// (!) find a position of new value in current node
		childNodeIdx := 0
		for i := range len(node.Data) {
			if newVal < node.Data[i] {
				break
			}

			childNodeIdx++
		}

		// check if this position has a child node
		if childNodeIdx < len(node.Childs) {
			previousNode := node
			node = node.Childs[childNodeIdx]

			// preventive split of full nodes during descent from root to leaf
			t.split(previousNode)
		}
	}

	// add new value to leaf and if necessary, split it
	node.Data = append(node.Data, newVal)
	SortSliceBubble(node.Data)

	if len(node.Data) > B_TREE_KEYS_MAX {
		t.split(node)
	}

	return true
}

func (t *BTree) split(node *BTreeNode) {
	// 1. exceeds in case of adding new value to full node, 2. equals in case of encountering full node during descent to leaf
	if len(node.Data) < B_TREE_KEYS_MAX {
		return
	}

	dataSize := len(node.Data)

	// (!)
	if node.Parent == nil { // this is root
		newRightPart := &BTreeNode{}

		newRootNode := &BTreeNode{}
		newRootNode.Data = append(newRootNode.Data, node.Data[dataSize/2])
		newRootNode.Childs = append(newRootNode.Childs, node, newRightPart)
		t.root = newRootNode

		node.Parent = newRootNode
		newRightPart.Parent = newRootNode

		// split data and childs between two nodes
		newRightPart.Data = slices.Clone(node.Data[dataSize/2+1:]) // '1' is middle element, which will go to the parent
		node.Data = node.Data[:dataSize/2]

		if len(node.Childs) > 0 { // number of childs in full node will always be B_TREE_CHILD_NODES_MAX, if it is not a leaf (because new key in internal node appears only after splitting of child)
			newRightPart.Childs = slices.Clone(node.Childs[B_TREE_CHILD_NODES_MAX/2:])
			node.Childs = node.Childs[:B_TREE_CHILD_NODES_MAX/2]
		}
	} else { // this is internal node or leaf (not necessarily outermost node of parent)
		parentNode := node.Parent
		newNode := &BTreeNode{}

		// new node in proper position of parent's childs, new data in proper position of parent's data
		parentNode.Data = append(parentNode.Data, node.Data[dataSize/2])
		SortSliceBubble(parentNode.Data)

		nodeIdx := slices.Index(parentNode.Childs, node)
		parentNode.Childs = slices.Insert(parentNode.Childs, nodeIdx+1, newNode)

		newNode.Parent = parentNode

		// split data and childs between two nodes
		newNode.Data = slices.Clone(node.Data[dataSize/2+1:]) // '1' is middle element, which will go to the parent
		node.Data = node.Data[:dataSize/2]

		if len(node.Childs) > 0 { // number of childs in full node will always be B_TREE_CHILD_NODES_MAX, if it is not a leaf (because new key in internal node appears only after splitting of child)
			newNode.Childs = slices.Clone(node.Childs[B_TREE_CHILD_NODES_MAX/2:])
			node.Childs = node.Childs[:B_TREE_CHILD_NODES_MAX/2]
		}
	}
}

func (t *BTree) RemoveTry2(value int) bool {

	// basic idea: deletion always from leaf (as deletion from regular binary tree)
	// 1. if LEAF:
	// 1.1 if it has more than B_TREE_KEYS_MIN or root - just delete target value
	// 1.2 otherwise borrow key from one of two siblings (via parent), which sitisfies to requirement of B_TREE_KEYS_MIN
	// 1.2.1 in case of internal node - borrow element with its edge child from sibling (sub-tree move)
	// 1.3 if both siblings are too small - merge current node with one of them along with middle value from parent
	// 1.4 make the same thing from 1.1 to 1.3 with parent if it became also deficient
	// 2. if INTERNAL node:
	// 2.1 find largest element in left subtree or smallest in right (will be leaf), and replace deleted element by it
	// 2.2 go to 'leaf' strategy (1.)
	// P.S. or just mark value as 'removed' :)

	var node *BTreeNode
	if node = t.findNode(value); node == nil {
		return false
	}

	if len(node.Childs) == 0 { // 1.
		t.removeFromLeaf(value, node)
	} else { // 2.
		var leaf *BTreeNode
		valueIdx := slices.Index(node.Data, value)

		if successorValue, successorNode := t.getNearestValue(value, true); successorNode != nil { // 2.1
			successorIdx := slices.Index(successorNode.Data, successorValue)
			successorNode.Data[successorIdx], node.Data[valueIdx] = node.Data[valueIdx], successorNode.Data[successorIdx]

			leaf = successorNode
		} else if predecessorValue, predecessorNode := t.getNearestValue(value, false); predecessorNode != nil { // 2.1
			predecessorIdx := slices.Index(predecessorNode.Data, predecessorValue)
			predecessorNode.Data[predecessorIdx], node.Data[valueIdx] = node.Data[valueIdx], predecessorNode.Data[predecessorIdx]

			leaf = predecessorNode
		} else {
			panic("no successor or predecessor was found")
		}

		t.removeFromLeaf(value, leaf) // 2.2
	}

	return true
}

func (t *BTree) removeFromLeaf(value int, node *BTreeNode) {

	if len(node.Childs) != 0 {
		panic("deletion is only allowed from leaf node")
	}

	valueIdx := slices.Index(node.Data, value) // delete target value
	node.Data = slices.Delete(node.Data, valueIdx, valueIdx+1)

	if len(node.Data) >= B_TREE_KEYS_MIN { // 1.1 (remove)
		return
	} else if node.Parent == nil {
		if len(node.Data) == 0 {
			t.root = nil
		}
		return
	} else {
		t.fixRemove(node, value)
	}
}

func (t *BTree) fixRemove(node *BTreeNode, valueForDebug int) {

	lSib, rSib := t.getImmediateSiblings(node) // determine siblings
	var largerSibling *BTreeNode               // need for 1.2
	var lesserSibling *BTreeNode               // need for 1.3

	// TODO: what if it is a parent ?

	if lSib == nil {
		largerSibling = rSib
		lesserSibling = rSib
	} else if rSib == nil {
		largerSibling = lSib
		lesserSibling = lSib
	} else {
		if len(lSib.Data) > len(rSib.Data) {
			largerSibling = lSib
			lesserSibling = rSib
		} else {
			largerSibling = rSib
			lesserSibling = lSib
		}
	}

	if len(largerSibling.Data) > B_TREE_KEYS_MIN { // 1.2 (borrow)
		fmt.Printf("borrow case after %d removal\n", valueForDebug)

		if largerSibling == lSib {
			if parentLeftValueIdx := t.getParentValueIdx(node, true); parentLeftValueIdx != -1 {
				node.Data = append(node.Data, node.Parent.Data[parentLeftValueIdx]) // from parent to deficient node
				SortSliceBubble(node.Data)

				node.Parent.Data[parentLeftValueIdx] = lSib.Data[len(lSib.Data)-1] // from sibling to parent
				lSib.Data = lSib.Data[:len(lSib.Data)-1]

				if len(lSib.Childs) > len(lSib.Data) { // move rightmost child of sibling to node's left side
					node.Childs = slices.Insert(node.Childs, 0, lSib.Childs[len(lSib.Childs)-1])
					lSib.Childs = lSib.Childs[:len(lSib.Childs)-1]
				}
			}
		} else {
			if parentRightValueIdx := t.getParentValueIdx(node, false); parentRightValueIdx != -1 {
				node.Data = append(node.Data, node.Parent.Data[parentRightValueIdx]) // from parent to deficient node
				SortSliceBubble(node.Data)

				node.Parent.Data[parentRightValueIdx] = rSib.Data[0] // from sibling to parent
				rSib.Data = rSib.Data[1:]

				if len(rSib.Childs) > len(rSib.Data) { // move leftmost child of sibling to node's right side
					node.Childs = append(node.Childs, rSib.Childs[0])
					rSib.Childs = rSib.Childs[1:]
				}
			}
		}
	} else { // 1.3 (merge)
		fmt.Printf("merge case after %d removal\n", valueForDebug)

		if lesserSibling == lSib {
			if parentLeftValueIdx := t.getParentValueIdx(node, true); parentLeftValueIdx != -1 {
				lesserSibling.Data = append(lesserSibling.Data, node.Parent.Data[parentLeftValueIdx]) // middle element from parent and deficient node's data to left sibling
				lesserSibling.Data = append(lesserSibling.Data, node.Data...)

				node.Parent.Data = slices.Delete(node.Parent.Data, parentLeftValueIdx, parentLeftValueIdx+1)
				if len(node.Parent.Data) == 0 {
					t.root = node
				} else {
					nodeIdx := slices.Index(node.Parent.Childs, node) // now parent has only left sibling as a child, deficient node is deleted
					node.Parent.Childs = slices.Delete(node.Parent.Childs, nodeIdx, nodeIdx+1)
				}
			}
		} else {
			if parentRightValueIdx := t.getParentValueIdx(node, false); parentRightValueIdx != -1 {
				node.Data = append(node.Data, node.Parent.Data[parentRightValueIdx]) // middle element from parent and right sibling's data to deficient node
				node.Data = append(node.Data, lesserSibling.Data...)

				node.Parent.Data = slices.Delete(node.Parent.Data, parentRightValueIdx, parentRightValueIdx+1)
				if len(node.Parent.Data) == 0 {
					t.root = node
				} else {
					lesserSiblingIdx := slices.Index(node.Parent.Childs, lesserSibling) // now parent has only deficient node as a child, right sibling is deleted
					node.Parent.Childs = slices.Delete(node.Parent.Childs, lesserSiblingIdx, lesserSiblingIdx+1)
				}
			}
		}

		// do the same for parent if it is also broken
		if len(node.Parent.Data) < B_TREE_KEYS_MIN {
			t.fixRemove(node.Parent, -1)
		}
	}
}

func (t *BTree) getNearestValue(value int, successor bool) (int, *BTreeNode) {
	leftChild, rightChild := t.getImmediateChilds(value)

	if successor {
		if rightChild == nil {
			return 0, nil
		}

		successorNode := rightChild
		for len(successorNode.Childs) > 0 { // until reach a leaf node
			successorNode = successorNode.Childs[0]
		}

		return successorNode.Data[0], successorNode
	} else {
		if leftChild == nil {
			return 0, nil
		}

		predecessorNode := leftChild
		for len(predecessorNode.Childs) > 0 { // until reach a leaf node
			predecessorNode = predecessorNode.Childs[len(predecessorNode.Childs)-1]
		}

		return predecessorNode.Data[len(predecessorNode.Data)-1], predecessorNode
	}
}

// utils ---------------------------------------------------------------------------------------------------------------

func NewBTree() *BTree {
	return &BTree{}
}

func count(node *BTreeNode) int {
	if node == nil {
		return 0
	}

	result := len(node.Data)
	for _, child := range node.Childs {
		result += count(child)
	}

	return result
}

func (t *BTree) GetSize() int {
	return count(t.root)
}

func (t *BTree) Contains(value int) bool {
	return t.findNode(value) != nil
}

func (t *BTree) findNode(value int) *BTreeNode {
	if t.root == nil {
		return nil
	}

	currNode := t.root
	for {
		// find a position of new value in current node, and if this value will be found during search, then it is done
		greaterValueIdx := 0
		for i := range len(currNode.Data) {
			if value < currNode.Data[i] {
				break
			}

			if value == currNode.Data[i] {
				return currNode
			}

			greaterValueIdx++
		}

		// check if this position has a child node
		if greaterValueIdx < len(currNode.Childs) && currNode.Childs[greaterValueIdx] != nil {
			currNode = currNode.Childs[greaterValueIdx]
			continue
		}

		break
	}

	return nil
}

func cloneBTreeNode(node *BTreeNode) *BTreeNode {
	if node == nil {
		return nil
	}

	newNode := &BTreeNode{}
	newNode.Data = slices.Clone(node.Data)

	for _, child := range node.Childs {
		newChild := cloneBTreeNode(child)
		newChild.Parent = newNode
		newNode.Childs = append(newNode.Childs, newChild)
	}

	return newNode
}

func (t *BTree) Clone() *BTree {
	clone := &BTree{}
	clone.root = cloneBTreeNode(t.root)

	return clone
}

func isNodeValid(node *BTreeNode) error {
	if len(node.Data) < B_TREE_KEYS_MIN || len(node.Data) > B_TREE_KEYS_MAX {
		return fmt.Errorf("incorrect number of keys in node: %v", node.Data)
	}

	if len(node.Childs) != 0 && (len(node.Childs) < B_TREE_CHILD_NODES_MIN || len(node.Childs) > B_TREE_CHILD_NODES_MAX) {
		return fmt.Errorf("incorrect number of childs in node: %v", node.Childs)
	}

	if len(node.Childs) != 0 && len(node.Data) != len(node.Childs)+1 {
		return fmt.Errorf("incorrect ratio of number of keys and childs: %v - %v", node.Data, node.Childs)
	}

	for _, child := range node.Childs {
		if err := isNodeValid(child); err != nil {
			return err
		}
	}

	return nil
}

func collectLeafs(node *BTreeNode) []*BTreeNode {

	leafs := []*BTreeNode{}

	if len(node.Childs) == 0 {
		leafs = append(leafs, node)
		return leafs
	}

	for _, child := range node.Childs {
		leafs = append(leafs, collectLeafs(child)...)
	}

	return leafs
}

func (t *BTree) IsValid() error {
	if t.root == nil {
		return nil
	}

	if err := isNodeValid(t.root); err != nil {
		return err
	}

	height := -1
	for _, leaf := range collectLeafs(t.root) {
		currHeight := 0
		node := leaf

		for node != nil {
			currHeight++
			node = node.Parent
		}

		if height == -1 {
			height = currHeight
		}

		if currHeight != height {
			return fmt.Errorf("height differs in leafs")
		}
	}

	return nil
}

// TODO: obsolete
func (t *BTree) getSibling(node *BTreeNode, larger bool) *BTreeNode {

	if node.Parent == nil {
		return nil
	}

	// find siblings around node
	nodeIdx := slices.Index(node.Parent.Childs, node)

	var leftSibling, rightSibling *BTreeNode
	if nodeIdx > 0 {
		leftSibling = node.Parent.Childs[nodeIdx-1]
	}

	if len(node.Parent.Childs) >= nodeIdx+2 {
		rightSibling = node.Parent.Childs[nodeIdx+1]
	}

	// choose appropriate one
	if rightSibling == nil {
		return leftSibling
	} else if leftSibling == nil {
		return rightSibling
	} else {
		if larger {
			if len(rightSibling.Data) > len(leftSibling.Data) {
				return rightSibling
			} else {
				return leftSibling
			}
		} else {
			if len(rightSibling.Data) < len(leftSibling.Data) {
				return rightSibling
			} else {
				return leftSibling
			}
		}
	}
}

func (t *BTree) getImmediateSiblings(node *BTreeNode) (*BTreeNode, *BTreeNode) {
	if node.Parent == nil {
		return nil, nil
	}

	nodeIdx := slices.Index(node.Parent.Childs, node)
	if nodeIdx == -1 {
		panic("node is not found in childs container")
	}

	var leftSibling, rightSibling *BTreeNode
	if nodeIdx > 0 {
		leftSibling = node.Parent.Childs[nodeIdx-1]
	}

	if len(node.Parent.Childs) >= nodeIdx+2 {
		rightSibling = node.Parent.Childs[nodeIdx+1]
	}

	return leftSibling, rightSibling
}

// TODO: obsolete
func (t *BTree) getChild(node *BTreeNode, value int, larger bool) *BTreeNode {

	valueIdx := slices.Index(node.Data, value)
	if valueIdx == -1 {
		panic("value is not found in container")
	}

	// find value's childs in node
	var leftChild, rightChild *BTreeNode
	if len(node.Childs) >= valueIdx+1 && node.Childs[valueIdx] != nil {
		leftChild = node.Childs[valueIdx]
	}

	if len(node.Childs) >= valueIdx+2 && node.Childs[valueIdx+1] != nil {
		rightChild = node.Childs[valueIdx+1]
	}

	// choose appropriate one
	if leftChild == nil {
		return rightChild
	} else if rightChild == nil {
		return leftChild
	} else {
		if larger {
			if len(leftChild.Data) > len(rightChild.Data) {
				return leftChild
			} else {
				return rightChild
			}
		} else {
			if len(leftChild.Data) < len(rightChild.Data) {
				return leftChild
			} else {
				return rightChild
			}
		}
	}
}

func (t *BTree) getImmediateChilds(value int) (*BTreeNode, *BTreeNode) {
	var node *BTreeNode
	if node = t.findNode(value); node == nil {
		return nil, nil
	}

	valueIdx := slices.Index(node.Data, value)
	if valueIdx == -1 {
		panic("value is not found in data container")
	}

	var leftChild, rightChild *BTreeNode
	if valueIdx < len(node.Childs) {
		leftChild = node.Childs[valueIdx]
	}

	if valueIdx+1 < len(node.Childs) {
		rightChild = node.Childs[valueIdx+1]
	}

	return leftChild, rightChild
}

func (t *BTree) getParentValueIdx(node *BTreeNode, leftValue bool) int {
	if node.Parent == nil {
		return -1
	}

	nodeIdx := slices.Index(node.Parent.Childs, node)
	if nodeIdx == -1 {
		panic("node is not found in childs container")
	}

	if leftValue {
		if nodeIdx == 0 {
			return -1
		}

		return nodeIdx - 1
	} else {
		if nodeIdx == len(node.Parent.Childs)-1 {
			return -1
		}

		return nodeIdx
	}
}

func (t *BTree) AcceptVisitor(v common.INodeVisitor) {
	t.traverse(t.root, v)
}

func (t *BTree) traverse(node *BTreeNode, v common.INodeVisitor) {
	if node == nil {
		return
	}

	serializeNodeData := func(node *BTreeNode) string {
		data := ""

		if node == nil {
			return data
		}

		for i := range len(node.Data) {
			data += strconv.Itoa(node.Data[i])
			if i+1 < len(node.Data) {
				data += "-"
			}
		}

		return data
	}

	// collect data
	dataStr := serializeNodeData(node)
	var childsData []string
	for _, childNode := range node.Childs {
		childsData = append(childsData, serializeNodeData(childNode))
	}

	// temporarily onle pre-order traverse
	v.VisitNode(dataStr, childsData)

	for _, childNode := range node.Childs {
		t.traverse(childNode, v)
	}
}

func (t *BTree) Clear() {
	for i := range B_TREE_CHILD_NODES_MAX {
		t.root.Childs[i] = nil // TODO: will it delete all children nodes by GC ?
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// O(n^2)
func SortSliceBubble(slice []int) {

	size := len(slice)
	if size < 2 {
		return
	}

	// idea: if element meets criteria, it remains in game up to the of container
	//           i <- (shrink unsorted range from the end)
	// +----------------------------------
	// | 5 4 1 2 3 6 7 9 ...
	// +------------------------
	//   j -> (compare two neighbour elements and swap them if needed)

	for i := size; i > 0; i-- { // upper bound to which the element will be moved
		for j := 0; j < i-1; j++ { // move the elements (current min/max will be picked up 'by wave')
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// O(n^2)
func SortSliceSelection(slice []int) {

	size := len(slice)
	if size < 2 {
		return
	}

	// idea: select the best candidate from unsorted range and put it at the end of sorted range
	//             i -> (shrink unsorted range from the begin)
	// +----------------------------------
	// | 2 4 5 6 7 9 8 ...
	// +------------------------
	//             j -> (search next min/max element by comparing with first unsorted one)

	for i := 0; i < size; i++ { // place where the element search starts
		minIdx := i
		for j := i + 1; j < size; j++ { // search element
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}

// O(n^2)
func SortSliceInsertion(slice []int) {

	size := len(slice)
	if size < 2 {
		return
	}

	// idea: pick the next element and find place for it in sorted range
	//           i -> (take next element from unsorted range)
	// +----------------------------------
	// | 5 6 7 8 9 2 ...
	// +------------------------
	//        <- j (move it by comparing to appropriate place in sorted range)

	for i := 0; i < size; i++ { // candidate for insertion
		for j := i; j > 0; j-- { // find its place in sorted range
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

// best case O(n*log(n))
func SortSliceShell(slice []int) {

	size := len(slice)
	if size < 2 {
		return
	}

	// idea: move smaller item many spaces to the left without shifting all intermediate items individually
	//                   i -> (move h-window with 1-cell step)
	// +--------------------------------------
	// | 5 - - - 4 - - - 3 ...
	// +----------------------------
	//                <- j (compare all element with h-step)

	h := 1
	for h*3 <= size {
		h = h*3 + 1
	}

	for h > 0 { // size of window
		for i := h; i < size; i++ { // move chain of windows forward by 1-cell step, from h-position to start comparing right away
			for j := i; j-h >= 0; j -= h { // check values of windows
				if slice[j] < slice[j-h] {
					slice[j-h], slice[j] = slice[j], slice[j-h]
				}
			}
		}

		h = (h - 1) / 3
	}
}

func Merge2(lBound, border, rBound int, buffer []int) {

	if lBound == rBound {
		return
	}

	lSize := border - lBound
	if lSize > 1 {
		Merge2(lBound, lBound+lSize/2, border, buffer)
	}

	rSize := rBound - border
	if rSize > 1 {
		Merge2(border, border+rSize/2, rBound, buffer)
	}

	// TODO: do
}

func Merge(left []int, right []int, buffer []int) {

	lSize := len(left)
	if lSize > 1 {
		Merge(left[:lSize/2], left[lSize/2:], buffer)
	}

	rSize := len(right)
	if rSize > 1 {
		Merge(right[:rSize/2], right[rSize/2:], buffer)
	}

	// pick from left, right in sorted order and put into buffer
	li, ri, bi := 0, 0, 0
	for li < lSize && ri < rSize {
		if left[li] < right[ri] {
			buffer[bi] = left[li]
			li++
			bi++
		} else {
			buffer[bi] = right[ri]
			ri++
			bi++
		}
	}

	for ; li < lSize; li++ {
		buffer[bi] = left[li]
		bi++
	}

	for ; ri < rSize; ri++ {
		buffer[bi] = right[ri]
		bi++
	}

	// sorted buffer back to left, right
	for i := range lSize {
		left[i] = buffer[i]
	}

	for i := range rSize {
		right[i] = buffer[lSize+i]
	}
}

// O(n * log(n))
// TODO: advantage of this sort algorithm
func SortSliceMerge(slice []int) {

	size := len(slice)
	if size < 2 {
		return
	}

	buffer := make([]int, size) // requires O(n) space complexity

	half := size / 2
	Merge(slice[:half], slice[half:], buffer)
}
