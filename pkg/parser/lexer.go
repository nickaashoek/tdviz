package formulaparser

import (
	"strings"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var (
	Literals []string
	Keywords []string
	Tokens   []string
	TokenIds map[string]int
	Lexer    *lexmachine.Lexer
)

func init() {
	var err error
	Lexer, err = setupLexer()
	if err != nil {
		panic(err)
	}
}

func initTokens() {
	Literals = []string{
		"!",
		"|",
		"&",
		"->",
		"(",
		")",
		"[",
		"]",
		"=",
		",",
		"<->",
	}

	Keywords = []string{}

	Tokens = []string{
		"ID",
	}

	Tokens = append(Tokens, Literals...)
	Tokens = append(Tokens, Keywords...)
	TokenIds = map[string]int{}
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

func token(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}

// No-op for ignoring whitespace
func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// SetupLexer should setup the lexer for handling transition formulas
func setupLexer() (*lexmachine.Lexer, error) {
	// Init stuff
	lexer := lexmachine.NewLexer()
	initTokens()
	// Lexmachine breaks ties by first added, so literals must come first
	for _, lit := range Literals {
		r := "\\" + strings.Join(strings.Split(lit, ""), "\\")
		lexer.Add([]byte(r), token(lit))
	}

	// Add Keywords. For our transition parser, there are no keywords
	for _, name := range Keywords {
		lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	// Add rule for variables. Single character, including subscript, and/or prime
	// Regular expression used: [A-Za-z](_[0-9]*)?'?
	// Can use any alpha char for variable, optional subscript, optional prime
	lexer.Add([]byte(`[A-Za-z](_[0-9]*)?'?`), token("ID"))
	lexer.Add([]byte("( |\t|\n|\r)+"), skip)

	err := lexer.Compile()
	if err != nil {
		return nil, err
	}
	return lexer, nil
}
