// Code generated from Transitions.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Transitions

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 42, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 25, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14, 3, 34, 11, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 3, 4, 7, 2, 4, 6, 8, 10, 2, 5,
	4, 2, 3, 4, 6, 7, 3, 2, 10, 11, 3, 2, 12, 13, 2, 40, 2, 12, 3, 2, 2, 2,
	4, 24, 3, 2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 37, 3, 2, 2, 2, 10, 39, 3, 2,
	2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7, 2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 16,
	8, 3, 1, 2, 16, 17, 7, 8, 2, 2, 17, 18, 5, 4, 3, 2, 18, 19, 7, 9, 2, 2,
	19, 25, 3, 2, 2, 2, 20, 21, 7, 5, 2, 2, 21, 25, 5, 4, 3, 6, 22, 25, 5,
	8, 5, 2, 23, 25, 5, 10, 6, 2, 24, 15, 3, 2, 2, 2, 24, 20, 3, 2, 2, 2, 24,
	22, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 32, 3, 2, 2, 2, 26, 27, 12, 5,
	2, 2, 27, 28, 5, 6, 4, 2, 28, 29, 5, 4, 3, 6, 29, 31, 3, 2, 2, 2, 30, 26,
	3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2,
	33, 5, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 36, 9, 2, 2, 2, 36, 7, 3, 2,
	2, 2, 37, 38, 9, 3, 2, 2, 38, 9, 3, 2, 2, 2, 39, 40, 9, 4, 2, 2, 40, 11,
	3, 2, 2, 2, 4, 24, 32,
}
var literalNames = []string{
	"", "'&'", "'|'", "'!'", "'->'", "'<->'", "'('", "')'", "", "", "'TRUE'",
	"'FALSE'",
}
var symbolicNames = []string{
	"", "AND", "OR", "NOT", "IMPL", "DOUBLIMPL", "LPAREN", "RPAREN", "START",
	"END", "TRUE", "FALSE", "WHITESPACE",
}

var ruleNames = []string{
	"start", "expression", "operator", "identifier", "bl",
}

type TransitionsParser struct {
	*antlr.BaseParser
}

// NewTransitionsParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *TransitionsParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTransitionsParser(input antlr.TokenStream) *TransitionsParser {
	this := new(TransitionsParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Transitions.g4"

	return this
}

// TransitionsParser tokens.
const (
	TransitionsParserEOF        = antlr.TokenEOF
	TransitionsParserAND        = 1
	TransitionsParserOR         = 2
	TransitionsParserNOT        = 3
	TransitionsParserIMPL       = 4
	TransitionsParserDOUBLIMPL  = 5
	TransitionsParserLPAREN     = 6
	TransitionsParserRPAREN     = 7
	TransitionsParserSTART      = 8
	TransitionsParserEND        = 9
	TransitionsParserTRUE       = 10
	TransitionsParserFALSE      = 11
	TransitionsParserWHITESPACE = 12
)

// TransitionsParser rules.
const (
	TransitionsParserRULE_start      = 0
	TransitionsParserRULE_expression = 1
	TransitionsParserRULE_operator   = 2
	TransitionsParserRULE_identifier = 3
	TransitionsParserRULE_bl         = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransitionsParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransitionsParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TransitionsParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TransitionsParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TransitionsParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.expression(0)
	}
	{
		p.SetState(11)
		p.Match(TransitionsParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransitionsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransitionsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExpressionContext struct {
	*ExpressionContext
}

func NewBoolExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) Bl() IBlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlContext)
}

func (s *BoolExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterBoolExpression(s)
	}
}

func (s *BoolExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitBoolExpression(s)
	}
}

type NestedExpressionContext struct {
	*ExpressionContext
}

func NewNestedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NestedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TransitionsParserLPAREN, 0)
}

func (s *NestedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TransitionsParserRPAREN, 0)
}

func (s *NestedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterNestedExpression(s)
	}
}

func (s *NestedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitNestedExpression(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

type NotExpressionContext struct {
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TransitionsParserNOT, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type OpExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	OP    IOperatorContext
	right IExpressionContext
}

func NewOpExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExpressionContext {
	var p = new(OpExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OpExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *OpExpressionContext) GetOP() IOperatorContext { return s.OP }

func (s *OpExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *OpExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *OpExpressionContext) SetOP(v IOperatorContext) { s.OP = v }

func (s *OpExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OpExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OpExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OpExpressionContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *OpExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterOpExpression(s)
	}
}

func (s *OpExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitOpExpression(s)
	}
}

func (p *TransitionsParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TransitionsParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TransitionsParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransitionsParserLPAREN:
		localctx = NewNestedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(14)
			p.Match(TransitionsParserLPAREN)
		}
		{
			p.SetState(15)
			p.expression(0)
		}
		{
			p.SetState(16)
			p.Match(TransitionsParserRPAREN)
		}

	case TransitionsParserNOT:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(TransitionsParserNOT)
		}
		{
			p.SetState(19)
			p.expression(4)
		}

	case TransitionsParserSTART, TransitionsParserEND:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Identifier()
		}

	case TransitionsParserTRUE, TransitionsParserFALSE:
		localctx = NewBoolExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Bl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOpExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*OpExpressionContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, TransitionsParserRULE_expression)
			p.SetState(24)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(25)

				var _x = p.Operator()

				localctx.(*OpExpressionContext).OP = _x
			}
			{
				p.SetState(26)

				var _x = p.expression(4)

				localctx.(*OpExpressionContext).right = _x
			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransitionsParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransitionsParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(TransitionsParserAND, 0)
}

func (s *OperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(TransitionsParserOR, 0)
}

func (s *OperatorContext) IMPL() antlr.TerminalNode {
	return s.GetToken(TransitionsParserIMPL, 0)
}

func (s *OperatorContext) DOUBLIMPL() antlr.TerminalNode {
	return s.GetToken(TransitionsParserDOUBLIMPL, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *TransitionsParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TransitionsParserRULE_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TransitionsParserAND)|(1<<TransitionsParserOR)|(1<<TransitionsParserIMPL)|(1<<TransitionsParserDOUBLIMPL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransitionsParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransitionsParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) START() antlr.TerminalNode {
	return s.GetToken(TransitionsParserSTART, 0)
}

func (s *IdentifierContext) END() antlr.TerminalNode {
	return s.GetToken(TransitionsParserEND, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *TransitionsParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TransitionsParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TransitionsParserSTART || _la == TransitionsParserEND) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlContext is an interface to support dynamic dispatch.
type IBlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlContext differentiates from other interfaces.
	IsBlContext()
}

type BlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlContext() *BlContext {
	var p = new(BlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransitionsParserRULE_bl
	return p
}

func (*BlContext) IsBlContext() {}

func NewBlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlContext {
	var p = new(BlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransitionsParserRULE_bl

	return p
}

func (s *BlContext) GetParser() antlr.Parser { return s.parser }

func (s *BlContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TransitionsParserTRUE, 0)
}

func (s *BlContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TransitionsParserFALSE, 0)
}

func (s *BlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.EnterBl(s)
	}
}

func (s *BlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransitionsListener); ok {
		listenerT.ExitBl(s)
	}
}

func (p *TransitionsParser) Bl() (localctx IBlContext) {
	localctx = NewBlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TransitionsParserRULE_bl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TransitionsParserTRUE || _la == TransitionsParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *TransitionsParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TransitionsParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
