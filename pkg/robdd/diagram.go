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
	return fmt.Sprintf("Node w/ Prop = %v, Lo = %v, Hi = %v", n.Prop, n.Lo, n.Hi)
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
	Nodes []ROBDDNode
	Cache map[string]int
	Lo    int
	Hi    int
}

func (r *ROBDD) InitROBDD() {
	r.Cache = make(map[string]int)
	r.MakeNode(-2, -2, 0)
	r.MakeNode(-1, -1, 0)
	r.Lo = 0
	r.Hi = 1
}

func (r *ROBDD) NewProp(p int) int {
	return r.MakeNode(p, r.Lo, r.Hi)
}

func (r *ROBDD) MakeNode(p, l, h int) int {
	// fmt.Printf("Making node %v %v %v\n", p, l, h)
	if l == h {
		return l
	}
	// fmt.Printf("PLHHash %v\n", PLHHash(p, l, h))
	if val, ok := r.Cache[PLHHash(p, l, h)]; ok {
		return val
	} else {
		res := ROBDDNode{
			Prop: p,
			Lo:   l,
			Hi:   h,
		}
		// fmt.Printf("Appending %v\n", res)
		r.Nodes = append(r.Nodes, res)
		r.Cache[res.Hash()] = len(r.Nodes) - 1
		return len(r.Nodes) - 1
	}
}

func (r *ROBDD) Not() {

}

func (r *ROBDD) And(left ROBDDNode, right ROBDDNode) int {
	if left.IsLo() || right.IsLo() {
		return r.Lo
	} else if left.IsHi() {
		return r.Cache[right.Hash()]
	} else if right.IsHi() {
		return r.Cache[left.Hash()]
	} else if left.Prop == right.Prop {
		return r.MakeNode(
			left.Prop,
			r.And(r.Nodes[left.Lo], r.Nodes[right.Lo]),
			r.And(r.Nodes[left.Hi], r.Nodes[right.Hi]),
		)
	} else if left.Prop < right.Prop {
		return r.MakeNode(
			left.Prop,
			r.And(r.Nodes[left.Lo], right),
			r.And(r.Nodes[left.Hi], right),
		)
	} else {
		return r.MakeNode(
			right.Prop,
			r.And(left, r.Nodes[right.Lo]),
			r.And(left, r.Nodes[right.Hi]),
		)
	}

}

// func (r *ROBDD) Or(root *ROBDDNode) {

// }

// func (r *ROBDD) Implies(r2 *ROBDD) {

// }

// func (r *ROBDD) Iff(r2 *ROBDD) {

// }

func (r ROBDD) String() string {
	return fmt.Sprintf("ROBDD Structure\n\tNodes: %v\n\tCache: %v", r.Nodes, r.Cache)
}
