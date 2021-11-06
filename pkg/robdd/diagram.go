package robdd

import "fmt"

type ROBDDNode struct {
	// Lo is -2
	// Hi is -1
	Prop int
	Lo   int
	Hi   int
}

func (n *ROBDDNode) Hash() string {
	return PLHHash(n.Prop, n.Lo, n.Hi)
}

func (n *ROBDDNode) IsHi() bool {
	return n.Prop == -1
}

func (n *ROBDDNode) IsLo() bool {
	return n.Prop == -2
}

func (n ROBDDNode) String() string {
	return fmt.Sprintf("| Node %v, %v,%v ", n.Prop, n.Lo, n.Hi)
}

func HiTerminal() ROBDDNode {
	return ROBDDNode{
		Prop: -1,
		Lo:   0,
		Hi:   0,
	}
}

func LoTerminal() ROBDDNode {
	return ROBDDNode{
		Prop: -2,
		Lo:   0,
		Hi:   0,
	}
}

func PLHHash(p, l, h int) string {
	return fmt.Sprintf("%v%v%v", p, l, h)
}

type ROBDD struct {
	Nodes    []ROBDDNode
	Cache    map[string]int
	RevOrder map[int]string
	Lo       int
	Hi       int
}

func (rbd *ROBDD) InitROBDD() {
	rbd.Cache = make(map[string]int)
	rbd.MakeNode(-2, -2, 0)
	rbd.MakeNode(-1, -1, 0)
	rbd.Lo = 0
	rbd.Hi = 1
}

func (rbd *ROBDD) NewProp(p int) int {
	return rbd.MakeNode(p, rbd.Lo, rbd.Hi)
}

func (rbd *ROBDD) MakeNode(p, l, h int) int {
	// fmt.Printf("Making node %v %v %v\n", p, l, h)
	if l == h {
		return l
	}
	// fmt.Printf("PLHHash %v\n", PLHHash(p, l, h))
	if val, ok := rbd.Cache[PLHHash(p, l, h)]; ok {
		return val
	} else {
		res := ROBDDNode{
			Prop: p,
			Lo:   l,
			Hi:   h,
		}
		// fmt.Printf("Appending %v\n", res)
		rbd.Nodes = append(rbd.Nodes, res)
		rbd.Cache[res.Hash()] = len(rbd.Nodes) - 1
		return len(rbd.Nodes) - 1
	}
}

func (rbd *ROBDD) Not(n int) int {
	node := rbd.Nodes[n]

	if node.IsHi() {
		return 0
	} else if node.IsLo() {
		return 1
	} else {
		return rbd.MakeNode(
			node.Prop,
			rbd.Not(node.Lo),
			rbd.Not(node.Hi),
		)
	}
}

func (rbd *ROBDD) And(l, r int) int {
	left := rbd.Nodes[l]
	right := rbd.Nodes[r]

	if left.IsLo() || right.IsLo() {
		return rbd.Lo
	} else if left.IsHi() {
		return rbd.Cache[right.Hash()]
	} else if right.IsHi() {
		return rbd.Cache[left.Hash()]
	} else if left.Prop == right.Prop {
		return rbd.MakeNode(
			left.Prop,
			rbd.And(left.Lo, right.Lo),
			rbd.And(left.Hi, right.Hi),
		)
	} else if left.Prop < right.Prop {
		return rbd.MakeNode(
			left.Prop,
			rbd.And(left.Lo, r),
			rbd.And(left.Hi, r),
		)
	} else {
		return rbd.MakeNode(
			right.Prop,
			rbd.And(l, right.Lo),
			rbd.And(l, right.Hi),
		)
	}
}

func (rbd *ROBDD) Or(l, r int) int {
	left := rbd.Nodes[l]
	right := rbd.Nodes[r]

	if left.IsHi() || right.IsHi() {
		return rbd.Hi
	} else if left.IsLo() {
		return rbd.Cache[right.Hash()]
	} else if right.IsLo() {
		return rbd.Cache[left.Hash()]
	} else if left.Prop == right.Prop {
		return rbd.MakeNode(
			left.Prop,
			rbd.Or(left.Lo, right.Lo),
			rbd.Or(left.Hi, right.Hi),
		)
	} else if left.Prop < right.Prop {
		return rbd.MakeNode(
			left.Prop,
			rbd.Or(left.Lo, r),
			rbd.Or(left.Hi, r),
		)
	} else {
		return rbd.MakeNode(
			right.Prop,
			rbd.Or(l, right.Lo),
			rbd.Or(l, right.Hi),
		)
	}
}

func (rbd *ROBDD) Implies(l, r int) int {
	not_l := rbd.Not(l)
	return rbd.Or(not_l, r)
}

func (rbd *ROBDD) Iff(l, r int) int {
	left_direction := rbd.Implies(l, r)
	right_direction := rbd.Implies(r, l)
	return rbd.And(left_direction, right_direction)
}

func (rbd ROBDD) String() string {
	nodeString := "Nodes: "
	for i, node := range rbd.Nodes {
		nodeString += fmt.Sprintf("\n\t\t%v: | %v, %v, %v", i, rbd.RevOrder[node.Prop], node.Lo, node.Hi)
	}

	return fmt.Sprintf("ROBDD Structure\n\tNodes: %v\n\tCache: %v", nodeString, rbd.Cache)
}
