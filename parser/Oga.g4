grammar Oga;

sourceFile: ((funcDecl | varDecl) EOS)* EOF;

varDecl: MAKE IDENTIFIER (ASSIGN expr)?;

funcDecl: FUNC IDENTIFIER '(' identifierList? ')' block;

identifierList: IDENTIFIER (COMMA IDENTIFIER)*;

stmtList: (stmt EOS+)+;

stmt:
	varDecl
	| ifStmt
	| returnStmt
	| expressionStmt
	| forStmt
	| assignStmt;

block: '{' EOS* stmtList? '}';

ifStmt: IF expr block (ELSE (ifStmt | block))?;

returnStmt: RETURN expr;

assignStmt: IDENTIFIER ASSIGN expr;

simpleStmt: varDecl | assignStmt;

forStmt:
	FOR (
		initStmt = simpleStmt? SEMI expr? SEMI postStmt = assignStmt?
	)? block;

expressionStmt: expr;

expr:
	IDENTIFIER '(' exprList? ')'								# FuncCall
	| expr op = (MUL | DIV) expr								# MultDivExpr
	| expr op = (PLUS | MINUS) expr								# AddSubExpr
	| expr rel_op = (GREATER | LESS | NOT_EQUAL | EQUALS) expr	# RelExpr
	| expr logical_op = (AND | OR) expr							# LogicalExpr
	| INT														# IntExpr
	| STR														# StrExpr
	| IDENTIFIER												# IDExpr
	| '(' expr ')'												# NestedExpr;

exprList: expr (COMMA expr)*;

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

// Logical operators
AND: 'and';
OR: 'or';

// Punctuation
SEMI: ';';
ASSIGN: '=';
COMMA: ',';

// Arithmetic
MUL: '*';
DIV: '/';
PLUS: '+';
MINUS: '-';

// Literals
INT: '0' | [1-9] [0-9]*;
STR: '"' (ESC | .)*? '"';
fragment ESC: '\\"' | '\\\\';

// Identifier
IDENTIFIER: [a-zA-Z0-9]+;

// Comments
LINE_COMMENT: '//' ~[\r\n]* EOS -> skip;

// Whitespace
EOS: [\r\n]+;
WS: [ \t]+ -> skip;

