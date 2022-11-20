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
L_BRACKET
    : '['
;

R_BRACKET
    : ']'
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

GTEQ:
    '>='
;

LTEQ:
    '<='
;

GT:
    '>'
;

LT:
    '<'
;

DIF:
    '!='
;

EQEQ:
    '=='
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

IF:
    'if'
;

OR:
    'or'
;

NOT:
    'not'
;

AND:
    'and'
;

LET:
    'let'
;

MUT:
    'mut'
;

ELSE:
    'else'
;

FOR_EACH:
    'forEach'
;

MAP:
    'map'
;

FILTER:
    'filter'
;

REDUCE:
    'reduce'
;

RETURN:
    'return'
;

PRINT_LN:
    'println'
;

//Var Types
TYPE_INT
    : 'int'
;

TYPE_FLOAT
    : 'float'
;

TYPE_BOOLEAN:
    'true' | 'false'
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
    : FN 'main'L_PAREN R_PAREN scope
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

ifScope:
    L_CURLY (bodyScope)* R_CURLY
;

elseExpression
    : ELSE ifScope
;

elseIfExpression
    : ELSE IF relExpression ifScope (elseIfExpression | elseExpression)?
;

ifExpression
   : IF relExpression ifScope (elseExpression | elseIfExpression )?
;

bodyScope: variable | assignment | ifExpression | function | funcCall | scope | println | forEach ;

println: PRINT_LN L_PAREN (variable | IDENTIFIER) R_PAREN;

variable: (MUT)? LET IDENTIFIER (dataTypes)?;

functionArgs: funcArg (COMMA funcArg)*;

funcArg: (MUT)? IDENTIFIER dataTypes | funcType;

dataTypes:  TYPE_INT | TYPE_FLOAT;

assignment
    : variable EQ anonimousFunc
    | variable EQ mathExpression
    | IDENTIFIER EQ mathExpression
    | variable EQ relExpression
    | IDENTIFIER EQ relExpression
    | variable EQ array
    | IDENTIFIER EQ array
    | variable EQ mapFilterReduce
;

array
    : dataTypes L_BRACKET ((INT_NUMBER | REAL_NUMBER) (COMMA (INT_NUMBER | REAL_NUMBER))*)* R_BRACKET
;

forEach
    : IDENTIFIER DOT FOR_EACH L_PAREN (anonimousFunc) R_PAREN
    | array DOT FOR_EACH L_PAREN (anonimousFunc | IDENTIFIER) R_PAREN
;

mapFilterReduce
    : mapFilterReduce DOT mapFilterReduce
    | IDENTIFIER
    | map
    | filter
    | reduce
    | array
;

map
    : MAP L_PAREN (anonimousFunc) R_PAREN
;

filter
    : FILTER L_PAREN (anonimousFunc) R_PAREN
;

reduce
    : REDUCE L_PAREN (anonimousFunc) R_PAREN
;

relExpression
    : relExpression (EQEQ | DIF | LT | GT | LTEQ | GTEQ | AND | OR ) relExpression
    | L_PAREN relExpression R_PAREN
    | NOT relExpression
    | atom
    | funcCall
    | TYPE_BOOLEAN
;

mathExpression
   :  mathExpression  (MULT | DIV)  mathExpression
   |  mathExpression  (PLUS | MINUS) mathExpression
   |  L_PAREN mathExpression R_PAREN
   |  atom
   |  funcCall
   ;

atom
   : INT_NUMBER
   | REAL_NUMBER
   | IDENTIFIER
   ;
