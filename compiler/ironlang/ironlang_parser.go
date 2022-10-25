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
		"", "'main'", "'.'", "','", "'+'", "'-'", "'*'", "'/'", "'='", "'('",
		"')'", "'{'", "'}'", "'++'", "'|'", "", "'->'", "", "", "'fn'", "'println'",
		"'let'", "'mut'", "'return'", "'int'", "'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "PIPE", "COMPOP", "ARROW",
		"INT_NUMBER", "REAL_NUMBER", "FN", "PRINT_LN", "LET", "MUT", "RETURN",
		"TYPE_INT", "TYPE_FLOAT", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"program", "funcMain", "function", "funcType", "return", "scope", "funcCall",
		"funcCallArg", "anonimousFunc", "bodyScope", "println", "variable",
		"functionArgs", "funcArg", "dataTypes", "assignment", "mathExpression",
		"atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 269, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 0, 1, 0, 5, 0, 45, 8, 0, 10, 0, 12, 0, 48, 9, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 54, 8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 68, 8, 2, 10, 2, 12, 2, 71, 9, 2, 3, 2,
		73, 8, 2, 1, 2, 1, 2, 3, 2, 77, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 86, 8, 3, 10, 3, 12, 3, 89, 9, 3, 3, 3, 91, 8, 3, 1, 3,
		1, 3, 3, 3, 95, 8, 3, 1, 4, 3, 4, 98, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5,
		5, 104, 8, 5, 10, 5, 12, 5, 107, 9, 5, 1, 5, 3, 5, 110, 8, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 119, 8, 6, 10, 6, 12, 6, 122, 9,
		6, 3, 6, 124, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 131, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 137, 8, 8, 10, 8, 12, 8, 140, 9, 8, 3, 8, 142,
		8, 8, 1, 8, 1, 8, 3, 8, 146, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 151, 8, 8, 1,
		8, 3, 8, 154, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 161, 8, 8, 10,
		8, 12, 8, 164, 9, 8, 3, 8, 166, 8, 8, 1, 8, 1, 8, 3, 8, 170, 8, 8, 1, 8,
		1, 8, 3, 8, 174, 8, 8, 1, 8, 3, 8, 177, 8, 8, 3, 8, 179, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 187, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 193, 8, 10, 1, 10, 1, 10, 1, 11, 3, 11, 198, 8, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 203, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 208, 8, 12, 10, 12,
		12, 12, 211, 9, 12, 1, 13, 3, 13, 214, 8, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		219, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 227, 8, 15,
		3, 15, 229, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 235, 8, 15, 3, 15,
		237, 8, 15, 3, 15, 239, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		3, 16, 247, 8, 16, 1, 16, 1, 16, 3, 16, 251, 8, 16, 1, 16, 3, 16, 254,
		8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 262, 8, 16, 10,
		16, 12, 16, 265, 9, 16, 1, 17, 1, 17, 1, 17, 0, 1, 32, 18, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 4, 1, 0, 24,
		25, 1, 0, 4, 5, 1, 0, 6, 7, 2, 0, 17, 18, 26, 26, 299, 0, 39, 1, 0, 0,
		0, 2, 49, 1, 0, 0, 0, 4, 61, 1, 0, 0, 0, 6, 80, 1, 0, 0, 0, 8, 97, 1, 0,
		0, 0, 10, 101, 1, 0, 0, 0, 12, 113, 1, 0, 0, 0, 14, 130, 1, 0, 0, 0, 16,
		178, 1, 0, 0, 0, 18, 186, 1, 0, 0, 0, 20, 188, 1, 0, 0, 0, 22, 197, 1,
		0, 0, 0, 24, 204, 1, 0, 0, 0, 26, 218, 1, 0, 0, 0, 28, 220, 1, 0, 0, 0,
		30, 238, 1, 0, 0, 0, 32, 253, 1, 0, 0, 0, 34, 266, 1, 0, 0, 0, 36, 38,
		3, 4, 2, 0, 37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 46, 3,
		2, 1, 0, 43, 45, 3, 4, 2, 0, 44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46,
		44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 1, 1, 0, 0, 0, 48, 46, 1, 0, 0,
		0, 49, 50, 5, 19, 0, 0, 50, 51, 5, 1, 0, 0, 51, 55, 5, 9, 0, 0, 52, 54,
		3, 24, 12, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0,
		0, 55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59,
		5, 10, 0, 0, 59, 60, 3, 10, 5, 0, 60, 3, 1, 0, 0, 0, 61, 62, 5, 19, 0,
		0, 62, 63, 5, 26, 0, 0, 63, 72, 5, 9, 0, 0, 64, 69, 3, 24, 12, 0, 65, 66,
		5, 3, 0, 0, 66, 68, 3, 24, 12, 0, 67, 65, 1, 0, 0, 0, 68, 71, 1, 0, 0,
		0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69,
		1, 0, 0, 0, 72, 64, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 76, 5, 10, 0, 0, 75, 77, 3, 28, 14, 0, 76, 75, 1, 0, 0, 0, 76, 77,
		1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 3, 10, 5, 0, 79, 5, 1, 0, 0, 0,
		80, 81, 5, 26, 0, 0, 81, 90, 5, 9, 0, 0, 82, 87, 3, 24, 12, 0, 83, 84,
		5, 3, 0, 0, 84, 86, 3, 24, 12, 0, 85, 83, 1, 0, 0, 0, 86, 89, 1, 0, 0,
		0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87,
		1, 0, 0, 0, 90, 82, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0,
		92, 94, 5, 10, 0, 0, 93, 95, 3, 28, 14, 0, 94, 93, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 7, 1, 0, 0, 0, 96, 98, 5, 23, 0, 0, 97, 96, 1, 0, 0, 0,
		97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 3, 32, 16, 0, 100, 9,
		1, 0, 0, 0, 101, 105, 5, 11, 0, 0, 102, 104, 3, 18, 9, 0, 103, 102, 1,
		0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0,
		0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 110, 3, 8, 4, 0, 109,
		108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112,
		5, 12, 0, 0, 112, 11, 1, 0, 0, 0, 113, 114, 5, 26, 0, 0, 114, 123, 5, 9,
		0, 0, 115, 120, 3, 14, 7, 0, 116, 117, 5, 3, 0, 0, 117, 119, 3, 14, 7,
		0, 118, 116, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 115,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 5, 10,
		0, 0, 126, 13, 1, 0, 0, 0, 127, 131, 3, 32, 16, 0, 128, 131, 3, 12, 6,
		0, 129, 131, 3, 16, 8, 0, 130, 127, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130,
		129, 1, 0, 0, 0, 131, 15, 1, 0, 0, 0, 132, 141, 5, 9, 0, 0, 133, 138, 3,
		24, 12, 0, 134, 135, 5, 3, 0, 0, 135, 137, 3, 24, 12, 0, 136, 134, 1, 0,
		0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 133, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 5, 10, 0, 0, 144, 146,
		3, 28, 14, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1,
		0, 0, 0, 147, 148, 5, 16, 0, 0, 148, 150, 5, 11, 0, 0, 149, 151, 3, 18,
		9, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0,
		152, 154, 3, 8, 4, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 179, 5, 12, 0, 0, 156, 165, 5, 9, 0, 0, 157, 162,
		3, 24, 12, 0, 158, 159, 5, 3, 0, 0, 159, 161, 3, 24, 12, 0, 160, 158, 1,
		0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0,
		0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 157, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 5, 10, 0, 0, 168, 170,
		3, 28, 14, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1,
		0, 0, 0, 171, 173, 5, 16, 0, 0, 172, 174, 3, 18, 9, 0, 173, 172, 1, 0,
		0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 177, 3, 8, 4, 0,
		176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178,
		132, 1, 0, 0, 0, 178, 156, 1, 0, 0, 0, 179, 17, 1, 0, 0, 0, 180, 187, 3,
		22, 11, 0, 181, 187, 3, 30, 15, 0, 182, 187, 3, 4, 2, 0, 183, 187, 3, 12,
		6, 0, 184, 187, 3, 10, 5, 0, 185, 187, 3, 20, 10, 0, 186, 180, 1, 0, 0,
		0, 186, 181, 1, 0, 0, 0, 186, 182, 1, 0, 0, 0, 186, 183, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 19, 1, 0, 0, 0, 188, 189, 5,
		20, 0, 0, 189, 192, 5, 9, 0, 0, 190, 193, 3, 22, 11, 0, 191, 193, 5, 26,
		0, 0, 192, 190, 1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 195, 5, 10, 0, 0, 195, 21, 1, 0, 0, 0, 196, 198, 5, 22, 0, 0, 197,
		196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200,
		5, 21, 0, 0, 200, 202, 5, 26, 0, 0, 201, 203, 3, 28, 14, 0, 202, 201, 1,
		0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 23, 1, 0, 0, 0, 204, 209, 3, 26, 13,
		0, 205, 206, 5, 3, 0, 0, 206, 208, 3, 26, 13, 0, 207, 205, 1, 0, 0, 0,
		208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210,
		25, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 214, 5, 22, 0, 0, 213, 212,
		1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 5, 26,
		0, 0, 216, 219, 3, 28, 14, 0, 217, 219, 3, 6, 3, 0, 218, 213, 1, 0, 0,
		0, 218, 217, 1, 0, 0, 0, 219, 27, 1, 0, 0, 0, 220, 221, 7, 0, 0, 0, 221,
		29, 1, 0, 0, 0, 222, 223, 3, 22, 11, 0, 223, 228, 5, 8, 0, 0, 224, 229,
		3, 32, 16, 0, 225, 227, 3, 16, 8, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1,
		0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 224, 1, 0, 0, 0, 228, 226, 1, 0, 0,
		0, 229, 239, 1, 0, 0, 0, 230, 231, 5, 26, 0, 0, 231, 236, 5, 8, 0, 0, 232,
		237, 3, 32, 16, 0, 233, 235, 3, 16, 8, 0, 234, 233, 1, 0, 0, 0, 234, 235,
		1, 0, 0, 0, 235, 237, 1, 0, 0, 0, 236, 232, 1, 0, 0, 0, 236, 234, 1, 0,
		0, 0, 237, 239, 1, 0, 0, 0, 238, 222, 1, 0, 0, 0, 238, 230, 1, 0, 0, 0,
		239, 31, 1, 0, 0, 0, 240, 241, 6, 16, -1, 0, 241, 242, 5, 9, 0, 0, 242,
		243, 3, 32, 16, 0, 243, 244, 5, 10, 0, 0, 244, 254, 1, 0, 0, 0, 245, 247,
		7, 1, 0, 0, 246, 245, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 1, 0,
		0, 0, 248, 254, 3, 34, 17, 0, 249, 251, 7, 1, 0, 0, 250, 249, 1, 0, 0,
		0, 250, 251, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 254, 3, 12, 6, 0, 253,
		240, 1, 0, 0, 0, 253, 246, 1, 0, 0, 0, 253, 250, 1, 0, 0, 0, 254, 263,
		1, 0, 0, 0, 255, 256, 10, 5, 0, 0, 256, 257, 7, 2, 0, 0, 257, 262, 3, 32,
		16, 6, 258, 259, 10, 4, 0, 0, 259, 260, 7, 1, 0, 0, 260, 262, 3, 32, 16,
		5, 261, 255, 1, 0, 0, 0, 261, 258, 1, 0, 0, 0, 262, 265, 1, 0, 0, 0, 263,
		261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 33, 1, 0, 0, 0, 265, 263, 1,
		0, 0, 0, 266, 267, 7, 3, 0, 0, 267, 35, 1, 0, 0, 0, 43, 39, 46, 55, 69,
		72, 76, 87, 90, 94, 97, 105, 109, 120, 123, 130, 138, 141, 145, 150, 153,
		162, 165, 169, 173, 176, 178, 186, 192, 197, 202, 209, 213, 218, 226, 228,
		234, 236, 238, 246, 250, 253, 261, 263,
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
	IronLangParserEOF         = antlr.TokenEOF
	IronLangParserT__0        = 1
	IronLangParserDOT         = 2
	IronLangParserCOMMA       = 3
	IronLangParserPLUS        = 4
	IronLangParserMINUS       = 5
	IronLangParserMULT        = 6
	IronLangParserDIV         = 7
	IronLangParserEQ          = 8
	IronLangParserL_PAREN     = 9
	IronLangParserR_PAREN     = 10
	IronLangParserL_CURLY     = 11
	IronLangParserR_CURLY     = 12
	IronLangParserPLUS_PLUS   = 13
	IronLangParserPIPE        = 14
	IronLangParserCOMPOP      = 15
	IronLangParserARROW       = 16
	IronLangParserINT_NUMBER  = 17
	IronLangParserREAL_NUMBER = 18
	IronLangParserFN          = 19
	IronLangParserPRINT_LN    = 20
	IronLangParserLET         = 21
	IronLangParserMUT         = 22
	IronLangParserRETURN      = 23
	IronLangParserTYPE_INT    = 24
	IronLangParserTYPE_FLOAT  = 25
	IronLangParserIDENTIFIER  = 26
	IronLangParserWHITE_SPACE = 27
)

// IronLangParser rules.
const (
	IronLangParserRULE_program        = 0
	IronLangParserRULE_funcMain       = 1
	IronLangParserRULE_function       = 2
	IronLangParserRULE_funcType       = 3
	IronLangParserRULE_return         = 4
	IronLangParserRULE_scope          = 5
	IronLangParserRULE_funcCall       = 6
	IronLangParserRULE_funcCallArg    = 7
	IronLangParserRULE_anonimousFunc  = 8
	IronLangParserRULE_bodyScope      = 9
	IronLangParserRULE_println        = 10
	IronLangParserRULE_variable       = 11
	IronLangParserRULE_functionArgs   = 12
	IronLangParserRULE_funcArg        = 13
	IronLangParserRULE_dataTypes      = 14
	IronLangParserRULE_assignment     = 15
	IronLangParserRULE_mathExpression = 16
	IronLangParserRULE_atom           = 17
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(36)
				p.Function()
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(42)
		p.FuncMain()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserFN {
		{
			p.SetState(43)
			p.Function()
		}

		p.SetState(48)
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

func (s *FuncMainContext) AllFunctionArgs() []IFunctionArgsContext {
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

func (s *FuncMainContext) FunctionArgs(i int) IFunctionArgsContext {
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
		p.SetState(49)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(50)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(51)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(52)
			p.FunctionArgs()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(IronLangParserR_PAREN)
	}
	{
		p.SetState(59)
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
		p.SetState(61)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(62)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(63)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(64)
			p.FunctionArgs()
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(65)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(66)
				p.FunctionArgs()
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(74)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(75)
			p.DataTypes()
		}

	}
	{
		p.SetState(78)
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
		p.SetState(80)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(81)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(82)
			p.FunctionArgs()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(83)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(84)
				p.FunctionArgs()
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(92)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(93)
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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserRETURN {
		{
			p.SetState(96)
			p.Match(IronLangParserRETURN)
		}

	}
	{
		p.SetState(99)
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
		p.SetState(101)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(102)
				p.BodyScope()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserPLUS)|(1<<IronLangParserMINUS)|(1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserRETURN)|(1<<IronLangParserIDENTIFIER))) != 0 {
		{
			p.SetState(108)
			p.Return()
		}

	}
	{
		p.SetState(111)
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
		p.SetState(113)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(114)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserPLUS)|(1<<IronLangParserMINUS)|(1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserIDENTIFIER))) != 0 {
		{
			p.SetState(115)
			p.FuncCallArg()
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(116)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(117)
				p.FuncCallArg()
			}

			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(125)
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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.mathExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.FuncCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(133)
				p.FunctionArgs()
			}
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(134)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(135)
					p.FunctionArgs()
				}

				p.SetState(140)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(143)
			p.Match(IronLangParserR_PAREN)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
			{
				p.SetState(144)
				p.DataTypes()
			}

		}
		{
			p.SetState(147)
			p.Match(IronLangParserARROW)
		}
		{
			p.SetState(148)
			p.Match(IronLangParserL_CURLY)
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(149)
				p.BodyScope()
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserPLUS)|(1<<IronLangParserMINUS)|(1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserRETURN)|(1<<IronLangParserIDENTIFIER))) != 0 {
			{
				p.SetState(152)
				p.Return()
			}

		}
		{
			p.SetState(155)
			p.Match(IronLangParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(IronLangParserL_PAREN)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
			{
				p.SetState(157)
				p.FunctionArgs()
			}
			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == IronLangParserCOMMA {
				{
					p.SetState(158)
					p.Match(IronLangParserCOMMA)
				}
				{
					p.SetState(159)
					p.FunctionArgs()
				}

				p.SetState(164)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(167)
			p.Match(IronLangParserR_PAREN)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
			{
				p.SetState(168)
				p.DataTypes()
			}

		}
		{
			p.SetState(171)
			p.Match(IronLangParserARROW)
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(172)
				p.BodyScope()
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(175)
				p.Return()
			}

		}

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
	p.EnterRule(localctx, 18, IronLangParserRULE_bodyScope)

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

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Function()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.FuncCall()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Scope()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)
			p.Println()
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
	p.EnterRule(localctx, 20, IronLangParserRULE_println)

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
		p.SetState(188)
		p.Match(IronLangParserPRINT_LN)
	}
	{
		p.SetState(189)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(190)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(191)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(194)
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
	p.EnterRule(localctx, 22, IronLangParserRULE_variable)
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

	if _la == IronLangParserMUT {
		{
			p.SetState(196)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(199)
		p.Match(IronLangParserLET)
	}
	{
		p.SetState(200)
		p.Match(IronLangParserIDENTIFIER)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(201)
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
	p.EnterRule(localctx, 24, IronLangParserRULE_functionArgs)

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
		p.SetState(204)
		p.FuncArg()
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(205)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(206)
				p.FuncArg()
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, IronLangParserRULE_funcArg)
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

	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserMUT {
			{
				p.SetState(212)
				p.Match(IronLangParserMUT)
			}

		}
		{
			p.SetState(215)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(216)
			p.DataTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
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
	p.EnterRule(localctx, 28, IronLangParserRULE_dataTypes)
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
		p.SetState(220)
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

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 30, IronLangParserRULE_assignment)

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

	p.SetState(238)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Variable()
		}
		{
			p.SetState(223)
			p.Match(IronLangParserEQ)
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(224)
				p.mathExpression(0)
			}

		case 2:
			p.SetState(226)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(225)
					p.AnonimousFunc()
				}

			}

		}

	case IronLangParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Match(IronLangParserIDENTIFIER)
		}
		{
			p.SetState(231)
			p.Match(IronLangParserEQ)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(232)
				p.mathExpression(0)
			}

		case 2:
			p.SetState(234)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(233)
					p.AnonimousFunc()
				}

			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *MathExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(IronLangParserPLUS, 0)
}

func (s *MathExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(IronLangParserMINUS, 0)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, IronLangParserRULE_mathExpression, _p)
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
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(241)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(242)
			p.mathExpression(0)
		}
		{
			p.SetState(243)
			p.Match(IronLangParserR_PAREN)
		}

	case 2:
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserPLUS || _la == IronLangParserMINUS {
			{
				p.SetState(245)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(248)
			p.Atom()
		}

	case 3:
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserPLUS || _la == IronLangParserMINUS {
			{
				p.SetState(249)
				_la = p.GetTokenStream().LA(1)

				if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(252)
			p.FuncCall()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(261)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(256)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserMULT || _la == IronLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(257)
					p.mathExpression(6)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(259)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(260)
					p.mathExpression(5)
				}

			}

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, IronLangParserRULE_atom)
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
		p.SetState(266)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserIDENTIFIER))) != 0) {
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
	case 16:
		var t *MathExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MathExpressionContext)
		}
		return p.MathExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *IronLangParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
