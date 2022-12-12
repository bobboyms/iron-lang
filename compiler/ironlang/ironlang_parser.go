// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package ironlang // IronLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type IronLangParser struct {
	*antlr.BaseParser
}

var ironlangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ironlangParserInit() {
	staticData := &ironlangParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'Constructor'", "'..'", "'::new'", "':'", "'.'", "','",
		"'+'", "'-'", "'*'", "'/'", "'='", "'('", "')'", "'{'", "'['", "']'",
		"'}'", "'++'", "'--'", "'|'", "'>='", "'<='", "'>'", "'<'", "'!='",
		"'=='", "'->'", "", "", "'do'", "'fn'", "'if'", "'or'", "'in'", "'for'",
		"'not'", "'and'", "'let'", "'mut'", "'break'", "'continue'", "'else'",
		"'while'", "'forEach'", "'map'", "'filter'", "'reduce'", "'return'",
		"'println'", "'struct'", "'int'", "'float'", "'string'", "'boolean'",
		"'double'", "'char'", "", "'trait'", "'impl'", "'this'", "'package'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV",
		"EQ", "L_PAREN", "R_PAREN", "L_CURLY", "L_BRACKET", "R_BRACKET", "R_CURLY",
		"PLUS_PLUS", "MINUS_MINUS", "PIPE", "GTEQ", "LTEQ", "GT", "LT", "DIF",
		"EQEQ", "ARROW", "INT_NUMBER", "REAL_NUMBER", "DO", "FN", "IF", "OR",
		"IN", "FOR", "NOT", "AND", "LET", "MUT", "BREAK", "CONTINUE", "ELSE",
		"WHILE", "FOR_EACH", "MAP", "FILTER", "REDUCE", "RETURN", "PRINT_LN",
		"STRUCT", "TYPE_INT", "TYPE_FLOAT", "TYPE_STRING", "TYPE_BOOLEAN", "TYPE_DOUBLE",
		"TYPE_CHAR", "BOOLEAN_VALUE", "TRAIT", "IMPL", "THIS", "PACKAGE", "MULT_LINE_COMMENT",
		"STRING_LITERAL", "CHAR_LITERAL", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"program", "bodyProgram", "functionReturn", "functionNoReturn", "funcType",
		"implConstructor", "returnDefinition", "funcCall", "funcCallArg", "anonimousFuncWithReturn",
		"anonimousFuncNoReturn", "ifScope", "elseExpression", "elseIfExpression",
		"ifExpression", "loopScope", "loopDoWhile", "loopWhile", "loopForIn",
		"loopForI", "bodyScope", "trait", "impl", "initImpl", "structDefinition",
		"definitionVariables", "initStruct", "initStructBody", "println", "variable",
		"functionArgs", "dataTypes", "assignment", "leftAssignment", "rigthAssignment",
		"array", "slice", "forEach", "map", "filter", "reduce", "relExpression",
		"expression", "mathExpression", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 681, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 95,
		8, 0, 10, 0, 12, 0, 98, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 105,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 113, 8, 2, 10, 2, 12, 2,
		116, 9, 2, 3, 2, 118, 8, 2, 1, 2, 1, 2, 3, 2, 122, 8, 2, 1, 2, 1, 2, 5,
		2, 126, 8, 2, 10, 2, 12, 2, 129, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 140, 8, 3, 10, 3, 12, 3, 143, 9, 3, 3, 3, 145,
		8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 150, 8, 3, 10, 3, 12, 3, 153, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 162, 8, 4, 10, 4, 12, 4, 165,
		9, 4, 3, 4, 167, 8, 4, 1, 4, 1, 4, 3, 4, 171, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 5, 5, 178, 8, 5, 10, 5, 12, 5, 181, 9, 5, 3, 5, 183, 8, 5, 1,
		5, 1, 5, 3, 5, 187, 8, 5, 1, 5, 1, 5, 4, 5, 191, 8, 5, 11, 5, 12, 5, 192,
		1, 5, 1, 5, 1, 6, 3, 6, 198, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 204, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 211, 8, 7, 10, 7, 12, 7, 214, 9,
		7, 3, 7, 216, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 229, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 235, 8, 9,
		10, 9, 12, 9, 238, 9, 9, 3, 9, 240, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 247, 8, 9, 10, 9, 12, 9, 250, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 5, 9, 259, 8, 9, 10, 9, 12, 9, 262, 9, 9, 3, 9, 264, 8, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 270, 8, 9, 1, 9, 1, 9, 3, 9, 274, 8, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 5, 10, 280, 8, 10, 10, 10, 12, 10, 283, 9, 10,
		3, 10, 285, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 291, 8, 10, 11, 10,
		12, 10, 292, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 301, 8, 10,
		10, 10, 12, 10, 304, 9, 10, 3, 10, 306, 8, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 311, 8, 10, 1, 11, 1, 11, 5, 11, 315, 8, 11, 10, 11, 12, 11, 318, 9,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 331, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 338,
		8, 14, 1, 15, 1, 15, 1, 15, 3, 15, 343, 8, 15, 1, 16, 1, 16, 1, 16, 5,
		16, 348, 8, 16, 10, 16, 12, 16, 351, 9, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		356, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 362, 8, 17, 10, 17, 12,
		17, 365, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 379, 8, 18, 1, 18, 1, 18, 5, 18, 383,
		8, 18, 10, 18, 12, 18, 386, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 398, 8, 19, 10, 19, 12, 19, 401,
		9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 417, 8, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 4, 21, 423, 8, 21, 11, 21, 12, 21, 424, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 433, 8, 22, 1, 22, 1, 22, 5, 22, 437, 8, 22, 10,
		22, 12, 22, 440, 9, 22, 1, 22, 5, 22, 443, 8, 22, 10, 22, 12, 22, 446,
		9, 22, 1, 22, 1, 22, 5, 22, 450, 8, 22, 10, 22, 12, 22, 453, 9, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 463, 8, 23,
		10, 23, 12, 23, 466, 9, 23, 3, 23, 468, 8, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 4, 24, 476, 8, 24, 11, 24, 12, 24, 477, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 25, 3, 25, 485, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 5, 26, 492, 8, 26, 10, 26, 12, 26, 495, 9, 26, 3, 26, 497, 8, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		3, 27, 510, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 516, 8, 28, 1, 28,
		1, 28, 1, 29, 3, 29, 521, 8, 29, 1, 29, 1, 29, 1, 29, 3, 29, 526, 8, 29,
		1, 30, 1, 30, 1, 30, 3, 30, 531, 8, 30, 1, 30, 3, 30, 534, 8, 30, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 545, 8,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 560, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 5, 35, 567, 8, 35, 10, 35, 12, 35, 570, 9, 35, 5, 35, 572, 8, 35, 10,
		35, 12, 35, 575, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 599, 8, 37, 1, 37, 1, 37, 3, 37,
		603, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 631, 8, 41, 1,
		41, 1, 41, 1, 41, 5, 41, 636, 8, 41, 10, 41, 12, 41, 639, 9, 41, 1, 42,
		1, 42, 1, 42, 1, 42, 3, 42, 645, 8, 42, 1, 42, 1, 42, 1, 42, 5, 42, 650,
		8, 42, 10, 42, 12, 42, 653, 9, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 3, 43, 662, 8, 43, 1, 43, 1, 43, 3, 43, 666, 8, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 674, 8, 43, 10, 43, 12, 43, 677,
		9, 43, 1, 44, 1, 44, 1, 44, 0, 3, 82, 84, 86, 45, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
		86, 88, 0, 6, 1, 0, 52, 57, 1, 0, 29, 30, 3, 0, 22, 27, 34, 34, 38, 38,
		1, 0, 19, 20, 1, 0, 10, 11, 1, 0, 8, 9, 762, 0, 90, 1, 0, 0, 0, 2, 104,
		1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6, 133, 1, 0, 0, 0, 8, 156, 1, 0, 0, 0,
		10, 172, 1, 0, 0, 0, 12, 197, 1, 0, 0, 0, 14, 205, 1, 0, 0, 0, 16, 228,
		1, 0, 0, 0, 18, 273, 1, 0, 0, 0, 20, 310, 1, 0, 0, 0, 22, 312, 1, 0, 0,
		0, 24, 321, 1, 0, 0, 0, 26, 324, 1, 0, 0, 0, 28, 332, 1, 0, 0, 0, 30, 342,
		1, 0, 0, 0, 32, 344, 1, 0, 0, 0, 34, 357, 1, 0, 0, 0, 36, 368, 1, 0, 0,
		0, 38, 389, 1, 0, 0, 0, 40, 416, 1, 0, 0, 0, 42, 418, 1, 0, 0, 0, 44, 428,
		1, 0, 0, 0, 46, 456, 1, 0, 0, 0, 48, 471, 1, 0, 0, 0, 50, 481, 1, 0, 0,
		0, 52, 486, 1, 0, 0, 0, 54, 500, 1, 0, 0, 0, 56, 511, 1, 0, 0, 0, 58, 520,
		1, 0, 0, 0, 60, 533, 1, 0, 0, 0, 62, 535, 1, 0, 0, 0, 64, 537, 1, 0, 0,
		0, 66, 544, 1, 0, 0, 0, 68, 559, 1, 0, 0, 0, 70, 561, 1, 0, 0, 0, 72, 578,
		1, 0, 0, 0, 74, 602, 1, 0, 0, 0, 76, 604, 1, 0, 0, 0, 78, 609, 1, 0, 0,
		0, 80, 614, 1, 0, 0, 0, 82, 630, 1, 0, 0, 0, 84, 644, 1, 0, 0, 0, 86, 665,
		1, 0, 0, 0, 88, 678, 1, 0, 0, 0, 90, 91, 5, 62, 0, 0, 91, 92, 5, 66, 0,
		0, 92, 96, 5, 1, 0, 0, 93, 95, 3, 2, 1, 0, 94, 93, 1, 0, 0, 0, 95, 98,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 1, 1, 0, 0, 0,
		98, 96, 1, 0, 0, 0, 99, 105, 3, 6, 3, 0, 100, 105, 3, 4, 2, 0, 101, 105,
		3, 42, 21, 0, 102, 105, 3, 44, 22, 0, 103, 105, 3, 48, 24, 0, 104, 99,
		1, 0, 0, 0, 104, 100, 1, 0, 0, 0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0,
		0, 0, 104, 103, 1, 0, 0, 0, 105, 3, 1, 0, 0, 0, 106, 107, 5, 32, 0, 0,
		107, 108, 5, 66, 0, 0, 108, 117, 5, 13, 0, 0, 109, 114, 3, 60, 30, 0, 110,
		111, 5, 7, 0, 0, 111, 113, 3, 60, 30, 0, 112, 110, 1, 0, 0, 0, 113, 116,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 118, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 121, 5, 14, 0, 0, 120, 122, 3, 62, 31, 0, 121,
		120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 127,
		5, 15, 0, 0, 124, 126, 3, 40, 20, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1,
		0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130, 1, 0, 0,
		0, 129, 127, 1, 0, 0, 0, 130, 131, 3, 12, 6, 0, 131, 132, 5, 18, 0, 0,
		132, 5, 1, 0, 0, 0, 133, 134, 5, 32, 0, 0, 134, 135, 5, 66, 0, 0, 135,
		144, 5, 13, 0, 0, 136, 141, 3, 60, 30, 0, 137, 138, 5, 7, 0, 0, 138, 140,
		3, 60, 30, 0, 139, 137, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1,
		0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0,
		0, 144, 136, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		147, 5, 14, 0, 0, 147, 151, 5, 15, 0, 0, 148, 150, 3, 40, 20, 0, 149, 148,
		1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0,
		0, 0, 152, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 155, 5, 18, 0, 0,
		155, 7, 1, 0, 0, 0, 156, 157, 5, 66, 0, 0, 157, 166, 5, 13, 0, 0, 158,
		163, 3, 60, 30, 0, 159, 160, 5, 7, 0, 0, 160, 162, 3, 60, 30, 0, 161, 159,
		1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 158, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 5, 14, 0, 0, 169,
		171, 3, 62, 31, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 9,
		1, 0, 0, 0, 172, 173, 5, 2, 0, 0, 173, 182, 5, 13, 0, 0, 174, 179, 3, 60,
		30, 0, 175, 176, 5, 7, 0, 0, 176, 178, 3, 60, 30, 0, 177, 175, 1, 0, 0,
		0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 174, 1, 0, 0, 0, 182, 183,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 5, 14, 0, 0, 185, 187, 3, 62,
		31, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 190, 5, 15, 0, 0, 189, 191, 3, 40, 20, 0, 190, 189, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194,
		1, 0, 0, 0, 194, 195, 5, 18, 0, 0, 195, 11, 1, 0, 0, 0, 196, 198, 5, 49,
		0, 0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 203, 1, 0, 0, 0,
		199, 204, 3, 86, 43, 0, 200, 204, 3, 82, 41, 0, 201, 204, 5, 66, 0, 0,
		202, 204, 3, 84, 42, 0, 203, 199, 1, 0, 0, 0, 203, 200, 1, 0, 0, 0, 203,
		201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 13, 1, 0, 0, 0, 205, 206, 5,
		66, 0, 0, 206, 215, 5, 13, 0, 0, 207, 212, 3, 16, 8, 0, 208, 209, 5, 7,
		0, 0, 209, 211, 3, 16, 8, 0, 210, 208, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0,
		212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 215, 207, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217,
		1, 0, 0, 0, 217, 218, 5, 14, 0, 0, 218, 15, 1, 0, 0, 0, 219, 229, 3, 86,
		43, 0, 220, 229, 3, 14, 7, 0, 221, 229, 3, 18, 9, 0, 222, 229, 3, 20, 10,
		0, 223, 229, 3, 52, 26, 0, 224, 229, 3, 46, 23, 0, 225, 229, 5, 64, 0,
		0, 226, 229, 5, 65, 0, 0, 227, 229, 5, 58, 0, 0, 228, 219, 1, 0, 0, 0,
		228, 220, 1, 0, 0, 0, 228, 221, 1, 0, 0, 0, 228, 222, 1, 0, 0, 0, 228,
		223, 1, 0, 0, 0, 228, 224, 1, 0, 0, 0, 228, 225, 1, 0, 0, 0, 228, 226,
		1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 17, 1, 0, 0, 0, 230, 239, 5, 13,
		0, 0, 231, 236, 3, 60, 30, 0, 232, 233, 5, 7, 0, 0, 233, 235, 3, 60, 30,
		0, 234, 232, 1, 0, 0, 0, 235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236,
		237, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 231,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 5, 14,
		0, 0, 242, 243, 3, 62, 31, 0, 243, 244, 5, 28, 0, 0, 244, 248, 5, 15, 0,
		0, 245, 247, 3, 40, 20, 0, 246, 245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0,
		248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 251, 1, 0, 0, 0, 250,
		248, 1, 0, 0, 0, 251, 252, 3, 12, 6, 0, 252, 253, 5, 18, 0, 0, 253, 274,
		1, 0, 0, 0, 254, 263, 5, 13, 0, 0, 255, 260, 3, 60, 30, 0, 256, 257, 5,
		7, 0, 0, 257, 259, 3, 60, 30, 0, 258, 256, 1, 0, 0, 0, 259, 262, 1, 0,
		0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0,
		262, 260, 1, 0, 0, 0, 263, 255, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264,
		265, 1, 0, 0, 0, 265, 266, 5, 14, 0, 0, 266, 267, 3, 62, 31, 0, 267, 269,
		5, 28, 0, 0, 268, 270, 3, 40, 20, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1,
		0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 3, 12, 6, 0, 272, 274, 1, 0, 0,
		0, 273, 230, 1, 0, 0, 0, 273, 254, 1, 0, 0, 0, 274, 19, 1, 0, 0, 0, 275,
		284, 5, 13, 0, 0, 276, 281, 3, 60, 30, 0, 277, 278, 5, 7, 0, 0, 278, 280,
		3, 60, 30, 0, 279, 277, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1,
		0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0,
		0, 284, 276, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286,
		287, 5, 14, 0, 0, 287, 288, 5, 28, 0, 0, 288, 290, 5, 15, 0, 0, 289, 291,
		3, 40, 20, 0, 290, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 290, 1,
		0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 295, 5, 18, 0,
		0, 295, 311, 1, 0, 0, 0, 296, 305, 5, 13, 0, 0, 297, 302, 3, 60, 30, 0,
		298, 299, 5, 7, 0, 0, 299, 301, 3, 60, 30, 0, 300, 298, 1, 0, 0, 0, 301,
		304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 306,
		1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 297, 1, 0, 0, 0, 305, 306, 1, 0,
		0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 5, 14, 0, 0, 308, 309, 5, 28, 0,
		0, 309, 311, 3, 40, 20, 0, 310, 275, 1, 0, 0, 0, 310, 296, 1, 0, 0, 0,
		311, 21, 1, 0, 0, 0, 312, 316, 5, 15, 0, 0, 313, 315, 3, 40, 20, 0, 314,
		313, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317,
		1, 0, 0, 0, 317, 319, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 320, 5, 18,
		0, 0, 320, 23, 1, 0, 0, 0, 321, 322, 5, 43, 0, 0, 322, 323, 3, 22, 11,
		0, 323, 25, 1, 0, 0, 0, 324, 325, 5, 43, 0, 0, 325, 326, 5, 33, 0, 0, 326,
		327, 3, 82, 41, 0, 327, 330, 3, 22, 11, 0, 328, 331, 3, 26, 13, 0, 329,
		331, 3, 24, 12, 0, 330, 328, 1, 0, 0, 0, 330, 329, 1, 0, 0, 0, 330, 331,
		1, 0, 0, 0, 331, 27, 1, 0, 0, 0, 332, 333, 5, 33, 0, 0, 333, 334, 3, 82,
		41, 0, 334, 337, 3, 22, 11, 0, 335, 338, 3, 24, 12, 0, 336, 338, 3, 26,
		13, 0, 337, 335, 1, 0, 0, 0, 337, 336, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0,
		338, 29, 1, 0, 0, 0, 339, 343, 5, 42, 0, 0, 340, 343, 5, 41, 0, 0, 341,
		343, 3, 40, 20, 0, 342, 339, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 341,
		1, 0, 0, 0, 343, 31, 1, 0, 0, 0, 344, 345, 5, 31, 0, 0, 345, 349, 5, 15,
		0, 0, 346, 348, 3, 30, 15, 0, 347, 346, 1, 0, 0, 0, 348, 351, 1, 0, 0,
		0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 352, 1, 0, 0, 0, 351,
		349, 1, 0, 0, 0, 352, 355, 5, 18, 0, 0, 353, 354, 5, 44, 0, 0, 354, 356,
		3, 82, 41, 0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 33, 1, 0,
		0, 0, 357, 358, 5, 44, 0, 0, 358, 359, 3, 82, 41, 0, 359, 363, 5, 15, 0,
		0, 360, 362, 3, 30, 15, 0, 361, 360, 1, 0, 0, 0, 362, 365, 1, 0, 0, 0,
		363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 366, 1, 0, 0, 0, 365,
		363, 1, 0, 0, 0, 366, 367, 5, 18, 0, 0, 367, 35, 1, 0, 0, 0, 368, 369,
		5, 36, 0, 0, 369, 370, 5, 66, 0, 0, 370, 378, 5, 35, 0, 0, 371, 379, 3,
		72, 36, 0, 372, 379, 5, 66, 0, 0, 373, 374, 5, 13, 0, 0, 374, 375, 5, 29,
		0, 0, 375, 376, 5, 3, 0, 0, 376, 377, 5, 29, 0, 0, 377, 379, 5, 14, 0,
		0, 378, 371, 1, 0, 0, 0, 378, 372, 1, 0, 0, 0, 378, 373, 1, 0, 0, 0, 379,
		380, 1, 0, 0, 0, 380, 384, 5, 15, 0, 0, 381, 383, 3, 30, 15, 0, 382, 381,
		1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0,
		0, 0, 385, 387, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 388, 5, 18, 0, 0,
		388, 37, 1, 0, 0, 0, 389, 390, 5, 36, 0, 0, 390, 391, 3, 64, 32, 0, 391,
		392, 5, 1, 0, 0, 392, 393, 3, 82, 41, 0, 393, 394, 5, 1, 0, 0, 394, 395,
		3, 86, 43, 0, 395, 399, 5, 15, 0, 0, 396, 398, 3, 30, 15, 0, 397, 396,
		1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0,
		0, 0, 400, 402, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 402, 403, 5, 18, 0, 0,
		403, 39, 1, 0, 0, 0, 404, 417, 3, 58, 29, 0, 405, 417, 3, 64, 32, 0, 406,
		417, 3, 28, 14, 0, 407, 417, 3, 14, 7, 0, 408, 417, 3, 56, 28, 0, 409,
		417, 3, 74, 37, 0, 410, 417, 3, 34, 17, 0, 411, 417, 3, 32, 16, 0, 412,
		417, 3, 36, 18, 0, 413, 417, 3, 38, 19, 0, 414, 417, 3, 86, 43, 0, 415,
		417, 3, 84, 42, 0, 416, 404, 1, 0, 0, 0, 416, 405, 1, 0, 0, 0, 416, 406,
		1, 0, 0, 0, 416, 407, 1, 0, 0, 0, 416, 408, 1, 0, 0, 0, 416, 409, 1, 0,
		0, 0, 416, 410, 1, 0, 0, 0, 416, 411, 1, 0, 0, 0, 416, 412, 1, 0, 0, 0,
		416, 413, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 416, 415, 1, 0, 0, 0, 417,
		41, 1, 0, 0, 0, 418, 419, 5, 59, 0, 0, 419, 420, 5, 66, 0, 0, 420, 422,
		5, 15, 0, 0, 421, 423, 3, 8, 4, 0, 422, 421, 1, 0, 0, 0, 423, 424, 1, 0,
		0, 0, 424, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0,
		426, 427, 5, 18, 0, 0, 427, 43, 1, 0, 0, 0, 428, 429, 5, 60, 0, 0, 429,
		432, 5, 66, 0, 0, 430, 431, 5, 36, 0, 0, 431, 433, 5, 66, 0, 0, 432, 430,
		1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 438, 5, 15,
		0, 0, 435, 437, 3, 50, 25, 0, 436, 435, 1, 0, 0, 0, 437, 440, 1, 0, 0,
		0, 438, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 444, 1, 0, 0, 0, 440,
		438, 1, 0, 0, 0, 441, 443, 3, 10, 5, 0, 442, 441, 1, 0, 0, 0, 443, 446,
		1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 445, 1, 0, 0, 0, 445, 451, 1, 0,
		0, 0, 446, 444, 1, 0, 0, 0, 447, 450, 3, 4, 2, 0, 448, 450, 3, 6, 3, 0,
		449, 447, 1, 0, 0, 0, 449, 448, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451,
		449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 454, 1, 0, 0, 0, 453, 451,
		1, 0, 0, 0, 454, 455, 5, 18, 0, 0, 455, 45, 1, 0, 0, 0, 456, 457, 5, 66,
		0, 0, 457, 458, 5, 4, 0, 0, 458, 467, 5, 13, 0, 0, 459, 464, 3, 16, 8,
		0, 460, 461, 5, 7, 0, 0, 461, 463, 3, 16, 8, 0, 462, 460, 1, 0, 0, 0, 463,
		466, 1, 0, 0, 0, 464, 462, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 468,
		1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 467, 459, 1, 0, 0, 0, 467, 468, 1, 0,
		0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 5, 14, 0, 0, 470, 47, 1, 0, 0, 0,
		471, 472, 5, 51, 0, 0, 472, 473, 5, 66, 0, 0, 473, 475, 5, 15, 0, 0, 474,
		476, 3, 50, 25, 0, 475, 474, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 475,
		1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479, 480, 5, 18,
		0, 0, 480, 49, 1, 0, 0, 0, 481, 484, 5, 66, 0, 0, 482, 485, 3, 62, 31,
		0, 483, 485, 5, 66, 0, 0, 484, 482, 1, 0, 0, 0, 484, 483, 1, 0, 0, 0, 485,
		51, 1, 0, 0, 0, 486, 487, 5, 66, 0, 0, 487, 496, 5, 15, 0, 0, 488, 493,
		3, 54, 27, 0, 489, 490, 5, 7, 0, 0, 490, 492, 3, 54, 27, 0, 491, 489, 1,
		0, 0, 0, 492, 495, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 493, 494, 1, 0, 0,
		0, 494, 497, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 496, 488, 1, 0, 0, 0, 496,
		497, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 499, 5, 18, 0, 0, 499, 53,
		1, 0, 0, 0, 500, 501, 5, 66, 0, 0, 501, 509, 5, 5, 0, 0, 502, 510, 3, 14,
		7, 0, 503, 510, 5, 58, 0, 0, 504, 510, 5, 29, 0, 0, 505, 510, 5, 30, 0,
		0, 506, 510, 5, 64, 0, 0, 507, 510, 5, 66, 0, 0, 508, 510, 3, 52, 26, 0,
		509, 502, 1, 0, 0, 0, 509, 503, 1, 0, 0, 0, 509, 504, 1, 0, 0, 0, 509,
		505, 1, 0, 0, 0, 509, 506, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 509, 508,
		1, 0, 0, 0, 510, 55, 1, 0, 0, 0, 511, 512, 5, 50, 0, 0, 512, 515, 5, 13,
		0, 0, 513, 516, 3, 58, 29, 0, 514, 516, 5, 66, 0, 0, 515, 513, 1, 0, 0,
		0, 515, 514, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 518, 5, 14, 0, 0, 518,
		57, 1, 0, 0, 0, 519, 521, 5, 40, 0, 0, 520, 519, 1, 0, 0, 0, 520, 521,
		1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 523, 5, 39, 0, 0, 523, 525, 5, 66,
		0, 0, 524, 526, 3, 62, 31, 0, 525, 524, 1, 0, 0, 0, 525, 526, 1, 0, 0,
		0, 526, 59, 1, 0, 0, 0, 527, 530, 5, 66, 0, 0, 528, 531, 3, 62, 31, 0,
		529, 531, 5, 66, 0, 0, 530, 528, 1, 0, 0, 0, 530, 529, 1, 0, 0, 0, 531,
		534, 1, 0, 0, 0, 532, 534, 3, 8, 4, 0, 533, 527, 1, 0, 0, 0, 533, 532,
		1, 0, 0, 0, 534, 61, 1, 0, 0, 0, 535, 536, 7, 0, 0, 0, 536, 63, 1, 0, 0,
		0, 537, 538, 3, 66, 33, 0, 538, 539, 5, 12, 0, 0, 539, 540, 3, 68, 34,
		0, 540, 65, 1, 0, 0, 0, 541, 545, 3, 58, 29, 0, 542, 545, 3, 84, 42, 0,
		543, 545, 5, 66, 0, 0, 544, 541, 1, 0, 0, 0, 544, 542, 1, 0, 0, 0, 544,
		543, 1, 0, 0, 0, 545, 67, 1, 0, 0, 0, 546, 560, 5, 64, 0, 0, 547, 560,
		5, 65, 0, 0, 548, 560, 5, 58, 0, 0, 549, 560, 5, 66, 0, 0, 550, 560, 3,
		72, 36, 0, 551, 560, 3, 70, 35, 0, 552, 560, 3, 18, 9, 0, 553, 560, 3,
		20, 10, 0, 554, 560, 3, 86, 43, 0, 555, 560, 3, 82, 41, 0, 556, 560, 3,
		52, 26, 0, 557, 560, 3, 46, 23, 0, 558, 560, 3, 84, 42, 0, 559, 546, 1,
		0, 0, 0, 559, 547, 1, 0, 0, 0, 559, 548, 1, 0, 0, 0, 559, 549, 1, 0, 0,
		0, 559, 550, 1, 0, 0, 0, 559, 551, 1, 0, 0, 0, 559, 552, 1, 0, 0, 0, 559,
		553, 1, 0, 0, 0, 559, 554, 1, 0, 0, 0, 559, 555, 1, 0, 0, 0, 559, 556,
		1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 559, 558, 1, 0, 0, 0, 560, 69, 1, 0,
		0, 0, 561, 562, 3, 62, 31, 0, 562, 573, 5, 16, 0, 0, 563, 568, 7, 1, 0,
		0, 564, 565, 5, 7, 0, 0, 565, 567, 7, 1, 0, 0, 566, 564, 1, 0, 0, 0, 567,
		570, 1, 0, 0, 0, 568, 566, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0, 569, 572,
		1, 0, 0, 0, 570, 568, 1, 0, 0, 0, 571, 563, 1, 0, 0, 0, 572, 575, 1, 0,
		0, 0, 573, 571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 576, 1, 0, 0, 0,
		575, 573, 1, 0, 0, 0, 576, 577, 5, 17, 0, 0, 577, 71, 1, 0, 0, 0, 578,
		579, 5, 66, 0, 0, 579, 580, 5, 16, 0, 0, 580, 581, 5, 29, 0, 0, 581, 582,
		5, 5, 0, 0, 582, 583, 5, 29, 0, 0, 583, 584, 5, 17, 0, 0, 584, 73, 1, 0,
		0, 0, 585, 586, 5, 66, 0, 0, 586, 587, 5, 6, 0, 0, 587, 588, 5, 45, 0,
		0, 588, 589, 5, 13, 0, 0, 589, 590, 3, 20, 10, 0, 590, 591, 5, 14, 0, 0,
		591, 603, 1, 0, 0, 0, 592, 593, 3, 70, 35, 0, 593, 594, 5, 6, 0, 0, 594,
		595, 5, 45, 0, 0, 595, 598, 5, 13, 0, 0, 596, 599, 3, 20, 10, 0, 597, 599,
		5, 66, 0, 0, 598, 596, 1, 0, 0, 0, 598, 597, 1, 0, 0, 0, 599, 600, 1, 0,
		0, 0, 600, 601, 5, 14, 0, 0, 601, 603, 1, 0, 0, 0, 602, 585, 1, 0, 0, 0,
		602, 592, 1, 0, 0, 0, 603, 75, 1, 0, 0, 0, 604, 605, 5, 46, 0, 0, 605,
		606, 5, 13, 0, 0, 606, 607, 3, 18, 9, 0, 607, 608, 5, 14, 0, 0, 608, 77,
		1, 0, 0, 0, 609, 610, 5, 47, 0, 0, 610, 611, 5, 13, 0, 0, 611, 612, 3,
		18, 9, 0, 612, 613, 5, 14, 0, 0, 613, 79, 1, 0, 0, 0, 614, 615, 5, 48,
		0, 0, 615, 616, 5, 13, 0, 0, 616, 617, 3, 18, 9, 0, 617, 618, 5, 14, 0,
		0, 618, 81, 1, 0, 0, 0, 619, 620, 6, 41, -1, 0, 620, 621, 5, 13, 0, 0,
		621, 622, 3, 82, 41, 0, 622, 623, 5, 14, 0, 0, 623, 631, 1, 0, 0, 0, 624,
		625, 5, 37, 0, 0, 625, 631, 3, 82, 41, 5, 626, 631, 5, 66, 0, 0, 627, 631,
		3, 88, 44, 0, 628, 631, 3, 14, 7, 0, 629, 631, 5, 58, 0, 0, 630, 619, 1,
		0, 0, 0, 630, 624, 1, 0, 0, 0, 630, 626, 1, 0, 0, 0, 630, 627, 1, 0, 0,
		0, 630, 628, 1, 0, 0, 0, 630, 629, 1, 0, 0, 0, 631, 637, 1, 0, 0, 0, 632,
		633, 10, 7, 0, 0, 633, 634, 7, 2, 0, 0, 634, 636, 3, 82, 41, 8, 635, 632,
		1, 0, 0, 0, 636, 639, 1, 0, 0, 0, 637, 635, 1, 0, 0, 0, 637, 638, 1, 0,
		0, 0, 638, 83, 1, 0, 0, 0, 639, 637, 1, 0, 0, 0, 640, 641, 6, 42, -1, 0,
		641, 645, 3, 14, 7, 0, 642, 645, 5, 66, 0, 0, 643, 645, 5, 61, 0, 0, 644,
		640, 1, 0, 0, 0, 644, 642, 1, 0, 0, 0, 644, 643, 1, 0, 0, 0, 645, 651,
		1, 0, 0, 0, 646, 647, 10, 4, 0, 0, 647, 648, 5, 6, 0, 0, 648, 650, 3, 84,
		42, 5, 649, 646, 1, 0, 0, 0, 650, 653, 1, 0, 0, 0, 651, 649, 1, 0, 0, 0,
		651, 652, 1, 0, 0, 0, 652, 85, 1, 0, 0, 0, 653, 651, 1, 0, 0, 0, 654, 655,
		6, 43, -1, 0, 655, 656, 5, 13, 0, 0, 656, 657, 3, 86, 43, 0, 657, 658,
		5, 14, 0, 0, 658, 666, 1, 0, 0, 0, 659, 661, 5, 66, 0, 0, 660, 662, 7,
		3, 0, 0, 661, 660, 1, 0, 0, 0, 661, 662, 1, 0, 0, 0, 662, 666, 1, 0, 0,
		0, 663, 666, 3, 88, 44, 0, 664, 666, 3, 14, 7, 0, 665, 654, 1, 0, 0, 0,
		665, 659, 1, 0, 0, 0, 665, 663, 1, 0, 0, 0, 665, 664, 1, 0, 0, 0, 666,
		675, 1, 0, 0, 0, 667, 668, 10, 6, 0, 0, 668, 669, 7, 4, 0, 0, 669, 674,
		3, 86, 43, 7, 670, 671, 10, 5, 0, 0, 671, 672, 7, 5, 0, 0, 672, 674, 3,
		86, 43, 6, 673, 667, 1, 0, 0, 0, 673, 670, 1, 0, 0, 0, 674, 677, 1, 0,
		0, 0, 675, 673, 1, 0, 0, 0, 675, 676, 1, 0, 0, 0, 676, 87, 1, 0, 0, 0,
		677, 675, 1, 0, 0, 0, 678, 679, 7, 1, 0, 0, 679, 89, 1, 0, 0, 0, 77, 96,
		104, 114, 117, 121, 127, 141, 144, 151, 163, 166, 170, 179, 182, 186, 192,
		197, 203, 212, 215, 228, 236, 239, 248, 260, 263, 269, 273, 281, 284, 292,
		302, 305, 310, 316, 330, 337, 342, 349, 355, 363, 378, 384, 399, 416, 424,
		432, 438, 444, 449, 451, 464, 467, 477, 484, 493, 496, 509, 515, 520, 525,
		530, 533, 544, 559, 568, 573, 598, 602, 630, 637, 644, 651, 661, 665, 673,
		675,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// IronLangParserInit initializes any static state used to implement IronLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIronLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IronLangParserInit() {
	staticData := &ironlangParserStaticData
	staticData.once.Do(ironlangParserInit)
}

// NewIronLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIronLangParser(input antlr.TokenStream) *IronLangParser {
	IronLangParserInit()
	this := new(IronLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ironlangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// IronLangParser tokens.
const (
	IronLangParserEOF               = antlr.TokenEOF
	IronLangParserT__0              = 1
	IronLangParserT__1              = 2
	IronLangParserT__2              = 3
	IronLangParserT__3              = 4
	IronLangParserT__4              = 5
	IronLangParserDOT               = 6
	IronLangParserCOMMA             = 7
	IronLangParserPLUS              = 8
	IronLangParserMINUS             = 9
	IronLangParserMULT              = 10
	IronLangParserDIV               = 11
	IronLangParserEQ                = 12
	IronLangParserL_PAREN           = 13
	IronLangParserR_PAREN           = 14
	IronLangParserL_CURLY           = 15
	IronLangParserL_BRACKET         = 16
	IronLangParserR_BRACKET         = 17
	IronLangParserR_CURLY           = 18
	IronLangParserPLUS_PLUS         = 19
	IronLangParserMINUS_MINUS       = 20
	IronLangParserPIPE              = 21
	IronLangParserGTEQ              = 22
	IronLangParserLTEQ              = 23
	IronLangParserGT                = 24
	IronLangParserLT                = 25
	IronLangParserDIF               = 26
	IronLangParserEQEQ              = 27
	IronLangParserARROW             = 28
	IronLangParserINT_NUMBER        = 29
	IronLangParserREAL_NUMBER       = 30
	IronLangParserDO                = 31
	IronLangParserFN                = 32
	IronLangParserIF                = 33
	IronLangParserOR                = 34
	IronLangParserIN                = 35
	IronLangParserFOR               = 36
	IronLangParserNOT               = 37
	IronLangParserAND               = 38
	IronLangParserLET               = 39
	IronLangParserMUT               = 40
	IronLangParserBREAK             = 41
	IronLangParserCONTINUE          = 42
	IronLangParserELSE              = 43
	IronLangParserWHILE             = 44
	IronLangParserFOR_EACH          = 45
	IronLangParserMAP               = 46
	IronLangParserFILTER            = 47
	IronLangParserREDUCE            = 48
	IronLangParserRETURN            = 49
	IronLangParserPRINT_LN          = 50
	IronLangParserSTRUCT            = 51
	IronLangParserTYPE_INT          = 52
	IronLangParserTYPE_FLOAT        = 53
	IronLangParserTYPE_STRING       = 54
	IronLangParserTYPE_BOOLEAN      = 55
	IronLangParserTYPE_DOUBLE       = 56
	IronLangParserTYPE_CHAR         = 57
	IronLangParserBOOLEAN_VALUE     = 58
	IronLangParserTRAIT             = 59
	IronLangParserIMPL              = 60
	IronLangParserTHIS              = 61
	IronLangParserPACKAGE           = 62
	IronLangParserMULT_LINE_COMMENT = 63
	IronLangParserSTRING_LITERAL    = 64
	IronLangParserCHAR_LITERAL      = 65
	IronLangParserIDENTIFIER        = 66
	IronLangParserWHITE_SPACE       = 67
)

// IronLangParser rules.
const (
	IronLangParserRULE_program                 = 0
	IronLangParserRULE_bodyProgram             = 1
	IronLangParserRULE_functionReturn          = 2
	IronLangParserRULE_functionNoReturn        = 3
	IronLangParserRULE_funcType                = 4
	IronLangParserRULE_implConstructor         = 5
	IronLangParserRULE_returnDefinition        = 6
	IronLangParserRULE_funcCall                = 7
	IronLangParserRULE_funcCallArg             = 8
	IronLangParserRULE_anonimousFuncWithReturn = 9
	IronLangParserRULE_anonimousFuncNoReturn   = 10
	IronLangParserRULE_ifScope                 = 11
	IronLangParserRULE_elseExpression          = 12
	IronLangParserRULE_elseIfExpression        = 13
	IronLangParserRULE_ifExpression            = 14
	IronLangParserRULE_loopScope               = 15
	IronLangParserRULE_loopDoWhile             = 16
	IronLangParserRULE_loopWhile               = 17
	IronLangParserRULE_loopForIn               = 18
	IronLangParserRULE_loopForI                = 19
	IronLangParserRULE_bodyScope               = 20
	IronLangParserRULE_trait                   = 21
	IronLangParserRULE_impl                    = 22
	IronLangParserRULE_initImpl                = 23
	IronLangParserRULE_structDefinition        = 24
	IronLangParserRULE_definitionVariables     = 25
	IronLangParserRULE_initStruct              = 26
	IronLangParserRULE_initStructBody          = 27
	IronLangParserRULE_println                 = 28
	IronLangParserRULE_variable                = 29
	IronLangParserRULE_functionArgs            = 30
	IronLangParserRULE_dataTypes               = 31
	IronLangParserRULE_assignment              = 32
	IronLangParserRULE_leftAssignment          = 33
	IronLangParserRULE_rigthAssignment         = 34
	IronLangParserRULE_array                   = 35
	IronLangParserRULE_slice                   = 36
	IronLangParserRULE_forEach                 = 37
	IronLangParserRULE_map                     = 38
	IronLangParserRULE_filter                  = 39
	IronLangParserRULE_reduce                  = 40
	IronLangParserRULE_relExpression           = 41
	IronLangParserRULE_expression              = 42
	IronLangParserRULE_mathExpression          = 43
	IronLangParserRULE_atom                    = 44
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPackageName returns the packageName token.
	GetPackageName() antlr.Token

	// SetPackageName sets the packageName token.
	SetPackageName(antlr.Token)

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	packageName antlr.Token
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) GetPackageName() antlr.Token { return s.packageName }

func (s *ProgramContext) SetPackageName(v antlr.Token) { s.packageName = v }

func (s *ProgramContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(IronLangParserPACKAGE, 0)
}

func (s *ProgramContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *ProgramContext) AllBodyProgram() []IBodyProgramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyProgramContext); ok {
			len++
		}
	}

	tst := make([]IBodyProgramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyProgramContext); ok {
			tst[i] = t.(IBodyProgramContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) BodyProgram(i int) IBodyProgramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyProgramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyProgramContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IronLangParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(IronLangParserPACKAGE)
	}
	{
		p.SetState(91)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*ProgramContext).packageName = _m
	}
	{
		p.SetState(92)
		p.Match(IronLangParserT__0)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1731634061018923008) != 0 {
		{
			p.SetState(93)
			p.BodyProgram()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBodyProgramContext is an interface to support dynamic dispatch.
type IBodyProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyProgramContext differentiates from other interfaces.
	IsBodyProgramContext()
}

type BodyProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyProgramContext() *BodyProgramContext {
	var p = new(BodyProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_bodyProgram
	return p
}

func (*BodyProgramContext) IsBodyProgramContext() {}

func NewBodyProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyProgramContext {
	var p = new(BodyProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_bodyProgram

	return p
}

func (s *BodyProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyProgramContext) FunctionNoReturn() IFunctionNoReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionNoReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionNoReturnContext)
}

func (s *BodyProgramContext) FunctionReturn() IFunctionReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnContext)
}

func (s *BodyProgramContext) Trait() ITraitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITraitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITraitContext)
}

func (s *BodyProgramContext) Impl() IImplContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplContext)
}

func (s *BodyProgramContext) StructDefinition() IStructDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefinitionContext)
}

func (s *BodyProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterBodyProgram(s)
	}
}

func (s *BodyProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitBodyProgram(s)
	}
}

func (s *BodyProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitBodyProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) BodyProgram() (localctx IBodyProgramContext) {
	this := p
	_ = this

	localctx = NewBodyProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IronLangParserRULE_bodyProgram)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.FunctionNoReturn()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.FunctionReturn()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.Trait()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Impl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(103)
			p.StructDefinition()
		}

	}

	return localctx
}

// IFunctionReturnContext is an interface to support dynamic dispatch.
type IFunctionReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionReturnContext differentiates from other interfaces.
	IsFunctionReturnContext()
}

type FunctionReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionReturnContext() *FunctionReturnContext {
	var p = new(FunctionReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_functionReturn
	return p
}

func (*FunctionReturnContext) IsFunctionReturnContext() {}

func NewFunctionReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnContext {
	var p = new(FunctionReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_functionReturn

	return p
}

func (s *FunctionReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnContext) FN() antlr.TerminalNode {
	return s.GetToken(IronLangParserFN, 0)
}

func (s *FunctionReturnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *FunctionReturnContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FunctionReturnContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FunctionReturnContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *FunctionReturnContext) ReturnDefinition() IReturnDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnDefinitionContext)
}

func (s *FunctionReturnContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *FunctionReturnContext) AllFunctionArgs() []IFunctionArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgsContext); ok {
			tst[i] = t.(IFunctionArgsContext)
			i++
		}
	}

	return tst
}

func (s *FunctionReturnContext) FunctionArgs(i int) IFunctionArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionReturnContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *FunctionReturnContext) AllBodyScope() []IBodyScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyScopeContext); ok {
			len++
		}
	}

	tst := make([]IBodyScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyScopeContext); ok {
			tst[i] = t.(IBodyScopeContext)
			i++
		}
	}

	return tst
}

func (s *FunctionReturnContext) BodyScope(i int) IBodyScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *FunctionReturnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *FunctionReturnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *FunctionReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFunctionReturn(s)
	}
}

func (s *FunctionReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFunctionReturn(s)
	}
}

func (s *FunctionReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFunctionReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FunctionReturn() (localctx IFunctionReturnContext) {
	this := p
	_ = this

	localctx = NewFunctionReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IronLangParserRULE_functionReturn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(107)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(108)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserIDENTIFIER {
		{
			p.SetState(109)
			p.FunctionArgs()
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(110)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(111)
				p.FunctionArgs()
			}

			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(119)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283726776524341248) != 0 {
		{
			p.SetState(120)
			p.DataTypes()
		}

	}
	{
		p.SetState(123)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.BodyScope()
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	{
		p.SetState(130)
		p.ReturnDefinition()
	}
	{
		p.SetState(131)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IFunctionNoReturnContext is an interface to support dynamic dispatch.
type IFunctionNoReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNoReturnContext differentiates from other interfaces.
	IsFunctionNoReturnContext()
}

type FunctionNoReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNoReturnContext() *FunctionNoReturnContext {
	var p = new(FunctionNoReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_functionNoReturn
	return p
}

func (*FunctionNoReturnContext) IsFunctionNoReturnContext() {}

func NewFunctionNoReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNoReturnContext {
	var p = new(FunctionNoReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_functionNoReturn

	return p
}

func (s *FunctionNoReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNoReturnContext) FN() antlr.TerminalNode {
	return s.GetToken(IronLangParserFN, 0)
}

func (s *FunctionNoReturnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *FunctionNoReturnContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FunctionNoReturnContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FunctionNoReturnContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *FunctionNoReturnContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *FunctionNoReturnContext) AllFunctionArgs() []IFunctionArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgsContext); ok {
			tst[i] = t.(IFunctionArgsContext)
			i++
		}
	}

	return tst
}

func (s *FunctionNoReturnContext) FunctionArgs(i int) IFunctionArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionNoReturnContext) AllBodyScope() []IBodyScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyScopeContext); ok {
			len++
		}
	}

	tst := make([]IBodyScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyScopeContext); ok {
			tst[i] = t.(IBodyScopeContext)
			i++
		}
	}

	return tst
}

func (s *FunctionNoReturnContext) BodyScope(i int) IBodyScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *FunctionNoReturnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *FunctionNoReturnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *FunctionNoReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNoReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNoReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFunctionNoReturn(s)
	}
}

func (s *FunctionNoReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFunctionNoReturn(s)
	}
}

func (s *FunctionNoReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFunctionNoReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FunctionNoReturn() (localctx IFunctionNoReturnContext) {
	this := p
	_ = this

	localctx = NewFunctionNoReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IronLangParserRULE_functionNoReturn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(134)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(135)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserIDENTIFIER {
		{
			p.SetState(136)
			p.FunctionArgs()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(137)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(138)
				p.FunctionArgs()
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(146)
		p.Match(IronLangParserR_PAREN)
	}
	{
		p.SetState(147)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323448645386241) != 0 {
		{
			p.SetState(148)
			p.BodyScope()
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(154)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IFuncTypeContext is an interface to support dynamic dispatch.
type IFuncTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncTypeContext differentiates from other interfaces.
	IsFuncTypeContext()
}

type FuncTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncTypeContext() *FuncTypeContext {
	var p = new(FuncTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_funcType
	return p
}

func (*FuncTypeContext) IsFuncTypeContext() {}

func NewFuncTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncTypeContext {
	var p = new(FuncTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_funcType

	return p
}

func (s *FuncTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *FuncTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FuncTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FuncTypeContext) AllFunctionArgs() []IFunctionArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgsContext); ok {
			tst[i] = t.(IFunctionArgsContext)
			i++
		}
	}

	return tst
}

func (s *FuncTypeContext) FunctionArgs(i int) IFunctionArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FuncTypeContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *FuncTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *FuncTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *FuncTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFuncType(s)
	}
}

func (s *FuncTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFuncType(s)
	}
}

func (s *FuncTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFuncType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FuncType() (localctx IFuncTypeContext) {
	this := p
	_ = this

	localctx = NewFuncTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IronLangParserRULE_funcType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(157)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserIDENTIFIER {
		{
			p.SetState(158)
			p.FunctionArgs()
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(159)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(160)
				p.FunctionArgs()
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(168)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283726776524341248) != 0 {
		{
			p.SetState(169)
			p.DataTypes()
		}

	}

	return localctx
}

// IImplConstructorContext is an interface to support dynamic dispatch.
type IImplConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplConstructorContext differentiates from other interfaces.
	IsImplConstructorContext()
}

type ImplConstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplConstructorContext() *ImplConstructorContext {
	var p = new(ImplConstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_implConstructor
	return p
}

func (*ImplConstructorContext) IsImplConstructorContext() {}

func NewImplConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplConstructorContext {
	var p = new(ImplConstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_implConstructor

	return p
}

func (s *ImplConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplConstructorContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *ImplConstructorContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *ImplConstructorContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *ImplConstructorContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *ImplConstructorContext) AllFunctionArgs() []IFunctionArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgsContext); ok {
			tst[i] = t.(IFunctionArgsContext)
			i++
		}
	}

	return tst
}

func (s *ImplConstructorContext) FunctionArgs(i int) IFunctionArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *ImplConstructorContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *ImplConstructorContext) AllBodyScope() []IBodyScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyScopeContext); ok {
			len++
		}
	}

	tst := make([]IBodyScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyScopeContext); ok {
			tst[i] = t.(IBodyScopeContext)
			i++
		}
	}

	return tst
}

func (s *ImplConstructorContext) BodyScope(i int) IBodyScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *ImplConstructorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *ImplConstructorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *ImplConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterImplConstructor(s)
	}
}

func (s *ImplConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitImplConstructor(s)
	}
}

func (s *ImplConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitImplConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) ImplConstructor() (localctx IImplConstructorContext) {
	this := p
	_ = this

	localctx = NewImplConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IronLangParserRULE_implConstructor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(IronLangParserT__1)
	}
	{
		p.SetState(173)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserIDENTIFIER {
		{
			p.SetState(174)
			p.FunctionArgs()
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(175)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(176)
				p.FunctionArgs()
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(184)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283726776524341248) != 0 {
		{
			p.SetState(185)
			p.DataTypes()
		}

	}
	{
		p.SetState(188)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323448645386241) != 0 {
		{
			p.SetState(189)
			p.BodyScope()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IReturnDefinitionContext is an interface to support dynamic dispatch.
type IReturnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnDefinitionContext differentiates from other interfaces.
	IsReturnDefinitionContext()
}

type ReturnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnDefinitionContext() *ReturnDefinitionContext {
	var p = new(ReturnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_returnDefinition
	return p
}

func (*ReturnDefinitionContext) IsReturnDefinitionContext() {}

func NewReturnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnDefinitionContext {
	var p = new(ReturnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_returnDefinition

	return p
}

func (s *ReturnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnDefinitionContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *ReturnDefinitionContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *ReturnDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *ReturnDefinitionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnDefinitionContext) RETURN() antlr.TerminalNode {
	return s.GetToken(IronLangParserRETURN, 0)
}

func (s *ReturnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterReturnDefinition(s)
	}
}

func (s *ReturnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitReturnDefinition(s)
	}
}

func (s *ReturnDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitReturnDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) ReturnDefinition() (localctx IReturnDefinitionContext) {
	this := p
	_ = this

	localctx = NewReturnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IronLangParserRULE_returnDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserRETURN {
		{
			p.SetState(196)
			p.Match(IronLangParserRETURN)
		}

	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(199)
			p.mathExpression(0)
		}

	case 2:
		{
			p.SetState(200)
			p.relExpression(0)
		}

	case 3:
		{
			p.SetState(201)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 4:
		{
			p.SetState(202)
			p.expression(0)
		}

	}

	return localctx
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_funcCall
	return p
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *FuncCallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FuncCallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FuncCallContext) AllFuncCallArg() []IFuncCallArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncCallArgContext); ok {
			len++
		}
	}

	tst := make([]IFuncCallArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncCallArgContext); ok {
			tst[i] = t.(IFuncCallArgContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallContext) FuncCallArg(i int) IFuncCallArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallArgContext)
}

func (s *FuncCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *FuncCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FuncCall() (localctx IFuncCallContext) {
	this := p
	_ = this

	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IronLangParserRULE_funcCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(206)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&15797783068082177) != 0 {
		{
			p.SetState(207)
			p.FuncCallArg()
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(208)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(209)
				p.FuncCallArg()
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(217)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IFuncCallArgContext is an interface to support dynamic dispatch.
type IFuncCallArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallArgContext differentiates from other interfaces.
	IsFuncCallArgContext()
}

type FuncCallArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallArgContext() *FuncCallArgContext {
	var p = new(FuncCallArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_funcCallArg
	return p
}

func (*FuncCallArgContext) IsFuncCallArgContext() {}

func NewFuncCallArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallArgContext {
	var p = new(FuncCallArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_funcCallArg

	return p
}

func (s *FuncCallArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallArgContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *FuncCallArgContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *FuncCallArgContext) AnonimousFuncWithReturn() IAnonimousFuncWithReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncWithReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncWithReturnContext)
}

func (s *FuncCallArgContext) AnonimousFuncNoReturn() IAnonimousFuncNoReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncNoReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncNoReturnContext)
}

func (s *FuncCallArgContext) InitStruct() IInitStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitStructContext)
}

func (s *FuncCallArgContext) InitImpl() IInitImplContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitImplContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitImplContext)
}

func (s *FuncCallArgContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(IronLangParserSTRING_LITERAL, 0)
}

func (s *FuncCallArgContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(IronLangParserCHAR_LITERAL, 0)
}

func (s *FuncCallArgContext) BOOLEAN_VALUE() antlr.TerminalNode {
	return s.GetToken(IronLangParserBOOLEAN_VALUE, 0)
}

func (s *FuncCallArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFuncCallArg(s)
	}
}

func (s *FuncCallArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFuncCallArg(s)
	}
}

func (s *FuncCallArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFuncCallArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FuncCallArg() (localctx IFuncCallArgContext) {
	this := p
	_ = this

	localctx = NewFuncCallArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IronLangParserRULE_funcCallArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.mathExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.FuncCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.AnonimousFuncWithReturn()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)
			p.AnonimousFuncNoReturn()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(223)
			p.InitStruct()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(224)
			p.InitImpl()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(225)
			p.Match(IronLangParserSTRING_LITERAL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(226)
			p.Match(IronLangParserCHAR_LITERAL)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(227)
			p.Match(IronLangParserBOOLEAN_VALUE)
		}

	}

	return localctx
}

// IAnonimousFuncWithReturnContext is an interface to support dynamic dispatch.
type IAnonimousFuncWithReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnonimousFuncWithReturnContext differentiates from other interfaces.
	IsAnonimousFuncWithReturnContext()
}

type AnonimousFuncWithReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnonimousFuncWithReturnContext() *AnonimousFuncWithReturnContext {
	var p = new(AnonimousFuncWithReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_anonimousFuncWithReturn
	return p
}

func (*AnonimousFuncWithReturnContext) IsAnonimousFuncWithReturnContext() {}

func NewAnonimousFuncWithReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonimousFuncWithReturnContext {
	var p = new(AnonimousFuncWithReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_anonimousFuncWithReturn

	return p
}

func (s *AnonimousFuncWithReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonimousFuncWithReturnContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *AnonimousFuncWithReturnContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *AnonimousFuncWithReturnContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *AnonimousFuncWithReturnContext) ARROW() antlr.TerminalNode {
	return s.GetToken(IronLangParserARROW, 0)
}

func (s *AnonimousFuncWithReturnContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *AnonimousFuncWithReturnContext) ReturnDefinition() IReturnDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnDefinitionContext)
}

func (s *AnonimousFuncWithReturnContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *AnonimousFuncWithReturnContext) AllFunctionArgs() []IFunctionArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgsContext); ok {
			tst[i] = t.(IFunctionArgsContext)
			i++
		}
	}

	return tst
}

func (s *AnonimousFuncWithReturnContext) FunctionArgs(i int) IFunctionArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *AnonimousFuncWithReturnContext) AllBodyScope() []IBodyScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyScopeContext); ok {
			len++
		}
	}

	tst := make([]IBodyScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyScopeContext); ok {
			tst[i] = t.(IBodyScopeContext)
			i++
		}
	}

	return tst
}

func (s *AnonimousFuncWithReturnContext) BodyScope(i int) IBodyScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *AnonimousFuncWithReturnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *AnonimousFuncWithReturnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *AnonimousFuncWithReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonimousFuncWithReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonimousFuncWithReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterAnonimousFuncWithReturn(s)
	}
}

func (s *AnonimousFuncWithReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitAnonimousFuncWithReturn(s)
	}
}

func (s *AnonimousFuncWithReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitAnonimousFuncWithReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) AnonimousFuncWithReturn() (localctx IAnonimousFuncWithReturnContext) {
	this := p
	_ = this

	localctx = NewAnonimousFuncWithReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IronLangParserRULE_anonimousFuncWithReturn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserIDENTIFIER {
			{
				p.SetState(231)
				p.FunctionArgs()
			}
			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(232)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(233)
					p.FunctionArgs()
				}

				p.SetState(238)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(241)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(242)
			p.DataTypes()
		}
		{
			p.SetState(243)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(244)
			p.Match(IronLangParserL_CURLY)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(245)
					p.BodyScope()
				}

			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
		}
		{
			p.SetState(251)
			p.ReturnDefinition()
		}
		{
			p.SetState(252)
			p.Match(IronLangParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserIDENTIFIER {
			{
				p.SetState(255)
				p.FunctionArgs()
			}
			p.SetState(260)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(256)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(257)
					p.FunctionArgs()
				}

				p.SetState(262)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(265)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(266)
			p.DataTypes()
		}
		{
			p.SetState(267)
			p.Match(IronLangParserARROW)
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(268)
				p.BodyScope()
			}

		}
		{
			p.SetState(271)
			p.ReturnDefinition()
		}

	}

	return localctx
}

// IAnonimousFuncNoReturnContext is an interface to support dynamic dispatch.
type IAnonimousFuncNoReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnonimousFuncNoReturnContext differentiates from other interfaces.
	IsAnonimousFuncNoReturnContext()
}

type AnonimousFuncNoReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnonimousFuncNoReturnContext() *AnonimousFuncNoReturnContext {
	var p = new(AnonimousFuncNoReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_anonimousFuncNoReturn
	return p
}

func (*AnonimousFuncNoReturnContext) IsAnonimousFuncNoReturnContext() {}

func NewAnonimousFuncNoReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonimousFuncNoReturnContext {
	var p = new(AnonimousFuncNoReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_anonimousFuncNoReturn

	return p
}

func (s *AnonimousFuncNoReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonimousFuncNoReturnContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *AnonimousFuncNoReturnContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *AnonimousFuncNoReturnContext) ARROW() antlr.TerminalNode {
	return s.GetToken(IronLangParserARROW, 0)
}

func (s *AnonimousFuncNoReturnContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *AnonimousFuncNoReturnContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *AnonimousFuncNoReturnContext) AllFunctionArgs() []IFunctionArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgsContext); ok {
			tst[i] = t.(IFunctionArgsContext)
			i++
		}
	}

	return tst
}

func (s *AnonimousFuncNoReturnContext) FunctionArgs(i int) IFunctionArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *AnonimousFuncNoReturnContext) AllBodyScope() []IBodyScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyScopeContext); ok {
			len++
		}
	}

	tst := make([]IBodyScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyScopeContext); ok {
			tst[i] = t.(IBodyScopeContext)
			i++
		}
	}

	return tst
}

func (s *AnonimousFuncNoReturnContext) BodyScope(i int) IBodyScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *AnonimousFuncNoReturnContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *AnonimousFuncNoReturnContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *AnonimousFuncNoReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonimousFuncNoReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonimousFuncNoReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterAnonimousFuncNoReturn(s)
	}
}

func (s *AnonimousFuncNoReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitAnonimousFuncNoReturn(s)
	}
}

func (s *AnonimousFuncNoReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitAnonimousFuncNoReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) AnonimousFuncNoReturn() (localctx IAnonimousFuncNoReturnContext) {
	this := p
	_ = this

	localctx = NewAnonimousFuncNoReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IronLangParserRULE_anonimousFuncNoReturn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserIDENTIFIER {
			{
				p.SetState(276)
				p.FunctionArgs()
			}
			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(277)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(278)
					p.FunctionArgs()
				}

				p.SetState(283)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(286)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(287)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(288)
			p.Match(IronLangParserL_CURLY)
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323448645386241) != 0 {
			{
				p.SetState(289)
				p.BodyScope()
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(294)
			p.Match(IronLangParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserIDENTIFIER {
			{
				p.SetState(297)
				p.FunctionArgs()
			}
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(298)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(299)
					p.FunctionArgs()
				}

				p.SetState(304)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(307)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(308)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(309)
			p.BodyScope()
		}

	}

	return localctx
}

// IIfScopeContext is an interface to support dynamic dispatch.
type IIfScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfScopeContext differentiates from other interfaces.
	IsIfScopeContext()
}

type IfScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfScopeContext() *IfScopeContext {
	var p = new(IfScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_ifScope
	return p
}

func (*IfScopeContext) IsIfScopeContext() {}

func NewIfScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfScopeContext {
	var p = new(IfScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_ifScope

	return p
}

func (s *IfScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *IfScopeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *IfScopeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *IfScopeContext) AllBodyScope() []IBodyScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBodyScopeContext); ok {
			len++
		}
	}

	tst := make([]IBodyScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBodyScopeContext); ok {
			tst[i] = t.(IBodyScopeContext)
			i++
		}
	}

	return tst
}

func (s *IfScopeContext) BodyScope(i int) IBodyScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *IfScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterIfScope(s)
	}
}

func (s *IfScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitIfScope(s)
	}
}

func (s *IfScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitIfScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) IfScope() (localctx IIfScopeContext) {
	this := p
	_ = this

	localctx = NewIfScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IronLangParserRULE_ifScope)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323448645386241) != 0 {
		{
			p.SetState(313)
			p.BodyScope()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IElseExpressionContext is an interface to support dynamic dispatch.
type IElseExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseExpressionContext differentiates from other interfaces.
	IsElseExpressionContext()
}

type ElseExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseExpressionContext() *ElseExpressionContext {
	var p = new(ElseExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_elseExpression
	return p
}

func (*ElseExpressionContext) IsElseExpressionContext() {}

func NewElseExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseExpressionContext {
	var p = new(ElseExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_elseExpression

	return p
}

func (s *ElseExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(IronLangParserELSE, 0)
}

func (s *ElseExpressionContext) IfScope() IIfScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfScopeContext)
}

func (s *ElseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterElseExpression(s)
	}
}

func (s *ElseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitElseExpression(s)
	}
}

func (s *ElseExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitElseExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) ElseExpression() (localctx IElseExpressionContext) {
	this := p
	_ = this

	localctx = NewElseExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IronLangParserRULE_elseExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(IronLangParserELSE)
	}
	{
		p.SetState(322)
		p.IfScope()
	}

	return localctx
}

// IElseIfExpressionContext is an interface to support dynamic dispatch.
type IElseIfExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfExpressionContext differentiates from other interfaces.
	IsElseIfExpressionContext()
}

type ElseIfExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfExpressionContext() *ElseIfExpressionContext {
	var p = new(ElseIfExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_elseIfExpression
	return p
}

func (*ElseIfExpressionContext) IsElseIfExpressionContext() {}

func NewElseIfExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfExpressionContext {
	var p = new(ElseIfExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_elseIfExpression

	return p
}

func (s *ElseIfExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(IronLangParserELSE, 0)
}

func (s *ElseIfExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(IronLangParserIF, 0)
}

func (s *ElseIfExpressionContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *ElseIfExpressionContext) IfScope() IIfScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfScopeContext)
}

func (s *ElseIfExpressionContext) ElseIfExpression() IElseIfExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfExpressionContext)
}

func (s *ElseIfExpressionContext) ElseExpression() IElseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseExpressionContext)
}

func (s *ElseIfExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterElseIfExpression(s)
	}
}

func (s *ElseIfExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitElseIfExpression(s)
	}
}

func (s *ElseIfExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitElseIfExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) ElseIfExpression() (localctx IElseIfExpressionContext) {
	this := p
	_ = this

	localctx = NewElseIfExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IronLangParserRULE_elseIfExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(IronLangParserELSE)
	}
	{
		p.SetState(325)
		p.Match(IronLangParserIF)
	}
	{
		p.SetState(326)
		p.relExpression(0)
	}
	{
		p.SetState(327)
		p.IfScope()
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(328)
			p.ElseIfExpression()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(329)
			p.ElseExpression()
		}

	}

	return localctx
}

// IIfExpressionContext is an interface to support dynamic dispatch.
type IIfExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExpressionContext differentiates from other interfaces.
	IsIfExpressionContext()
}

type IfExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExpressionContext() *IfExpressionContext {
	var p = new(IfExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_ifExpression
	return p
}

func (*IfExpressionContext) IsIfExpressionContext() {}

func NewIfExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExpressionContext {
	var p = new(IfExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_ifExpression

	return p
}

func (s *IfExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(IronLangParserIF, 0)
}

func (s *IfExpressionContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *IfExpressionContext) IfScope() IIfScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfScopeContext)
}

func (s *IfExpressionContext) ElseExpression() IElseExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseExpressionContext)
}

func (s *IfExpressionContext) ElseIfExpression() IElseIfExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfExpressionContext)
}

func (s *IfExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterIfExpression(s)
	}
}

func (s *IfExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitIfExpression(s)
	}
}

func (s *IfExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitIfExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) IfExpression() (localctx IIfExpressionContext) {
	this := p
	_ = this

	localctx = NewIfExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, IronLangParserRULE_ifExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(IronLangParserIF)
	}
	{
		p.SetState(333)
		p.relExpression(0)
	}
	{
		p.SetState(334)
		p.IfScope()
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(335)
			p.ElseExpression()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(336)
			p.ElseIfExpression()
		}

	}

	return localctx
}

// ILoopScopeContext is an interface to support dynamic dispatch.
type ILoopScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopScopeContext differentiates from other interfaces.
	IsLoopScopeContext()
}

type LoopScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopScopeContext() *LoopScopeContext {
	var p = new(LoopScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_loopScope
	return p
}

func (*LoopScopeContext) IsLoopScopeContext() {}

func NewLoopScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopScopeContext {
	var p = new(LoopScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_loopScope

	return p
}

func (s *LoopScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopScopeContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(IronLangParserCONTINUE, 0)
}

func (s *LoopScopeContext) BREAK() antlr.TerminalNode {
	return s.GetToken(IronLangParserBREAK, 0)
}

func (s *LoopScopeContext) BodyScope() IBodyScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyScopeContext)
}

func (s *LoopScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterLoopScope(s)
	}
}

func (s *LoopScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitLoopScope(s)
	}
}

func (s *LoopScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitLoopScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) LoopScope() (localctx ILoopScopeContext) {
	this := p
	_ = this

	localctx = NewLoopScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, IronLangParserRULE_loopScope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserCONTINUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.Match(IronLangParserCONTINUE)
		}

	case IronLangParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.Match(IronLangParserBREAK)
		}

	case IronLangParserL_PAREN, IronLangParserINT_NUMBER, IronLangParserREAL_NUMBER, IronLangParserDO, IronLangParserIF, IronLangParserFOR, IronLangParserLET, IronLangParserMUT, IronLangParserWHILE, IronLangParserPRINT_LN, IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT, IronLangParserTYPE_STRING, IronLangParserTYPE_BOOLEAN, IronLangParserTYPE_DOUBLE, IronLangParserTYPE_CHAR, IronLangParserTHIS, IronLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
			p.BodyScope()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoopDoWhileContext is an interface to support dynamic dispatch.
type ILoopDoWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopDoWhileContext differentiates from other interfaces.
	IsLoopDoWhileContext()
}

type LoopDoWhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopDoWhileContext() *LoopDoWhileContext {
	var p = new(LoopDoWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_loopDoWhile
	return p
}

func (*LoopDoWhileContext) IsLoopDoWhileContext() {}

func NewLoopDoWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopDoWhileContext {
	var p = new(LoopDoWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_loopDoWhile

	return p
}

func (s *LoopDoWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopDoWhileContext) DO() antlr.TerminalNode {
	return s.GetToken(IronLangParserDO, 0)
}

func (s *LoopDoWhileContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *LoopDoWhileContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *LoopDoWhileContext) AllLoopScope() []ILoopScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILoopScopeContext); ok {
			len++
		}
	}

	tst := make([]ILoopScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILoopScopeContext); ok {
			tst[i] = t.(ILoopScopeContext)
			i++
		}
	}

	return tst
}

func (s *LoopDoWhileContext) LoopScope(i int) ILoopScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopScopeContext)
}

func (s *LoopDoWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(IronLangParserWHILE, 0)
}

func (s *LoopDoWhileContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *LoopDoWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopDoWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopDoWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterLoopDoWhile(s)
	}
}

func (s *LoopDoWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitLoopDoWhile(s)
	}
}

func (s *LoopDoWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitLoopDoWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) LoopDoWhile() (localctx ILoopDoWhileContext) {
	this := p
	_ = this

	localctx = NewLoopDoWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, IronLangParserRULE_loopDoWhile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(IronLangParserDO)
	}
	{
		p.SetState(345)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323449450692609) != 0 {
		{
			p.SetState(346)
			p.LoopScope()
		}

		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(352)
		p.Match(IronLangParserR_CURLY)
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(353)
			p.Match(IronLangParserWHILE)
		}
		{
			p.SetState(354)
			p.relExpression(0)
		}

	}

	return localctx
}

// ILoopWhileContext is an interface to support dynamic dispatch.
type ILoopWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopWhileContext differentiates from other interfaces.
	IsLoopWhileContext()
}

type LoopWhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopWhileContext() *LoopWhileContext {
	var p = new(LoopWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_loopWhile
	return p
}

func (*LoopWhileContext) IsLoopWhileContext() {}

func NewLoopWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopWhileContext {
	var p = new(LoopWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_loopWhile

	return p
}

func (s *LoopWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(IronLangParserWHILE, 0)
}

func (s *LoopWhileContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *LoopWhileContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *LoopWhileContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *LoopWhileContext) AllLoopScope() []ILoopScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILoopScopeContext); ok {
			len++
		}
	}

	tst := make([]ILoopScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILoopScopeContext); ok {
			tst[i] = t.(ILoopScopeContext)
			i++
		}
	}

	return tst
}

func (s *LoopWhileContext) LoopScope(i int) ILoopScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopScopeContext)
}

func (s *LoopWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterLoopWhile(s)
	}
}

func (s *LoopWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitLoopWhile(s)
	}
}

func (s *LoopWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitLoopWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) LoopWhile() (localctx ILoopWhileContext) {
	this := p
	_ = this

	localctx = NewLoopWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, IronLangParserRULE_loopWhile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(IronLangParserWHILE)
	}
	{
		p.SetState(358)
		p.relExpression(0)
	}
	{
		p.SetState(359)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323449450692609) != 0 {
		{
			p.SetState(360)
			p.LoopScope()
		}

		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(366)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// ILoopForInContext is an interface to support dynamic dispatch.
type ILoopForInContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopForInContext differentiates from other interfaces.
	IsLoopForInContext()
}

type LoopForInContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopForInContext() *LoopForInContext {
	var p = new(LoopForInContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_loopForIn
	return p
}

func (*LoopForInContext) IsLoopForInContext() {}

func NewLoopForInContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopForInContext {
	var p = new(LoopForInContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_loopForIn

	return p
}

func (s *LoopForInContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopForInContext) FOR() antlr.TerminalNode {
	return s.GetToken(IronLangParserFOR, 0)
}

func (s *LoopForInContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserIDENTIFIER)
}

func (s *LoopForInContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, i)
}

func (s *LoopForInContext) IN() antlr.TerminalNode {
	return s.GetToken(IronLangParserIN, 0)
}

func (s *LoopForInContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *LoopForInContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *LoopForInContext) Slice() ISliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *LoopForInContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *LoopForInContext) AllINT_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserINT_NUMBER)
}

func (s *LoopForInContext) INT_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserINT_NUMBER, i)
}

func (s *LoopForInContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *LoopForInContext) AllLoopScope() []ILoopScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILoopScopeContext); ok {
			len++
		}
	}

	tst := make([]ILoopScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILoopScopeContext); ok {
			tst[i] = t.(ILoopScopeContext)
			i++
		}
	}

	return tst
}

func (s *LoopForInContext) LoopScope(i int) ILoopScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopScopeContext)
}

func (s *LoopForInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopForInContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopForInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterLoopForIn(s)
	}
}

func (s *LoopForInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitLoopForIn(s)
	}
}

func (s *LoopForInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitLoopForIn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) LoopForIn() (localctx ILoopForInContext) {
	this := p
	_ = this

	localctx = NewLoopForInContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, IronLangParserRULE_loopForIn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(IronLangParserFOR)
	}
	{
		p.SetState(369)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(370)
		p.Match(IronLangParserIN)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(371)
			p.Slice()
		}

	case 2:
		{
			p.SetState(372)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 3:
		{
			p.SetState(373)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(374)
			p.Match(IronLangParserINT_NUMBER)
		}
		{
			p.SetState(375)
			p.Match(IronLangParserT__2)
		}
		{
			p.SetState(376)
			p.Match(IronLangParserINT_NUMBER)
		}
		{
			p.SetState(377)
			p.Match(IronLangParserR_PAREN)
		}

	}
	{
		p.SetState(380)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323449450692609) != 0 {
		{
			p.SetState(381)
			p.LoopScope()
		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(387)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// ILoopForIContext is an interface to support dynamic dispatch.
type ILoopForIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopForIContext differentiates from other interfaces.
	IsLoopForIContext()
}

type LoopForIContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopForIContext() *LoopForIContext {
	var p = new(LoopForIContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_loopForI
	return p
}

func (*LoopForIContext) IsLoopForIContext() {}

func NewLoopForIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopForIContext {
	var p = new(LoopForIContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_loopForI

	return p
}

func (s *LoopForIContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopForIContext) FOR() antlr.TerminalNode {
	return s.GetToken(IronLangParserFOR, 0)
}

func (s *LoopForIContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *LoopForIContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *LoopForIContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *LoopForIContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *LoopForIContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *LoopForIContext) AllLoopScope() []ILoopScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILoopScopeContext); ok {
			len++
		}
	}

	tst := make([]ILoopScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILoopScopeContext); ok {
			tst[i] = t.(ILoopScopeContext)
			i++
		}
	}

	return tst
}

func (s *LoopForIContext) LoopScope(i int) ILoopScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopScopeContext)
}

func (s *LoopForIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopForIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopForIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterLoopForI(s)
	}
}

func (s *LoopForIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitLoopForI(s)
	}
}

func (s *LoopForIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitLoopForI(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) LoopForI() (localctx ILoopForIContext) {
	this := p
	_ = this

	localctx = NewLoopForIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, IronLangParserRULE_loopForI)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(IronLangParserFOR)
	}
	{
		p.SetState(390)
		p.Assignment()
	}
	{
		p.SetState(391)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(392)
		p.relExpression(0)
	}
	{
		p.SetState(393)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(394)
		p.mathExpression(0)
	}
	{
		p.SetState(395)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&9323449450692609) != 0 {
		{
			p.SetState(396)
			p.LoopScope()
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(402)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IBodyScopeContext is an interface to support dynamic dispatch.
type IBodyScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyScopeContext differentiates from other interfaces.
	IsBodyScopeContext()
}

type BodyScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyScopeContext() *BodyScopeContext {
	var p = new(BodyScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_bodyScope
	return p
}

func (*BodyScopeContext) IsBodyScopeContext() {}

func NewBodyScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyScopeContext {
	var p = new(BodyScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_bodyScope

	return p
}

func (s *BodyScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyScopeContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *BodyScopeContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *BodyScopeContext) IfExpression() IIfExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExpressionContext)
}

func (s *BodyScopeContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *BodyScopeContext) Println_() IPrintlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnContext)
}

func (s *BodyScopeContext) ForEach() IForEachContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForEachContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForEachContext)
}

func (s *BodyScopeContext) LoopWhile() ILoopWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopWhileContext)
}

func (s *BodyScopeContext) LoopDoWhile() ILoopDoWhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopDoWhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopDoWhileContext)
}

func (s *BodyScopeContext) LoopForIn() ILoopForInContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopForInContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopForInContext)
}

func (s *BodyScopeContext) LoopForI() ILoopForIContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopForIContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopForIContext)
}

func (s *BodyScopeContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *BodyScopeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BodyScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterBodyScope(s)
	}
}

func (s *BodyScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitBodyScope(s)
	}
}

func (s *BodyScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitBodyScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) BodyScope() (localctx IBodyScopeContext) {
	this := p
	_ = this

	localctx = NewBodyScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, IronLangParserRULE_bodyScope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(406)
			p.IfExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(407)
			p.FuncCall()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(408)
			p.Println_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(409)
			p.ForEach()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(410)
			p.LoopWhile()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(411)
			p.LoopDoWhile()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(412)
			p.LoopForIn()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(413)
			p.LoopForI()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(414)
			p.mathExpression(0)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(415)
			p.expression(0)
		}

	}

	return localctx
}

// ITraitContext is an interface to support dynamic dispatch.
type ITraitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInterfaceName returns the interfaceName token.
	GetInterfaceName() antlr.Token

	// SetInterfaceName sets the interfaceName token.
	SetInterfaceName(antlr.Token)

	// IsTraitContext differentiates from other interfaces.
	IsTraitContext()
}

type TraitContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	interfaceName antlr.Token
}

func NewEmptyTraitContext() *TraitContext {
	var p = new(TraitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_trait
	return p
}

func (*TraitContext) IsTraitContext() {}

func NewTraitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TraitContext {
	var p = new(TraitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_trait

	return p
}

func (s *TraitContext) GetParser() antlr.Parser { return s.parser }

func (s *TraitContext) GetInterfaceName() antlr.Token { return s.interfaceName }

func (s *TraitContext) SetInterfaceName(v antlr.Token) { s.interfaceName = v }

func (s *TraitContext) TRAIT() antlr.TerminalNode {
	return s.GetToken(IronLangParserTRAIT, 0)
}

func (s *TraitContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *TraitContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *TraitContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *TraitContext) AllFuncType() []IFuncTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncTypeContext); ok {
			len++
		}
	}

	tst := make([]IFuncTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncTypeContext); ok {
			tst[i] = t.(IFuncTypeContext)
			i++
		}
	}

	return tst
}

func (s *TraitContext) FuncType(i int) IFuncTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncTypeContext)
}

func (s *TraitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TraitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TraitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterTrait(s)
	}
}

func (s *TraitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitTrait(s)
	}
}

func (s *TraitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitTrait(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Trait() (localctx ITraitContext) {
	this := p
	_ = this

	localctx = NewTraitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, IronLangParserRULE_trait)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(IronLangParserTRAIT)
	}
	{
		p.SetState(419)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*TraitContext).interfaceName = _m
	}
	{
		p.SetState(420)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == IronLangParserIDENTIFIER {
		{
			p.SetState(421)
			p.FuncType()
		}

		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(426)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IImplContext is an interface to support dynamic dispatch.
type IImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImplName returns the implName token.
	GetImplName() antlr.Token

	// GetTraitName returns the traitName token.
	GetTraitName() antlr.Token

	// SetImplName sets the implName token.
	SetImplName(antlr.Token)

	// SetTraitName sets the traitName token.
	SetTraitName(antlr.Token)

	// IsImplContext differentiates from other interfaces.
	IsImplContext()
}

type ImplContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	implName  antlr.Token
	traitName antlr.Token
}

func NewEmptyImplContext() *ImplContext {
	var p = new(ImplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_impl
	return p
}

func (*ImplContext) IsImplContext() {}

func NewImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplContext {
	var p = new(ImplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_impl

	return p
}

func (s *ImplContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplContext) GetImplName() antlr.Token { return s.implName }

func (s *ImplContext) GetTraitName() antlr.Token { return s.traitName }

func (s *ImplContext) SetImplName(v antlr.Token) { s.implName = v }

func (s *ImplContext) SetTraitName(v antlr.Token) { s.traitName = v }

func (s *ImplContext) IMPL() antlr.TerminalNode {
	return s.GetToken(IronLangParserIMPL, 0)
}

func (s *ImplContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *ImplContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *ImplContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserIDENTIFIER)
}

func (s *ImplContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, i)
}

func (s *ImplContext) FOR() antlr.TerminalNode {
	return s.GetToken(IronLangParserFOR, 0)
}

func (s *ImplContext) AllDefinitionVariables() []IDefinitionVariablesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionVariablesContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionVariablesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionVariablesContext); ok {
			tst[i] = t.(IDefinitionVariablesContext)
			i++
		}
	}

	return tst
}

func (s *ImplContext) DefinitionVariables(i int) IDefinitionVariablesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionVariablesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionVariablesContext)
}

func (s *ImplContext) AllImplConstructor() []IImplConstructorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImplConstructorContext); ok {
			len++
		}
	}

	tst := make([]IImplConstructorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImplConstructorContext); ok {
			tst[i] = t.(IImplConstructorContext)
			i++
		}
	}

	return tst
}

func (s *ImplContext) ImplConstructor(i int) IImplConstructorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplConstructorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplConstructorContext)
}

func (s *ImplContext) AllFunctionReturn() []IFunctionReturnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionReturnContext); ok {
			len++
		}
	}

	tst := make([]IFunctionReturnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionReturnContext); ok {
			tst[i] = t.(IFunctionReturnContext)
			i++
		}
	}

	return tst
}

func (s *ImplContext) FunctionReturn(i int) IFunctionReturnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionReturnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnContext)
}

func (s *ImplContext) AllFunctionNoReturn() []IFunctionNoReturnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionNoReturnContext); ok {
			len++
		}
	}

	tst := make([]IFunctionNoReturnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionNoReturnContext); ok {
			tst[i] = t.(IFunctionNoReturnContext)
			i++
		}
	}

	return tst
}

func (s *ImplContext) FunctionNoReturn(i int) IFunctionNoReturnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionNoReturnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionNoReturnContext)
}

func (s *ImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterImpl(s)
	}
}

func (s *ImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitImpl(s)
	}
}

func (s *ImplContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitImpl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Impl() (localctx IImplContext) {
	this := p
	_ = this

	localctx = NewImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, IronLangParserRULE_impl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(428)
		p.Match(IronLangParserIMPL)
	}
	{
		p.SetState(429)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*ImplContext).implName = _m
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserFOR {
		{
			p.SetState(430)
			p.Match(IronLangParserFOR)
		}
		{
			p.SetState(431)

			var _m = p.Match(IronLangParserIDENTIFIER)

			localctx.(*ImplContext).traitName = _m
		}

	}
	{
		p.SetState(434)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserIDENTIFIER {
		{
			p.SetState(435)
			p.DefinitionVariables()
		}

		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserT__1 {
		{
			p.SetState(441)
			p.ImplConstructor()
		}

		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserFN {
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(447)
				p.FunctionReturn()
			}

		case 2:
			{
				p.SetState(448)
				p.FunctionNoReturn()
			}

		}

		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(454)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IInitImplContext is an interface to support dynamic dispatch.
type IInitImplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImplName returns the implName token.
	GetImplName() antlr.Token

	// SetImplName sets the implName token.
	SetImplName(antlr.Token)

	// IsInitImplContext differentiates from other interfaces.
	IsInitImplContext()
}

type InitImplContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	implName antlr.Token
}

func NewEmptyInitImplContext() *InitImplContext {
	var p = new(InitImplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_initImpl
	return p
}

func (*InitImplContext) IsInitImplContext() {}

func NewInitImplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitImplContext {
	var p = new(InitImplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_initImpl

	return p
}

func (s *InitImplContext) GetParser() antlr.Parser { return s.parser }

func (s *InitImplContext) GetImplName() antlr.Token { return s.implName }

func (s *InitImplContext) SetImplName(v antlr.Token) { s.implName = v }

func (s *InitImplContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *InitImplContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *InitImplContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *InitImplContext) AllFuncCallArg() []IFuncCallArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncCallArgContext); ok {
			len++
		}
	}

	tst := make([]IFuncCallArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncCallArgContext); ok {
			tst[i] = t.(IFuncCallArgContext)
			i++
		}
	}

	return tst
}

func (s *InitImplContext) FuncCallArg(i int) IFuncCallArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallArgContext)
}

func (s *InitImplContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *InitImplContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *InitImplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitImplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitImplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterInitImpl(s)
	}
}

func (s *InitImplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitInitImpl(s)
	}
}

func (s *InitImplContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitInitImpl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) InitImpl() (localctx IInitImplContext) {
	this := p
	_ = this

	localctx = NewInitImplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, IronLangParserRULE_initImpl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*InitImplContext).implName = _m
	}
	{
		p.SetState(457)
		p.Match(IronLangParserT__3)
	}
	{
		p.SetState(458)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-13)) & ^0x3f) == 0 && ((int64(1)<<(_la-13))&15797783068082177) != 0 {
		{
			p.SetState(459)
			p.FuncCallArg()
		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(460)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(461)
				p.FuncCallArg()
			}

			p.SetState(466)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(469)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IStructDefinitionContext is an interface to support dynamic dispatch.
type IStructDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// IsStructDefinitionContext differentiates from other interfaces.
	IsStructDefinitionContext()
}

type StructDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	structName antlr.Token
}

func NewEmptyStructDefinitionContext() *StructDefinitionContext {
	var p = new(StructDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_structDefinition
	return p
}

func (*StructDefinitionContext) IsStructDefinitionContext() {}

func NewStructDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefinitionContext {
	var p = new(StructDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_structDefinition

	return p
}

func (s *StructDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefinitionContext) GetStructName() antlr.Token { return s.structName }

func (s *StructDefinitionContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *StructDefinitionContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(IronLangParserSTRUCT, 0)
}

func (s *StructDefinitionContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *StructDefinitionContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *StructDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *StructDefinitionContext) AllDefinitionVariables() []IDefinitionVariablesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionVariablesContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionVariablesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionVariablesContext); ok {
			tst[i] = t.(IDefinitionVariablesContext)
			i++
		}
	}

	return tst
}

func (s *StructDefinitionContext) DefinitionVariables(i int) IDefinitionVariablesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionVariablesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionVariablesContext)
}

func (s *StructDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterStructDefinition(s)
	}
}

func (s *StructDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitStructDefinition(s)
	}
}

func (s *StructDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitStructDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) StructDefinition() (localctx IStructDefinitionContext) {
	this := p
	_ = this

	localctx = NewStructDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, IronLangParserRULE_structDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Match(IronLangParserSTRUCT)
	}
	{
		p.SetState(472)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*StructDefinitionContext).structName = _m
	}
	{
		p.SetState(473)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == IronLangParserIDENTIFIER {
		{
			p.SetState(474)
			p.DefinitionVariables()
		}

		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(479)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IDefinitionVariablesContext is an interface to support dynamic dispatch.
type IDefinitionVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariableName returns the variableName token.
	GetVariableName() antlr.Token

	// GetGenericType returns the genericType token.
	GetGenericType() antlr.Token

	// SetVariableName sets the variableName token.
	SetVariableName(antlr.Token)

	// SetGenericType sets the genericType token.
	SetGenericType(antlr.Token)

	// IsDefinitionVariablesContext differentiates from other interfaces.
	IsDefinitionVariablesContext()
}

type DefinitionVariablesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	variableName antlr.Token
	genericType  antlr.Token
}

func NewEmptyDefinitionVariablesContext() *DefinitionVariablesContext {
	var p = new(DefinitionVariablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_definitionVariables
	return p
}

func (*DefinitionVariablesContext) IsDefinitionVariablesContext() {}

func NewDefinitionVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionVariablesContext {
	var p = new(DefinitionVariablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_definitionVariables

	return p
}

func (s *DefinitionVariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionVariablesContext) GetVariableName() antlr.Token { return s.variableName }

func (s *DefinitionVariablesContext) GetGenericType() antlr.Token { return s.genericType }

func (s *DefinitionVariablesContext) SetVariableName(v antlr.Token) { s.variableName = v }

func (s *DefinitionVariablesContext) SetGenericType(v antlr.Token) { s.genericType = v }

func (s *DefinitionVariablesContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserIDENTIFIER)
}

func (s *DefinitionVariablesContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, i)
}

func (s *DefinitionVariablesContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *DefinitionVariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionVariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionVariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterDefinitionVariables(s)
	}
}

func (s *DefinitionVariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitDefinitionVariables(s)
	}
}

func (s *DefinitionVariablesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitDefinitionVariables(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) DefinitionVariables() (localctx IDefinitionVariablesContext) {
	this := p
	_ = this

	localctx = NewDefinitionVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, IronLangParserRULE_definitionVariables)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*DefinitionVariablesContext).variableName = _m
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT, IronLangParserTYPE_STRING, IronLangParserTYPE_BOOLEAN, IronLangParserTYPE_DOUBLE, IronLangParserTYPE_CHAR:
		{
			p.SetState(482)
			p.DataTypes()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(483)

			var _m = p.Match(IronLangParserIDENTIFIER)

			localctx.(*DefinitionVariablesContext).genericType = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInitStructContext is an interface to support dynamic dispatch.
type IInitStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// IsInitStructContext differentiates from other interfaces.
	IsInitStructContext()
}

type InitStructContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	structName antlr.Token
}

func NewEmptyInitStructContext() *InitStructContext {
	var p = new(InitStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_initStruct
	return p
}

func (*InitStructContext) IsInitStructContext() {}

func NewInitStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitStructContext {
	var p = new(InitStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_initStruct

	return p
}

func (s *InitStructContext) GetParser() antlr.Parser { return s.parser }

func (s *InitStructContext) GetStructName() antlr.Token { return s.structName }

func (s *InitStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *InitStructContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *InitStructContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *InitStructContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *InitStructContext) AllInitStructBody() []IInitStructBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInitStructBodyContext); ok {
			len++
		}
	}

	tst := make([]IInitStructBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInitStructBodyContext); ok {
			tst[i] = t.(IInitStructBodyContext)
			i++
		}
	}

	return tst
}

func (s *InitStructContext) InitStructBody(i int) IInitStructBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitStructBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitStructBodyContext)
}

func (s *InitStructContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *InitStructContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *InitStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterInitStruct(s)
	}
}

func (s *InitStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitInitStruct(s)
	}
}

func (s *InitStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitInitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) InitStruct() (localctx IInitStructContext) {
	this := p
	_ = this

	localctx = NewInitStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, IronLangParserRULE_initStruct)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*InitStructContext).structName = _m
	}
	{
		p.SetState(487)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserIDENTIFIER {
		{
			p.SetState(488)
			p.InitStructBody()
		}
		p.SetState(493)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(489)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(490)
				p.InitStructBody()
			}

			p.SetState(495)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(498)
		p.Match(IronLangParserR_CURLY)
	}

	return localctx
}

// IInitStructBodyContext is an interface to support dynamic dispatch.
type IInitStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructKey returns the structKey token.
	GetStructKey() antlr.Token

	// GetAsIdent returns the asIdent token.
	GetAsIdent() antlr.Token

	// SetStructKey sets the structKey token.
	SetStructKey(antlr.Token)

	// SetAsIdent sets the asIdent token.
	SetAsIdent(antlr.Token)

	// IsInitStructBodyContext differentiates from other interfaces.
	IsInitStructBodyContext()
}

type InitStructBodyContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	structKey antlr.Token
	asIdent   antlr.Token
}

func NewEmptyInitStructBodyContext() *InitStructBodyContext {
	var p = new(InitStructBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_initStructBody
	return p
}

func (*InitStructBodyContext) IsInitStructBodyContext() {}

func NewInitStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitStructBodyContext {
	var p = new(InitStructBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_initStructBody

	return p
}

func (s *InitStructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *InitStructBodyContext) GetStructKey() antlr.Token { return s.structKey }

func (s *InitStructBodyContext) GetAsIdent() antlr.Token { return s.asIdent }

func (s *InitStructBodyContext) SetStructKey(v antlr.Token) { s.structKey = v }

func (s *InitStructBodyContext) SetAsIdent(v antlr.Token) { s.asIdent = v }

func (s *InitStructBodyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserIDENTIFIER)
}

func (s *InitStructBodyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, i)
}

func (s *InitStructBodyContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *InitStructBodyContext) BOOLEAN_VALUE() antlr.TerminalNode {
	return s.GetToken(IronLangParserBOOLEAN_VALUE, 0)
}

func (s *InitStructBodyContext) INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(IronLangParserINT_NUMBER, 0)
}

func (s *InitStructBodyContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(IronLangParserREAL_NUMBER, 0)
}

func (s *InitStructBodyContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(IronLangParserSTRING_LITERAL, 0)
}

func (s *InitStructBodyContext) InitStruct() IInitStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitStructContext)
}

func (s *InitStructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitStructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitStructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterInitStructBody(s)
	}
}

func (s *InitStructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitInitStructBody(s)
	}
}

func (s *InitStructBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitInitStructBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) InitStructBody() (localctx IInitStructBodyContext) {
	this := p
	_ = this

	localctx = NewInitStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, IronLangParserRULE_initStructBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*InitStructBodyContext).structKey = _m
	}
	{
		p.SetState(501)
		p.Match(IronLangParserT__4)
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(502)
			p.FuncCall()
		}

	case 2:
		{
			p.SetState(503)
			p.Match(IronLangParserBOOLEAN_VALUE)
		}

	case 3:
		{
			p.SetState(504)
			p.Match(IronLangParserINT_NUMBER)
		}

	case 4:
		{
			p.SetState(505)
			p.Match(IronLangParserREAL_NUMBER)
		}

	case 5:
		{
			p.SetState(506)
			p.Match(IronLangParserSTRING_LITERAL)
		}

	case 6:
		{
			p.SetState(507)

			var _m = p.Match(IronLangParserIDENTIFIER)

			localctx.(*InitStructBodyContext).asIdent = _m
		}

	case 7:
		{
			p.SetState(508)
			p.InitStruct()
		}

	}

	return localctx
}

// IPrintlnContext is an interface to support dynamic dispatch.
type IPrintlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintlnContext differentiates from other interfaces.
	IsPrintlnContext()
}

type PrintlnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnContext() *PrintlnContext {
	var p = new(PrintlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_println
	return p
}

func (*PrintlnContext) IsPrintlnContext() {}

func NewPrintlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnContext {
	var p = new(PrintlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_println

	return p
}

func (s *PrintlnContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnContext) PRINT_LN() antlr.TerminalNode {
	return s.GetToken(IronLangParserPRINT_LN, 0)
}

func (s *PrintlnContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *PrintlnContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *PrintlnContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PrintlnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *PrintlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterPrintln(s)
	}
}

func (s *PrintlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitPrintln(s)
	}
}

func (s *PrintlnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitPrintln(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Println_() (localctx IPrintlnContext) {
	this := p
	_ = this

	localctx = NewPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, IronLangParserRULE_println)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
		p.Match(IronLangParserPRINT_LN)
	}
	{
		p.SetState(512)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(513)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(514)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(517)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariableName returns the variableName token.
	GetVariableName() antlr.Token

	// SetVariableName sets the variableName token.
	SetVariableName(antlr.Token)

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	variableName antlr.Token
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) GetVariableName() antlr.Token { return s.variableName }

func (s *VariableContext) SetVariableName(v antlr.Token) { s.variableName = v }

func (s *VariableContext) LET() antlr.TerminalNode {
	return s.GetToken(IronLangParserLET, 0)
}

func (s *VariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *VariableContext) MUT() antlr.TerminalNode {
	return s.GetToken(IronLangParserMUT, 0)
}

func (s *VariableContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, IronLangParserRULE_variable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(519)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(522)
		p.Match(IronLangParserLET)
	}
	{
		p.SetState(523)

		var _m = p.Match(IronLangParserIDENTIFIER)

		localctx.(*VariableContext).variableName = _m
	}
	p.SetState(525)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(524)
			p.DataTypes()
		}

	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgName returns the argName token.
	GetArgName() antlr.Token

	// GetIdentType returns the identType token.
	GetIdentType() antlr.Token

	// SetArgName sets the argName token.
	SetArgName(antlr.Token)

	// SetIdentType sets the identType token.
	SetIdentType(antlr.Token)

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	argName   antlr.Token
	identType antlr.Token
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) GetArgName() antlr.Token { return s.argName }

func (s *FunctionArgsContext) GetIdentType() antlr.Token { return s.identType }

func (s *FunctionArgsContext) SetArgName(v antlr.Token) { s.argName = v }

func (s *FunctionArgsContext) SetIdentType(v antlr.Token) { s.identType = v }

func (s *FunctionArgsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserIDENTIFIER)
}

func (s *FunctionArgsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, i)
}

func (s *FunctionArgsContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *FunctionArgsContext) FuncType() IFuncTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncTypeContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FunctionArgs() (localctx IFunctionArgsContext) {
	this := p
	_ = this

	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, IronLangParserRULE_functionArgs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(527)

			var _m = p.Match(IronLangParserIDENTIFIER)

			localctx.(*FunctionArgsContext).argName = _m
		}
		p.SetState(530)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT, IronLangParserTYPE_STRING, IronLangParserTYPE_BOOLEAN, IronLangParserTYPE_DOUBLE, IronLangParserTYPE_CHAR:
			{
				p.SetState(528)
				p.DataTypes()
			}

		case IronLangParserIDENTIFIER:
			{
				p.SetState(529)

				var _m = p.Match(IronLangParserIDENTIFIER)

				localctx.(*FunctionArgsContext).identType = _m
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.FuncType()
		}

	}

	return localctx
}

// IDataTypesContext is an interface to support dynamic dispatch.
type IDataTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataTypesContext differentiates from other interfaces.
	IsDataTypesContext()
}

type DataTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypesContext() *DataTypesContext {
	var p = new(DataTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_dataTypes
	return p
}

func (*DataTypesContext) IsDataTypesContext() {}

func NewDataTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypesContext {
	var p = new(DataTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_dataTypes

	return p
}

func (s *DataTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypesContext) TYPE_INT() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_INT, 0)
}

func (s *DataTypesContext) TYPE_FLOAT() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_FLOAT, 0)
}

func (s *DataTypesContext) TYPE_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_BOOLEAN, 0)
}

func (s *DataTypesContext) TYPE_STRING() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_STRING, 0)
}

func (s *DataTypesContext) TYPE_DOUBLE() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_DOUBLE, 0)
}

func (s *DataTypesContext) TYPE_CHAR() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_CHAR, 0)
}

func (s *DataTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterDataTypes(s)
	}
}

func (s *DataTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitDataTypes(s)
	}
}

func (s *DataTypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitDataTypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) DataTypes() (localctx IDataTypesContext) {
	this := p
	_ = this

	localctx = NewDataTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, IronLangParserRULE_dataTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&283726776524341248) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) LeftAssignment() ILeftAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftAssignmentContext)
}

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(IronLangParserEQ, 0)
}

func (s *AssignmentContext) RigthAssignment() IRigthAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRigthAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRigthAssignmentContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, IronLangParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(537)
		p.LeftAssignment()
	}
	{
		p.SetState(538)
		p.Match(IronLangParserEQ)
	}
	{
		p.SetState(539)
		p.RigthAssignment()
	}

	return localctx
}

// ILeftAssignmentContext is an interface to support dynamic dispatch.
type ILeftAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariableName returns the variableName token.
	GetVariableName() antlr.Token

	// SetVariableName sets the variableName token.
	SetVariableName(antlr.Token)

	// GetLeftExpression returns the leftExpression rule contexts.
	GetLeftExpression() IExpressionContext

	// SetLeftExpression sets the leftExpression rule contexts.
	SetLeftExpression(IExpressionContext)

	// IsLeftAssignmentContext differentiates from other interfaces.
	IsLeftAssignmentContext()
}

type LeftAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	leftExpression IExpressionContext
	variableName   antlr.Token
}

func NewEmptyLeftAssignmentContext() *LeftAssignmentContext {
	var p = new(LeftAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_leftAssignment
	return p
}

func (*LeftAssignmentContext) IsLeftAssignmentContext() {}

func NewLeftAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftAssignmentContext {
	var p = new(LeftAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_leftAssignment

	return p
}

func (s *LeftAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftAssignmentContext) GetVariableName() antlr.Token { return s.variableName }

func (s *LeftAssignmentContext) SetVariableName(v antlr.Token) { s.variableName = v }

func (s *LeftAssignmentContext) GetLeftExpression() IExpressionContext { return s.leftExpression }

func (s *LeftAssignmentContext) SetLeftExpression(v IExpressionContext) { s.leftExpression = v }

func (s *LeftAssignmentContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *LeftAssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LeftAssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *LeftAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterLeftAssignment(s)
	}
}

func (s *LeftAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitLeftAssignment(s)
	}
}

func (s *LeftAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitLeftAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) LeftAssignment() (localctx ILeftAssignmentContext) {
	this := p
	_ = this

	localctx = NewLeftAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, IronLangParserRULE_leftAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(542)

			var _x = p.expression(0)

			localctx.(*LeftAssignmentContext).leftExpression = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(543)

			var _m = p.Match(IronLangParserIDENTIFIER)

			localctx.(*LeftAssignmentContext).variableName = _m
		}

	}

	return localctx
}

// IRigthAssignmentContext is an interface to support dynamic dispatch.
type IRigthAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentType returns the identType token.
	GetIdentType() antlr.Token

	// SetIdentType sets the identType token.
	SetIdentType(antlr.Token)

	// GetRightExpression returns the rightExpression rule contexts.
	GetRightExpression() IExpressionContext

	// SetRightExpression sets the rightExpression rule contexts.
	SetRightExpression(IExpressionContext)

	// IsRigthAssignmentContext differentiates from other interfaces.
	IsRigthAssignmentContext()
}

type RigthAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	identType       antlr.Token
	rightExpression IExpressionContext
}

func NewEmptyRigthAssignmentContext() *RigthAssignmentContext {
	var p = new(RigthAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_rigthAssignment
	return p
}

func (*RigthAssignmentContext) IsRigthAssignmentContext() {}

func NewRigthAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RigthAssignmentContext {
	var p = new(RigthAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_rigthAssignment

	return p
}

func (s *RigthAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *RigthAssignmentContext) GetIdentType() antlr.Token { return s.identType }

func (s *RigthAssignmentContext) SetIdentType(v antlr.Token) { s.identType = v }

func (s *RigthAssignmentContext) GetRightExpression() IExpressionContext { return s.rightExpression }

func (s *RigthAssignmentContext) SetRightExpression(v IExpressionContext) { s.rightExpression = v }

func (s *RigthAssignmentContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(IronLangParserSTRING_LITERAL, 0)
}

func (s *RigthAssignmentContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(IronLangParserCHAR_LITERAL, 0)
}

func (s *RigthAssignmentContext) BOOLEAN_VALUE() antlr.TerminalNode {
	return s.GetToken(IronLangParserBOOLEAN_VALUE, 0)
}

func (s *RigthAssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *RigthAssignmentContext) Slice() ISliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *RigthAssignmentContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *RigthAssignmentContext) AnonimousFuncWithReturn() IAnonimousFuncWithReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncWithReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncWithReturnContext)
}

func (s *RigthAssignmentContext) AnonimousFuncNoReturn() IAnonimousFuncNoReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncNoReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncNoReturnContext)
}

func (s *RigthAssignmentContext) MathExpression() IMathExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *RigthAssignmentContext) RelExpression() IRelExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *RigthAssignmentContext) InitStruct() IInitStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitStructContext)
}

func (s *RigthAssignmentContext) InitImpl() IInitImplContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitImplContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitImplContext)
}

func (s *RigthAssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RigthAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RigthAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RigthAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterRigthAssignment(s)
	}
}

func (s *RigthAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitRigthAssignment(s)
	}
}

func (s *RigthAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitRigthAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) RigthAssignment() (localctx IRigthAssignmentContext) {
	this := p
	_ = this

	localctx = NewRigthAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, IronLangParserRULE_rigthAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(559)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(546)
			p.Match(IronLangParserSTRING_LITERAL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(547)
			p.Match(IronLangParserCHAR_LITERAL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(548)
			p.Match(IronLangParserBOOLEAN_VALUE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(549)

			var _m = p.Match(IronLangParserIDENTIFIER)

			localctx.(*RigthAssignmentContext).identType = _m
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(550)
			p.Slice()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(551)
			p.Array()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(552)
			p.AnonimousFuncWithReturn()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(553)
			p.AnonimousFuncNoReturn()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(554)
			p.mathExpression(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(555)
			p.relExpression(0)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(556)
			p.InitStruct()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(557)
			p.InitImpl()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(558)

			var _x = p.expression(0)

			localctx.(*RigthAssignmentContext).rightExpression = _x
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) DataTypes() IDataTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypesContext)
}

func (s *ArrayContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_BRACKET, 0)
}

func (s *ArrayContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_BRACKET, 0)
}

func (s *ArrayContext) AllINT_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserINT_NUMBER)
}

func (s *ArrayContext) INT_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserINT_NUMBER, i)
}

func (s *ArrayContext) AllREAL_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserREAL_NUMBER)
}

func (s *ArrayContext) REAL_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserREAL_NUMBER, i)
}

func (s *ArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *ArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, IronLangParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		p.DataTypes()
	}
	{
		p.SetState(562)
		p.Match(IronLangParserL_BRACKET)
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER {
		{
			p.SetState(563)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(568)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(564)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(565)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(570)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(576)
		p.Match(IronLangParserR_BRACKET)
	}

	return localctx
}

// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_slice
	return p
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *SliceContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_BRACKET, 0)
}

func (s *SliceContext) AllINT_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserINT_NUMBER)
}

func (s *SliceContext) INT_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserINT_NUMBER, i)
}

func (s *SliceContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_BRACKET, 0)
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitSlice(s)
	}
}

func (s *SliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Slice() (localctx ISliceContext) {
	this := p
	_ = this

	localctx = NewSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, IronLangParserRULE_slice)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(578)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(579)
		p.Match(IronLangParserL_BRACKET)
	}
	{
		p.SetState(580)
		p.Match(IronLangParserINT_NUMBER)
	}
	{
		p.SetState(581)
		p.Match(IronLangParserT__4)
	}
	{
		p.SetState(582)
		p.Match(IronLangParserINT_NUMBER)
	}
	{
		p.SetState(583)
		p.Match(IronLangParserR_BRACKET)
	}

	return localctx
}

// IForEachContext is an interface to support dynamic dispatch.
type IForEachContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForEachContext differentiates from other interfaces.
	IsForEachContext()
}

type ForEachContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForEachContext() *ForEachContext {
	var p = new(ForEachContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_forEach
	return p
}

func (*ForEachContext) IsForEachContext() {}

func NewForEachContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForEachContext {
	var p = new(ForEachContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_forEach

	return p
}

func (s *ForEachContext) GetParser() antlr.Parser { return s.parser }

func (s *ForEachContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *ForEachContext) DOT() antlr.TerminalNode {
	return s.GetToken(IronLangParserDOT, 0)
}

func (s *ForEachContext) FOR_EACH() antlr.TerminalNode {
	return s.GetToken(IronLangParserFOR_EACH, 0)
}

func (s *ForEachContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *ForEachContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *ForEachContext) AnonimousFuncNoReturn() IAnonimousFuncNoReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncNoReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncNoReturnContext)
}

func (s *ForEachContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ForEachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForEachContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForEachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterForEach(s)
	}
}

func (s *ForEachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitForEach(s)
	}
}

func (s *ForEachContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitForEach(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) ForEach() (localctx IForEachContext) {
	this := p
	_ = this

	localctx = NewForEachContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, IronLangParserRULE_forEach)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(602)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(585)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(586)
			p.Match(IronLangParserDOT)
		}
		{
			p.SetState(587)
			p.Match(IronLangParserFOR_EACH)
		}
		{
			p.SetState(588)
			p.Match(IronLangParserL_PAREN)
		}

		{
			p.SetState(589)
			p.AnonimousFuncNoReturn()
		}

		{
			p.SetState(590)
			p.Match(IronLangParserR_PAREN)
		}

	case IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT, IronLangParserTYPE_STRING, IronLangParserTYPE_BOOLEAN, IronLangParserTYPE_DOUBLE, IronLangParserTYPE_CHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(592)
			p.Array()
		}
		{
			p.SetState(593)
			p.Match(IronLangParserDOT)
		}
		{
			p.SetState(594)
			p.Match(IronLangParserFOR_EACH)
		}
		{
			p.SetState(595)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(598)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IronLangParserL_PAREN:
			{
				p.SetState(596)
				p.AnonimousFuncNoReturn()
			}

		case IronLangParserIDENTIFIER:
			{
				p.SetState(597)
				p.Match(IronLangParserIDENTIFIER)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(600)
			p.Match(IronLangParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMapContext is an interface to support dynamic dispatch.
type IMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapContext differentiates from other interfaces.
	IsMapContext()
}

type MapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapContext() *MapContext {
	var p = new(MapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_map
	return p
}

func (*MapContext) IsMapContext() {}

func NewMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapContext {
	var p = new(MapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_map

	return p
}

func (s *MapContext) GetParser() antlr.Parser { return s.parser }

func (s *MapContext) MAP() antlr.TerminalNode {
	return s.GetToken(IronLangParserMAP, 0)
}

func (s *MapContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *MapContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *MapContext) AnonimousFuncWithReturn() IAnonimousFuncWithReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncWithReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncWithReturnContext)
}

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitMap(s)
	}
}

func (s *MapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitMap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Map_() (localctx IMapContext) {
	this := p
	_ = this

	localctx = NewMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, IronLangParserRULE_map)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(604)
		p.Match(IronLangParserMAP)
	}
	{
		p.SetState(605)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(606)
		p.AnonimousFuncWithReturn()
	}

	{
		p.SetState(607)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(IronLangParserFILTER, 0)
}

func (s *FilterContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FilterContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FilterContext) AnonimousFuncWithReturn() IAnonimousFuncWithReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncWithReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncWithReturnContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (s *FilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Filter() (localctx IFilterContext) {
	this := p
	_ = this

	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, IronLangParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
		p.Match(IronLangParserFILTER)
	}
	{
		p.SetState(610)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(611)
		p.AnonimousFuncWithReturn()
	}

	{
		p.SetState(612)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IReduceContext is an interface to support dynamic dispatch.
type IReduceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReduceContext differentiates from other interfaces.
	IsReduceContext()
}

type ReduceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReduceContext() *ReduceContext {
	var p = new(ReduceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_reduce
	return p
}

func (*ReduceContext) IsReduceContext() {}

func NewReduceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReduceContext {
	var p = new(ReduceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_reduce

	return p
}

func (s *ReduceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReduceContext) REDUCE() antlr.TerminalNode {
	return s.GetToken(IronLangParserREDUCE, 0)
}

func (s *ReduceContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *ReduceContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *ReduceContext) AnonimousFuncWithReturn() IAnonimousFuncWithReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncWithReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncWithReturnContext)
}

func (s *ReduceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReduceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReduceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterReduce(s)
	}
}

func (s *ReduceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitReduce(s)
	}
}

func (s *ReduceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitReduce(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Reduce() (localctx IReduceContext) {
	this := p
	_ = this

	localctx = NewReduceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, IronLangParserRULE_reduce)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(614)
		p.Match(IronLangParserREDUCE)
	}
	{
		p.SetState(615)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(616)
		p.AnonimousFuncWithReturn()
	}

	{
		p.SetState(617)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IRelExpressionContext is an interface to support dynamic dispatch.
type IRelExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelExpressionContext differentiates from other interfaces.
	IsRelExpressionContext()
}

type RelExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelExpressionContext() *RelExpressionContext {
	var p = new(RelExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_relExpression
	return p
}

func (*RelExpressionContext) IsRelExpressionContext() {}

func NewRelExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelExpressionContext {
	var p = new(RelExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_relExpression

	return p
}

func (s *RelExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *RelExpressionContext) AllRelExpression() []IRelExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelExpressionContext); ok {
			len++
		}
	}

	tst := make([]IRelExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelExpressionContext); ok {
			tst[i] = t.(IRelExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelExpressionContext) RelExpression(i int) IRelExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *RelExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *RelExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(IronLangParserNOT, 0)
}

func (s *RelExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *RelExpressionContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *RelExpressionContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *RelExpressionContext) BOOLEAN_VALUE() antlr.TerminalNode {
	return s.GetToken(IronLangParserBOOLEAN_VALUE, 0)
}

func (s *RelExpressionContext) EQEQ() antlr.TerminalNode {
	return s.GetToken(IronLangParserEQEQ, 0)
}

func (s *RelExpressionContext) DIF() antlr.TerminalNode {
	return s.GetToken(IronLangParserDIF, 0)
}

func (s *RelExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(IronLangParserLT, 0)
}

func (s *RelExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(IronLangParserGT, 0)
}

func (s *RelExpressionContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(IronLangParserLTEQ, 0)
}

func (s *RelExpressionContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(IronLangParserGTEQ, 0)
}

func (s *RelExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(IronLangParserAND, 0)
}

func (s *RelExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(IronLangParserOR, 0)
}

func (s *RelExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterRelExpression(s)
	}
}

func (s *RelExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitRelExpression(s)
	}
}

func (s *RelExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitRelExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) RelExpression() (localctx IRelExpressionContext) {
	return p.relExpression(0)
}

func (p *IronLangParser) relExpression(_p int) (localctx IRelExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, IronLangParserRULE_relExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(630)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(620)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(621)
			p.relExpression(0)
		}
		{
			p.SetState(622)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		{
			p.SetState(624)
			p.Match(IronLangParserNOT)
		}
		{
			p.SetState(625)
			p.relExpression(5)
		}

	case 3:
		{
			p.SetState(626)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 4:
		{
			p.SetState(627)
			p.Atom()
		}

	case 5:
		{
			p.SetState(628)
			p.FuncCall()
		}

	case 6:
		{
			p.SetState(629)
			p.Match(IronLangParserBOOLEAN_VALUE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 70, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_relExpression)
			p.SetState(632)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
			}
			{
				p.SetState(633)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&292322017280) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(634)
				p.relExpression(8)
			}

		}
		p.SetState(639)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 70, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeftExpression returns the leftExpression rule contexts.
	GetLeftExpression() IExpressionContext

	// GetRigthExpression returns the rigthExpression rule contexts.
	GetRigthExpression() IExpressionContext

	// SetLeftExpression sets the leftExpression rule contexts.
	SetLeftExpression(IExpressionContext)

	// SetRigthExpression sets the rigthExpression rule contexts.
	SetRigthExpression(IExpressionContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	leftExpression  IExpressionContext
	rigthExpression IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetLeftExpression() IExpressionContext { return s.leftExpression }

func (s *ExpressionContext) GetRigthExpression() IExpressionContext { return s.rigthExpression }

func (s *ExpressionContext) SetLeftExpression(v IExpressionContext) { s.leftExpression = v }

func (s *ExpressionContext) SetRigthExpression(v IExpressionContext) { s.rigthExpression = v }

func (s *ExpressionContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *ExpressionContext) THIS() antlr.TerminalNode {
	return s.GetToken(IronLangParserTHIS, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(IronLangParserDOT, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *IronLangParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 84
	p.EnterRecursionRule(localctx, 84, IronLangParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(644)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(641)
			p.FuncCall()
		}

	case 2:
		{
			p.SetState(642)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 3:
		{
			p.SetState(643)
			p.Match(IronLangParserTHIS)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(651)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).leftExpression = _prevctx
			p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_expression)
			p.SetState(646)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(647)
				p.Match(IronLangParserDOT)
			}
			{
				p.SetState(648)

				var _x = p.expression(5)

				localctx.(*ExpressionContext).rigthExpression = _x
			}

		}
		p.SetState(653)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())
	}

	return localctx
}

// IMathExpressionContext is an interface to support dynamic dispatch.
type IMathExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathExpressionContext differentiates from other interfaces.
	IsMathExpressionContext()
}

type MathExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpressionContext() *MathExpressionContext {
	var p = new(MathExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_mathExpression
	return p
}

func (*MathExpressionContext) IsMathExpressionContext() {}

func NewMathExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpressionContext {
	var p = new(MathExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_mathExpression

	return p
}

func (s *MathExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *MathExpressionContext) AllMathExpression() []IMathExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMathExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMathExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMathExpressionContext); ok {
			tst[i] = t.(IMathExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MathExpressionContext) MathExpression(i int) IMathExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *MathExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *MathExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *MathExpressionContext) PLUS_PLUS() antlr.TerminalNode {
	return s.GetToken(IronLangParserPLUS_PLUS, 0)
}

func (s *MathExpressionContext) MINUS_MINUS() antlr.TerminalNode {
	return s.GetToken(IronLangParserMINUS_MINUS, 0)
}

func (s *MathExpressionContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *MathExpressionContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *MathExpressionContext) MULT() antlr.TerminalNode {
	return s.GetToken(IronLangParserMULT, 0)
}

func (s *MathExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(IronLangParserDIV, 0)
}

func (s *MathExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(IronLangParserPLUS, 0)
}

func (s *MathExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(IronLangParserMINUS, 0)
}

func (s *MathExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterMathExpression(s)
	}
}

func (s *MathExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitMathExpression(s)
	}
}

func (s *MathExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitMathExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) MathExpression() (localctx IMathExpressionContext) {
	return p.mathExpression(0)
}

func (p *IronLangParser) mathExpression(_p int) (localctx IMathExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 86
	p.EnterRecursionRule(localctx, 86, IronLangParserRULE_mathExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(665)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(655)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(656)
			p.mathExpression(0)
		}
		{
			p.SetState(657)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		{
			p.SetState(659)
			p.Match(IronLangParserIDENTIFIER)
		}
		p.SetState(661)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(660)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IronLangParserPLUS_PLUS || _la == IronLangParserMINUS_MINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case 3:
		{
			p.SetState(663)
			p.Atom()
		}

	case 4:
		{
			p.SetState(664)
			p.FuncCall()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(675)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(673)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 75, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(667)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(668)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserMULT || _la == IronLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(669)
					p.mathExpression(7)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(670)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(671)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(672)
					p.mathExpression(6)
				}

			}

		}
		p.SetState(677)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) INT_NUMBER() antlr.TerminalNode {
	return s.GetToken(IronLangParserINT_NUMBER, 0)
}

func (s *AtomContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(IronLangParserREAL_NUMBER, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, IronLangParserRULE_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(678)
		_la = p.GetTokenStream().LA(1)

		if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *IronLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 41:
		var t *RelExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelExpressionContext)
		}
		return p.RelExpression_Sempred(t, predIndex)

	case 42:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 43:
		var t *MathExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MathExpressionContext)
		}
		return p.MathExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *IronLangParser) RelExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *IronLangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *IronLangParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
