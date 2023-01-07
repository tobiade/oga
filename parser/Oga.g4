grammar Oga;

sourceFile: ((funcDecl | varDecl) EOS)* EOF;

varDecl: MAKE IDENTIFIER (ASSIGN expr)?;

funcDecl: FUNC IDENTIFIER '(' exprList? ')' block;

stmtList: EOS | (stmt EOS)+;

stmt: varDecl | ifStmt | returnStmt | expressionStmt | forStmt;

block: '{' stmtList? '}';

ifStmt: IF condition block (ELSE (ifStmt | block))?;

returnStmt: RETURN expr;

assignStmt: IDENTIFIER ASSIGN expr;

forStmt:
	FOR (assignStmt? SEMI condition? SEMI assignStmt?)? block;

condition: expr relOp expr;

relOp: GREATER | LESS | NOT_EQUAL | EQUALS;

expressionStmt: expr;

expr:
	IDENTIFIER '(' exprList? ')'	# FuncCall
	| expr ('*' | '/') expr			# MultDivExpr
	| expr ('+' | '-') expr			# AddSubExpr
	| expr relOp expr				# RelExpr
	| INT							# IntExpr
	| IDENTIFIER					# IDExpr
	| '(' expr ')'					# NestedExpr;

exprList: expr (',' expr)*;

// Keywords
MAKE: 'make';
FUNC: 'funke';
IF: 'suppose say';
RETURN: 'dapada';
ELSE: 'otherwise';
FOR: 'dey play';

// Relational operators
GREATER: 'big pass';
EQUALS: 'resemble';
LESS: 'small pass';
NOT_EQUAL: 'no resemble';

// Punctuation
SEMI: ';';
ASSIGN: '=';

// Num literals
INT: '0' | [1-9] [0-9]*;

// Identifier
IDENTIFIER: [a-zA-Z0-9]+;

// Comments
LINE_COMMENT: '//' ~[\r\n]* EOS -> skip;

// Whitespace
EOS: [\r\n]+;
WS: [ \t\n\r]+ -> skip;

