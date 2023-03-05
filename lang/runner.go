package lang

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
)

func RunSourceCode(provider SourceProvider) {
	source, err := provider.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	input := antlr.NewInputStream(source)
	lexer := parser.NewOgaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgaParser(stream)
	p.BuildParseTrees = true
	tree := p.SourceFile()

	global := NewDefaultScope("global", nil)
	defineNativeFunctions(&global)
	defV := ScopeDefineVisitor{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		CurrentScope:   &global,
		NodeMetadata:   make(map[antlr.ParserRuleContext]*Metadata),
	}
	defV.Visit(tree)

	resV := ScopeResVisitor{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		NodeMetadata:   defV.NodeMetadata,
		Errors:         make([]error, 0),
	}
	resV.Visit(tree)
	for _, e := range resV.Errors {
		fmt.Println(e)
	}

	globalMem := MemorySpace{}
	stack := []MemorySpace{globalMem}
	interpreter := &Interpreter{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		Stack:          stack,
		NodeMetadata:   resV.NodeMetadata,
	}

	defV.MainFunc.Accept(interpreter)
}

func defineNativeFunctions(scope Scope) {
	print := &PrintFunc{
		FuncSymbolBase: FuncSymbolBase{
			FuncName: "print",
		},
	}
	scope.Define(print)
}
