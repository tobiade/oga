package lang

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
)

const RETURN = "return"

type MemorySpace map[string]any

type Interpreter struct {
	*parser.BaseOgaVisitor
	Global       MemorySpace
	Stack        []MemorySpace
	NodeMetadata map[antlr.ParserRuleContext]*Metadata
}

func (v *Interpreter) top() MemorySpace {
	return v.Stack[len(v.Stack)-1]
}

func (v *Interpreter) push(mem MemorySpace) {
	v.Stack = append(v.Stack, mem)
}

func (v *Interpreter) pop() MemorySpace {
	top := v.top()
	v.Stack = v.Stack[:len(v.Stack)-1]
	return top
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
	return v.VisitChildren(ctx)
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

func (v *Interpreter) VisitIDExpr(ctx *parser.IDExprContext) interface{} {
	return v.top()[ctx.IDENTIFIER().GetText()]
}

func (v *Interpreter) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	return ctx.STR().GetText()
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

func (v *Interpreter) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	ret := ctx.Expr().Accept(v)
	mem := MemorySpace{}
	mem[RETURN] = ret
	v.push(mem)
	return ret
}

func (v *Interpreter) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	mem := MemorySpace{}
	sym := v.NodeMetadata[ctx].symbol.(*FuncSymbol)
	if ctx.ExprList() != nil {
		exprResults := ctx.ExprList().Accept(v).([]any)
		for idx, p := range sym.Params {
			mem[p.Name()] = exprResults[idx]
		}
	}
	v.Stack = append(v.Stack, mem)
	sym.Node.Accept(v)
	// Pop return value
	r := v.pop()
	// Pop stack frame for function
	v.pop()
	return r[RETURN]
}

func (v *Interpreter) VisitExprList(ctx *parser.ExprListContext) interface{} {
	res := make([]any, 0)
	for _, e := range ctx.AllExpr() {
		res = append(res, e.Accept(v))
	}
	return res
}
