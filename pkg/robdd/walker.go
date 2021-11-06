package robdd

import (
	"fmt"
	"nickaashoek/tdviz/pkg/parser"
)

type ROBDDTransitionWalker struct {
	*parser.BaseTransitionsListener

	PropOrder map[string]int

	BddManager ROBDD
	Result     ROBDDNode

	OperandQueue []int
}

func (tw *ROBDDTransitionWalker) push(node int) {
	tw.OperandQueue = append(tw.OperandQueue, node)
}

func (tw *ROBDDTransitionWalker) pop() int {
	if len(tw.OperandQueue) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := tw.OperandQueue[len(tw.OperandQueue)-1]

	// Remove the last element from the stack.
	tw.OperandQueue = tw.OperandQueue[:len(tw.OperandQueue)-1]

	return result
}

// EnterStart is called when entering the start production.
func (r *ROBDDTransitionWalker) EnterStart(c *parser.StartContext) {
	revOrder := make(map[int]string)
	for k, v := range r.PropOrder {
		revOrder[v] = k
	}

	r.BddManager.RevOrder = revOrder
	r.BddManager.InitROBDD()
	fmt.Printf("ROBDD at start of walk %v\n", r.BddManager)
}

// ExitStart is called when exiting the start production.
func (r *ROBDDTransitionWalker) ExitStart(c *parser.StartContext) {
	fmt.Println("Exiting Start")
	r.Result = r.BddManager.Nodes[r.pop()]
}

func (r *ROBDDTransitionWalker) ExitNotExpression(c *parser.NotExpressionContext) {
	fmt.Printf("Exiting the not expression %v\n", c.GetText())
	argument := r.pop()
	result := r.BddManager.Not(argument)
	r.push(result)
}

// ExitExpression is called when exiting the expression production.
func (r *ROBDDTransitionWalker) ExitOpExpression(c *parser.OpExpressionContext) {
	fmt.Printf("Hit op expression %v\n", c.GetText())
	operator := c.GetOP().(*parser.OperatorContext)

	right := r.pop()
	left := r.pop()

	if operator.AND() != nil {
		result := r.BddManager.And(left, right)
		r.push(result)
	} else if operator.OR() != nil {
		result := r.BddManager.Or(left, right)
		r.push(result)
	} else if operator.IMPL() != nil {
		result := r.BddManager.Implies(left, right)
		r.push(result)
	} else if operator.DOUBLIMPL() != nil {
		result := r.BddManager.Iff(left, right)
		r.push(result)
	}
}

// ExitIdentifier is called when exiting the identifier production.
func (r *ROBDDTransitionWalker) ExitIdentifier(c *parser.IdentifierContext) {
	propOrder := r.PropOrder[c.GetText()]
	fmt.Printf("Hit identifier %v (%v)\n", c.GetText(), propOrder)
	result := r.BddManager.MakeNode(propOrder, r.BddManager.Lo, r.BddManager.Hi)
	fmt.Printf("Pushed node %v onto stack\n", r.BddManager.Nodes[result])
	r.push(result)
}
