package lang

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/tobiade/oga/parser"
)

// ScopeDefineVistor defines symbols in relevant scopes
type ScopeDefineVisitor struct {
	*parser.BaseOgaVisitor
	CurrentScope Scope
	MainFunc     antlr.ParserRuleContext
	NodeMetadata map[antlr.ParserRuleContext]*Metadata
}

func (v *ScopeDefineVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ScopeDefineVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *ScopeDefineVisitor) VisitSourceFile(ctx *parser.SourceFileContext) interface{} {
	for _, varDecl := range ctx.AllVarDecl() {
		varDecl.Accept(v)
	}
	for _, funcDecl := range ctx.AllFuncDecl() {
		funcDecl.Accept(v)
	}
	return nil
}

func (v *ScopeDefineVisitor) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	varSym := &VarSymbol{
		VarName: ctx.IDENTIFIER().GetText(),
		Node:    ctx,
	}
	v.CurrentScope.Define(varSym)
	ctx.Expr().Accept(v)
	return nil
}

func (v *ScopeDefineVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	funSym := &UserFuncSymbol{
		FuncSymbolBase: FuncSymbolBase{
			FuncName: ctx.IDENTIFIER().GetText(),
		},
		Node: ctx,
	}
	v.CurrentScope.Define(funSym)
	funcScope := NewDefaultScope(funSym.FuncName, v.CurrentScope)
	v.CurrentScope = &funcScope

	if ctx.IdentifierList() != nil {
		symbols := ctx.IdentifierList().Accept(v).([]*VarSymbol)
		funSym.FuncParams = symbols
	}

	ctx.Block().Accept(v)

	// Record node for main function (or should I say 'mehn' function? ha ha)
	if funSym.FuncName == MAIN {
		v.MainFunc = ctx
	}

	v.CurrentScope = funcScope.EnclosingScope()
	return nil
}

func (v *ScopeDefineVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitStmtList(ctx *parser.StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	ctx.Expr().Accept(v)
	for _, b := range ctx.AllBlock() {
		ifScope := NewDefaultScope("if-else-scope", v.CurrentScope)
		v.CurrentScope = &ifScope
		b.Accept(v)
		v.CurrentScope = ifScope.EnclosingScope()
	}
	if ctx.IfStmt() != nil {
		ctx.IfStmt().Accept(v)
	}
	return nil
}

func (v *ScopeDefineVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	forScope := NewDefaultScope("for-scope", v.CurrentScope)
	v.CurrentScope = &forScope
	if ctx.VarDecl() != nil {
		ctx.VarDecl().Accept(v)
	}
	ctx.Block().Accept(v)
	v.CurrentScope = forScope.EnclosingScope()
	return nil
}

func (v *ScopeDefineVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	symbols := make([]*VarSymbol, 0)
	for _, id := range ctx.AllIDENTIFIER() {
		varSym := &VarSymbol{
			VarName: id.GetText(),
			Node:    id,
		}
		v.CurrentScope.Define(varSym)
		symbols = append(symbols, varSym)
	}
	return symbols
}

func (v *ScopeDefineVisitor) VisitIDExpr(ctx *parser.IDExprContext) interface{} {
	v.NodeMetadata[ctx] = &Metadata{scope: v.CurrentScope}
	return nil
}

func (v *ScopeDefineVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *ScopeDefineVisitor) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	v.NodeMetadata[ctx] = &Metadata{scope: v.CurrentScope}
	if ctx.ExprList() != nil {
		ctx.ExprList().Accept(v)
	}
	return nil
}

func (v *ScopeDefineVisitor) VisitExprList(ctx *parser.ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitRelExpr(ctx *parser.RelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitNestedExpr(ctx *parser.NestedExprContext) interface{} {
	return ctx.Expr().Accept(v)
}

func (v *ScopeDefineVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeDefineVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	v.NodeMetadata[ctx] = &Metadata{scope: v.CurrentScope}
	return ctx.Expr().Accept(v)
}

// ScopeResVisitor resolves symbols that might have beeen defined by ScopeDefineVistor
type ScopeResVisitor struct {
	*parser.BaseOgaVisitor
	NodeMetadata map[antlr.ParserRuleContext]*Metadata
	Errors       []error
}

func (v *ScopeResVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ScopeResVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *ScopeResVisitor) VisitIDExpr(ctx *parser.IDExprContext) interface{} {
	m := v.NodeMetadata[ctx]
	sym, err := m.scope.Resolve(ctx.IDENTIFIER().GetText())
	if err != nil {
		v.Errors = append(v.Errors, err)
		return nil
	}
	m.symbol = sym
	return nil
}

func (v *ScopeResVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	m := v.NodeMetadata[ctx]
	sym, err := m.scope.Resolve(ctx.IDENTIFIER().GetText())
	if err != nil {
		v.Errors = append(v.Errors, err)
		return nil
	}
	m.symbol = sym
	return ctx.Expr().Accept(v)
}

func (v *ScopeResVisitor) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	m := v.NodeMetadata[ctx]
	sym, err := m.scope.Resolve(ctx.IDENTIFIER().GetText())
	if err != nil {
		v.Errors = append(v.Errors, err)
		return nil
	}
	m.symbol = sym
	if ctx.ExprList() != nil {
		ctx.ExprList().Accept(v)
	}
	return nil
}

func (v *ScopeResVisitor) VisitSourceFile(ctx *parser.SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitStmtList(ctx *parser.StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitRelExpr(ctx *parser.RelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitNestedExpr(ctx *parser.NestedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeResVisitor) VisitExprList(ctx *parser.ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}
