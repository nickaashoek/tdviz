// Code generated from Transitions.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Transitions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTransitionsListener is a complete listener for a parse tree produced by TransitionsParser.
type BaseTransitionsListener struct{}

var _ TransitionsListener = &BaseTransitionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTransitionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTransitionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTransitionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTransitionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTransitionsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTransitionsListener) ExitStart(ctx *StartContext) {}

// EnterBoolExpression is called when production boolExpression is entered.
func (s *BaseTransitionsListener) EnterBoolExpression(ctx *BoolExpressionContext) {}

// ExitBoolExpression is called when production boolExpression is exited.
func (s *BaseTransitionsListener) ExitBoolExpression(ctx *BoolExpressionContext) {}

// EnterNestedExpression is called when production nestedExpression is entered.
func (s *BaseTransitionsListener) EnterNestedExpression(ctx *NestedExpressionContext) {}

// ExitNestedExpression is called when production nestedExpression is exited.
func (s *BaseTransitionsListener) ExitNestedExpression(ctx *NestedExpressionContext) {}

// EnterIdentifierExpression is called when production identifierExpression is entered.
func (s *BaseTransitionsListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production identifierExpression is exited.
func (s *BaseTransitionsListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseTransitionsListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseTransitionsListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterOpExpression is called when production opExpression is entered.
func (s *BaseTransitionsListener) EnterOpExpression(ctx *OpExpressionContext) {}

// ExitOpExpression is called when production opExpression is exited.
func (s *BaseTransitionsListener) ExitOpExpression(ctx *OpExpressionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseTransitionsListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseTransitionsListener) ExitOperator(ctx *OperatorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseTransitionsListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseTransitionsListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterBl is called when production bl is entered.
func (s *BaseTransitionsListener) EnterBl(ctx *BlContext) {}

// ExitBl is called when production bl is exited.
func (s *BaseTransitionsListener) ExitBl(ctx *BlContext) {}
