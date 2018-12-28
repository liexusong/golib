package heap

type HeapValue interface {
	GetSortValue() int
}

type Heap struct {
	bucket []HeapValue
	hType  int
}

const (
	heapInitSize = 64
	HeapTypeMax  = 1
	HeapTypeMin  = 2
)

func New(hType int) *Heap {
	return &Heap{
		bucket: make([]HeapValue, 0, heapInitSize),
		hType:  hType,
	}
}

func heapComparePeer(v1 HeapValue, v2 HeapValue) int {
	sortVal1, sortVal2 := v1.GetSortValue(), v2.GetSortValue()

	if sortVal1 > sortVal2 {
		return 1
	} else if sortVal1 == sortVal2 {
		return 0
	}
	return -1
}

func (h *Heap) Push(v HeapValue) {
	currentIdx := len(h.bucket)

	h.bucket = append(h.bucket, v)

	for ; currentIdx > 0; currentIdx = currentIdx / 2 {
		parentIdx := currentIdx / 2

		result := heapComparePeer(h.bucket[currentIdx], h.bucket[parentIdx])

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

func (h *Heap) Pop() HeapValue {
	if len(h.bucket) == 0 {
		return nil
	}

	retVal := h.bucket[0]

	h.bucket[0] = h.bucket[len(h.bucket)-1]
	h.bucket = h.bucket[0 : len(h.bucket)-1] // delete last element

loop:
	for lastIdx := 0; lastIdx < len(h.bucket)-1; {
		leftIdx, rightIdx := lastIdx*2, lastIdx*2+1

		winner := leftIdx

		if rightIdx <= len(h.bucket)-1 {
			result := heapComparePeer(h.bucket[rightIdx], h.bucket[leftIdx])
			switch h.hType {
			case HeapTypeMax:
				if result > 0 {
					winner = rightIdx
				}
			case HeapTypeMin:
				if result < 0 {
					winner = rightIdx
				}
			}
		}

		result := heapComparePeer(h.bucket[winner], h.bucket[lastIdx])

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

		h.bucket[winner], h.bucket[lastIdx] = h.bucket[lastIdx], h.bucket[winner]

		lastIdx = winner
	}

	return retVal
}

func (h *Heap) Top() HeapValue {
	if len(h.bucket) == 0 {
		return nil
	}
	return h.bucket[0]
}
