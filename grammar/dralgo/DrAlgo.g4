grammar DrAlgo;

// ============================================================
// Parser rules
// ============================================================

program
  : sep* stmtList? EOF
  ;

stmtList
  : stmt (sep+ stmt)* sep*
  ;

stmt
  : assignStmt
  | incDecStmt
  | printStmt
  | readStmt
  | readIntStmt
  | ifStmt
  | whileStmt
  ;

// --- Statements ---

assignStmt
  : IDENT ASSIGN expr
  ;

incDecStmt
  : IDENT (INCR | DECR)
  ;

printStmt
  : AFFICHER (printArgs)?
  ;

printArgs
  : expr (COMMA expr)*
  ;

readStmt
  : LIRE IDENT
  ;

readIntStmt
  : LIRE_ENTIER IDENT
  ;

ifStmt
  : SI expr (ALORS)? sep+ block (elseIfPart)* (elsePart)? FIN_SI
  ;

elseIfPart
  : SINON_SI expr (ALORS)? sep+ block
  ;

elsePart
  : SINON sep+ block
  ;

whileStmt
  : TANT_QUE expr sep+ block FIN_TANT_QUE
  ;

block
  : stmtList
  ;

sep
  : NEWLINE
  ;

// --- Expressions (precedence) ---
//
// Goal:
// - unary +/-
// - ** right-associative (Python-like)
// - *, /, %
// - +, -
// - comparisons
// - ==, !=
// - list literals [a, b, c]

expr
  : equality
  ;

equality
  : comparison ((EQ | NEQ) comparison)*
  ;

comparison
  : addition ((LT | LTE | GT | GTE) addition)*
  ;

addition
  : multiplication ((PLUS | MINUS) multiplication)*
  ;

multiplication
  : power ((MUL | DIV | MOD) power)*
  ;

// ** is right-associative: 2**3**2 = 2**(3**2)
power
  : unary (POW power)?
  ;

unary
  : (PLUS | MINUS) unary
  | primary
  ;

primary
  : listLiteral
  | literal
  | IDENT
  | LPAREN expr RPAREN
  ;

listLiteral
  : LBRACK (expr (COMMA expr)*)? RBRACK
  ;

literal
  : FLOAT
  | INT
  | STRING
  ;

// ============================================================
// Lexer rules
// ============================================================

// --- Keywords (case-insensitive) ---

SI        : [sS] [iI] ;
SINON     : [sS] [iI] [nN] [oO] [nN] ;
ALORS     : [aA] [lL] [oO] [rR] [sS] ;

AFFICHER  : [aA] [fF] [fF] [iI] [cC] [hH] [eE] [rR] ;

LIRE
  : [lL] [iI] [rR] [eE]
  ;

LIRE_ENTIER
  : [lL] [iI] [rR] [eE] HWS+ [eE] [nN] [tT] [iI] [eE] [rR]
  ;

SINON_SI
  : SINON HWS+ SI
  ;

FIN_SI
  : [fF] [iI] [nN] HWS+ SI
  ;

TANT_QUE
  : [tT] [aA] [nN] [tT] HWS+ [qQ] [uU] [eE]
  ;

FIN_TANT_QUE
  : [fF] [iI] [nN] HWS+ TANT_QUE
  ;

// --- Operators / punctuation ---

INCR   : '++' ;
DECR   : '--' ;

POW    : '**' ;
MUL    : '*' ;
DIV    : '/' ;
MOD    : '%' ;
PLUS   : '+' ;
MINUS  : '-' ;

LTE    : '<=' ;
GTE    : '>=' ;
LT     : '<' ;
GT     : '>' ;

EQ     : '==' ;
NEQ    : '!=' ;

// Keep ASSIGN after EQ so '==' is not tokenized as '=' '='
ASSIGN : '=' ;

LPAREN : '(' ;
RPAREN : ')' ;
LBRACK : '[' ;
RBRACK : ']' ;
COMMA  : ',' ;
SEMI   : ';' ;

// --- Identifiers (Unicode letters supported; optional trailing '?') ---
// Examples: nom, age, température, est_admin?
IDENT
  : ID_START ID_CONT* '?'?
  ;

fragment ID_START
  : [_\p{L}]
  ;

fragment ID_CONT
  : [_\p{L}\p{Nd}]
  ;

// --- Numbers (strict underscores) ---
// No scientific notation.
// Floats require digits on both sides of the dot.
// Underscores only allowed between digits (no leading, trailing, or doubled underscores).
// Examples:
//   INT:   1, 1_000_000
//   FLOAT: 1.23, 1_000.5, 1.0_5, 1_000.0_5

FLOAT
  : DIGITS '.' DIGITS
  ;

INT
  : DIGITS
  ;

fragment DIGITS
  : DIGIT (DIGIT_OR_US_DIGIT)*
  ;

fragment DIGIT_OR_US_DIGIT
  : DIGIT
  | '_' DIGIT
  ;

fragment DIGIT
  : [0-9]
  ;

// --- Strings ---
STRING
  : '"' ( '\\"' | '\\\\' | ~["\r\n\\] )* '"'
  ;

// --- Comments / whitespace / newlines ---

LINE_COMMENT
  : '//' ~[\r\n]* -> skip
  ;

// NEWLINE is significant (statement separator)
NEWLINE
  : '\r'? '\n'
  ;

// Horizontal whitespace (space/tab) skipped
WS
  : HWS+ -> skip
  ;

fragment HWS
  : [ \t]
  ;