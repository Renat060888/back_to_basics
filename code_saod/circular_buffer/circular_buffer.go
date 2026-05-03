package circular_buffer

import (
	"fmt"
)

type CircularBuffer struct {
	data []int
	mask int

	// 1st approach
	idx  int
	size int

	// 2nd approach
	beginIdx int
	endIdx   int
}

func NewCircularBuffer(size int) (*CircularBuffer, error) {

	quotent := size
	for quotent > 1 {
		if quotent%2 != 0 {
			return nil, fmt.Errorf("size must be power of 2")
		}

		quotent /= 2
	}

	result := &CircularBuffer{}
	result.mask = size - 1 // 0b0100 - 1 = 0b0011 (4 is number of elements, 0-3 is range of indexing)
	result.idx = -1
	result.data = make([]int, size)

	return result, nil
}

func (cb *CircularBuffer) Push(newVal int) {

	cb.idx++
	cb.idx = cb.idx & cb.mask

	cb.data[cb.idx] = newVal

	if cb.size < (cb.mask + 1) {
		cb.size++
	}

	fmt.Printf("ADD size: %d, idx: %d, data; %v\n", cb.GetSize(), cb.idx, cb.data)
}

func (cb *CircularBuffer) Pop() *int {

	if cb.size == 0 {
		return nil
	}

	var result *int

	if cb.idx+1 >= cb.size {
		result = &cb.data[cb.idx+1-cb.size]
	} else {
		result = &cb.data[cap(cb.data)-(cb.size-(cb.idx+1))]
	}

	cb.size--

	fmt.Printf("SUB size: %d, idx: %d, data; %v\n", cb.GetSize(), cb.idx, cb.data)

	return result
}

func (cb *CircularBuffer) GetSize() int {
	return cb.size
}

func (cb *CircularBuffer) Clear() {

	cb.data = cb.data[:0] // clear slice, but do not touch allocated memory

	cb.idx = -1 // moves before push
	cb.size = 0

	cb.beginIdx = 0
	cb.endIdx = 0 // moves after push
}

func (cb *CircularBuffer) Push2(newVal int) {

	if cb.size == len(cb.data) && cb.endIdx == cb.beginIdx {
		cb.beginIdx++
		cb.beginIdx = cb.beginIdx & cb.mask
	}

	cb.data[cb.endIdx] = newVal

	cb.endIdx++
	cb.endIdx = cb.endIdx & cb.mask

	if cb.size < cb.mask+1 {
		cb.size++
	}

	fmt.Printf("ADD size: %d, begin idx: %d, end idx: %d, data; %v\n", cb.GetSize(), cb.beginIdx, cb.endIdx, cb.data)
}

func (cb *CircularBuffer) Pop2() *int {

	if cb.size == 0 {
		return nil
	}

	result := &cb.data[cb.beginIdx]

	cb.beginIdx++
	cb.beginIdx = cb.beginIdx & cb.mask

	cb.size--

	fmt.Printf("SUB size: %d, begin idx: %d end idx: %d, data; %v\n", cb.GetSize(), cb.beginIdx, cb.endIdx, cb.data)

	return result
}
