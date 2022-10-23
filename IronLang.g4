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

return
    : (RETURN)? mathExpression
;

scope:
    L_CURLY (bodyScope)* (return)? R_CURLY
;

funcCall
    : IDENTIFIER L_PAREN (funcCallArg (COMMA funcCallArg)*)? R_PAREN
;

funcCallArg: mathExpression | funcCall;

//() int -> {}
anonimousFunc
    : L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)?
      ARROW (bodyScope)*
;

bodyScope: anonimousFunc | variable | assignment | function | funcCall | scope | println ;


println: PRINT_LN L_PAREN (variable | IDENTIFIER) R_PAREN;

variable: (MUT)? LET IDENTIFIER (dataTypes)?;

functionArgs: funcArg (COMMA funcArg)*;

funcArg: (MUT)? IDENTIFIER dataTypes;

dataTypes: TYPE_INT | TYPE_FLOAT;

assignment:  IDENTIFIER | variable EQ (anonimousFunc)? | mathExpression;

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
