package transitions

import (
	"fmt"
	"nickaashoek/tdviz/pkg/robdd"
)

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

func BFSInverseIndex(rbd *robdd.ROBDD, root int) map[int][]int {
	visited := make(map[int]bool)

	inverseIndex := make(map[int][]int)
	inverseIndex[root] = make([]int, 0)

	queue := intQueue{}
	queue.Push(root)
	for queue.Full() {
		elem := queue.Pop()
		if _, visited := visited[elem]; visited {
			continue
		}
		visited[elem] = true
		if elem > 1 {
			node := rbd.Nodes[elem]

			if _, p := inverseIndex[node.Lo]; !p {
				inverseIndex[node.Lo] = make([]int, 0)
			}
			inverseIndex[node.Lo] = append(inverseIndex[node.Lo], -elem)

			if _, p := inverseIndex[node.Hi]; !p {
				inverseIndex[node.Hi] = make([]int, 0)
			}
			inverseIndex[node.Hi] = append(inverseIndex[node.Hi], elem)

			queue.Push(node.Lo)
			queue.Push(node.Hi)
		}
	}
	return inverseIndex
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func _generateValidTransitions(rbd *robdd.ROBDD, node int, inverseIndex map[int][]int) [][]int {
	result := make([][]int, 0)
	if len(inverseIndex[node]) == 0 {
		result = append(result, make([]int, 0))
	}
	for _, parent := range inverseIndex[node] {
		for _, partial := range _generateValidTransitions(rbd, abs(parent), inverseIndex) {
			partial = append(partial, parent)
			result = append(result, partial)
		}
	}
	return result
}

func GenerateAllValidTransitions(rbd *robdd.ROBDD, root int) [][]int {
	inverseIndex := BFSInverseIndex(rbd, root)
	results := _generateValidTransitions(rbd, 1, inverseIndex)
	for i, _ := range results {
		results[i] = append(results[i], 1)
	}
	fmt.Printf("Results: %v\n", results)
	return results
}
