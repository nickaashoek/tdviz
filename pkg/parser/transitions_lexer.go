// Code generated from Transitions.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 67, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 6, 13, 62, 10, 13, 13, 13, 14, 13, 63, 3, 13, 3, 13, 2, 2,
	14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 3, 2, 4, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 67, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2,
	2, 9, 33, 3, 2, 2, 2, 11, 36, 3, 2, 2, 2, 13, 40, 3, 2, 2, 2, 15, 42, 3,
	2, 2, 2, 17, 44, 3, 2, 2, 2, 19, 46, 3, 2, 2, 2, 21, 49, 3, 2, 2, 2, 23,
	54, 3, 2, 2, 2, 25, 61, 3, 2, 2, 2, 27, 28, 7, 40, 2, 2, 28, 4, 3, 2, 2,
	2, 29, 30, 7, 126, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 35, 2, 2, 32, 8,
	3, 2, 2, 2, 33, 34, 7, 47, 2, 2, 34, 35, 7, 64, 2, 2, 35, 10, 3, 2, 2,
	2, 36, 37, 7, 62, 2, 2, 37, 38, 7, 47, 2, 2, 38, 39, 7, 64, 2, 2, 39, 12,
	3, 2, 2, 2, 40, 41, 7, 42, 2, 2, 41, 14, 3, 2, 2, 2, 42, 43, 7, 43, 2,
	2, 43, 16, 3, 2, 2, 2, 44, 45, 9, 2, 2, 2, 45, 18, 3, 2, 2, 2, 46, 47,
	9, 2, 2, 2, 47, 48, 7, 41, 2, 2, 48, 20, 3, 2, 2, 2, 49, 50, 7, 86, 2,
	2, 50, 51, 7, 84, 2, 2, 51, 52, 7, 87, 2, 2, 52, 53, 7, 71, 2, 2, 53, 22,
	3, 2, 2, 2, 54, 55, 7, 72, 2, 2, 55, 56, 7, 67, 2, 2, 56, 57, 7, 78, 2,
	2, 57, 58, 7, 85, 2, 2, 58, 59, 7, 71, 2, 2, 59, 24, 3, 2, 2, 2, 60, 62,
	9, 3, 2, 2, 61, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2,
	63, 64, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 8, 13, 2, 2, 66, 26, 3,
	2, 2, 2, 4, 2, 63, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'&'", "'|'", "'!'", "'->'", "'<->'", "'('", "')'", "", "", "'TRUE'",
	"'FALSE'",
}

var lexerSymbolicNames = []string{
	"", "AND", "OR", "NOT", "IMPL", "DOUBLIMPL", "LPAREN", "RPAREN", "START",
	"END", "TRUE", "FALSE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"AND", "OR", "NOT", "IMPL", "DOUBLIMPL", "LPAREN", "RPAREN", "START", "END",
	"TRUE", "FALSE", "WHITESPACE",
}

type TransitionsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewTransitionsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *TransitionsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTransitionsLexer(input antlr.CharStream) *TransitionsLexer {
	l := new(TransitionsLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Transitions.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TransitionsLexer tokens.
const (
	TransitionsLexerAND        = 1
	TransitionsLexerOR         = 2
	TransitionsLexerNOT        = 3
	TransitionsLexerIMPL       = 4
	TransitionsLexerDOUBLIMPL  = 5
	TransitionsLexerLPAREN     = 6
	TransitionsLexerRPAREN     = 7
	TransitionsLexerSTART      = 8
	TransitionsLexerEND        = 9
	TransitionsLexerTRUE       = 10
	TransitionsLexerFALSE      = 11
	TransitionsLexerWHITESPACE = 12
)
