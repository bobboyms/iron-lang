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

STRUCT:
    'struct'
;

//Var Types
TYPE_INT
    : 'int'
;

TYPE_FLOAT
    : 'float'
;

TYPE_STRING
    : 'string'
;

TYPE_BOOLEAN
    : 'boolean'
;

TYPE_DOUBLE
    : 'double'
;

TYPE_CHAR
    : 'char'
;

BOOLEAN_VALUE:
    'true' | 'false'
;

TRAIT
    : 'trait'
;

IMPL
    : 'impl'
;

THIS
    : 'this'
;

PACKAGE
    : 'package'
;

MULT_LINE_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
;

fragment ESC_NEWLINE: '\\' '\n';
fragment OCT_DIGIT: [0-7];
fragment QUOTE_ESCAPE: '\\' ['"];
fragment HEX_DIGIT: [0-9a-fA-F];
fragment COMMON_ESCAPE: '\\' [nrt\\0];
fragment ASCII_ESCAPE: '\\x' OCT_DIGIT HEX_DIGIT | COMMON_ESCAPE;
fragment UNICODE_ESCAPE
   : '\\u{' HEX_DIGIT HEX_DIGIT? HEX_DIGIT? HEX_DIGIT? HEX_DIGIT? HEX_DIGIT? '}'
;

STRING_LITERAL
   : '"'
   (
      ~["]
      | QUOTE_ESCAPE
      | ASCII_ESCAPE
      | UNICODE_ESCAPE
      | ESC_NEWLINE
   )* '"'
   ;

CHAR_LITERAL
: '\''
   (
      ~["]
      | QUOTE_ESCAPE
      | ASCII_ESCAPE
      | UNICODE_ESCAPE
      | ESC_NEWLINE
   ) '\''
   ;

IDENTIFIER:
    ('a'..'z'|'A'..'Z'|'_') ('a'..'z'|'A'..'Z'|'0'..'9'|'_')*;

WHITE_SPACE:
    ( ' ' |'\t' | '\r' | '\n') -> skip
;

program
    : PACKAGE packageName=IDENTIFIER ';' (bodyProgram)*
;

bodyProgram
    : functionNoReturn
    | functionReturn
    | trait
    | impl
    | structDefinition
;

//funcMain
//    : FN 'main'L_PAREN R_PAREN L_CURLY (bodyScope)* R_CURLY
//;

functionReturn
    : FN IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)? L_CURLY (bodyScope)* returnDefinition R_CURLY
;

functionNoReturn
    : FN IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN L_CURLY (bodyScope)* R_CURLY
;

funcType
    : IDENTIFIER L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)?
;

implConstructor
    : 'Constructor' L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN (dataTypes)? L_CURLY (bodyScope)+ R_CURLY
;

returnDefinition
    : (RETURN)? (mathExpression | relExpression | IDENTIFIER | expression)
;

funcCall
    : IDENTIFIER L_PAREN (funcCallArg (COMMA funcCallArg)*)? R_PAREN
;

funcCallArg
    : mathExpression
    | funcCall
    | anonimousFuncWithReturn
    | anonimousFuncNoReturn
    | initStruct
    | initImpl
    | STRING_LITERAL
    | CHAR_LITERAL
    | BOOLEAN_VALUE
;

anonimousFuncWithReturn
    : L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN dataTypes ARROW L_CURLY bodyScope* returnDefinition R_CURLY
    | L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN dataTypes ARROW bodyScope? returnDefinition
;

anonimousFuncNoReturn
    : L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN ARROW L_CURLY bodyScope+ R_CURLY
    | L_PAREN (functionArgs (COMMA functionArgs)*)? R_PAREN ARROW bodyScope
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
    | funcCall
    | println
    | forEach
    | loopWhile
    | loopDoWhile
    | loopForIn
    | loopForI
    | mathExpression
    | expression
;

trait
    : TRAIT interfaceName=IDENTIFIER L_CURLY (funcType)+ R_CURLY
;

impl
    : IMPL implName=IDENTIFIER (FOR traitName=IDENTIFIER)? L_CURLY (definitionVariables)* (implConstructor)* (functionReturn | functionNoReturn)* R_CURLY
;

initImpl
    : implName=IDENTIFIER '::new'  L_PAREN (funcCallArg (COMMA funcCallArg)*)? R_PAREN
;

structDefinition
    : STRUCT structName=IDENTIFIER L_CURLY (definitionVariables)+ R_CURLY;

definitionVariables
    : variableName=IDENTIFIER (dataTypes | genericType=IDENTIFIER)
;

initStruct
    : structName=IDENTIFIER L_CURLY (initStructBody (COMMA initStructBody)*)? R_CURLY
;

initStructBody
    : structKey=IDENTIFIER ':' (funcCall | BOOLEAN_VALUE | INT_NUMBER | REAL_NUMBER | STRING_LITERAL | asIdent=IDENTIFIER | initStruct)
;

println: PRINT_LN L_PAREN (variable | IDENTIFIER) R_PAREN;

variable: (MUT)? LET variableName=IDENTIFIER (dataTypes)?;

functionArgs
    : argName=IDENTIFIER (dataTypes | identType=IDENTIFIER)
    | funcType
;

dataTypes:  TYPE_INT | TYPE_FLOAT | TYPE_BOOLEAN | TYPE_STRING | TYPE_DOUBLE | TYPE_CHAR;

//assignment
//    : variable EQ (STRING_LITERAL | BOOLEAN_VALUE | slice | array | anonimousFuncWithReturn | anonimousFuncNoReturn | mathExpression | relExpression | initStruct | initImpl)
//    | leftExpression=expression EQ (STRING_LITERAL | BOOLEAN_VALUE | slice | array | mathExpression | relExpression | initStruct | rightExpression=expression)
//    | variableName=IDENTIFIER EQ (STRING_LITERAL | BOOLEAN_VALUE | slice | array | anonimousFuncWithReturn | anonimousFuncNoReturn | mathExpression | relExpression | identType=IDENTIFIER)
//;

assignment
    : leftAssignment EQ rigthAssignment
;

leftAssignment
    : variable
    | leftExpression=expression
    | variableName=IDENTIFIER
;

rigthAssignment
    : STRING_LITERAL
    | CHAR_LITERAL
    | BOOLEAN_VALUE
    | identType=IDENTIFIER
    | slice
    | array
    | anonimousFuncWithReturn
    | anonimousFuncNoReturn
    | mathExpression
    | relExpression
    | initStruct
    | initImpl
    | rightExpression=expression
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

map
    : MAP L_PAREN (anonimousFuncWithReturn) R_PAREN
;

filter
    : FILTER L_PAREN (anonimousFuncWithReturn) R_PAREN
;

reduce
    : REDUCE L_PAREN (anonimousFuncWithReturn) R_PAREN
;

//expression
//    : expression DOT expression
//    | slice
//    | IDENTIFIER
//    | map
//    | filter
//    | reduce
//    | array
//    | funcCall
//;

relExpression
    : relExpression (EQEQ | DIF | LT | GT | LTEQ | GTEQ | AND | OR) relExpression
    | L_PAREN relExpression R_PAREN
    | NOT relExpression
    | IDENTIFIER
    | atom
    | funcCall
    | BOOLEAN_VALUE
;

expression
    : leftExpression=expression '.' rigthExpression=expression
    | funcCall
    | IDENTIFIER
    | THIS
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
