CC=go

UTIL = $(shell find pkg/util/*.go)
PARSE = $(shell find pkg/parser/*.go)

lexer: $(UTIL) $(PARSE)
	go build -o build/lexer cmd/formula/main.go
