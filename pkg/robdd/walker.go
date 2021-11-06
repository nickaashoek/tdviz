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
	r.BddManager.InitROBDD()
	fmt.Printf("ROBDD at start of walk %v\n", r.BddManager)
}

// EnterExpression is called when entering the expression production.
func (r *ROBDDTransitionWalker) EnterExpression(c *parser.ExpressionContext) {
}

// EnterOperator is called when entering the operator production.
func (r *ROBDDTransitionWalker) EnterOperator(c *parser.OperatorContext) {

}

// EnterIdentifier is called when entering the identifier production.
func (r *ROBDDTransitionWalker) EnterIdentifier(c *parser.IdentifierContext) {
}

// EnterBl is called when entering the bl production.
func (r *ROBDDTransitionWalker) EnterBl(c *parser.BlContext) {
}

// ExitStart is called when exiting the start production.
func (r *ROBDDTransitionWalker) ExitStart(c *parser.StartContext) {
	fmt.Println("Exiting Start")
	r.Result = r.BddManager.Nodes[0]
}

// ExitExpression is called when exiting the expression production.
func (r *ROBDDTransitionWalker) ExitExpression(c *parser.ExpressionContext) {
	fmt.Printf("Hit expression %v\n", c.GetText())
}

// ExitOperator is called when exiting the operator production.
func (r *ROBDDTransitionWalker) ExitOperator(c *parser.OperatorContext) {
}

// ExitIdentifier is called when exiting the identifier production.
func (r *ROBDDTransitionWalker) ExitIdentifier(c *parser.IdentifierContext) {
	propOrder := r.PropOrder[c.GetText()]
	fmt.Printf("Hit identifier %v (%v)\n", c.GetText(), propOrder)
	result := r.BddManager.MakeNode(propOrder, r.BddManager.Lo, r.BddManager.Hi)
	fmt.Printf("Pushed node %v onto stack\n", r.BddManager.Nodes[result])
	r.push(result)
}

// ExitBl is called when exiting the bl production.
func (r *ROBDDTransitionWalker) ExitBl(c *parser.BlContext) {
	fmt.Printf("Hit boolean endpoint %v\n", c.GetText())
}
