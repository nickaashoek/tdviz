package main

import (
	"fmt"
	"nickaashoek/tdviz/pkg/parser"
	"nickaashoek/tdviz/pkg/robdd"
	"regexp"
	"sort"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	testFormula := `(p' <-> (p | q)) & (q | q') & !(q' & q) & !(q & r & r') & (r' -> (p | q | r))`

	testFormula = `(p & q)`

	varPattern := `[A-Za-z]`

	varRE := regexp.MustCompile(varPattern)

	props := varRE.FindAllString(testFormula, -1)

	sort.Strings(props)
	fmt.Printf("Starts: %v\n", props)

	order := make(map[string]int)

	for _, prop := range props {
		if _, ok := order[prop]; !ok {
			order[prop] = len(order)
			order[prop+"'"] = len(order)
		}
	}

	fmt.Printf("Parsing formula %v with order %v\n", testFormula, order)

	input := antlr.NewInputStream(testFormula)
	lexer := parser.NewTransitionsLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTransitionsParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	// antlr.ParseTreeWalkerDefault.Walk(&robdd.ROBDDTransitionWalker{}, p.Start())
	walker := &robdd.ROBDDTransitionWalker{PropOrder: order}
	antlr.ParseTreeWalkerDefault.Walk(walker, p.Start())
	fmt.Printf("After parsing, result is %v\n", walker.Result)
	fmt.Printf("Full ROBDD Structure is %v\n", walker.BddManager)
}
