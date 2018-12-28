// Copyright 2018 LieXuSong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package heap

type HeapElement interface {
	GetHeapCompareIndex() int64
}

type Heap struct {
	bucket []HeapElement
	hType  int
}

const (
	heapInitSize = 128
	HeapTypeMax  = 1
	HeapTypeMin  = 2
)

func NewMaxHeap() *Heap {
	return New(HeapTypeMax)
}

func NewMinHeap() *Heap {
	return New(HeapTypeMin)
}

func New(hType int) *Heap {
	return &Heap{
		bucket: make([]HeapElement, 1, heapInitSize),
		hType:  hType,
	}
}

func heapCompareFunc(v1 HeapElement, v2 HeapElement) int {
	idx1, idx2 := v1.GetHeapCompareIndex(), v2.GetHeapCompareIndex()

	if idx1 > idx2 {
		return 1
	} else if idx1 == idx2 {
		return 0
	}
	return -1
}

func (h *Heap) Empty() bool {
	if len(h.bucket) <= 1 {
		return true
	}
	return false
}

func (h *Heap) Push(v HeapElement) {
	currentIdx := len(h.bucket)

	h.bucket = append(h.bucket, v) // append the heap buckets tail

	for ; currentIdx > 1; currentIdx = currentIdx / 2 {
		parentIdx := currentIdx / 2

		result := heapCompareFunc(h.bucket[currentIdx], h.bucket[parentIdx])

		switch h.hType {
		case HeapTypeMax:
			if result <= 0 {
				return
			}
		case HeapTypeMin:
			if result >= 0 {
				return
			}
		}

		// swap child and parent values
		h.bucket[currentIdx], h.bucket[parentIdx] = h.bucket[parentIdx], h.bucket[currentIdx]
	}
}

func (h *Heap) Pop() HeapElement {
	if h.Empty() {
		return nil
	}

	topVal := h.bucket[1] // index 1 is the top element

	if len(h.bucket) <= 2 {
		h.bucket = h.bucket[0:1:cap(h.bucket)] // delete all elements
	} else {
		h.bucket[1] = h.bucket[len(h.bucket)-1]  // save last element to top position
		h.bucket = h.bucket[0 : len(h.bucket)-1] // delete last element
	}

	length := len(h.bucket)

loop:
	for lastIdx := 1; lastIdx < length-1; {
		leftIdx, rightIdx := lastIdx*2, lastIdx*2+1

		if leftIdx > length-1 {
			break loop
		}

		swapIdx := leftIdx

		if rightIdx <= length-1 {
			result := heapCompareFunc(h.bucket[rightIdx], h.bucket[leftIdx])
			switch h.hType {
			case HeapTypeMax:
				if result > 0 {
					swapIdx = rightIdx
				}
			case HeapTypeMin:
				if result < 0 {
					swapIdx = rightIdx
				}
			}
		}

		result := heapCompareFunc(h.bucket[swapIdx], h.bucket[lastIdx])

		switch h.hType {
		case HeapTypeMax:
			if result <= 0 {
				break loop
			}
		case HeapTypeMin:
			if result >= 0 {
				break loop
			}
		}

		h.bucket[swapIdx], h.bucket[lastIdx] = h.bucket[lastIdx], h.bucket[swapIdx]

		lastIdx = swapIdx
	}

	return topVal
}

func (h *Heap) Top() HeapElement {
	if h.Empty() {
		return nil
	}
	return h.bucket[1]
}
