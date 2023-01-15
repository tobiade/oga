package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
	"github.com/tobiade/oga/sem"
)

func main() {
	code := `funke dostuff(){

		make x = 4

		suppose say x small pass 5 {
			
		} otherwise suppose say z resemble 4 {
	
		}
	
		
	}
	`
	input := antlr.NewInputStream(code)
	lexer := parser.NewOgaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgaParser(stream)
	p.BuildParseTrees = true
	tree := p.SourceFile()

	global := sem.NewDefaultScope("global", nil)
	defV := sem.ScopeDefineVisitor{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		CurrentScope:   &global,
		NodeMetadata:   make(map[antlr.ParserRuleContext]sem.Metadata),
	}
	defV.Visit(tree)

	resV := sem.ScopeResVisitor{
		BaseOgaVisitor: &parser.BaseOgaVisitor{},
		NodeMetadata:   defV.NodeMetadata,
		Errors:         make([]error, 0),
	}
	resV.Visit(tree)
	for _, e := range resV.Errors {
		fmt.Println(e)
	}

}
