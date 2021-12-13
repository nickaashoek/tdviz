CC=go

UTIL = $(shell find pkg/util/*.go)
PARSE = $(shell find pkg/parser/*.go)
FETEST = $(shell find cmd/fetest/*.go)
BACKEND = $(shell find cmd/backend/*.go)

lexer: $(UTIL) $(PARSE)
	go build -o build/lexer cmd/formula/main.go

fetest: $(UTIL) $(PARSE) $(FETEST)
	go build -o build/fetest cmd/fetest/main.go

backend: $(UTIL) $(PARSE) $(BACKEND)
	go build -o build/backend cmd/backend/main.go