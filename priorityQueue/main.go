// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	Value    string // The value of the item; arbitrary.
	Priority int    // The priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)
	bs, _ := json.Marshal(pq)
	fmt.Println("After Init:\t", string(bs))

	// Insert a new item and then modify its priority.
	item := &Item{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	bs, _ = json.Marshal(pq)
	fmt.Println("After Push:\t", string(bs))
	pq.update(item, item.Value, 5)
	bs, _ = json.Marshal(pq)
	fmt.Println("After Update:\t", string(bs))

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("\nPop %.2d:%s \n", item.Priority, item.Value)
		bs, _ = json.Marshal(pq)
		fmt.Println("After Pop:\t", string(bs))
	}
}
