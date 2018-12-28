Max/Min heap implements by Golang

Usage:
------
```go
package main

import(
	"fmt"
	"math/rand"
	"github/liexusong/golib/heap"
)

type HeapValue struct {
	Key int64
	Val int64
}

// must implement GetHeapCompareIndex() method
func (e *HeapValue) GetHeapCompareIndex() int64 {
	return e.Key
}

func (e *HeapValue) GetValue() int64 {
	return e.Val
}

func main() {
	h := heap.NewMaxHeap()
	
	for i := 0; i < 100; i++ {
        randVal := rand.Int63n(100)
        elem := &HeapValue{
            Key: randVal,
            Val: randVal,
        }
        h.Push(elem)
    }

    for {
        elem := h.Pop()
        if elem == nil {
            break
        }
        fmt.Println(elem.(*HeapValue).GetValue())
    }
}
```