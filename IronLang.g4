grammar IronLang;

//LEXICAL SYMBOLS
DOT:
    '.'
;

COMMA:
    ','
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
PLUS_PLUS
    : '++'
;

PIPE
    : '|'
;

COMPOP
    : '=='
    | '!='
    | '<'
    | '<='
    | '>'
    | '>='
;

ARROW
    : '->'
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

PRINT_LN:
    'println'
;

LET:
    'let'
;

MUT:
    'mut'
;

RETURN:
    'return'
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
program
    : (function)* funcMain (function)*
;


funcMain
    : FN 'main'L_PAREN (functionArgs)* R_PAREN scope
;

function
    : FN IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)? scope
;

funcType
    : IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)?
;

return
    : (RETURN)? mathExpression
;

scope:
    L_CURLY (bodyScope)* (return)? R_CURLY
;

funcCall
    : IDENTIFIER L_PAREN (funcCallArg (COMMA funcCallArg)*)? R_PAREN
;

funcCallArg: mathExpression | funcCall | anonimousFunc;

//() int -> {}
anonimousFunc
    : L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)? ARROW L_CURLY (bodyScope)? (return)? R_CURLY
    | L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)? ARROW (bodyScope)? (return)?
;

bodyScope: variable | assignment | function | funcCall | scope | println ;

println: PRINT_LN L_PAREN (variable | IDENTIFIER) R_PAREN;

variable: (MUT)? LET IDENTIFIER (dataTypes)?;

functionArgs: funcArg (COMMA funcArg)*;

funcArg: (MUT)? IDENTIFIER dataTypes | funcType;

dataTypes:  TYPE_INT | TYPE_FLOAT;

assignment
    : variable EQ ( mathExpression | (anonimousFunc)?)
    | IDENTIFIER EQ ( mathExpression | (anonimousFunc)?)
;

mathExpression
   :  mathExpression  (MULT | DIV)  mathExpression
   |  mathExpression  (PLUS | MINUS) mathExpression
   |  L_PAREN mathExpression R_PAREN
   |  (PLUS | MINUS)? atom
   |  (PLUS | MINUS)? funcCall
   ;

atom
   : INT_NUMBER
   | REAL_NUMBER
   | IDENTIFIER
   ;
