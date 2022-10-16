grammar IronLang;

//LEXICAL SYMBOLS
DOT:
    '.'
;
PLUS:
    '+'
    ;
MINUS:
    '-'
;
MULT:
    '*'
;
DIV:
    '/'
;
EQ:
    '='
;
L_PAREN
    : '('
;
R_PAREN
    : ')'
;
L_CURLY
    : '{'
;
R_CURLY
    : '}'
;
SIGN
   : ('+' | '-')
   ;

PLUS_PLUS
    : '++'
;

INT_NUMBER
    : [0-9]+
;

REAL_NUMBER
    : [0-9]+
    | [0-9]+ DOT [0-9]+
;

//Keywords
FN:
    'fn'
;

LET:
    'let'
;

MUT:
    'mut'
;

//Var Types
TYPE_INT
    : 'int'
;

TYPE_FLOAT
    : 'float'
;

IDENTIFIER:
    ('a'..'z'|'A'..'Z'|'_') ('a'..'z'|'A'..'Z'|'0'..'9'|'_')*;

WHITE_SPACE:
    ( ' ' |'\t' | '\r' | '\n') -> skip
;

//PARSER RULES
program:
    funcMain
;

funcMain
    : FN 'main'L_PAREN (functionArgs)* R_PAREN scope
;

scope:
    L_CURLY (assignment | variable | scope)* R_CURLY
;

variable: (MUT)? LET IDENTIFIER (dataTypes)?;

functionArgs: funcArg (',' funcArg)*;

funcArg: IDENTIFIER dataTypes;

dataTypes: TYPE_INT | TYPE_FLOAT;

assignment: (variable | IDENTIFIER) EQ (mathExpression);

mathExpression
   :  mathExpression  (MULT | DIV)  mathExpression
   |  mathExpression  (PLUS | MINUS) mathExpression
   |  L_PAREN mathExpression R_PAREN
   |  (PLUS | MINUS)? atom
   ;

atom
   : INT_NUMBER
   | REAL_NUMBER
   | IDENTIFIER
   ;
