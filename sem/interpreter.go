package sem

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
)

type MemorySpace map[string]any

type Interpreter struct {
	*parser.BaseOgaVisitor
	Global       MemorySpace
	Stack        []MemorySpace
	NodeMetadata map[antlr.ParserRuleContext]Metadata
}

func (v *Interpreter) top() MemorySpace {
	return v.Stack[len(v.Stack)-1]
}

func (v *Interpreter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Interpreter) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *Interpreter) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *Interpreter) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Interpreter) VisitStmtList(ctx *parser.StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Interpreter) VisitStmt(ctx *parser.StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Interpreter) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Interpreter) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
	left := ctx.Expr(0).Accept(v).(int)
	right := ctx.Expr(1).Accept(v).(int)
	var res int
	switch ctx.GetOp().GetText() {
	case "*":
		res = left * right
	case "/":
		res = left / right
	default:
		panic(fmt.Sprintf("unknown operator: %s", ctx.GetOp().GetText()))
	}
	return res
}

func (v *Interpreter) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	num, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		panic(fmt.Sprintf("unable to convert number: %s", ctx.INT().GetText()))
	}
	return num
}

func (v *Interpreter) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	if ctx.Expr() != nil {
		varName := ctx.IDENTIFIER().GetText()
		res := ctx.Expr().Accept(v)
		v.top()[varName] = res
		return nil
	}
	return nil
}

func (v *Interpreter) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	v.top()[ctx.IDENTIFIER().GetText()] = ctx.Expr().Accept(v)
	return nil
}
