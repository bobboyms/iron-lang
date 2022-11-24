// Code generated from /Users/thiagoluizrodrigues/go/src/iron-lang/IronLang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package ironlang // IronLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		"", "'main'", "'..'", "';'", "':'", "'.'", "','", "'+'", "'-'", "'*'",
		"'/'", "'='", "'('", "')'", "'{'", "'['", "']'", "'}'", "'++'", "'--'",
		"'|'", "'>='", "'<='", "'>'", "'<'", "'!='", "'=='", "'->'", "", "",
		"'do'", "'fn'", "'if'", "'or'", "'in'", "'for'", "'not'", "'and'", "'let'",
		"'mut'", "'break'", "'continue'", "'else'", "'while'", "'forEach'",
		"'map'", "'filter'", "'reduce'", "'return'", "'println'", "'int'", "'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV",
		"EQ", "L_PAREN", "R_PAREN", "L_CURLY", "L_BRACKET", "R_BRACKET", "R_CURLY",
		"PLUS_PLUS", "MINUS_MINUS", "PIPE", "GTEQ", "LTEQ", "GT", "LT", "DIF",
		"EQEQ", "ARROW", "INT_NUMBER", "REAL_NUMBER", "DO", "FN", "IF", "OR",
		"IN", "FOR", "NOT", "AND", "LET", "MUT", "BREAK", "CONTINUE", "ELSE",
		"WHILE", "FOR_EACH", "MAP", "FILTER", "REDUCE", "RETURN", "PRINT_LN",
		"TYPE_INT", "TYPE_FLOAT", "TYPE_BOOLEAN", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"program", "funcMain", "function", "funcType", "return", "funcScope",
		"funcCall", "funcCallArg", "anonimousFuncWithReturn", "anonimousFuncNoReturn",
		"ifScope", "elseExpression", "elseIfExpression", "ifExpression", "loopScope",
		"loopDoWhile", "loopWhile", "loopForIn", "loopForI", "bodyScope", "println",
		"variable", "functionArgs", "funcArg", "dataTypes", "assignment", "array",
		"slice", "forEach", "mapFilterReduce", "map", "filter", "reduce", "relExpression",
		"mathExpression", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 528, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 5, 0,
		74, 8, 0, 10, 0, 12, 0, 77, 9, 0, 1, 0, 1, 0, 5, 0, 81, 8, 0, 10, 0, 12,
		0, 84, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 98, 8, 2, 10, 2, 12, 2, 101, 9, 2, 3, 2, 103, 8, 2, 1,
		2, 1, 2, 3, 2, 107, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 116, 8, 3, 10, 3, 12, 3, 119, 9, 3, 3, 3, 121, 8, 3, 1, 3, 1, 3, 3,
		3, 125, 8, 3, 1, 4, 3, 4, 128, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 133, 8, 4,
		1, 5, 1, 5, 5, 5, 137, 8, 5, 10, 5, 12, 5, 140, 9, 5, 1, 5, 3, 5, 143,
		8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 152, 8, 6, 10, 6,
		12, 6, 155, 9, 6, 3, 6, 157, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		3, 7, 165, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 171, 8, 8, 10, 8, 12, 8,
		174, 9, 8, 3, 8, 176, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 183, 8,
		8, 10, 8, 12, 8, 186, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 195, 8, 8, 10, 8, 12, 8, 198, 9, 8, 3, 8, 200, 8, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 206, 8, 8, 1, 8, 1, 8, 3, 8, 210, 8, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 216, 8, 9, 10, 9, 12, 9, 219, 9, 9, 3, 9, 221, 8, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 227, 8, 9, 10, 9, 12, 9, 230, 9, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 5, 9, 237, 8, 9, 10, 9, 12, 9, 240, 9, 9, 3, 9, 242,
		8, 9, 1, 9, 1, 9, 1, 9, 5, 9, 247, 8, 9, 10, 9, 12, 9, 250, 9, 9, 3, 9,
		252, 8, 9, 1, 10, 1, 10, 5, 10, 256, 8, 10, 10, 10, 12, 10, 259, 9, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 272, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 279, 8,
		13, 1, 14, 1, 14, 1, 14, 3, 14, 284, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15,
		289, 8, 15, 10, 15, 12, 15, 292, 9, 15, 1, 15, 1, 15, 1, 15, 3, 15, 297,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 303, 8, 16, 10, 16, 12, 16, 306,
		9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 320, 8, 17, 1, 17, 1, 17, 5, 17, 324, 8, 17, 10,
		17, 12, 17, 327, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 5, 18, 339, 8, 18, 10, 18, 12, 18, 342, 9, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 359, 8, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 365, 8, 20, 1, 20, 1, 20, 1, 21, 3, 21, 370, 8, 21, 1, 21, 1,
		21, 1, 21, 3, 21, 375, 8, 21, 1, 22, 1, 22, 1, 22, 5, 22, 380, 8, 22, 10,
		22, 12, 22, 383, 9, 22, 1, 23, 3, 23, 386, 8, 23, 1, 23, 1, 23, 1, 23,
		3, 23, 391, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 3, 25, 404, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 5, 26, 411, 8, 26, 10, 26, 12, 26, 414, 9, 26, 5, 26, 416, 8, 26,
		10, 26, 12, 26, 419, 9, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 443, 8, 28, 1, 28, 1, 28, 3,
		28, 447, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29,
		456, 8, 29, 1, 29, 1, 29, 1, 29, 5, 29, 461, 8, 29, 10, 29, 12, 29, 464,
		9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 492, 8, 33, 1, 33, 1,
		33, 1, 33, 5, 33, 497, 8, 33, 10, 33, 12, 33, 500, 9, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 509, 8, 34, 1, 34, 1, 34, 3,
		34, 513, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 521, 8,
		34, 10, 34, 12, 34, 524, 9, 34, 1, 35, 1, 35, 1, 35, 0, 3, 58, 66, 68,
		36, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		0, 6, 1, 0, 50, 51, 1, 0, 28, 29, 3, 0, 21, 26, 33, 33, 37, 37, 1, 0, 18,
		19, 1, 0, 9, 10, 1, 0, 7, 8, 583, 0, 75, 1, 0, 0, 0, 2, 85, 1, 0, 0, 0,
		4, 91, 1, 0, 0, 0, 6, 110, 1, 0, 0, 0, 8, 127, 1, 0, 0, 0, 10, 134, 1,
		0, 0, 0, 12, 146, 1, 0, 0, 0, 14, 164, 1, 0, 0, 0, 16, 209, 1, 0, 0, 0,
		18, 251, 1, 0, 0, 0, 20, 253, 1, 0, 0, 0, 22, 262, 1, 0, 0, 0, 24, 265,
		1, 0, 0, 0, 26, 273, 1, 0, 0, 0, 28, 283, 1, 0, 0, 0, 30, 285, 1, 0, 0,
		0, 32, 298, 1, 0, 0, 0, 34, 309, 1, 0, 0, 0, 36, 330, 1, 0, 0, 0, 38, 358,
		1, 0, 0, 0, 40, 360, 1, 0, 0, 0, 42, 369, 1, 0, 0, 0, 44, 376, 1, 0, 0,
		0, 46, 390, 1, 0, 0, 0, 48, 392, 1, 0, 0, 0, 50, 394, 1, 0, 0, 0, 52, 405,
		1, 0, 0, 0, 54, 422, 1, 0, 0, 0, 56, 446, 1, 0, 0, 0, 58, 455, 1, 0, 0,
		0, 60, 465, 1, 0, 0, 0, 62, 470, 1, 0, 0, 0, 64, 475, 1, 0, 0, 0, 66, 491,
		1, 0, 0, 0, 68, 512, 1, 0, 0, 0, 70, 525, 1, 0, 0, 0, 72, 74, 3, 4, 2,
		0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76,
		1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 82, 3, 2, 1, 0,
		79, 81, 3, 4, 2, 0, 80, 79, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 1, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85,
		86, 5, 31, 0, 0, 86, 87, 5, 1, 0, 0, 87, 88, 5, 12, 0, 0, 88, 89, 5, 13,
		0, 0, 89, 90, 3, 10, 5, 0, 90, 3, 1, 0, 0, 0, 91, 92, 5, 31, 0, 0, 92,
		93, 5, 53, 0, 0, 93, 102, 5, 12, 0, 0, 94, 99, 3, 44, 22, 0, 95, 96, 5,
		6, 0, 0, 96, 98, 3, 44, 22, 0, 97, 95, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0,
		99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99,
		1, 0, 0, 0, 102, 94, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 106, 5, 13, 0, 0, 105, 107, 3, 48, 24, 0, 106, 105, 1, 0, 0,
		0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 3, 10, 5, 0, 109,
		5, 1, 0, 0, 0, 110, 111, 5, 53, 0, 0, 111, 120, 5, 12, 0, 0, 112, 117,
		3, 44, 22, 0, 113, 114, 5, 6, 0, 0, 114, 116, 3, 44, 22, 0, 115, 113, 1,
		0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0,
		0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 112, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 5, 13, 0, 0, 123, 125,
		3, 48, 24, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 7, 1, 0,
		0, 0, 126, 128, 5, 48, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0,
		128, 132, 1, 0, 0, 0, 129, 133, 3, 68, 34, 0, 130, 133, 3, 66, 33, 0, 131,
		133, 5, 53, 0, 0, 132, 129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 131,
		1, 0, 0, 0, 133, 9, 1, 0, 0, 0, 134, 138, 5, 14, 0, 0, 135, 137, 3, 38,
		19, 0, 136, 135, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141,
		143, 3, 8, 4, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 145, 5, 17, 0, 0, 145, 11, 1, 0, 0, 0, 146, 147, 5, 53,
		0, 0, 147, 156, 5, 12, 0, 0, 148, 153, 3, 14, 7, 0, 149, 150, 5, 6, 0,
		0, 150, 152, 3, 14, 7, 0, 151, 149, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153,
		151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153,
		1, 0, 0, 0, 156, 148, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0,
		0, 0, 158, 159, 5, 13, 0, 0, 159, 13, 1, 0, 0, 0, 160, 165, 3, 68, 34,
		0, 161, 165, 3, 12, 6, 0, 162, 165, 3, 16, 8, 0, 163, 165, 3, 18, 9, 0,
		164, 160, 1, 0, 0, 0, 164, 161, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164,
		163, 1, 0, 0, 0, 165, 15, 1, 0, 0, 0, 166, 175, 5, 12, 0, 0, 167, 172,
		3, 44, 22, 0, 168, 169, 5, 6, 0, 0, 169, 171, 3, 44, 22, 0, 170, 168, 1,
		0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0,
		0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 167, 1, 0, 0, 0, 175,
		176, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 5, 13, 0, 0, 178, 179,
		3, 48, 24, 0, 179, 180, 5, 27, 0, 0, 180, 184, 5, 14, 0, 0, 181, 183, 3,
		38, 19, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0,
		0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0,
		187, 188, 3, 8, 4, 0, 188, 189, 5, 17, 0, 0, 189, 210, 1, 0, 0, 0, 190,
		199, 5, 12, 0, 0, 191, 196, 3, 44, 22, 0, 192, 193, 5, 6, 0, 0, 193, 195,
		3, 44, 22, 0, 194, 192, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1,
		0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0,
		0, 199, 191, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		202, 5, 13, 0, 0, 202, 203, 3, 48, 24, 0, 203, 205, 5, 27, 0, 0, 204, 206,
		3, 38, 19, 0, 205, 204, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207, 1,
		0, 0, 0, 207, 208, 3, 8, 4, 0, 208, 210, 1, 0, 0, 0, 209, 166, 1, 0, 0,
		0, 209, 190, 1, 0, 0, 0, 210, 17, 1, 0, 0, 0, 211, 220, 5, 12, 0, 0, 212,
		217, 3, 44, 22, 0, 213, 214, 5, 6, 0, 0, 214, 216, 3, 44, 22, 0, 215, 213,
		1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0,
		0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 212, 1, 0, 0, 0,
		220, 221, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 5, 13, 0, 0, 223,
		224, 5, 27, 0, 0, 224, 228, 5, 14, 0, 0, 225, 227, 3, 38, 19, 0, 226, 225,
		1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0,
		0, 0, 229, 231, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 252, 5, 17, 0, 0,
		232, 241, 5, 12, 0, 0, 233, 238, 3, 44, 22, 0, 234, 235, 5, 6, 0, 0, 235,
		237, 3, 44, 22, 0, 236, 234, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236,
		1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0,
		0, 0, 241, 233, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0,
		243, 244, 5, 13, 0, 0, 244, 248, 5, 27, 0, 0, 245, 247, 3, 38, 19, 0, 246,
		245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 211, 1, 0,
		0, 0, 251, 232, 1, 0, 0, 0, 252, 19, 1, 0, 0, 0, 253, 257, 5, 14, 0, 0,
		254, 256, 3, 38, 19, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257,
		255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0, 259, 257,
		1, 0, 0, 0, 260, 261, 5, 17, 0, 0, 261, 21, 1, 0, 0, 0, 262, 263, 5, 42,
		0, 0, 263, 264, 3, 20, 10, 0, 264, 23, 1, 0, 0, 0, 265, 266, 5, 42, 0,
		0, 266, 267, 5, 32, 0, 0, 267, 268, 3, 66, 33, 0, 268, 271, 3, 20, 10,
		0, 269, 272, 3, 24, 12, 0, 270, 272, 3, 22, 11, 0, 271, 269, 1, 0, 0, 0,
		271, 270, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 25, 1, 0, 0, 0, 273, 274,
		5, 32, 0, 0, 274, 275, 3, 66, 33, 0, 275, 278, 3, 20, 10, 0, 276, 279,
		3, 22, 11, 0, 277, 279, 3, 24, 12, 0, 278, 276, 1, 0, 0, 0, 278, 277, 1,
		0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 27, 1, 0, 0, 0, 280, 284, 5, 41, 0,
		0, 281, 284, 5, 40, 0, 0, 282, 284, 3, 38, 19, 0, 283, 280, 1, 0, 0, 0,
		283, 281, 1, 0, 0, 0, 283, 282, 1, 0, 0, 0, 284, 29, 1, 0, 0, 0, 285, 286,
		5, 30, 0, 0, 286, 290, 5, 14, 0, 0, 287, 289, 3, 28, 14, 0, 288, 287, 1,
		0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0,
		0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 296, 5, 17, 0, 0, 294,
		295, 5, 43, 0, 0, 295, 297, 3, 66, 33, 0, 296, 294, 1, 0, 0, 0, 296, 297,
		1, 0, 0, 0, 297, 31, 1, 0, 0, 0, 298, 299, 5, 43, 0, 0, 299, 300, 3, 66,
		33, 0, 300, 304, 5, 14, 0, 0, 301, 303, 3, 28, 14, 0, 302, 301, 1, 0, 0,
		0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305,
		307, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 308, 5, 17, 0, 0, 308, 33,
		1, 0, 0, 0, 309, 310, 5, 35, 0, 0, 310, 311, 5, 53, 0, 0, 311, 319, 5,
		34, 0, 0, 312, 320, 3, 54, 27, 0, 313, 320, 5, 53, 0, 0, 314, 315, 5, 12,
		0, 0, 315, 316, 5, 28, 0, 0, 316, 317, 5, 2, 0, 0, 317, 318, 5, 28, 0,
		0, 318, 320, 5, 13, 0, 0, 319, 312, 1, 0, 0, 0, 319, 313, 1, 0, 0, 0, 319,
		314, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 325, 5, 14, 0, 0, 322, 324,
		3, 28, 14, 0, 323, 322, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1,
		0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 328, 1, 0, 0, 0, 327, 325, 1, 0, 0,
		0, 328, 329, 5, 17, 0, 0, 329, 35, 1, 0, 0, 0, 330, 331, 5, 35, 0, 0, 331,
		332, 3, 50, 25, 0, 332, 333, 5, 3, 0, 0, 333, 334, 3, 66, 33, 0, 334, 335,
		5, 3, 0, 0, 335, 336, 3, 68, 34, 0, 336, 340, 5, 14, 0, 0, 337, 339, 3,
		28, 14, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0,
		0, 0, 340, 341, 1, 0, 0, 0, 341, 343, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0,
		343, 344, 5, 17, 0, 0, 344, 37, 1, 0, 0, 0, 345, 359, 3, 42, 21, 0, 346,
		359, 3, 50, 25, 0, 347, 359, 3, 26, 13, 0, 348, 359, 3, 4, 2, 0, 349, 359,
		3, 12, 6, 0, 350, 359, 3, 40, 20, 0, 351, 359, 3, 56, 28, 0, 352, 359,
		3, 32, 16, 0, 353, 359, 3, 30, 15, 0, 354, 359, 3, 34, 17, 0, 355, 359,
		3, 36, 18, 0, 356, 359, 3, 68, 34, 0, 357, 359, 3, 58, 29, 0, 358, 345,
		1, 0, 0, 0, 358, 346, 1, 0, 0, 0, 358, 347, 1, 0, 0, 0, 358, 348, 1, 0,
		0, 0, 358, 349, 1, 0, 0, 0, 358, 350, 1, 0, 0, 0, 358, 351, 1, 0, 0, 0,
		358, 352, 1, 0, 0, 0, 358, 353, 1, 0, 0, 0, 358, 354, 1, 0, 0, 0, 358,
		355, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 39, 1,
		0, 0, 0, 360, 361, 5, 49, 0, 0, 361, 364, 5, 12, 0, 0, 362, 365, 3, 42,
		21, 0, 363, 365, 5, 53, 0, 0, 364, 362, 1, 0, 0, 0, 364, 363, 1, 0, 0,
		0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 13, 0, 0, 367, 41, 1, 0, 0, 0, 368,
		370, 5, 39, 0, 0, 369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371,
		1, 0, 0, 0, 371, 372, 5, 38, 0, 0, 372, 374, 5, 53, 0, 0, 373, 375, 3,
		48, 24, 0, 374, 373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 43, 1, 0, 0,
		0, 376, 381, 3, 46, 23, 0, 377, 378, 5, 6, 0, 0, 378, 380, 3, 46, 23, 0,
		379, 377, 1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381,
		382, 1, 0, 0, 0, 382, 45, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 386, 5,
		39, 0, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 1, 0, 0,
		0, 387, 388, 5, 53, 0, 0, 388, 391, 3, 48, 24, 0, 389, 391, 3, 6, 3, 0,
		390, 385, 1, 0, 0, 0, 390, 389, 1, 0, 0, 0, 391, 47, 1, 0, 0, 0, 392, 393,
		7, 0, 0, 0, 393, 49, 1, 0, 0, 0, 394, 395, 3, 42, 21, 0, 395, 403, 5, 11,
		0, 0, 396, 404, 3, 54, 27, 0, 397, 404, 3, 52, 26, 0, 398, 404, 3, 16,
		8, 0, 399, 404, 3, 18, 9, 0, 400, 404, 3, 68, 34, 0, 401, 404, 3, 66, 33,
		0, 402, 404, 3, 58, 29, 0, 403, 396, 1, 0, 0, 0, 403, 397, 1, 0, 0, 0,
		403, 398, 1, 0, 0, 0, 403, 399, 1, 0, 0, 0, 403, 400, 1, 0, 0, 0, 403,
		401, 1, 0, 0, 0, 403, 402, 1, 0, 0, 0, 404, 51, 1, 0, 0, 0, 405, 406, 3,
		48, 24, 0, 406, 417, 5, 15, 0, 0, 407, 412, 7, 1, 0, 0, 408, 409, 5, 6,
		0, 0, 409, 411, 7, 1, 0, 0, 410, 408, 1, 0, 0, 0, 411, 414, 1, 0, 0, 0,
		412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 416, 1, 0, 0, 0, 414,
		412, 1, 0, 0, 0, 415, 407, 1, 0, 0, 0, 416, 419, 1, 0, 0, 0, 417, 415,
		1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 420, 1, 0, 0, 0, 419, 417, 1, 0,
		0, 0, 420, 421, 5, 16, 0, 0, 421, 53, 1, 0, 0, 0, 422, 423, 5, 53, 0, 0,
		423, 424, 5, 15, 0, 0, 424, 425, 5, 28, 0, 0, 425, 426, 5, 4, 0, 0, 426,
		427, 5, 28, 0, 0, 427, 428, 5, 16, 0, 0, 428, 55, 1, 0, 0, 0, 429, 430,
		5, 53, 0, 0, 430, 431, 5, 5, 0, 0, 431, 432, 5, 44, 0, 0, 432, 433, 5,
		12, 0, 0, 433, 434, 3, 18, 9, 0, 434, 435, 5, 13, 0, 0, 435, 447, 1, 0,
		0, 0, 436, 437, 3, 52, 26, 0, 437, 438, 5, 5, 0, 0, 438, 439, 5, 44, 0,
		0, 439, 442, 5, 12, 0, 0, 440, 443, 3, 18, 9, 0, 441, 443, 5, 53, 0, 0,
		442, 440, 1, 0, 0, 0, 442, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444,
		445, 5, 13, 0, 0, 445, 447, 1, 0, 0, 0, 446, 429, 1, 0, 0, 0, 446, 436,
		1, 0, 0, 0, 447, 57, 1, 0, 0, 0, 448, 449, 6, 29, -1, 0, 449, 456, 3, 54,
		27, 0, 450, 456, 5, 53, 0, 0, 451, 456, 3, 60, 30, 0, 452, 456, 3, 62,
		31, 0, 453, 456, 3, 64, 32, 0, 454, 456, 3, 52, 26, 0, 455, 448, 1, 0,
		0, 0, 455, 450, 1, 0, 0, 0, 455, 451, 1, 0, 0, 0, 455, 452, 1, 0, 0, 0,
		455, 453, 1, 0, 0, 0, 455, 454, 1, 0, 0, 0, 456, 462, 1, 0, 0, 0, 457,
		458, 10, 7, 0, 0, 458, 459, 5, 5, 0, 0, 459, 461, 3, 58, 29, 8, 460, 457,
		1, 0, 0, 0, 461, 464, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 462, 463, 1, 0,
		0, 0, 463, 59, 1, 0, 0, 0, 464, 462, 1, 0, 0, 0, 465, 466, 5, 45, 0, 0,
		466, 467, 5, 12, 0, 0, 467, 468, 3, 16, 8, 0, 468, 469, 5, 13, 0, 0, 469,
		61, 1, 0, 0, 0, 470, 471, 5, 46, 0, 0, 471, 472, 5, 12, 0, 0, 472, 473,
		3, 16, 8, 0, 473, 474, 5, 13, 0, 0, 474, 63, 1, 0, 0, 0, 475, 476, 5, 47,
		0, 0, 476, 477, 5, 12, 0, 0, 477, 478, 3, 16, 8, 0, 478, 479, 5, 13, 0,
		0, 479, 65, 1, 0, 0, 0, 480, 481, 6, 33, -1, 0, 481, 482, 5, 12, 0, 0,
		482, 483, 3, 66, 33, 0, 483, 484, 5, 13, 0, 0, 484, 492, 1, 0, 0, 0, 485,
		486, 5, 36, 0, 0, 486, 492, 3, 66, 33, 5, 487, 492, 5, 53, 0, 0, 488, 492,
		3, 70, 35, 0, 489, 492, 3, 12, 6, 0, 490, 492, 5, 52, 0, 0, 491, 480, 1,
		0, 0, 0, 491, 485, 1, 0, 0, 0, 491, 487, 1, 0, 0, 0, 491, 488, 1, 0, 0,
		0, 491, 489, 1, 0, 0, 0, 491, 490, 1, 0, 0, 0, 492, 498, 1, 0, 0, 0, 493,
		494, 10, 7, 0, 0, 494, 495, 7, 2, 0, 0, 495, 497, 3, 66, 33, 8, 496, 493,
		1, 0, 0, 0, 497, 500, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 498, 499, 1, 0,
		0, 0, 499, 67, 1, 0, 0, 0, 500, 498, 1, 0, 0, 0, 501, 502, 6, 34, -1, 0,
		502, 503, 5, 12, 0, 0, 503, 504, 3, 68, 34, 0, 504, 505, 5, 13, 0, 0, 505,
		513, 1, 0, 0, 0, 506, 508, 5, 53, 0, 0, 507, 509, 7, 3, 0, 0, 508, 507,
		1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 513, 1, 0, 0, 0, 510, 513, 3, 70,
		35, 0, 511, 513, 3, 12, 6, 0, 512, 501, 1, 0, 0, 0, 512, 506, 1, 0, 0,
		0, 512, 510, 1, 0, 0, 0, 512, 511, 1, 0, 0, 0, 513, 522, 1, 0, 0, 0, 514,
		515, 10, 6, 0, 0, 515, 516, 7, 4, 0, 0, 516, 521, 3, 68, 34, 7, 517, 518,
		10, 5, 0, 0, 518, 519, 7, 5, 0, 0, 519, 521, 3, 68, 34, 6, 520, 514, 1,
		0, 0, 0, 520, 517, 1, 0, 0, 0, 521, 524, 1, 0, 0, 0, 522, 520, 1, 0, 0,
		0, 522, 523, 1, 0, 0, 0, 523, 69, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0, 525,
		526, 7, 1, 0, 0, 526, 71, 1, 0, 0, 0, 59, 75, 82, 99, 102, 106, 117, 120,
		124, 127, 132, 138, 142, 153, 156, 164, 172, 175, 184, 196, 199, 205, 209,
		217, 220, 228, 238, 241, 248, 251, 257, 271, 278, 283, 290, 296, 304, 319,
		325, 340, 358, 364, 369, 374, 381, 385, 390, 403, 412, 417, 442, 446, 455,
		462, 491, 498, 508, 512, 520, 522,
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
	this.GrammarFileName = "IronLang.g4"

	return this
}

// IronLangParser tokens.
const (
	IronLangParserEOF          = antlr.TokenEOF
	IronLangParserT__0         = 1
	IronLangParserT__1         = 2
	IronLangParserT__2         = 3
	IronLangParserT__3         = 4
	IronLangParserDOT          = 5
	IronLangParserCOMMA        = 6
	IronLangParserPLUS         = 7
	IronLangParserMINUS        = 8
	IronLangParserMULT         = 9
	IronLangParserDIV          = 10
	IronLangParserEQ           = 11
	IronLangParserL_PAREN      = 12
	IronLangParserR_PAREN      = 13
	IronLangParserL_CURLY      = 14
	IronLangParserL_BRACKET    = 15
	IronLangParserR_BRACKET    = 16
	IronLangParserR_CURLY      = 17
	IronLangParserPLUS_PLUS    = 18
	IronLangParserMINUS_MINUS  = 19
	IronLangParserPIPE         = 20
	IronLangParserGTEQ         = 21
	IronLangParserLTEQ         = 22
	IronLangParserGT           = 23
	IronLangParserLT           = 24
	IronLangParserDIF          = 25
	IronLangParserEQEQ         = 26
	IronLangParserARROW        = 27
	IronLangParserINT_NUMBER   = 28
	IronLangParserREAL_NUMBER  = 29
	IronLangParserDO           = 30
	IronLangParserFN           = 31
	IronLangParserIF           = 32
	IronLangParserOR           = 33
	IronLangParserIN           = 34
	IronLangParserFOR          = 35
	IronLangParserNOT          = 36
	IronLangParserAND          = 37
	IronLangParserLET          = 38
	IronLangParserMUT          = 39
	IronLangParserBREAK        = 40
	IronLangParserCONTINUE     = 41
	IronLangParserELSE         = 42
	IronLangParserWHILE        = 43
	IronLangParserFOR_EACH     = 44
	IronLangParserMAP          = 45
	IronLangParserFILTER       = 46
	IronLangParserREDUCE       = 47
	IronLangParserRETURN       = 48
	IronLangParserPRINT_LN     = 49
	IronLangParserTYPE_INT     = 50
	IronLangParserTYPE_FLOAT   = 51
	IronLangParserTYPE_BOOLEAN = 52
	IronLangParserIDENTIFIER   = 53
	IronLangParserWHITE_SPACE  = 54
)

// IronLangParser rules.
const (
	IronLangParserRULE_program                 = 0
	IronLangParserRULE_funcMain                = 1
	IronLangParserRULE_function                = 2
	IronLangParserRULE_funcType                = 3
	IronLangParserRULE_return                  = 4
	IronLangParserRULE_funcScope               = 5
	IronLangParserRULE_funcCall                = 6
	IronLangParserRULE_funcCallArg             = 7
	IronLangParserRULE_anonimousFuncWithReturn = 8
	IronLangParserRULE_anonimousFuncNoReturn   = 9
	IronLangParserRULE_ifScope                 = 10
	IronLangParserRULE_elseExpression          = 11
	IronLangParserRULE_elseIfExpression        = 12
	IronLangParserRULE_ifExpression            = 13
	IronLangParserRULE_loopScope               = 14
	IronLangParserRULE_loopDoWhile             = 15
	IronLangParserRULE_loopWhile               = 16
	IronLangParserRULE_loopForIn               = 17
	IronLangParserRULE_loopForI                = 18
	IronLangParserRULE_bodyScope               = 19
	IronLangParserRULE_println                 = 20
	IronLangParserRULE_variable                = 21
	IronLangParserRULE_functionArgs            = 22
	IronLangParserRULE_funcArg                 = 23
	IronLangParserRULE_dataTypes               = 24
	IronLangParserRULE_assignment              = 25
	IronLangParserRULE_array                   = 26
	IronLangParserRULE_slice                   = 27
	IronLangParserRULE_forEach                 = 28
	IronLangParserRULE_mapFilterReduce         = 29
	IronLangParserRULE_map                     = 30
	IronLangParserRULE_filter                  = 31
	IronLangParserRULE_reduce                  = 32
	IronLangParserRULE_relExpression           = 33
	IronLangParserRULE_mathExpression          = 34
	IronLangParserRULE_atom                    = 35
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ProgramContext) FuncMain() IFuncMainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncMainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncMainContext)
}

func (s *ProgramContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
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

	return t.(IFunctionContext)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(72)
				p.Function()
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(78)
		p.FuncMain()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserFN {
		{
			p.SetState(79)
			p.Function()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncMainContext is an interface to support dynamic dispatch.
type IFuncMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncMainContext differentiates from other interfaces.
	IsFuncMainContext()
}

type FuncMainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncMainContext() *FuncMainContext {
	var p = new(FuncMainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_funcMain
	return p
}

func (*FuncMainContext) IsFuncMainContext() {}

func NewFuncMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncMainContext {
	var p = new(FuncMainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_funcMain

	return p
}

func (s *FuncMainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncMainContext) FN() antlr.TerminalNode {
	return s.GetToken(IronLangParserFN, 0)
}

func (s *FuncMainContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FuncMainContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FuncMainContext) FuncScope() IFuncScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncScopeContext)
}

func (s *FuncMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncMainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFuncMain(s)
	}
}

func (s *FuncMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFuncMain(s)
	}
}

func (s *FuncMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFuncMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FuncMain() (localctx IFuncMainContext) {
	this := p
	_ = this

	localctx = NewFuncMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IronLangParserRULE_funcMain)

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
		p.SetState(85)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(86)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(87)
		p.Match(IronLangParserL_PAREN)
	}
	{
		p.SetState(88)
		p.Match(IronLangParserR_PAREN)
	}
	{
		p.SetState(89)
		p.FuncScope()
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FN() antlr.TerminalNode {
	return s.GetToken(IronLangParserFN, 0)
}

func (s *FunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *FunctionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *FunctionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *FunctionContext) FuncScope() IFuncScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncScopeContext)
}

func (s *FunctionContext) AllFunctionArgs() []IFunctionArgsContext {
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

func (s *FunctionContext) FunctionArgs(i int) IFunctionArgsContext {
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

func (s *FunctionContext) DataTypes() IDataTypesContext {
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

func (s *FunctionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *FunctionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IronLangParserRULE_function)
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
		p.SetState(91)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(92)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(93)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(94)
			p.FunctionArgs()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(95)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(96)
				p.FunctionArgs()
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(104)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(105)
			p.DataTypes()
		}

	}
	{
		p.SetState(108)
		p.FuncScope()
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
	p.EnterRule(localctx, 6, IronLangParserRULE_funcType)
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
		p.SetState(110)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(111)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(112)
			p.FunctionArgs()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(113)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(114)
				p.FunctionArgs()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(122)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(123)
			p.DataTypes()
		}

	}

	return localctx
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_return
	return p
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) MathExpression() IMathExpressionContext {
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

func (s *ReturnContext) RelExpression() IRelExpressionContext {
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

func (s *ReturnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(IronLangParserRETURN, 0)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Return() (localctx IReturnContext) {
	this := p
	_ = this

	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IronLangParserRULE_return)
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
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserRETURN {
		{
			p.SetState(126)
			p.Match(IronLangParserRETURN)
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(129)
			p.mathExpression(0)
		}

	case 2:
		{
			p.SetState(130)
			p.relExpression(0)
		}

	case 3:
		{
			p.SetState(131)
			p.Match(IronLangParserIDENTIFIER)
		}

	}

	return localctx
}

// IFuncScopeContext is an interface to support dynamic dispatch.
type IFuncScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncScopeContext differentiates from other interfaces.
	IsFuncScopeContext()
}

type FuncScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncScopeContext() *FuncScopeContext {
	var p = new(FuncScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_funcScope
	return p
}

func (*FuncScopeContext) IsFuncScopeContext() {}

func NewFuncScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncScopeContext {
	var p = new(FuncScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_funcScope

	return p
}

func (s *FuncScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncScopeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *FuncScopeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *FuncScopeContext) AllBodyScope() []IBodyScopeContext {
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

func (s *FuncScopeContext) BodyScope(i int) IBodyScopeContext {
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

func (s *FuncScopeContext) Return() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *FuncScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFuncScope(s)
	}
}

func (s *FuncScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFuncScope(s)
	}
}

func (s *FuncScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFuncScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FuncScope() (localctx IFuncScopeContext) {
	this := p
	_ = this

	localctx = NewFuncScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IronLangParserRULE_funcScope)
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
		p.SetState(134)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(135)
				p.BodyScope()
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(IronLangParserNOT-36))|(1<<(IronLangParserRETURN-36))|(1<<(IronLangParserTYPE_BOOLEAN-36))|(1<<(IronLangParserIDENTIFIER-36)))) != 0) {
		{
			p.SetState(141)
			p.Return()
		}

	}
	{
		p.SetState(144)
		p.Match(IronLangParserR_CURLY)
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
	p.EnterRule(localctx, 12, IronLangParserRULE_funcCall)
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
		p.SetState(146)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(147)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER))) != 0) || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(148)
			p.FuncCallArg()
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(149)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(150)
				p.FuncCallArg()
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(158)
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
	p.EnterRule(localctx, 14, IronLangParserRULE_funcCallArg)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.mathExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.FuncCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.AnonimousFuncWithReturn()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
			p.AnonimousFuncNoReturn()
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

func (s *AnonimousFuncWithReturnContext) Return() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
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
	p.EnterRule(localctx, 16, IronLangParserRULE_anonimousFuncWithReturn)
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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(167)
				p.FunctionArgs()
			}
			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(168)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(169)
					p.FunctionArgs()
				}

				p.SetState(174)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(177)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(178)
			p.DataTypes()
		}
		{
			p.SetState(179)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(180)
			p.Match(IronLangParserL_CURLY)
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(181)
					p.BodyScope()
				}

			}
			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}
		{
			p.SetState(187)
			p.Return()
		}
		{
			p.SetState(188)
			p.Match(IronLangParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(191)
				p.FunctionArgs()
			}
			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(192)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(193)
					p.FunctionArgs()
				}

				p.SetState(198)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(201)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(202)
			p.DataTypes()
		}
		{
			p.SetState(203)
			p.Match(IronLangParserARROW)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(204)
				p.BodyScope()
			}

		}
		{
			p.SetState(207)
			p.Return()
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
	p.EnterRule(localctx, 18, IronLangParserRULE_anonimousFuncNoReturn)
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

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(212)
				p.FunctionArgs()
			}
			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(213)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(214)
					p.FunctionArgs()
				}

				p.SetState(219)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(222)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(223)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(224)
			p.Match(IronLangParserL_CURLY)
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserDO)|(1<<IronLangParserFN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IronLangParserIF-32))|(1<<(IronLangParserFOR-32))|(1<<(IronLangParserLET-32))|(1<<(IronLangParserMUT-32))|(1<<(IronLangParserWHILE-32))|(1<<(IronLangParserMAP-32))|(1<<(IronLangParserFILTER-32))|(1<<(IronLangParserREDUCE-32))|(1<<(IronLangParserPRINT_LN-32))|(1<<(IronLangParserTYPE_INT-32))|(1<<(IronLangParserTYPE_FLOAT-32))|(1<<(IronLangParserIDENTIFIER-32)))) != 0) {
			{
				p.SetState(225)
				p.BodyScope()
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(231)
			p.Match(IronLangParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(233)
				p.FunctionArgs()
			}
			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(234)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(235)
					p.FunctionArgs()
				}

				p.SetState(240)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(243)
			p.Match(IronLangParserR_PAREN)
		}
		{
			p.SetState(244)
			p.Match(IronLangParserARROW)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(245)
					p.BodyScope()
				}

			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 20, IronLangParserRULE_ifScope)
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
		p.SetState(253)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserDO)|(1<<IronLangParserFN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IronLangParserIF-32))|(1<<(IronLangParserFOR-32))|(1<<(IronLangParserLET-32))|(1<<(IronLangParserMUT-32))|(1<<(IronLangParserWHILE-32))|(1<<(IronLangParserMAP-32))|(1<<(IronLangParserFILTER-32))|(1<<(IronLangParserREDUCE-32))|(1<<(IronLangParserPRINT_LN-32))|(1<<(IronLangParserTYPE_INT-32))|(1<<(IronLangParserTYPE_FLOAT-32))|(1<<(IronLangParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(254)
			p.BodyScope()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
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
	p.EnterRule(localctx, 22, IronLangParserRULE_elseExpression)

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
		p.SetState(262)
		p.Match(IronLangParserELSE)
	}
	{
		p.SetState(263)
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
	p.EnterRule(localctx, 24, IronLangParserRULE_elseIfExpression)

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
		p.SetState(265)
		p.Match(IronLangParserELSE)
	}
	{
		p.SetState(266)
		p.Match(IronLangParserIF)
	}
	{
		p.SetState(267)
		p.relExpression(0)
	}
	{
		p.SetState(268)
		p.IfScope()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(269)
			p.ElseIfExpression()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(270)
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
	p.EnterRule(localctx, 26, IronLangParserRULE_ifExpression)

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
		p.SetState(273)
		p.Match(IronLangParserIF)
	}
	{
		p.SetState(274)
		p.relExpression(0)
	}
	{
		p.SetState(275)
		p.IfScope()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(276)
			p.ElseExpression()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(277)
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
	p.EnterRule(localctx, 28, IronLangParserRULE_loopScope)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserCONTINUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Match(IronLangParserCONTINUE)
		}

	case IronLangParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(IronLangParserBREAK)
		}

	case IronLangParserL_PAREN, IronLangParserINT_NUMBER, IronLangParserREAL_NUMBER, IronLangParserDO, IronLangParserFN, IronLangParserIF, IronLangParserFOR, IronLangParserLET, IronLangParserMUT, IronLangParserWHILE, IronLangParserMAP, IronLangParserFILTER, IronLangParserREDUCE, IronLangParserPRINT_LN, IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT, IronLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
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
	p.EnterRule(localctx, 30, IronLangParserRULE_loopDoWhile)
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
		p.SetState(285)
		p.Match(IronLangParserDO)
	}
	{
		p.SetState(286)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserDO)|(1<<IronLangParserFN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IronLangParserIF-32))|(1<<(IronLangParserFOR-32))|(1<<(IronLangParserLET-32))|(1<<(IronLangParserMUT-32))|(1<<(IronLangParserBREAK-32))|(1<<(IronLangParserCONTINUE-32))|(1<<(IronLangParserWHILE-32))|(1<<(IronLangParserMAP-32))|(1<<(IronLangParserFILTER-32))|(1<<(IronLangParserREDUCE-32))|(1<<(IronLangParserPRINT_LN-32))|(1<<(IronLangParserTYPE_INT-32))|(1<<(IronLangParserTYPE_FLOAT-32))|(1<<(IronLangParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(287)
			p.LoopScope()
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(293)
		p.Match(IronLangParserR_CURLY)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(294)
			p.Match(IronLangParserWHILE)
		}
		{
			p.SetState(295)
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
	p.EnterRule(localctx, 32, IronLangParserRULE_loopWhile)
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
		p.SetState(298)
		p.Match(IronLangParserWHILE)
	}
	{
		p.SetState(299)
		p.relExpression(0)
	}
	{
		p.SetState(300)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserDO)|(1<<IronLangParserFN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IronLangParserIF-32))|(1<<(IronLangParserFOR-32))|(1<<(IronLangParserLET-32))|(1<<(IronLangParserMUT-32))|(1<<(IronLangParserBREAK-32))|(1<<(IronLangParserCONTINUE-32))|(1<<(IronLangParserWHILE-32))|(1<<(IronLangParserMAP-32))|(1<<(IronLangParserFILTER-32))|(1<<(IronLangParserREDUCE-32))|(1<<(IronLangParserPRINT_LN-32))|(1<<(IronLangParserTYPE_INT-32))|(1<<(IronLangParserTYPE_FLOAT-32))|(1<<(IronLangParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(301)
			p.LoopScope()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
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
	p.EnterRule(localctx, 34, IronLangParserRULE_loopForIn)
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
		p.SetState(309)
		p.Match(IronLangParserFOR)
	}
	{
		p.SetState(310)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(311)
		p.Match(IronLangParserIN)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(312)
			p.Slice()
		}

	case 2:
		{
			p.SetState(313)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 3:
		{
			p.SetState(314)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(315)
			p.Match(IronLangParserINT_NUMBER)
		}
		{
			p.SetState(316)
			p.Match(IronLangParserT__1)
		}
		{
			p.SetState(317)
			p.Match(IronLangParserINT_NUMBER)
		}
		{
			p.SetState(318)
			p.Match(IronLangParserR_PAREN)
		}

	}
	{
		p.SetState(321)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserDO)|(1<<IronLangParserFN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IronLangParserIF-32))|(1<<(IronLangParserFOR-32))|(1<<(IronLangParserLET-32))|(1<<(IronLangParserMUT-32))|(1<<(IronLangParserBREAK-32))|(1<<(IronLangParserCONTINUE-32))|(1<<(IronLangParserWHILE-32))|(1<<(IronLangParserMAP-32))|(1<<(IronLangParserFILTER-32))|(1<<(IronLangParserREDUCE-32))|(1<<(IronLangParserPRINT_LN-32))|(1<<(IronLangParserTYPE_INT-32))|(1<<(IronLangParserTYPE_FLOAT-32))|(1<<(IronLangParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(322)
			p.LoopScope()
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(328)
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
	p.EnterRule(localctx, 36, IronLangParserRULE_loopForI)
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
		p.SetState(330)
		p.Match(IronLangParserFOR)
	}
	{
		p.SetState(331)
		p.Assignment()
	}
	{
		p.SetState(332)
		p.Match(IronLangParserT__2)
	}
	{
		p.SetState(333)
		p.relExpression(0)
	}
	{
		p.SetState(334)
		p.Match(IronLangParserT__2)
	}
	{
		p.SetState(335)
		p.mathExpression(0)
	}
	{
		p.SetState(336)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserDO)|(1<<IronLangParserFN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(IronLangParserIF-32))|(1<<(IronLangParserFOR-32))|(1<<(IronLangParserLET-32))|(1<<(IronLangParserMUT-32))|(1<<(IronLangParserBREAK-32))|(1<<(IronLangParserCONTINUE-32))|(1<<(IronLangParserWHILE-32))|(1<<(IronLangParserMAP-32))|(1<<(IronLangParserFILTER-32))|(1<<(IronLangParserREDUCE-32))|(1<<(IronLangParserPRINT_LN-32))|(1<<(IronLangParserTYPE_INT-32))|(1<<(IronLangParserTYPE_FLOAT-32))|(1<<(IronLangParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(337)
			p.LoopScope()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(343)
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

func (s *BodyScopeContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
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

func (s *BodyScopeContext) Println() IPrintlnContext {
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

func (s *BodyScopeContext) MapFilterReduce() IMapFilterReduceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFilterReduceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFilterReduceContext)
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
	p.EnterRule(localctx, 38, IronLangParserRULE_bodyScope)

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

	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(347)
			p.IfExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(348)
			p.Function()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(349)
			p.FuncCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(350)
			p.Println()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(351)
			p.ForEach()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(352)
			p.LoopWhile()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(353)
			p.LoopDoWhile()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(354)
			p.LoopForIn()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(355)
			p.LoopForI()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(356)
			p.mathExpression(0)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(357)
			p.mapFilterReduce(0)
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

func (p *IronLangParser) Println() (localctx IPrintlnContext) {
	this := p
	_ = this

	localctx = NewPrintlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, IronLangParserRULE_println)

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
		p.SetState(360)
		p.Match(IronLangParserPRINT_LN)
	}
	{
		p.SetState(361)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(362)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(363)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(366)
		p.Match(IronLangParserR_PAREN)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 42, IronLangParserRULE_variable)
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
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(368)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(371)
		p.Match(IronLangParserLET)
	}
	{
		p.SetState(372)
		p.Match(IronLangParserIDENTIFIER)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(373)
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

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *FunctionArgsContext) AllFuncArg() []IFuncArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncArgContext); ok {
			len++
		}
	}

	tst := make([]IFuncArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncArgContext); ok {
			tst[i] = t.(IFuncArgContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) FuncArg(i int) IFuncArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgContext); ok {
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

	return t.(IFuncArgContext)
}

func (s *FunctionArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *FunctionArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
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
	p.EnterRule(localctx, 44, IronLangParserRULE_functionArgs)

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
		p.SetState(376)
		p.FuncArg()
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(377)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(378)
				p.FuncArg()
			}

		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncArgContext is an interface to support dynamic dispatch.
type IFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgContext differentiates from other interfaces.
	IsFuncArgContext()
}

type FuncArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgContext() *FuncArgContext {
	var p = new(FuncArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_funcArg
	return p
}

func (*FuncArgContext) IsFuncArgContext() {}

func NewFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgContext {
	var p = new(FuncArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_funcArg

	return p
}

func (s *FuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *FuncArgContext) DataTypes() IDataTypesContext {
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

func (s *FuncArgContext) MUT() antlr.TerminalNode {
	return s.GetToken(IronLangParserMUT, 0)
}

func (s *FuncArgContext) FuncType() IFuncTypeContext {
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

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) FuncArg() (localctx IFuncArgContext) {
	this := p
	_ = this

	localctx = NewFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, IronLangParserRULE_funcArg)
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

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT {
			{
				p.SetState(384)
				p.Match(IronLangParserMUT)
			}

		}
		{
			p.SetState(387)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(388)
			p.DataTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
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
	p.EnterRule(localctx, 48, IronLangParserRULE_dataTypes)
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
		p.SetState(392)
		_la = p.GetTokenStream().LA(1)

		if !(_la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT) {
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

func (s *AssignmentContext) Variable() IVariableContext {
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

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(IronLangParserEQ, 0)
}

func (s *AssignmentContext) Slice() ISliceContext {
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

func (s *AssignmentContext) Array() IArrayContext {
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

func (s *AssignmentContext) AnonimousFuncWithReturn() IAnonimousFuncWithReturnContext {
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

func (s *AssignmentContext) AnonimousFuncNoReturn() IAnonimousFuncNoReturnContext {
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

func (s *AssignmentContext) MathExpression() IMathExpressionContext {
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

func (s *AssignmentContext) RelExpression() IRelExpressionContext {
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

func (s *AssignmentContext) MapFilterReduce() IMapFilterReduceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFilterReduceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFilterReduceContext)
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
	p.EnterRule(localctx, 50, IronLangParserRULE_assignment)

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
		p.SetState(394)
		p.Variable()
	}
	{
		p.SetState(395)
		p.Match(IronLangParserEQ)
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(396)
			p.Slice()
		}

	case 2:
		{
			p.SetState(397)
			p.Array()
		}

	case 3:
		{
			p.SetState(398)
			p.AnonimousFuncWithReturn()
		}

	case 4:
		{
			p.SetState(399)
			p.AnonimousFuncNoReturn()
		}

	case 5:
		{
			p.SetState(400)
			p.mathExpression(0)
		}

	case 6:
		{
			p.SetState(401)
			p.relExpression(0)
		}

	case 7:
		{
			p.SetState(402)
			p.mapFilterReduce(0)
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
	p.EnterRule(localctx, 52, IronLangParserRULE_array)
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
		p.SetState(405)
		p.DataTypes()
	}
	{
		p.SetState(406)
		p.Match(IronLangParserL_BRACKET)
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER {
		{
			p.SetState(407)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(408)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(409)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(414)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(420)
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
	p.EnterRule(localctx, 54, IronLangParserRULE_slice)

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
		p.SetState(422)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(423)
		p.Match(IronLangParserL_BRACKET)
	}
	{
		p.SetState(424)
		p.Match(IronLangParserINT_NUMBER)
	}
	{
		p.SetState(425)
		p.Match(IronLangParserT__3)
	}
	{
		p.SetState(426)
		p.Match(IronLangParserINT_NUMBER)
	}
	{
		p.SetState(427)
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
	p.EnterRule(localctx, 56, IronLangParserRULE_forEach)

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

	p.SetState(446)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(430)
			p.Match(IronLangParserDOT)
		}
		{
			p.SetState(431)
			p.Match(IronLangParserFOR_EACH)
		}
		{
			p.SetState(432)
			p.Match(IronLangParserL_PAREN)
		}

		{
			p.SetState(433)
			p.AnonimousFuncNoReturn()
		}

		{
			p.SetState(434)
			p.Match(IronLangParserR_PAREN)
		}

	case IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.Array()
		}
		{
			p.SetState(437)
			p.Match(IronLangParserDOT)
		}
		{
			p.SetState(438)
			p.Match(IronLangParserFOR_EACH)
		}
		{
			p.SetState(439)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IronLangParserL_PAREN:
			{
				p.SetState(440)
				p.AnonimousFuncNoReturn()
			}

		case IronLangParserIDENTIFIER:
			{
				p.SetState(441)
				p.Match(IronLangParserIDENTIFIER)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(444)
			p.Match(IronLangParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMapFilterReduceContext is an interface to support dynamic dispatch.
type IMapFilterReduceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapFilterReduceContext differentiates from other interfaces.
	IsMapFilterReduceContext()
}

type MapFilterReduceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFilterReduceContext() *MapFilterReduceContext {
	var p = new(MapFilterReduceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_mapFilterReduce
	return p
}

func (*MapFilterReduceContext) IsMapFilterReduceContext() {}

func NewMapFilterReduceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFilterReduceContext {
	var p = new(MapFilterReduceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_mapFilterReduce

	return p
}

func (s *MapFilterReduceContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFilterReduceContext) Slice() ISliceContext {
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

func (s *MapFilterReduceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

func (s *MapFilterReduceContext) Map() IMapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapContext)
}

func (s *MapFilterReduceContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *MapFilterReduceContext) Reduce() IReduceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReduceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReduceContext)
}

func (s *MapFilterReduceContext) Array() IArrayContext {
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

func (s *MapFilterReduceContext) AllMapFilterReduce() []IMapFilterReduceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapFilterReduceContext); ok {
			len++
		}
	}

	tst := make([]IMapFilterReduceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapFilterReduceContext); ok {
			tst[i] = t.(IMapFilterReduceContext)
			i++
		}
	}

	return tst
}

func (s *MapFilterReduceContext) MapFilterReduce(i int) IMapFilterReduceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFilterReduceContext); ok {
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

	return t.(IMapFilterReduceContext)
}

func (s *MapFilterReduceContext) DOT() antlr.TerminalNode {
	return s.GetToken(IronLangParserDOT, 0)
}

func (s *MapFilterReduceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFilterReduceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFilterReduceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterMapFilterReduce(s)
	}
}

func (s *MapFilterReduceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitMapFilterReduce(s)
	}
}

func (s *MapFilterReduceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitMapFilterReduce(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) MapFilterReduce() (localctx IMapFilterReduceContext) {
	return p.mapFilterReduce(0)
}

func (p *IronLangParser) mapFilterReduce(_p int) (localctx IMapFilterReduceContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMapFilterReduceContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMapFilterReduceContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, IronLangParserRULE_mapFilterReduce, _p)

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
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(449)
			p.Slice()
		}

	case 2:
		{
			p.SetState(450)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 3:
		{
			p.SetState(451)
			p.Map()
		}

	case 4:
		{
			p.SetState(452)
			p.Filter()
		}

	case 5:
		{
			p.SetState(453)
			p.Reduce()
		}

	case 6:
		{
			p.SetState(454)
			p.Array()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMapFilterReduceContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mapFilterReduce)
			p.SetState(457)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
			}
			{
				p.SetState(458)
				p.Match(IronLangParserDOT)
			}
			{
				p.SetState(459)
				p.mapFilterReduce(8)
			}

		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
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

func (p *IronLangParser) Map() (localctx IMapContext) {
	this := p
	_ = this

	localctx = NewMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, IronLangParserRULE_map)

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
		p.SetState(465)
		p.Match(IronLangParserMAP)
	}
	{
		p.SetState(466)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(467)
		p.AnonimousFuncWithReturn()
	}

	{
		p.SetState(468)
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
	p.EnterRule(localctx, 62, IronLangParserRULE_filter)

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
		p.SetState(470)
		p.Match(IronLangParserFILTER)
	}
	{
		p.SetState(471)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(472)
		p.AnonimousFuncWithReturn()
	}

	{
		p.SetState(473)
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
	p.EnterRule(localctx, 64, IronLangParserRULE_reduce)

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
		p.SetState(475)
		p.Match(IronLangParserREDUCE)
	}
	{
		p.SetState(476)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(477)
		p.AnonimousFuncWithReturn()
	}

	{
		p.SetState(478)
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

func (s *RelExpressionContext) TYPE_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(IronLangParserTYPE_BOOLEAN, 0)
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
	_startState := 66
	p.EnterRecursionRule(localctx, 66, IronLangParserRULE_relExpression, _p)
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
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(481)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(482)
			p.relExpression(0)
		}
		{
			p.SetState(483)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		{
			p.SetState(485)
			p.Match(IronLangParserNOT)
		}
		{
			p.SetState(486)
			p.relExpression(5)
		}

	case 3:
		{
			p.SetState(487)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 4:
		{
			p.SetState(488)
			p.Atom()
		}

	case 5:
		{
			p.SetState(489)
			p.FuncCall()
		}

	case 6:
		{
			p.SetState(490)
			p.Match(IronLangParserTYPE_BOOLEAN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_relExpression)
			p.SetState(493)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
			}
			{
				p.SetState(494)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(IronLangParserGTEQ-21))|(1<<(IronLangParserLTEQ-21))|(1<<(IronLangParserGT-21))|(1<<(IronLangParserLT-21))|(1<<(IronLangParserDIF-21))|(1<<(IronLangParserEQEQ-21))|(1<<(IronLangParserOR-21))|(1<<(IronLangParserAND-21)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(495)
				p.relExpression(8)
			}

		}
		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, IronLangParserRULE_mathExpression, _p)
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
	p.SetState(512)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(502)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(503)
			p.mathExpression(0)
		}
		{
			p.SetState(504)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		{
			p.SetState(506)
			p.Match(IronLangParserIDENTIFIER)
		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(507)
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
			p.SetState(510)
			p.Atom()
		}

	case 4:
		{
			p.SetState(511)
			p.FuncCall()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(520)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(515)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserMULT || _la == IronLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(516)
					p.mathExpression(7)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(517)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(518)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(519)
					p.mathExpression(6)
				}

			}

		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 70, IronLangParserRULE_atom)
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
		p.SetState(525)
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
	case 29:
		var t *MapFilterReduceContext = nil
		if localctx != nil {
			t = localctx.(*MapFilterReduceContext)
		}
		return p.MapFilterReduce_Sempred(t, predIndex)

	case 33:
		var t *RelExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelExpressionContext)
		}
		return p.RelExpression_Sempred(t, predIndex)

	case 34:
		var t *MathExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MathExpressionContext)
		}
		return p.MathExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *IronLangParser) MapFilterReduce_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *IronLangParser) RelExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

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
