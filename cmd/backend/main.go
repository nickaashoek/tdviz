package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nickaashoek/tdviz/pkg/parser"
	"nickaashoek/tdviz/pkg/robdd"
	"nickaashoek/tdviz/pkg/transitions"
	"regexp"
	"sort"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gorilla/mux"
)

const PORT = 8000

const varPattern = `[A-Za-z]`

var (
	variableRegex *regexp.Regexp
)

type ResponseStruct struct {
	ROBDD       robdd.JsonIntermediate `json:"ROBDD"`
	Transitions map[string][]string    `json:"Transitions"`
}

func findPropsInFormula(formula string) map[string]int {
	order := make(map[string]int)

	props := variableRegex.FindAllString(formula, -1)
	sort.Strings(props)

	for _, prop := range props {
		if _, ok := order[prop]; !ok {
			order[prop] = len(order)
			order[prop+"'"] = len(order)
		}
	}
	return order
}

func HandleFormula(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("[BACKEND] Got a request for handling a formula")

	body, _ := ioutil.ReadAll(r.Body)
	formula := string(body)
	order := findPropsInFormula(formula)

	parserInput := antlr.NewInputStream(formula)
	lexer := parser.NewTransitionsLexer(parserInput)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTransitionsParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	fmt.Printf("[BACKEND] Doing parse for formula %s\n", formula)

	walker := &robdd.ROBDDTransitionWalker{PropOrder: order}
	antlr.ParseTreeWalkerDefault.Walk(walker, p.Start())

	robddJSON := robdd.BddToIntermediate(&walker.BddManager, walker.Result, order)

	fmt.Printf("[BACKEND] Generating transitions for formula %s\n", formula)

	allTransitions := transitions.GenerateAllValidTransitions(&walker.BddManager, walker.Result, order)

	fmt.Printf("[BACKEND] Formatting result as json\n")

	resp := ResponseStruct{ROBDD: robddJSON, Transitions: allTransitions}

	jsonified, _ := json.Marshal(resp)

	w.Write(jsonified)
}

func main() {
	variableRegex = regexp.MustCompile(varPattern)

	r := mux.NewRouter()
	r.HandleFunc("/formula", HandleFormula).
		Methods("POST")

	go http.ListenAndServe(fmt.Sprintf(":%v", PORT), r)

	fmt.Printf("[BACKEND] Running server on localhost:%d\n", PORT)

	select {}
}
