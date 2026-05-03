package main

import (
	"testing"

	"saod/heap"
)

func TestHeapInitialState(t *testing.T) {

	h := heap.NewHeap(heap.HEAP_POLICY_MAX)

	if h.GetSize() != 0 || h.GetCapacity() != 0 {
		t.Errorf("newly created heap has non-zero size or capacity")
	}

	if h.TakeValue() != nil {
		t.Errorf("newly created heap has some data")
	}
}

func TestHeapIndexing(t *testing.T) {

	h := heap.NewHeap(heap.HEAP_POLICY_MAX)

	if h.GetParentIdx(0) != heap.HEAP_NIL_NODE_IDX {
		t.Errorf("heap's root does not have NIL parent")
	}

	if h.GetParentIdx(1) != 0 {
		t.Errorf("parent idx of element[1] is not 0")
	}

	if h.GetParentIdx(3) != 1 {
		t.Errorf("parent idx of element[3] is not 1")
	}

	if h.GetParentIdx(5) != 2 {
		t.Errorf("parent idx of element[5] is not 2")
	}

	if h.GetParentIdx(7) != 3 {
		t.Errorf("parent idx of element[7] is not 3")
	}
}

func TestHeapDataInsertion(t *testing.T) {

	h := heap.NewHeap(heap.HEAP_POLICY_MAX)

	/* expected result
		       100
		    30     50
	      10  20 15  40
	*/

	// 3rd level
	h.AddValue(10)
	h.AddValue(15)

	h.AddValue(20)
	h.AddValue(30)

	// 2nd level
	h.AddValue(40)

	h.AddValue(50)

	// 1st level
	h.AddValue(100)

	if val := h.TakeValue(); val == nil || *val != 100 {
		t.Errorf("maximum number in heap is not found or is not 100")
	}
}

func TestHeapDataRemoval(t *testing.T) {

	h := heap.NewHeap(heap.HEAP_POLICY_MAX)

	// 3rd level
	h.AddValue(10)
	h.AddValue(15)

	h.AddValue(20)
	h.AddValue(30)

	// 2nd level
	h.AddValue(40)

	h.AddValue(50)

	// 1st level
	h.AddValue(100)

	if val := h.TakeValue(); val == nil || *val != 50 {
		t.Errorf("maximum number in heap is not found or is not 50 (after removing 100)")
	}
}
