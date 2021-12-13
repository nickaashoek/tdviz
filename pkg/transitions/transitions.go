package transitions

import (
	"fmt"
	"math"
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

type StateTransition struct {
	Start string `json:"Start"`
	End   string `json:"End"`
}

func (st StateTransition) String() string {
	return fmt.Sprintf("%s -> %s", st.Start, st.End)
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

func PropValuesFromPath(rbd *robdd.ROBDD, path []int) map[string]int {
	result := make(map[string]int)
	// Last item in path is always 1 terminal, we can ignore that
	for _, item := range path[:len(path)-1] {
		// Item may be negative
		node := rbd.Nodes[abs(item)]
		value := 1
		if abs(item) != item {
			// If negative, represents a lo branch
			value = 0
		}
		result[rbd.RevOrder[node.Prop]] = value
	}
	return result
}

func IsEnd(prop string) bool {
	return prop[len(prop)-1] == '\''
}

func NodeToValue(node int) int {
	value := 0
	if abs(node) == node {
		value = 1
	}
	return value
}

// NodeToProp returns the string representing the prop
func NodeToProp(rbd *robdd.ROBDD, node int) string {
	return rbd.RevOrder[rbd.Nodes[abs(node)].Prop]
}

func _fillMissing(index int, item []int, buildup string) []string {
	if index == len(item) {
		return []string{buildup}
	}
	value := item[index]
	if value != -1 {
		next := buildup + fmt.Sprintf("%d", value)
		return _fillMissing(index+1, item, next)
	} else {
		left := buildup + "0"
		right := buildup + "1"
		leftResult := append(_fillMissing(index+1, item, left), _fillMissing(index+1, item, right)...)
		return leftResult
	}
}

// PathToStates converts a given path into a set of start states and a set of end states
func PathToStates(rbd *robdd.ROBDD, path []int, order map[string]int) ([]string, []string) {
	// dbgInfo := PropValuesFromPath(rbd, path)

	/*
		Skeleton of what we want to do
		Set array of start and end values, init all to -1
		Find fixed propositions by path
			Add non-primed to start set
			Add primed to end set
		Then, iterate through start set, building array of strings
		Do the same for end set
		return those results
	*/

	numProps := len(order) / 2
	start := make([]int, numProps)
	end := make([]int, numProps)
	for i := range start {
		start[i] = -1
		end[i] = -1
	}

	// Find the values that are fixed
	for _, node := range path {
		value := NodeToValue(node)
		prop := NodeToProp(rbd, node)
		if len(prop) == 0 {
			continue
		}
		if IsEnd(prop) {
			end[order[prop]/2] = value
		} else {
			start[order[prop]/2] = value
		}
	}

	// fmt.Printf("Fixed values for path %v are %v -> %v\n", dbgInfo, start, end)

	// Call fill missing to generate arrays of strings of all possible states
	return _fillMissing(0, start, ""), _fillMissing(0, end, "")
}

func CreateStateTransitions(rbd *robdd.ROBDD, paths [][]int, order map[string]int) map[string][]string {
	result := make(map[string][]string)

	num_props := len(order) / 2

	// First generate all possible states
	for i := 0.0; i < math.Pow(2, float64(num_props)); i++ {
		format := fmt.Sprintf("%%0%vb", num_props)
		// fmt.Printf("%s, %d\n", format, int(i))
		state := fmt.Sprintf(format, int(i))
		// fmt.Printf("State %v is %v\n", i, state)
		result[state] = make([]string, 0)
	}

	for _, path := range paths {
		starts, ends := PathToStates(rbd, path, order)
		// For every start state, add all end states
		for _, start := range starts {
			result[start] = append(result[start], ends...)
		}
	}
	return result
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

func GenerateAllValidTransitions(rbd *robdd.ROBDD, root int, order map[string]int) map[string][]string {
	inverseIndex := BFSInverseIndex(rbd, root)
	results := _generateValidTransitions(rbd, 1, inverseIndex)
	for i, _ := range results {
		results[i] = append(results[i], 1)
	}
	// fmt.Printf("Results of all valid transitions: %v\n", results)
	paths := make([]map[string]int, 0)
	for _, result := range results {
		values := PropValuesFromPath(rbd, result)
		// fmt.Printf("Path %v represents %v\n", result, values)
		// fmt.Printf("\tWhich corresponds to state transition %v\n", PropValuesToStateTransition(values, order))
		paths = append(paths, values)
	}
	transitions := CreateStateTransitions(rbd, results, order)
	// fmt.Printf("All transitions: %v\n", transitions)
	return transitions
}
