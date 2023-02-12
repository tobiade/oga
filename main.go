package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/lang"
	"github.com/tobiade/oga/parser"
)

func main() {
	code := `funke mehn(){

		make x = doIt(2)
		dey play make y = 0; y small pass 5; y = y + 1 {
			print(y)
		}
		
	}
	funke doIt(num) {
		make x = (num*(5+5)) + 2 + (num*2)
		dapada x
	}
	`
	input := antlr.NewInputStream(code)
	lexer := parser.NewOgaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgaParser(stream)
	p.BuildParseTrees = true
	tree := p.SourceFile()

	global := lang.NewDefaultScope("global", nil)
	defineNativeFunctions(&global)
	defV := lang.ScopeDefineVisitor{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		CurrentScope:   &global,
		NodeMetadata:   make(map[antlr.ParserRuleContext]*lang.Metadata),
	}
	defV.Visit(tree)

	resV := lang.ScopeResVisitor{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		NodeMetadata:   defV.NodeMetadata,
		Errors:         make([]error, 0),
	}
	resV.Visit(tree)
	for _, e := range resV.Errors {
		fmt.Println(e)
	}

	globalMem := lang.MemorySpace{}
	stack := []lang.MemorySpace{globalMem}
	interpreter := &lang.Interpreter{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		Stack:          stack,
		NodeMetadata:   resV.NodeMetadata,
	}

	defV.MainFunc.Accept(interpreter)
	if len(interpreter.Stack) > 0 {
		fmt.Println(interpreter.Stack[0])
	}
}

func defineNativeFunctions(scope lang.Scope) {
	print := &lang.PrintFunc{
		FuncSymbolBase: lang.FuncSymbolBase{
			FuncName: "print",
		},
	}
	scope.Define(print)
}
