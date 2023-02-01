package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/lang"
	"github.com/tobiade/oga/parser"
)

func main() {
	code := `funke mehn(){

		make x = doIt(2)*doIt(2)
		suppose say x big pass 100 {
			suppose say x big pass 10 {
				make x = 1
			}
		}
		
	}
	funke doIt(num) {
		make x = num*10
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
	fmt.Println(interpreter.Stack[0])

}
