package robdd

import (
	"encoding/json"
	"os"
)

type JsonIntermediate struct {
	// Nodes    []ROBDDNode    `json:"Nodes"`
	Root     int            `json:"Root"`
	Order    map[string]int `json:"Order"`
	RevOrder map[int]string `json:"ReverseOrder"`
	Cache    map[string]int `json:"Cache"`
	AllNodes []ROBDDNode    `json:"AllNodes"`
}

type intQueue struct {
	queue []int
}

func (i *intQueue) Push(x int) {
	i.queue = append(i.queue, x)
}

func (i *intQueue) Pop() int {
	x := i.queue[0]
	i.queue = i.queue[1:]
	return x
}

func (i intQueue) Full() bool {
	return len(i.queue) > 0
}

func (rbd *ROBDD) BFSRoot(root int) {
	queue := intQueue{}
	queue.Push(root)
	for queue.Full() {
		elem := queue.Pop()
		if elem > 1 {
			node := rbd.Nodes[elem]
			queue.Push(node.Lo)
			queue.Push(node.Hi)
		}
	}
}

func BddToIntermediate(rbd *ROBDD, root int, order map[string]int) JsonIntermediate {
	result := JsonIntermediate{}
	// result.Nodes = make([]ROBDDNode, 0)

	// rbd.BFSRoot(root, &result)
	result.Root = root

	result.AllNodes = rbd.Nodes
	result.RevOrder = rbd.RevOrder
	result.Order = order
	result.Cache = rbd.Cache

	return result
}

func DumpIntermediate(intermediate *JsonIntermediate) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(intermediate)
}
