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

MINUS_MINUS:
    '--'
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

DO:
    'do'
;

FN:
    'fn'
;

IF:
    'if'
;

OR:
    'or'
;

IN:
    'in'
;

FOR:
    'for'
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

BREAK:
    'break'
;

CONTINUE:
    'continue'
;

ELSE:
    'else'
;

WHILE:
    'while'
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
    : FN 'main'L_PAREN R_PAREN funcScope
;

function
    : FN IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)? funcScope
;

funcType
    : IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)?
;

return
    : (RETURN)? (mathExpression | relExpression | IDENTIFIER)
;

funcScope:
    L_CURLY (bodyScope)* (return)? R_CURLY
;

funcCall
    : IDENTIFIER L_PAREN (funcCallArg (COMMA funcCallArg)*)? R_PAREN
;

funcCallArg: mathExpression | funcCall | anonimousFuncWithReturn | anonimousFuncNoReturn;

//() int -> {}
anonimousFuncWithReturn
    : L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN dataTypes ARROW L_CURLY (bodyScope)* return R_CURLY
    | L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN dataTypes ARROW (bodyScope)? return
;

anonimousFuncNoReturn
    : L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN ARROW L_CURLY (bodyScope)* R_CURLY
    | L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN ARROW (bodyScope)*
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

loopScope
    : CONTINUE | BREAK | bodyScope
;

loopDoWhile
    : DO L_CURLY (loopScope)* R_CURLY (WHILE relExpression)?
;

loopWhile
    : WHILE relExpression L_CURLY (loopScope)* R_CURLY
;

loopForIn
    : FOR IDENTIFIER IN (slice | IDENTIFIER |  L_PAREN INT_NUMBER '..' INT_NUMBER R_PAREN) L_CURLY (loopScope)* R_CURLY
;

loopForI
    : FOR assignment ';' relExpression ';' mathExpression L_CURLY (loopScope)* R_CURLY
;

bodyScope
    : variable
    | assignment
    | ifExpression
    | function
    | funcCall
    | println
    | forEach
    | loopWhile
    | loopDoWhile
    | loopForIn
    | loopForI
    | mathExpression
    | mapFilterReduce
;

println: PRINT_LN L_PAREN (variable | IDENTIFIER) R_PAREN;

variable: (MUT)? LET IDENTIFIER (dataTypes)?;

functionArgs: funcArg (COMMA funcArg)*;

funcArg: (MUT)? IDENTIFIER dataTypes | funcType;

dataTypes:  TYPE_INT | TYPE_FLOAT;

assignment
    : variable EQ (slice | array | anonimousFuncWithReturn | anonimousFuncNoReturn | mathExpression | relExpression | mapFilterReduce)
;

array
    : dataTypes L_BRACKET ((INT_NUMBER | REAL_NUMBER) (COMMA (INT_NUMBER | REAL_NUMBER))*)* R_BRACKET
;

slice
    : IDENTIFIER L_BRACKET INT_NUMBER ':' INT_NUMBER R_BRACKET
;

forEach
    : IDENTIFIER DOT FOR_EACH L_PAREN (anonimousFuncNoReturn) R_PAREN
    | array DOT FOR_EACH L_PAREN (anonimousFuncNoReturn | IDENTIFIER) R_PAREN
;

mapFilterReduce
    : mapFilterReduce DOT mapFilterReduce
    | slice
    | IDENTIFIER
    | map
    | filter
    | reduce
    | array
;

map
    : MAP L_PAREN (anonimousFuncWithReturn) R_PAREN
;

filter
    : FILTER L_PAREN (anonimousFuncWithReturn) R_PAREN
;

reduce
    : REDUCE L_PAREN (anonimousFuncWithReturn) R_PAREN
;

relExpression
    : relExpression (EQEQ | DIF | LT | GT | LTEQ | GTEQ | AND | OR) relExpression
    | L_PAREN relExpression R_PAREN
    | NOT relExpression
    | IDENTIFIER
    | atom
    | funcCall
    | TYPE_BOOLEAN
;

mathExpression
   : mathExpression  (MULT | DIV)  mathExpression
   | mathExpression  (PLUS | MINUS) mathExpression
   | L_PAREN mathExpression R_PAREN
   | IDENTIFIER (PLUS_PLUS | MINUS_MINUS)?
   | atom
   | funcCall
   ;

atom
   : INT_NUMBER
   | REAL_NUMBER
   ;
