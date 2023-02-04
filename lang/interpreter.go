package lang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
)

// TODO: Implement another mechanism for returning results from function?
const RETURN = "return"

type MemorySpace map[string]any

type Interpreter struct {
	*parser.BaseOgaVisitor
	Global       MemorySpace
	Stack        []MemorySpace
	NodeMetadata map[antlr.ParserRuleContext]*Metadata
}

type MathExprContext interface {
	GetOp() antlr.Token
	Expr(i int) parser.IExprContext
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

func (v *Interpreter) findMemorySpace(name string) MemorySpace {
	for i := len(v.Stack) - 1; i >= 0; i-- {
		mem := v.Stack[i]
		if _, ok := mem[name]; ok {
			return mem
		}
	}
	return nil
}

func (v *Interpreter) load(name string) any {
	mem := v.findMemorySpace(name)
	return mem[name]
}

func (v *Interpreter) set(name string, val any) {
	mem := v.findMemorySpace(name)
	mem[name] = val
}

func (v *Interpreter) detonate(errMsg string, ctx antlr.ParserRuleContext) {
	panic(fmt.Errorf("%s: line %d, col %d", errMsg, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()))
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
	// Push a new instance of MemorySpace unto the stack if we're not invoking this from a function
	_, parentIsFunction := ctx.GetParent().(*parser.FuncDeclContext)
	if !parentIsFunction {
		v.push(MemorySpace{})
	}
	v.VisitChildren(ctx)
	if !parentIsFunction {
		v.pop()
	}
	return nil
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

func (v *Interpreter) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	return v.doMath(ctx)
}

func (v *Interpreter) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
	return v.doMath(ctx)
}

func (v *Interpreter) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	num, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		panic(fmt.Sprintf("unable to convert number: %s", ctx.INT().GetText()))
	}
	return num
}

func (v *Interpreter) VisitIDExpr(ctx *parser.IDExprContext) interface{} {
	return v.load(ctx.IDENTIFIER().GetText())
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

func (v *Interpreter) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := ctx.Expr().Accept(v).(bool)
	if cond {
		ctx.Block(0).Accept(v)
		return nil
	}

	if !cond && ctx.ELSE() != nil {
		if ctx.IfStmt() != nil {
			ctx.IfStmt().Accept(v)
		} else {
			ctx.Block(1).Accept(v)
		}
	}
	return nil
}

func (v *Interpreter) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	left, leftOk := ctx.Expr(0).Accept(v).(bool)
	right, rightOk := ctx.Expr(1).Accept(v).(bool)
	if !leftOk {
		v.detonate("non-boolean operands present: line %d, col %d", ctx.Expr(0))
	}
	if !rightOk {
		v.detonate("non-boolean operands present: line %d, col %d", ctx.Expr(1))
	}

	var res bool
	switch ctx.GetLogical_op().GetTokenType() {
	case parser.OgaParserAND:
		res = left && right
	case parser.OgaParserOR:
		res = left || right
	default:
		panic(fmt.Errorf("unrecognised logical operator: %v", ctx.GetLogical_op().GetText()))
	}
	return res
}

func (v *Interpreter) VisitRelExpr(ctx *parser.RelExprContext) interface{} {
	left := ctx.Expr(0).Accept(v)
	right := ctx.Expr(1).Accept(v)
	leftKind := reflect.TypeOf(left).Kind()
	rightKind := reflect.TypeOf(right).Kind()

	var res bool
	var err error
	if leftKind == reflect.Int && rightKind == reflect.Int {
		res, err = evalRelExpr(left.(int), right.(int), ctx.GetRel_op())
	} else if leftKind == reflect.String && rightKind == reflect.String {
		res, err = evalRelExpr(left.(string), right.(string), ctx.GetRel_op())
	} else {
		err = fmt.Errorf("unexpected operands for relational op: %s and %s", leftKind, rightKind)
	}

	if err != nil {
		panic(err)
	}

	return res

}

func evalRelExpr[T string | int](left, right T, op antlr.Token) (bool, error) {
	var res bool
	switch op.GetTokenType() {
	case parser.OgaParserGREATER:
		res = left > right
	case parser.OgaParserLESS:
		res = left < right
	case parser.OgaParserEQUALS:
		res = left == right
	case parser.OgaParserNOT_EQUAL:
		res = left != right
	default:
		return false, fmt.Errorf("unable to recognise operator: %v", op.GetText())
	}
	return res, nil
}

func (v *Interpreter) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	v.set(ctx.IDENTIFIER().GetText(), ctx.Expr().Accept(v))
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
	v.push(mem)
	// Call function
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

func (v *Interpreter) VisitNestedExpr(ctx *parser.NestedExprContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *Interpreter) doMath(ctx MathExprContext) int {
	left := ctx.Expr(0).Accept(v).(int)
	right := ctx.Expr(1).Accept(v).(int)
	var res int
	switch ctx.GetOp().GetText() {
	case "*":
		res = left * right
	case "/":
		res = left / right
	case "+":
		res = left + right
	case "-":
		res = left - right
	default:
		panic(fmt.Sprintf("unknown operator: %s", ctx.GetOp().GetText()))
	}
	return res
}
