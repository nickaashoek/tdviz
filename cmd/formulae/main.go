package main

import (
	"fmt"
	"log"
	"nickaashoek/tdviz/pkg/formula-parser"

	"github.com/timtadh/lexmachine"
)

func main() {
	testFormula := `(p' <-> (p | q)) & (q | q') & !(q' & q) & !(q & r & r') & (r' -> (p | q | r))`
	fmt.Println("Parsing formula", testFormula)

	// insert formula here
	s, err := formula.Lexer.Scanner([]byte(testFormula))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Type    | Lexeme     | Position")
	fmt.Println("--------+------------+------------")
	for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
		if err != nil {
			log.Fatal(err)
		}
		token := tok.(*lexmachine.Token)
		fmt.Printf("%-7v | %-10v | %v:%v-%v:%v\n",
			formula.Tokens[token.Type],
			string(token.Lexeme),
			token.StartLine,
			token.StartColumn,
			token.EndLine,
			token.EndColumn)
	}
}
