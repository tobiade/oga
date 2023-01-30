// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Oga

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by OgaParser.
type OgaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by OgaParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by OgaParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by OgaParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by OgaParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by OgaParser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by OgaParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by OgaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by OgaParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by OgaParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by OgaParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by OgaParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by OgaParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by OgaParser#MultDivExpr.
	VisitMultDivExpr(ctx *MultDivExprContext) interface{}

	// Visit a parse tree produced by OgaParser#FuncCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by OgaParser#IDExpr.
	VisitIDExpr(ctx *IDExprContext) interface{}

	// Visit a parse tree produced by OgaParser#RelExpr.
	VisitRelExpr(ctx *RelExprContext) interface{}

	// Visit a parse tree produced by OgaParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by OgaParser#NestedExpr.
	VisitNestedExpr(ctx *NestedExprContext) interface{}

	// Visit a parse tree produced by OgaParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by OgaParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by OgaParser#LogicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by OgaParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}
}
