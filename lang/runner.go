package lang

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
)

// RunSource code takes a provider which provides the source code, and a forceRun flag which forces oga to run the code.
// By default, oga will flip a coin before running code. If it's heads oga will run it.
func RunSourceCode(provider SourceProvider, forceRun bool) error {
	if !forceRun {
		// 0 is heads, 1 is tails
		if face := flipCoin(); face == 1 {
			return fmt.Errorf("You dey mad")
		}
	}

	source, err := provider.Get()
	if err != nil {
		return err
	}
	input := antlr.NewInputStream(source)
	lexer := parser.NewOgaLexer(input)
	errListener := &OgaErrorListener{Errors: make([]error, 0)}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgaParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	p.BuildParseTrees = true
	tree := p.SourceFile()

	if errListener.HasErrors() {
		return errors.Join(errListener.Errors...)
	}

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
	if len(resV.Errors) > 0 {
		return errors.Join(resV.Errors...)
	}

	globalMem := MemorySpace{}
	stack := []MemorySpace{globalMem}
	interpreter := &Interpreter{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		Stack:          stack,
		NodeMetadata:   resV.NodeMetadata,
	}

	defV.MainFunc.Accept(interpreter)

	return nil
}

func defineNativeFunctions(scope Scope) {
	print := &PrintFunc{
		FuncSymbolBase: FuncSymbolBase{
			FuncName: "print",
		},
	}
	scope.Define(print)
}

func flipCoin() int {
	return rand.Intn(2)
}
