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
		"", "'main'", "'..'", "'.'", "','", "'+'", "'-'", "'*'", "'/'", "'='",
		"'('", "')'", "'{'", "'['", "']'", "'}'", "'++'", "'|'", "'>='", "'<='",
		"'>'", "'<'", "'!='", "'=='", "'->'", "", "", "'fn'", "'if'", "'or'",
		"'in'", "'for'", "'not'", "'and'", "'let'", "'mut'", "'else'", "'while'",
		"'forEach'", "'map'", "'filter'", "'reduce'", "'return'", "'println'",
		"'int'", "'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "L_BRACKET", "R_BRACKET", "R_CURLY", "PLUS_PLUS",
		"PIPE", "GTEQ", "LTEQ", "GT", "LT", "DIF", "EQEQ", "ARROW", "INT_NUMBER",
		"REAL_NUMBER", "FN", "IF", "OR", "IN", "FOR", "NOT", "AND", "LET", "MUT",
		"ELSE", "WHILE", "FOR_EACH", "MAP", "FILTER", "REDUCE", "RETURN", "PRINT_LN",
		"TYPE_INT", "TYPE_FLOAT", "TYPE_BOOLEAN", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"program", "funcMain", "function", "funcType", "return", "scope", "funcCall",
		"funcCallArg", "anonimousFunc", "ifScope", "elseExpression", "elseIfExpression",
		"ifExpression", "loopWhile", "loopForIn", "bodyScope", "println", "variable",
		"functionArgs", "funcArg", "dataTypes", "assignment", "array", "forEach",
		"mapFilterReduce", "map", "filter", "reduce", "relExpression", "mathExpression",
		"atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 445, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 5,
		0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 5, 0, 71, 8, 0, 10, 0,
		12, 0, 74, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 88, 8, 2, 10, 2, 12, 2, 91, 9, 2, 3, 2, 93, 8,
		2, 1, 2, 1, 2, 3, 2, 97, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 106, 8, 3, 10, 3, 12, 3, 109, 9, 3, 3, 3, 111, 8, 3, 1, 3, 1, 3,
		3, 3, 115, 8, 3, 1, 4, 3, 4, 118, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 124,
		8, 5, 10, 5, 12, 5, 127, 9, 5, 1, 5, 3, 5, 130, 8, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 139, 8, 6, 10, 6, 12, 6, 142, 9, 6, 3, 6,
		144, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 151, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 5, 8, 157, 8, 8, 10, 8, 12, 8, 160, 9, 8, 3, 8, 162, 8, 8, 1,
		8, 1, 8, 3, 8, 166, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 171, 8, 8, 1, 8, 3, 8,
		174, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 181, 8, 8, 10, 8, 12, 8,
		184, 9, 8, 3, 8, 186, 8, 8, 1, 8, 1, 8, 3, 8, 190, 8, 8, 1, 8, 1, 8, 3,
		8, 194, 8, 8, 1, 8, 3, 8, 197, 8, 8, 3, 8, 199, 8, 8, 1, 9, 1, 9, 5, 9,
		203, 8, 9, 10, 9, 12, 9, 206, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 219, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 226, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 232, 8, 13, 10, 13, 12, 13, 235, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 248, 8, 14, 1,
		14, 1, 14, 5, 14, 252, 8, 14, 10, 14, 12, 14, 255, 9, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 269, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 275, 8, 16, 1, 16, 1,
		16, 1, 17, 3, 17, 280, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 285, 8, 17, 1,
		18, 1, 18, 1, 18, 5, 18, 290, 8, 18, 10, 18, 12, 18, 293, 9, 18, 1, 19,
		3, 19, 296, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 301, 8, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 334, 8,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 341, 8, 22, 10, 22, 12, 22,
		344, 9, 22, 5, 22, 346, 8, 22, 10, 22, 12, 22, 349, 9, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 366, 8, 23, 1, 23, 1, 23, 3, 23, 370, 8, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 378, 8, 24, 1, 24, 1, 24,
		1, 24, 5, 24, 383, 8, 24, 10, 24, 12, 24, 386, 9, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 413, 8, 28, 1, 28, 1, 28, 1, 28, 5, 28, 418, 8, 28, 10,
		28, 12, 28, 421, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		3, 29, 430, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 438,
		8, 29, 10, 29, 12, 29, 441, 9, 29, 1, 30, 1, 30, 1, 30, 0, 3, 48, 56, 58,
		31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 6, 1, 0, 44, 45,
		1, 0, 25, 26, 3, 0, 18, 23, 29, 29, 33, 33, 1, 0, 7, 8, 1, 0, 5, 6, 2,
		0, 25, 26, 47, 47, 487, 0, 65, 1, 0, 0, 0, 2, 75, 1, 0, 0, 0, 4, 81, 1,
		0, 0, 0, 6, 100, 1, 0, 0, 0, 8, 117, 1, 0, 0, 0, 10, 121, 1, 0, 0, 0, 12,
		133, 1, 0, 0, 0, 14, 150, 1, 0, 0, 0, 16, 198, 1, 0, 0, 0, 18, 200, 1,
		0, 0, 0, 20, 209, 1, 0, 0, 0, 22, 212, 1, 0, 0, 0, 24, 220, 1, 0, 0, 0,
		26, 227, 1, 0, 0, 0, 28, 238, 1, 0, 0, 0, 30, 268, 1, 0, 0, 0, 32, 270,
		1, 0, 0, 0, 34, 279, 1, 0, 0, 0, 36, 286, 1, 0, 0, 0, 38, 300, 1, 0, 0,
		0, 40, 302, 1, 0, 0, 0, 42, 333, 1, 0, 0, 0, 44, 335, 1, 0, 0, 0, 46, 369,
		1, 0, 0, 0, 48, 377, 1, 0, 0, 0, 50, 387, 1, 0, 0, 0, 52, 392, 1, 0, 0,
		0, 54, 397, 1, 0, 0, 0, 56, 412, 1, 0, 0, 0, 58, 429, 1, 0, 0, 0, 60, 442,
		1, 0, 0, 0, 62, 64, 3, 4, 2, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1,
		0, 0, 0, 68, 72, 3, 2, 1, 0, 69, 71, 3, 4, 2, 0, 70, 69, 1, 0, 0, 0, 71,
		74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 1, 1, 0, 0,
		0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 27, 0, 0, 76, 77, 5, 1, 0, 0, 77, 78,
		5, 10, 0, 0, 78, 79, 5, 11, 0, 0, 79, 80, 3, 10, 5, 0, 80, 3, 1, 0, 0,
		0, 81, 82, 5, 27, 0, 0, 82, 83, 5, 47, 0, 0, 83, 92, 5, 10, 0, 0, 84, 89,
		3, 36, 18, 0, 85, 86, 5, 4, 0, 0, 86, 88, 3, 36, 18, 0, 87, 85, 1, 0, 0,
		0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 94, 1, 0, 0, 0, 94, 96, 5, 11, 0, 0, 95, 97, 3, 40, 20, 0, 96, 95,
		1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 3, 10, 5, 0,
		99, 5, 1, 0, 0, 0, 100, 101, 5, 47, 0, 0, 101, 110, 5, 10, 0, 0, 102, 107,
		3, 36, 18, 0, 103, 104, 5, 4, 0, 0, 104, 106, 3, 36, 18, 0, 105, 103, 1,
		0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0,
		0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 102, 1, 0, 0, 0, 110,
		111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 5, 11, 0, 0, 113, 115,
		3, 40, 20, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 7, 1, 0,
		0, 0, 116, 118, 5, 42, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 120, 3, 58, 29, 0, 120, 9, 1, 0, 0, 0, 121,
		125, 5, 12, 0, 0, 122, 124, 3, 30, 15, 0, 123, 122, 1, 0, 0, 0, 124, 127,
		1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 129, 1, 0,
		0, 0, 127, 125, 1, 0, 0, 0, 128, 130, 3, 8, 4, 0, 129, 128, 1, 0, 0, 0,
		129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 5, 15, 0, 0, 132,
		11, 1, 0, 0, 0, 133, 134, 5, 47, 0, 0, 134, 143, 5, 10, 0, 0, 135, 140,
		3, 14, 7, 0, 136, 137, 5, 4, 0, 0, 137, 139, 3, 14, 7, 0, 138, 136, 1,
		0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0,
		0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 135, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 5, 11, 0, 0, 146, 13,
		1, 0, 0, 0, 147, 151, 3, 58, 29, 0, 148, 151, 3, 12, 6, 0, 149, 151, 3,
		16, 8, 0, 150, 147, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 149, 1, 0, 0,
		0, 151, 15, 1, 0, 0, 0, 152, 161, 5, 10, 0, 0, 153, 158, 3, 36, 18, 0,
		154, 155, 5, 4, 0, 0, 155, 157, 3, 36, 18, 0, 156, 154, 1, 0, 0, 0, 157,
		160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 162,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 153, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 5, 11, 0, 0, 164, 166, 3, 40, 20,
		0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		168, 5, 24, 0, 0, 168, 170, 5, 12, 0, 0, 169, 171, 3, 30, 15, 0, 170, 169,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0, 0, 172, 174, 3, 8,
		4, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0,
		175, 199, 5, 15, 0, 0, 176, 185, 5, 10, 0, 0, 177, 182, 3, 36, 18, 0, 178,
		179, 5, 4, 0, 0, 179, 181, 3, 36, 18, 0, 180, 178, 1, 0, 0, 0, 181, 184,
		1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 186, 1, 0,
		0, 0, 184, 182, 1, 0, 0, 0, 185, 177, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 187, 1, 0, 0, 0, 187, 189, 5, 11, 0, 0, 188, 190, 3, 40, 20, 0, 189,
		188, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193,
		5, 24, 0, 0, 192, 194, 3, 30, 15, 0, 193, 192, 1, 0, 0, 0, 193, 194, 1,
		0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 197, 3, 8, 4, 0, 196, 195, 1, 0, 0,
		0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0, 198, 152, 1, 0, 0, 0, 198,
		176, 1, 0, 0, 0, 199, 17, 1, 0, 0, 0, 200, 204, 5, 12, 0, 0, 201, 203,
		3, 30, 15, 0, 202, 201, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1,
		0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 207, 1, 0, 0, 0, 206, 204, 1, 0, 0,
		0, 207, 208, 5, 15, 0, 0, 208, 19, 1, 0, 0, 0, 209, 210, 5, 36, 0, 0, 210,
		211, 3, 18, 9, 0, 211, 21, 1, 0, 0, 0, 212, 213, 5, 36, 0, 0, 213, 214,
		5, 28, 0, 0, 214, 215, 3, 56, 28, 0, 215, 218, 3, 18, 9, 0, 216, 219, 3,
		22, 11, 0, 217, 219, 3, 20, 10, 0, 218, 216, 1, 0, 0, 0, 218, 217, 1, 0,
		0, 0, 218, 219, 1, 0, 0, 0, 219, 23, 1, 0, 0, 0, 220, 221, 5, 28, 0, 0,
		221, 222, 3, 56, 28, 0, 222, 225, 3, 18, 9, 0, 223, 226, 3, 20, 10, 0,
		224, 226, 3, 22, 11, 0, 225, 223, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 225,
		226, 1, 0, 0, 0, 226, 25, 1, 0, 0, 0, 227, 228, 5, 37, 0, 0, 228, 229,
		3, 56, 28, 0, 229, 233, 5, 12, 0, 0, 230, 232, 3, 30, 15, 0, 231, 230,
		1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0,
		0, 0, 234, 236, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 237, 5, 15, 0, 0,
		237, 27, 1, 0, 0, 0, 238, 239, 5, 31, 0, 0, 239, 240, 5, 47, 0, 0, 240,
		247, 5, 30, 0, 0, 241, 248, 5, 47, 0, 0, 242, 243, 5, 10, 0, 0, 243, 244,
		5, 25, 0, 0, 244, 245, 5, 2, 0, 0, 245, 246, 5, 25, 0, 0, 246, 248, 5,
		11, 0, 0, 247, 241, 1, 0, 0, 0, 247, 242, 1, 0, 0, 0, 248, 249, 1, 0, 0,
		0, 249, 253, 5, 12, 0, 0, 250, 252, 3, 30, 15, 0, 251, 250, 1, 0, 0, 0,
		252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		256, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 257, 5, 15, 0, 0, 257, 29,
		1, 0, 0, 0, 258, 269, 3, 34, 17, 0, 259, 269, 3, 42, 21, 0, 260, 269, 3,
		24, 12, 0, 261, 269, 3, 4, 2, 0, 262, 269, 3, 12, 6, 0, 263, 269, 3, 10,
		5, 0, 264, 269, 3, 32, 16, 0, 265, 269, 3, 46, 23, 0, 266, 269, 3, 26,
		13, 0, 267, 269, 3, 28, 14, 0, 268, 258, 1, 0, 0, 0, 268, 259, 1, 0, 0,
		0, 268, 260, 1, 0, 0, 0, 268, 261, 1, 0, 0, 0, 268, 262, 1, 0, 0, 0, 268,
		263, 1, 0, 0, 0, 268, 264, 1, 0, 0, 0, 268, 265, 1, 0, 0, 0, 268, 266,
		1, 0, 0, 0, 268, 267, 1, 0, 0, 0, 269, 31, 1, 0, 0, 0, 270, 271, 5, 43,
		0, 0, 271, 274, 5, 10, 0, 0, 272, 275, 3, 34, 17, 0, 273, 275, 5, 47, 0,
		0, 274, 272, 1, 0, 0, 0, 274, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276,
		277, 5, 11, 0, 0, 277, 33, 1, 0, 0, 0, 278, 280, 5, 35, 0, 0, 279, 278,
		1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 282, 5, 34,
		0, 0, 282, 284, 5, 47, 0, 0, 283, 285, 3, 40, 20, 0, 284, 283, 1, 0, 0,
		0, 284, 285, 1, 0, 0, 0, 285, 35, 1, 0, 0, 0, 286, 291, 3, 38, 19, 0, 287,
		288, 5, 4, 0, 0, 288, 290, 3, 38, 19, 0, 289, 287, 1, 0, 0, 0, 290, 293,
		1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 37, 1, 0,
		0, 0, 293, 291, 1, 0, 0, 0, 294, 296, 5, 35, 0, 0, 295, 294, 1, 0, 0, 0,
		295, 296, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298, 5, 47, 0, 0, 298,
		301, 3, 40, 20, 0, 299, 301, 3, 6, 3, 0, 300, 295, 1, 0, 0, 0, 300, 299,
		1, 0, 0, 0, 301, 39, 1, 0, 0, 0, 302, 303, 7, 0, 0, 0, 303, 41, 1, 0, 0,
		0, 304, 305, 3, 34, 17, 0, 305, 306, 5, 9, 0, 0, 306, 307, 3, 16, 8, 0,
		307, 334, 1, 0, 0, 0, 308, 309, 3, 34, 17, 0, 309, 310, 5, 9, 0, 0, 310,
		311, 3, 58, 29, 0, 311, 334, 1, 0, 0, 0, 312, 313, 5, 47, 0, 0, 313, 314,
		5, 9, 0, 0, 314, 334, 3, 58, 29, 0, 315, 316, 3, 34, 17, 0, 316, 317, 5,
		9, 0, 0, 317, 318, 3, 56, 28, 0, 318, 334, 1, 0, 0, 0, 319, 320, 5, 47,
		0, 0, 320, 321, 5, 9, 0, 0, 321, 334, 3, 56, 28, 0, 322, 323, 3, 34, 17,
		0, 323, 324, 5, 9, 0, 0, 324, 325, 3, 44, 22, 0, 325, 334, 1, 0, 0, 0,
		326, 327, 5, 47, 0, 0, 327, 328, 5, 9, 0, 0, 328, 334, 3, 44, 22, 0, 329,
		330, 3, 34, 17, 0, 330, 331, 5, 9, 0, 0, 331, 332, 3, 48, 24, 0, 332, 334,
		1, 0, 0, 0, 333, 304, 1, 0, 0, 0, 333, 308, 1, 0, 0, 0, 333, 312, 1, 0,
		0, 0, 333, 315, 1, 0, 0, 0, 333, 319, 1, 0, 0, 0, 333, 322, 1, 0, 0, 0,
		333, 326, 1, 0, 0, 0, 333, 329, 1, 0, 0, 0, 334, 43, 1, 0, 0, 0, 335, 336,
		3, 40, 20, 0, 336, 347, 5, 13, 0, 0, 337, 342, 7, 1, 0, 0, 338, 339, 5,
		4, 0, 0, 339, 341, 7, 1, 0, 0, 340, 338, 1, 0, 0, 0, 341, 344, 1, 0, 0,
		0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344,
		342, 1, 0, 0, 0, 345, 337, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345,
		1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 350, 1, 0, 0, 0, 349, 347, 1, 0,
		0, 0, 350, 351, 5, 14, 0, 0, 351, 45, 1, 0, 0, 0, 352, 353, 5, 47, 0, 0,
		353, 354, 5, 3, 0, 0, 354, 355, 5, 38, 0, 0, 355, 356, 5, 10, 0, 0, 356,
		357, 3, 16, 8, 0, 357, 358, 5, 11, 0, 0, 358, 370, 1, 0, 0, 0, 359, 360,
		3, 44, 22, 0, 360, 361, 5, 3, 0, 0, 361, 362, 5, 38, 0, 0, 362, 365, 5,
		10, 0, 0, 363, 366, 3, 16, 8, 0, 364, 366, 5, 47, 0, 0, 365, 363, 1, 0,
		0, 0, 365, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 5, 11, 0, 0,
		368, 370, 1, 0, 0, 0, 369, 352, 1, 0, 0, 0, 369, 359, 1, 0, 0, 0, 370,
		47, 1, 0, 0, 0, 371, 372, 6, 24, -1, 0, 372, 378, 5, 47, 0, 0, 373, 378,
		3, 50, 25, 0, 374, 378, 3, 52, 26, 0, 375, 378, 3, 54, 27, 0, 376, 378,
		3, 44, 22, 0, 377, 371, 1, 0, 0, 0, 377, 373, 1, 0, 0, 0, 377, 374, 1,
		0, 0, 0, 377, 375, 1, 0, 0, 0, 377, 376, 1, 0, 0, 0, 378, 384, 1, 0, 0,
		0, 379, 380, 10, 6, 0, 0, 380, 381, 5, 3, 0, 0, 381, 383, 3, 48, 24, 7,
		382, 379, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384,
		385, 1, 0, 0, 0, 385, 49, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 388, 5,
		39, 0, 0, 388, 389, 5, 10, 0, 0, 389, 390, 3, 16, 8, 0, 390, 391, 5, 11,
		0, 0, 391, 51, 1, 0, 0, 0, 392, 393, 5, 40, 0, 0, 393, 394, 5, 10, 0, 0,
		394, 395, 3, 16, 8, 0, 395, 396, 5, 11, 0, 0, 396, 53, 1, 0, 0, 0, 397,
		398, 5, 41, 0, 0, 398, 399, 5, 10, 0, 0, 399, 400, 3, 16, 8, 0, 400, 401,
		5, 11, 0, 0, 401, 55, 1, 0, 0, 0, 402, 403, 6, 28, -1, 0, 403, 404, 5,
		10, 0, 0, 404, 405, 3, 56, 28, 0, 405, 406, 5, 11, 0, 0, 406, 413, 1, 0,
		0, 0, 407, 408, 5, 32, 0, 0, 408, 413, 3, 56, 28, 4, 409, 413, 3, 60, 30,
		0, 410, 413, 3, 12, 6, 0, 411, 413, 5, 46, 0, 0, 412, 402, 1, 0, 0, 0,
		412, 407, 1, 0, 0, 0, 412, 409, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412,
		411, 1, 0, 0, 0, 413, 419, 1, 0, 0, 0, 414, 415, 10, 6, 0, 0, 415, 416,
		7, 2, 0, 0, 416, 418, 3, 56, 28, 7, 417, 414, 1, 0, 0, 0, 418, 421, 1,
		0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 57, 1, 0, 0,
		0, 421, 419, 1, 0, 0, 0, 422, 423, 6, 29, -1, 0, 423, 424, 5, 10, 0, 0,
		424, 425, 3, 58, 29, 0, 425, 426, 5, 11, 0, 0, 426, 430, 1, 0, 0, 0, 427,
		430, 3, 60, 30, 0, 428, 430, 3, 12, 6, 0, 429, 422, 1, 0, 0, 0, 429, 427,
		1, 0, 0, 0, 429, 428, 1, 0, 0, 0, 430, 439, 1, 0, 0, 0, 431, 432, 10, 5,
		0, 0, 432, 433, 7, 3, 0, 0, 433, 438, 3, 58, 29, 6, 434, 435, 10, 4, 0,
		0, 435, 436, 7, 4, 0, 0, 436, 438, 3, 58, 29, 5, 437, 431, 1, 0, 0, 0,
		437, 434, 1, 0, 0, 0, 438, 441, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439,
		440, 1, 0, 0, 0, 440, 59, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0, 442, 443, 7,
		5, 0, 0, 443, 61, 1, 0, 0, 0, 50, 65, 72, 89, 92, 96, 107, 110, 114, 117,
		125, 129, 140, 143, 150, 158, 161, 165, 170, 173, 182, 185, 189, 193, 196,
		198, 204, 218, 225, 233, 247, 253, 268, 274, 279, 284, 291, 295, 300, 333,
		342, 347, 365, 369, 377, 384, 412, 419, 429, 437, 439,
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
	IronLangParserDOT          = 3
	IronLangParserCOMMA        = 4
	IronLangParserPLUS         = 5
	IronLangParserMINUS        = 6
	IronLangParserMULT         = 7
	IronLangParserDIV          = 8
	IronLangParserEQ           = 9
	IronLangParserL_PAREN      = 10
	IronLangParserR_PAREN      = 11
	IronLangParserL_CURLY      = 12
	IronLangParserL_BRACKET    = 13
	IronLangParserR_BRACKET    = 14
	IronLangParserR_CURLY      = 15
	IronLangParserPLUS_PLUS    = 16
	IronLangParserPIPE         = 17
	IronLangParserGTEQ         = 18
	IronLangParserLTEQ         = 19
	IronLangParserGT           = 20
	IronLangParserLT           = 21
	IronLangParserDIF          = 22
	IronLangParserEQEQ         = 23
	IronLangParserARROW        = 24
	IronLangParserINT_NUMBER   = 25
	IronLangParserREAL_NUMBER  = 26
	IronLangParserFN           = 27
	IronLangParserIF           = 28
	IronLangParserOR           = 29
	IronLangParserIN           = 30
	IronLangParserFOR          = 31
	IronLangParserNOT          = 32
	IronLangParserAND          = 33
	IronLangParserLET          = 34
	IronLangParserMUT          = 35
	IronLangParserELSE         = 36
	IronLangParserWHILE        = 37
	IronLangParserFOR_EACH     = 38
	IronLangParserMAP          = 39
	IronLangParserFILTER       = 40
	IronLangParserREDUCE       = 41
	IronLangParserRETURN       = 42
	IronLangParserPRINT_LN     = 43
	IronLangParserTYPE_INT     = 44
	IronLangParserTYPE_FLOAT   = 45
	IronLangParserTYPE_BOOLEAN = 46
	IronLangParserIDENTIFIER   = 47
	IronLangParserWHITE_SPACE  = 48
)

// IronLangParser rules.
const (
	IronLangParserRULE_program          = 0
	IronLangParserRULE_funcMain         = 1
	IronLangParserRULE_function         = 2
	IronLangParserRULE_funcType         = 3
	IronLangParserRULE_return           = 4
	IronLangParserRULE_scope            = 5
	IronLangParserRULE_funcCall         = 6
	IronLangParserRULE_funcCallArg      = 7
	IronLangParserRULE_anonimousFunc    = 8
	IronLangParserRULE_ifScope          = 9
	IronLangParserRULE_elseExpression   = 10
	IronLangParserRULE_elseIfExpression = 11
	IronLangParserRULE_ifExpression     = 12
	IronLangParserRULE_loopWhile        = 13
	IronLangParserRULE_loopForIn        = 14
	IronLangParserRULE_bodyScope        = 15
	IronLangParserRULE_println          = 16
	IronLangParserRULE_variable         = 17
	IronLangParserRULE_functionArgs     = 18
	IronLangParserRULE_funcArg          = 19
	IronLangParserRULE_dataTypes        = 20
	IronLangParserRULE_assignment       = 21
	IronLangParserRULE_array            = 22
	IronLangParserRULE_forEach          = 23
	IronLangParserRULE_mapFilterReduce  = 24
	IronLangParserRULE_map              = 25
	IronLangParserRULE_filter           = 26
	IronLangParserRULE_reduce           = 27
	IronLangParserRULE_relExpression    = 28
	IronLangParserRULE_mathExpression   = 29
	IronLangParserRULE_atom             = 30
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(62)
				p.Function()
			}

		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(68)
		p.FuncMain()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserFN {
		{
			p.SetState(69)
			p.Function()
		}

		p.SetState(74)
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

func (s *FuncMainContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
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
		p.SetState(75)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(76)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(77)
		p.Match(IronLangParserL_PAREN)
	}
	{
		p.SetState(78)
		p.Match(IronLangParserR_PAREN)
	}
	{
		p.SetState(79)
		p.Scope()
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

func (s *FunctionContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
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
		p.SetState(81)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(82)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(83)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(84)
			p.FunctionArgs()
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(85)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(86)
				p.FunctionArgs()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(94)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(95)
			p.DataTypes()
		}

	}
	{
		p.SetState(98)
		p.Scope()
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
		p.SetState(100)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(101)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(102)
			p.FunctionArgs()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(103)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(104)
				p.FunctionArgs()
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(112)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(113)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserRETURN {
		{
			p.SetState(116)
			p.Match(IronLangParserRETURN)
		}

	}
	{
		p.SetState(119)
		p.mathExpression(0)
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *ScopeContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *ScopeContext) AllBodyScope() []IBodyScopeContext {
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

func (s *ScopeContext) BodyScope(i int) IBodyScopeContext {
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

func (s *ScopeContext) Return() IReturnContext {
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

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitScope(s)
	}
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IronLangParserRULE_scope)
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
		p.SetState(121)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.BodyScope()
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER))) != 0) || _la == IronLangParserRETURN || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(128)
			p.Return()
		}

	}
	{
		p.SetState(131)
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
		p.SetState(133)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(134)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER))) != 0) || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(135)
			p.FuncCallArg()
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(136)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(137)
				p.FuncCallArg()
			}

			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(145)
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

func (s *FuncCallArgContext) AnonimousFunc() IAnonimousFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncContext)
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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.mathExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.FuncCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.AnonimousFunc()
		}

	}

	return localctx
}

// IAnonimousFuncContext is an interface to support dynamic dispatch.
type IAnonimousFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnonimousFuncContext differentiates from other interfaces.
	IsAnonimousFuncContext()
}

type AnonimousFuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnonimousFuncContext() *AnonimousFuncContext {
	var p = new(AnonimousFuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronLangParserRULE_anonimousFunc
	return p
}

func (*AnonimousFuncContext) IsAnonimousFuncContext() {}

func NewAnonimousFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonimousFuncContext {
	var p = new(AnonimousFuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronLangParserRULE_anonimousFunc

	return p
}

func (s *AnonimousFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonimousFuncContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_PAREN, 0)
}

func (s *AnonimousFuncContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_PAREN, 0)
}

func (s *AnonimousFuncContext) ARROW() antlr.TerminalNode {
	return s.GetToken(IronLangParserARROW, 0)
}

func (s *AnonimousFuncContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserL_CURLY, 0)
}

func (s *AnonimousFuncContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(IronLangParserR_CURLY, 0)
}

func (s *AnonimousFuncContext) AllFunctionArgs() []IFunctionArgsContext {
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

func (s *AnonimousFuncContext) FunctionArgs(i int) IFunctionArgsContext {
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

func (s *AnonimousFuncContext) DataTypes() IDataTypesContext {
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

func (s *AnonimousFuncContext) BodyScope() IBodyScopeContext {
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

func (s *AnonimousFuncContext) Return() IReturnContext {
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

func (s *AnonimousFuncContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IronLangParserCOMMA)
}

func (s *AnonimousFuncContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IronLangParserCOMMA, i)
}

func (s *AnonimousFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonimousFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonimousFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.EnterAnonimousFunc(s)
	}
}

func (s *AnonimousFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronLangListener); ok {
		listenerT.ExitAnonimousFunc(s)
	}
}

func (s *AnonimousFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case IronLangVisitor:
		return t.VisitAnonimousFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *IronLangParser) AnonimousFunc() (localctx IAnonimousFuncContext) {
	this := p
	_ = this

	localctx = NewAnonimousFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IronLangParserRULE_anonimousFunc)
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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(153)
				p.FunctionArgs()
			}
			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(154)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(155)
					p.FunctionArgs()
				}

				p.SetState(160)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(163)
			p.Match(IronLangParserR_PAREN)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
			{
				p.SetState(164)
				p.DataTypes()
			}

		}
		{
			p.SetState(167)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(168)
			p.Match(IronLangParserL_CURLY)
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(169)
				p.BodyScope()
			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER))) != 0) || _la == IronLangParserRETURN || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(172)
				p.Return()
			}

		}
		{
			p.SetState(175)
			p.Match(IronLangParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(177)
				p.FunctionArgs()
			}
			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(178)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(179)
					p.FunctionArgs()
				}

				p.SetState(184)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(187)
			p.Match(IronLangParserR_PAREN)
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
			{
				p.SetState(188)
				p.DataTypes()
			}

		}
		{
			p.SetState(191)
			p.Match(IronLangParserARROW)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(192)
				p.BodyScope()
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(195)
				p.Return()
			}

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
	p.EnterRule(localctx, 18, IronLangParserRULE_ifScope)
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
		p.SetState(200)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_CURLY)|(1<<IronLangParserFN)|(1<<IronLangParserIF)|(1<<IronLangParserFOR))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(IronLangParserLET-34))|(1<<(IronLangParserMUT-34))|(1<<(IronLangParserWHILE-34))|(1<<(IronLangParserPRINT_LN-34))|(1<<(IronLangParserTYPE_INT-34))|(1<<(IronLangParserTYPE_FLOAT-34))|(1<<(IronLangParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(201)
			p.BodyScope()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
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
	p.EnterRule(localctx, 20, IronLangParserRULE_elseExpression)

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
		p.SetState(209)
		p.Match(IronLangParserELSE)
	}
	{
		p.SetState(210)
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
	p.EnterRule(localctx, 22, IronLangParserRULE_elseIfExpression)

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
		p.SetState(212)
		p.Match(IronLangParserELSE)
	}
	{
		p.SetState(213)
		p.Match(IronLangParserIF)
	}
	{
		p.SetState(214)
		p.relExpression(0)
	}
	{
		p.SetState(215)
		p.IfScope()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(216)
			p.ElseIfExpression()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(217)
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
	p.EnterRule(localctx, 24, IronLangParserRULE_ifExpression)

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
		p.SetState(220)
		p.Match(IronLangParserIF)
	}
	{
		p.SetState(221)
		p.relExpression(0)
	}
	{
		p.SetState(222)
		p.IfScope()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(223)
			p.ElseExpression()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(224)
			p.ElseIfExpression()
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

func (s *LoopWhileContext) AllBodyScope() []IBodyScopeContext {
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

func (s *LoopWhileContext) BodyScope(i int) IBodyScopeContext {
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
	p.EnterRule(localctx, 26, IronLangParserRULE_loopWhile)
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
		p.SetState(227)
		p.Match(IronLangParserWHILE)
	}
	{
		p.SetState(228)
		p.relExpression(0)
	}
	{
		p.SetState(229)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_CURLY)|(1<<IronLangParserFN)|(1<<IronLangParserIF)|(1<<IronLangParserFOR))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(IronLangParserLET-34))|(1<<(IronLangParserMUT-34))|(1<<(IronLangParserWHILE-34))|(1<<(IronLangParserPRINT_LN-34))|(1<<(IronLangParserTYPE_INT-34))|(1<<(IronLangParserTYPE_FLOAT-34))|(1<<(IronLangParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(230)
			p.BodyScope()
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(236)
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

func (s *LoopForInContext) AllBodyScope() []IBodyScopeContext {
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

func (s *LoopForInContext) BodyScope(i int) IBodyScopeContext {
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
	p.EnterRule(localctx, 28, IronLangParserRULE_loopForIn)
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
		p.SetState(238)
		p.Match(IronLangParserFOR)
	}
	{
		p.SetState(239)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(240)
		p.Match(IronLangParserIN)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserIDENTIFIER:
		{
			p.SetState(241)
			p.Match(IronLangParserIDENTIFIER)
		}

	case IronLangParserL_PAREN:
		{
			p.SetState(242)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(243)
			p.Match(IronLangParserINT_NUMBER)
		}
		{
			p.SetState(244)
			p.Match(IronLangParserT__1)
		}
		{
			p.SetState(245)
			p.Match(IronLangParserINT_NUMBER)
		}
		{
			p.SetState(246)
			p.Match(IronLangParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(249)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_CURLY)|(1<<IronLangParserFN)|(1<<IronLangParserIF)|(1<<IronLangParserFOR))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(IronLangParserLET-34))|(1<<(IronLangParserMUT-34))|(1<<(IronLangParserWHILE-34))|(1<<(IronLangParserPRINT_LN-34))|(1<<(IronLangParserTYPE_INT-34))|(1<<(IronLangParserTYPE_FLOAT-34))|(1<<(IronLangParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(250)
			p.BodyScope()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(256)
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

func (s *BodyScopeContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
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
	p.EnterRule(localctx, 30, IronLangParserRULE_bodyScope)

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

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.IfExpression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(261)
			p.Function()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(262)
			p.FuncCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(263)
			p.Scope()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(264)
			p.Println()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(265)
			p.ForEach()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(266)
			p.LoopWhile()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(267)
			p.LoopForIn()
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
	p.EnterRule(localctx, 32, IronLangParserRULE_println)

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
		p.SetState(270)
		p.Match(IronLangParserPRINT_LN)
	}
	{
		p.SetState(271)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(272)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(273)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(276)
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
	p.EnterRule(localctx, 34, IronLangParserRULE_variable)
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
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(278)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(281)
		p.Match(IronLangParserLET)
	}
	{
		p.SetState(282)
		p.Match(IronLangParserIDENTIFIER)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(283)
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
	p.EnterRule(localctx, 36, IronLangParserRULE_functionArgs)

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
		p.SetState(286)
		p.FuncArg()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(287)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(288)
				p.FuncArg()
			}

		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, IronLangParserRULE_funcArg)
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

	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT {
			{
				p.SetState(294)
				p.Match(IronLangParserMUT)
			}

		}
		{
			p.SetState(297)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(298)
			p.DataTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
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
	p.EnterRule(localctx, 40, IronLangParserRULE_dataTypes)
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
		p.SetState(302)
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

func (s *AssignmentContext) AnonimousFunc() IAnonimousFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncContext)
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

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 42, IronLangParserRULE_assignment)

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

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Variable()
		}
		{
			p.SetState(305)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(306)
			p.AnonimousFunc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Variable()
		}
		{
			p.SetState(309)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(310)
			p.mathExpression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(312)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(313)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(314)
			p.mathExpression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(315)
			p.Variable()
		}
		{
			p.SetState(316)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(317)
			p.relExpression(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(319)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(320)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(321)
			p.relExpression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(322)
			p.Variable()
		}
		{
			p.SetState(323)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(324)
			p.Array()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(326)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(327)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(328)
			p.Array()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(329)
			p.Variable()
		}
		{
			p.SetState(330)
			p.Match(IronLangParserEQ)
		}
		{
			p.SetState(331)
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
	p.EnterRule(localctx, 44, IronLangParserRULE_array)
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
		p.SetState(335)
		p.DataTypes()
	}
	{
		p.SetState(336)
		p.Match(IronLangParserL_BRACKET)
	}
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER {
		{
			p.SetState(337)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(338)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(339)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IronLangParserINT_NUMBER || _la == IronLangParserREAL_NUMBER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(350)
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

func (s *ForEachContext) AnonimousFunc() IAnonimousFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncContext)
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
	p.EnterRule(localctx, 46, IronLangParserRULE_forEach)

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

	p.SetState(369)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(353)
			p.Match(IronLangParserDOT)
		}
		{
			p.SetState(354)
			p.Match(IronLangParserFOR_EACH)
		}
		{
			p.SetState(355)
			p.Match(IronLangParserL_PAREN)
		}

		{
			p.SetState(356)
			p.AnonimousFunc()
		}

		{
			p.SetState(357)
			p.Match(IronLangParserR_PAREN)
		}

	case IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.Array()
		}
		{
			p.SetState(360)
			p.Match(IronLangParserDOT)
		}
		{
			p.SetState(361)
			p.Match(IronLangParserFOR_EACH)
		}
		{
			p.SetState(362)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case IronLangParserL_PAREN:
			{
				p.SetState(363)
				p.AnonimousFunc()
			}

		case IronLangParserIDENTIFIER:
			{
				p.SetState(364)
				p.Match(IronLangParserIDENTIFIER)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(367)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, IronLangParserRULE_mapFilterReduce, _p)

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
	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserIDENTIFIER:
		{
			p.SetState(372)
			p.Match(IronLangParserIDENTIFIER)
		}

	case IronLangParserMAP:
		{
			p.SetState(373)
			p.Map()
		}

	case IronLangParserFILTER:
		{
			p.SetState(374)
			p.Filter()
		}

	case IronLangParserREDUCE:
		{
			p.SetState(375)
			p.Reduce()
		}

	case IronLangParserTYPE_INT, IronLangParserTYPE_FLOAT:
		{
			p.SetState(376)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMapFilterReduceContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mapFilterReduce)
			p.SetState(379)

			if !(p.Precpred(p.GetParserRuleContext(), 6)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
			}
			{
				p.SetState(380)
				p.Match(IronLangParserDOT)
			}
			{
				p.SetState(381)
				p.mapFilterReduce(7)
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
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

func (s *MapContext) AnonimousFunc() IAnonimousFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncContext)
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
	p.EnterRule(localctx, 50, IronLangParserRULE_map)

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
		p.SetState(387)
		p.Match(IronLangParserMAP)
	}
	{
		p.SetState(388)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(389)
		p.AnonimousFunc()
	}

	{
		p.SetState(390)
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

func (s *FilterContext) AnonimousFunc() IAnonimousFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncContext)
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
	p.EnterRule(localctx, 52, IronLangParserRULE_filter)

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
		p.Match(IronLangParserFILTER)
	}
	{
		p.SetState(393)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(394)
		p.AnonimousFunc()
	}

	{
		p.SetState(395)
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

func (s *ReduceContext) AnonimousFunc() IAnonimousFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonimousFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonimousFuncContext)
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
	p.EnterRule(localctx, 54, IronLangParserRULE_reduce)

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
		p.SetState(397)
		p.Match(IronLangParserREDUCE)
	}
	{
		p.SetState(398)
		p.Match(IronLangParserL_PAREN)
	}

	{
		p.SetState(399)
		p.AnonimousFunc()
	}

	{
		p.SetState(400)
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
	_startState := 56
	p.EnterRecursionRule(localctx, 56, IronLangParserRULE_relExpression, _p)
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
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(403)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(404)
			p.relExpression(0)
		}
		{
			p.SetState(405)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		{
			p.SetState(407)
			p.Match(IronLangParserNOT)
		}
		{
			p.SetState(408)
			p.relExpression(4)
		}

	case 3:
		{
			p.SetState(409)
			p.Atom()
		}

	case 4:
		{
			p.SetState(410)
			p.FuncCall()
		}

	case 5:
		{
			p.SetState(411)
			p.Match(IronLangParserTYPE_BOOLEAN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_relExpression)
			p.SetState(414)

			if !(p.Precpred(p.GetParserRuleContext(), 6)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
			}
			{
				p.SetState(415)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-18)&-(0x1f+1)) == 0 && ((1<<uint((_la-18)))&((1<<(IronLangParserGTEQ-18))|(1<<(IronLangParserLTEQ-18))|(1<<(IronLangParserGT-18))|(1<<(IronLangParserLT-18))|(1<<(IronLangParserDIF-18))|(1<<(IronLangParserEQEQ-18))|(1<<(IronLangParserOR-18))|(1<<(IronLangParserAND-18)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(416)
				p.relExpression(7)
			}

		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
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
	_startState := 58
	p.EnterRecursionRule(localctx, 58, IronLangParserRULE_mathExpression, _p)
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
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(423)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(424)
			p.mathExpression(0)
		}
		{
			p.SetState(425)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		{
			p.SetState(427)
			p.Atom()
		}

	case 3:
		{
			p.SetState(428)
			p.FuncCall()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(437)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(432)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserMULT || _la == IronLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(433)
					p.mathExpression(6)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(435)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(436)
					p.mathExpression(5)
				}

			}

		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext())
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

func (s *AtomContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 60, IronLangParserRULE_atom)
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
		p.SetState(442)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(IronLangParserINT_NUMBER-25))|(1<<(IronLangParserREAL_NUMBER-25))|(1<<(IronLangParserIDENTIFIER-25)))) != 0) {
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
	case 24:
		var t *MapFilterReduceContext = nil
		if localctx != nil {
			t = localctx.(*MapFilterReduceContext)
		}
		return p.MapFilterReduce_Sempred(t, predIndex)

	case 28:
		var t *RelExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelExpressionContext)
		}
		return p.RelExpression_Sempred(t, predIndex)

	case 29:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *IronLangParser) RelExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *IronLangParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
