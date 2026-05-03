package heap

import (
	"fmt"
	"math"
	"saod/common"
	"slices"
	"strconv"
)

// heap
const (
	HEAP_NIL_NODE_IDX = 0xFFFFFFFF
	HEAP_POLICY_MAX   = 1
	HEAP_POLICY_MIN   = 0
)

// heap as weak-ordered binary tree placed into array
//      ______
//  _  |     |
// | |\|     |\
// 0 1 2 3 4 5 6 ...
//   |   |/
//   |___|
//

// basic idea:
// - add new value to the end of array, and then trickle it up, until greater value will be encountered
// - trickle is performed by formula idx*2+1 (+2) to move down, and idx/2-0 (-1 for even idx) to move up
// - reason of multiplying and dividing by 2 is that each layer of tree has the same number of nodes as all previous layers combined
// - so the number of nodes before target node and number of them between this one and its child node is the same
// my opinion about ordering:
// - it is necessary to make sure that main rule is not violated - at the top of heap we have max or min value depending on policy
// - but this ordering is weak because we, roughly speaking, do not care about the rest of elements in container

type Heap struct {
	data   []int
	policy int
}

// essense -------------------------------------------------------------------------------------------------------------

func (h *Heap) GetChildIdx(parentIdx int, left bool) int {

	childIdx := parentIdx * 2
	if left {
		childIdx += 1
	} else {
		childIdx += 2
	}

	if childIdx >= h.GetSize() {
		return HEAP_NIL_NODE_IDX
	}

	return childIdx
}

func (h *Heap) GetParentIdx(childIdx int) int {

	if childIdx <= 0 {
		return HEAP_NIL_NODE_IDX
	}

	parentIdx := childIdx / 2
	if childIdx%2 == 0 {
		parentIdx -= 1
	}

	return parentIdx
}
func (h *Heap) AddValue(newVal int) error {

	if slices.Contains(h.data, newVal) {
		return fmt.Errorf("such value already exists: %d", newVal)
	}

	h.data = append(h.data, newVal)

	if h.GetSize() == 1 { // nowhere to trickle
		return nil
	}

	trickledCurrentIdx := h.GetSize() - 1
	parentIdx := h.GetParentIdx(trickledCurrentIdx)

	// trickle new value from the end of array to up
	for parentIdx != HEAP_NIL_NODE_IDX {
		if h.GetBestCandidate(h.data[trickledCurrentIdx], h.data[parentIdx]) {
			h.data[trickledCurrentIdx], h.data[parentIdx] = h.data[parentIdx], h.data[trickledCurrentIdx]
			trickledCurrentIdx = parentIdx

			parentIdx = h.GetParentIdx(trickledCurrentIdx)
		} else {
			// next element turned out to be greater or lesser depending on policy
			break
		}
	}

	return nil
}

func (h *Heap) GetBestCandidate(leftChildData, rightChildData int) bool {
	if h.policy == HEAP_POLICY_MAX {
		return leftChildData > rightChildData
	} else {
		return leftChildData < rightChildData
	}
}

func (h *Heap) GetChildIdxForTrickle(parentIdx int) int {

	leftChildIdx := h.GetChildIdx(parentIdx, true)
	rightChildIdx := h.GetChildIdx(parentIdx, false)

	if leftChildIdx == HEAP_NIL_NODE_IDX && rightChildIdx == HEAP_NIL_NODE_IDX {
		return HEAP_NIL_NODE_IDX
	}

	if leftChildIdx != HEAP_NIL_NODE_IDX && rightChildIdx != HEAP_NIL_NODE_IDX {
		if h.GetBestCandidate(h.data[leftChildIdx], h.data[rightChildIdx]) {
			// TODO: what if both childs are smaller than trickled 'min' value ?
			return leftChildIdx
		} else {
			return rightChildIdx
		}
	} else if leftChildIdx != HEAP_NIL_NODE_IDX {
		if !h.GetBestCandidate(h.data[parentIdx], h.data[leftChildIdx]) { // when comparing Parent and Child - don't exchange them
			return leftChildIdx
		} else {
			return HEAP_NIL_NODE_IDX // heap is FULL weak-balanced tree - if there is only one child - it is left
		}
	} else {
		panic("heap cannot have only right child without left one")
	}
}

func (h *Heap) TakeValue() *int {

	if h.GetSize() == 0 {
		return nil
	}

	value := h.data[0]

	if h.GetSize() == 1 {
		h.data = h.data[:0] // clear slice, but do not touch allocated memory
		return &value
	}

	// move last element to the root and trickle it down
	trickledCurrentIdx := 0
	h.data[trickledCurrentIdx] = h.data[h.GetSize()-1]
	h.data = h.data[:h.GetSize()-1]

	for {
		childIdxForTrickle := h.GetChildIdxForTrickle(trickledCurrentIdx)
		if childIdxForTrickle == HEAP_NIL_NODE_IDX {
			break
		}

		h.data[childIdxForTrickle], h.data[trickledCurrentIdx] = h.data[trickledCurrentIdx], h.data[childIdxForTrickle]
		trickledCurrentIdx = childIdxForTrickle
	}

	return &value
}

// utils ---------------------------------------------------------------------------------------------------------------

func NewHeap(policy int) *Heap {

	return &Heap{policy: policy}
}

func (h *Heap) GetSize() int {
	return len(h.data)
}

func (h *Heap) GetCapacity() int {
	return cap(h.data)
}

func (h *Heap) Clear() {
	if h.GetSize() > 0 {
		h.data = h.data[:0] // clear slice, but do not touch allocated memory
	}
}

func (h *Heap) AcceptVisitor(v common.INodeVisitor) {

	for parentIdx, parentData := range h.data {

		var childsData []string
		if childIdx := h.GetChildIdx(parentIdx, true); childIdx != HEAP_NIL_NODE_IDX {
			childsData = append(childsData, strconv.Itoa(h.data[childIdx]))
		}

		if childIdx := h.GetChildIdx(parentIdx, false); childIdx != HEAP_NIL_NODE_IDX {
			childsData = append(childsData, strconv.Itoa(h.data[childIdx]))
		}

		v.VisitNode(strconv.Itoa(parentData), childsData)
	}
}

func (h *Heap) ConvertToStr() string {

	powerOfTwo := 0.0
	entitiesToProcess := math.Pow(2.0, powerOfTwo)

	str := ""
	for i, v := range h.data {
		str += strconv.Itoa(v)

		if i != h.GetSize()-1 {
			entitiesToProcess--
			if entitiesToProcess == 0 {
				str += "\n"
				powerOfTwo++
				entitiesToProcess = math.Pow(2, float64(powerOfTwo))
			} else {
				if i%2 == 0 {
					str += "|"
				} else {
					str += " "
				}
			}
		}
	}

	return str
}
