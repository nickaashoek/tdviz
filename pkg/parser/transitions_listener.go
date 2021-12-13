// Code generated from Transitions.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Transitions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TransitionsListener is a complete listener for a parse tree produced by TransitionsParser.
type TransitionsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBoolExpression is called when entering the boolExpression production.
	EnterBoolExpression(c *BoolExpressionContext)

	// EnterNestedExpression is called when entering the nestedExpression production.
	EnterNestedExpression(c *NestedExpressionContext)

	// EnterIdentifierExpression is called when entering the identifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterOpExpression is called when entering the opExpression production.
	EnterOpExpression(c *OpExpressionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterBl is called when entering the bl production.
	EnterBl(c *BlContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBoolExpression is called when exiting the boolExpression production.
	ExitBoolExpression(c *BoolExpressionContext)

	// ExitNestedExpression is called when exiting the nestedExpression production.
	ExitNestedExpression(c *NestedExpressionContext)

	// ExitIdentifierExpression is called when exiting the identifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitOpExpression is called when exiting the opExpression production.
	ExitOpExpression(c *OpExpressionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitBl is called when exiting the bl production.
	ExitBl(c *BlContext)
}
