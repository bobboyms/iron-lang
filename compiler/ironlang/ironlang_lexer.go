// Code generated from /Users/thiagoluizrodrigues/go/src/iron-lang/IronLang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package ironlang

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type IronLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ironlanglexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ironlanglexerLexerInit() {
	staticData := &ironlanglexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ",
		"L_PAREN", "R_PAREN", "L_CURLY", "L_BRACKET", "R_BRACKET", "R_CURLY",
		"PLUS_PLUS", "PIPE", "GTEQ", "LTEQ", "GT", "LT", "DIF", "EQEQ", "ARROW",
		"INT_NUMBER", "REAL_NUMBER", "FN", "IF", "OR", "IN", "FOR", "NOT", "AND",
		"LET", "MUT", "ELSE", "WHILE", "FOR_EACH", "MAP", "FILTER", "REDUCE",
		"RETURN", "PRINT_LN", "TYPE_INT", "TYPE_FLOAT", "TYPE_BOOLEAN", "IDENTIFIER",
		"WHITE_SPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 48, 294, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 157, 8, 24, 11, 24, 12,
		24, 158, 1, 25, 4, 25, 162, 8, 25, 11, 25, 12, 25, 163, 1, 25, 4, 25, 167,
		8, 25, 11, 25, 12, 25, 168, 1, 25, 1, 25, 4, 25, 173, 8, 25, 11, 25, 12,
		25, 174, 3, 25, 177, 8, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 3, 45, 282, 8, 45, 1, 46, 1, 46, 5, 46, 286, 8, 46, 10, 46,
		12, 46, 289, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 0, 0, 48, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 1,
		0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 300, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0,
		0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0,
		0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0,
		0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1,
		0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95,
		1, 0, 0, 0, 1, 97, 1, 0, 0, 0, 3, 102, 1, 0, 0, 0, 5, 105, 1, 0, 0, 0,
		7, 107, 1, 0, 0, 0, 9, 109, 1, 0, 0, 0, 11, 111, 1, 0, 0, 0, 13, 113, 1,
		0, 0, 0, 15, 115, 1, 0, 0, 0, 17, 117, 1, 0, 0, 0, 19, 119, 1, 0, 0, 0,
		21, 121, 1, 0, 0, 0, 23, 123, 1, 0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 127,
		1, 0, 0, 0, 29, 129, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 134, 1, 0, 0,
		0, 35, 136, 1, 0, 0, 0, 37, 139, 1, 0, 0, 0, 39, 142, 1, 0, 0, 0, 41, 144,
		1, 0, 0, 0, 43, 146, 1, 0, 0, 0, 45, 149, 1, 0, 0, 0, 47, 152, 1, 0, 0,
		0, 49, 156, 1, 0, 0, 0, 51, 176, 1, 0, 0, 0, 53, 178, 1, 0, 0, 0, 55, 181,
		1, 0, 0, 0, 57, 184, 1, 0, 0, 0, 59, 187, 1, 0, 0, 0, 61, 190, 1, 0, 0,
		0, 63, 194, 1, 0, 0, 0, 65, 198, 1, 0, 0, 0, 67, 202, 1, 0, 0, 0, 69, 206,
		1, 0, 0, 0, 71, 210, 1, 0, 0, 0, 73, 215, 1, 0, 0, 0, 75, 221, 1, 0, 0,
		0, 77, 229, 1, 0, 0, 0, 79, 233, 1, 0, 0, 0, 81, 240, 1, 0, 0, 0, 83, 247,
		1, 0, 0, 0, 85, 254, 1, 0, 0, 0, 87, 262, 1, 0, 0, 0, 89, 266, 1, 0, 0,
		0, 91, 281, 1, 0, 0, 0, 93, 283, 1, 0, 0, 0, 95, 290, 1, 0, 0, 0, 97, 98,
		5, 109, 0, 0, 98, 99, 5, 97, 0, 0, 99, 100, 5, 105, 0, 0, 100, 101, 5,
		110, 0, 0, 101, 2, 1, 0, 0, 0, 102, 103, 5, 46, 0, 0, 103, 104, 5, 46,
		0, 0, 104, 4, 1, 0, 0, 0, 105, 106, 5, 46, 0, 0, 106, 6, 1, 0, 0, 0, 107,
		108, 5, 44, 0, 0, 108, 8, 1, 0, 0, 0, 109, 110, 5, 43, 0, 0, 110, 10, 1,
		0, 0, 0, 111, 112, 5, 45, 0, 0, 112, 12, 1, 0, 0, 0, 113, 114, 5, 42, 0,
		0, 114, 14, 1, 0, 0, 0, 115, 116, 5, 47, 0, 0, 116, 16, 1, 0, 0, 0, 117,
		118, 5, 61, 0, 0, 118, 18, 1, 0, 0, 0, 119, 120, 5, 40, 0, 0, 120, 20,
		1, 0, 0, 0, 121, 122, 5, 41, 0, 0, 122, 22, 1, 0, 0, 0, 123, 124, 5, 123,
		0, 0, 124, 24, 1, 0, 0, 0, 125, 126, 5, 91, 0, 0, 126, 26, 1, 0, 0, 0,
		127, 128, 5, 93, 0, 0, 128, 28, 1, 0, 0, 0, 129, 130, 5, 125, 0, 0, 130,
		30, 1, 0, 0, 0, 131, 132, 5, 43, 0, 0, 132, 133, 5, 43, 0, 0, 133, 32,
		1, 0, 0, 0, 134, 135, 5, 124, 0, 0, 135, 34, 1, 0, 0, 0, 136, 137, 5, 62,
		0, 0, 137, 138, 5, 61, 0, 0, 138, 36, 1, 0, 0, 0, 139, 140, 5, 60, 0, 0,
		140, 141, 5, 61, 0, 0, 141, 38, 1, 0, 0, 0, 142, 143, 5, 62, 0, 0, 143,
		40, 1, 0, 0, 0, 144, 145, 5, 60, 0, 0, 145, 42, 1, 0, 0, 0, 146, 147, 5,
		33, 0, 0, 147, 148, 5, 61, 0, 0, 148, 44, 1, 0, 0, 0, 149, 150, 5, 61,
		0, 0, 150, 151, 5, 61, 0, 0, 151, 46, 1, 0, 0, 0, 152, 153, 5, 45, 0, 0,
		153, 154, 5, 62, 0, 0, 154, 48, 1, 0, 0, 0, 155, 157, 7, 0, 0, 0, 156,
		155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159,
		1, 0, 0, 0, 159, 50, 1, 0, 0, 0, 160, 162, 7, 0, 0, 0, 161, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0,
		164, 177, 1, 0, 0, 0, 165, 167, 7, 0, 0, 0, 166, 165, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 172, 3, 5, 2, 0, 171, 173, 7, 0, 0, 0, 172, 171, 1, 0,
		0, 0, 173, 174, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0,
		175, 177, 1, 0, 0, 0, 176, 161, 1, 0, 0, 0, 176, 166, 1, 0, 0, 0, 177,
		52, 1, 0, 0, 0, 178, 179, 5, 102, 0, 0, 179, 180, 5, 110, 0, 0, 180, 54,
		1, 0, 0, 0, 181, 182, 5, 105, 0, 0, 182, 183, 5, 102, 0, 0, 183, 56, 1,
		0, 0, 0, 184, 185, 5, 111, 0, 0, 185, 186, 5, 114, 0, 0, 186, 58, 1, 0,
		0, 0, 187, 188, 5, 105, 0, 0, 188, 189, 5, 110, 0, 0, 189, 60, 1, 0, 0,
		0, 190, 191, 5, 102, 0, 0, 191, 192, 5, 111, 0, 0, 192, 193, 5, 114, 0,
		0, 193, 62, 1, 0, 0, 0, 194, 195, 5, 110, 0, 0, 195, 196, 5, 111, 0, 0,
		196, 197, 5, 116, 0, 0, 197, 64, 1, 0, 0, 0, 198, 199, 5, 97, 0, 0, 199,
		200, 5, 110, 0, 0, 200, 201, 5, 100, 0, 0, 201, 66, 1, 0, 0, 0, 202, 203,
		5, 108, 0, 0, 203, 204, 5, 101, 0, 0, 204, 205, 5, 116, 0, 0, 205, 68,
		1, 0, 0, 0, 206, 207, 5, 109, 0, 0, 207, 208, 5, 117, 0, 0, 208, 209, 5,
		116, 0, 0, 209, 70, 1, 0, 0, 0, 210, 211, 5, 101, 0, 0, 211, 212, 5, 108,
		0, 0, 212, 213, 5, 115, 0, 0, 213, 214, 5, 101, 0, 0, 214, 72, 1, 0, 0,
		0, 215, 216, 5, 119, 0, 0, 216, 217, 5, 104, 0, 0, 217, 218, 5, 105, 0,
		0, 218, 219, 5, 108, 0, 0, 219, 220, 5, 101, 0, 0, 220, 74, 1, 0, 0, 0,
		221, 222, 5, 102, 0, 0, 222, 223, 5, 111, 0, 0, 223, 224, 5, 114, 0, 0,
		224, 225, 5, 69, 0, 0, 225, 226, 5, 97, 0, 0, 226, 227, 5, 99, 0, 0, 227,
		228, 5, 104, 0, 0, 228, 76, 1, 0, 0, 0, 229, 230, 5, 109, 0, 0, 230, 231,
		5, 97, 0, 0, 231, 232, 5, 112, 0, 0, 232, 78, 1, 0, 0, 0, 233, 234, 5,
		102, 0, 0, 234, 235, 5, 105, 0, 0, 235, 236, 5, 108, 0, 0, 236, 237, 5,
		116, 0, 0, 237, 238, 5, 101, 0, 0, 238, 239, 5, 114, 0, 0, 239, 80, 1,
		0, 0, 0, 240, 241, 5, 114, 0, 0, 241, 242, 5, 101, 0, 0, 242, 243, 5, 100,
		0, 0, 243, 244, 5, 117, 0, 0, 244, 245, 5, 99, 0, 0, 245, 246, 5, 101,
		0, 0, 246, 82, 1, 0, 0, 0, 247, 248, 5, 114, 0, 0, 248, 249, 5, 101, 0,
		0, 249, 250, 5, 116, 0, 0, 250, 251, 5, 117, 0, 0, 251, 252, 5, 114, 0,
		0, 252, 253, 5, 110, 0, 0, 253, 84, 1, 0, 0, 0, 254, 255, 5, 112, 0, 0,
		255, 256, 5, 114, 0, 0, 256, 257, 5, 105, 0, 0, 257, 258, 5, 110, 0, 0,
		258, 259, 5, 116, 0, 0, 259, 260, 5, 108, 0, 0, 260, 261, 5, 110, 0, 0,
		261, 86, 1, 0, 0, 0, 262, 263, 5, 105, 0, 0, 263, 264, 5, 110, 0, 0, 264,
		265, 5, 116, 0, 0, 265, 88, 1, 0, 0, 0, 266, 267, 5, 102, 0, 0, 267, 268,
		5, 108, 0, 0, 268, 269, 5, 111, 0, 0, 269, 270, 5, 97, 0, 0, 270, 271,
		5, 116, 0, 0, 271, 90, 1, 0, 0, 0, 272, 273, 5, 116, 0, 0, 273, 274, 5,
		114, 0, 0, 274, 275, 5, 117, 0, 0, 275, 282, 5, 101, 0, 0, 276, 277, 5,
		102, 0, 0, 277, 278, 5, 97, 0, 0, 278, 279, 5, 108, 0, 0, 279, 280, 5,
		115, 0, 0, 280, 282, 5, 101, 0, 0, 281, 272, 1, 0, 0, 0, 281, 276, 1, 0,
		0, 0, 282, 92, 1, 0, 0, 0, 283, 287, 7, 1, 0, 0, 284, 286, 7, 2, 0, 0,
		285, 284, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287,
		288, 1, 0, 0, 0, 288, 94, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 291, 7,
		3, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 6, 47, 0, 0, 293, 96, 1, 0, 0,
		0, 8, 0, 158, 163, 168, 174, 176, 281, 287, 1, 6, 0, 0,
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

// IronLangLexerInit initializes any static state used to implement IronLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewIronLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func IronLangLexerInit() {
	staticData := &ironlanglexerLexerStaticData
	staticData.once.Do(ironlanglexerLexerInit)
}

// NewIronLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewIronLangLexer(input antlr.CharStream) *IronLangLexer {
	IronLangLexerInit()
	l := new(IronLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ironlanglexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "IronLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// IronLangLexer tokens.
const (
	IronLangLexerT__0         = 1
	IronLangLexerT__1         = 2
	IronLangLexerDOT          = 3
	IronLangLexerCOMMA        = 4
	IronLangLexerPLUS         = 5
	IronLangLexerMINUS        = 6
	IronLangLexerMULT         = 7
	IronLangLexerDIV          = 8
	IronLangLexerEQ           = 9
	IronLangLexerL_PAREN      = 10
	IronLangLexerR_PAREN      = 11
	IronLangLexerL_CURLY      = 12
	IronLangLexerL_BRACKET    = 13
	IronLangLexerR_BRACKET    = 14
	IronLangLexerR_CURLY      = 15
	IronLangLexerPLUS_PLUS    = 16
	IronLangLexerPIPE         = 17
	IronLangLexerGTEQ         = 18
	IronLangLexerLTEQ         = 19
	IronLangLexerGT           = 20
	IronLangLexerLT           = 21
	IronLangLexerDIF          = 22
	IronLangLexerEQEQ         = 23
	IronLangLexerARROW        = 24
	IronLangLexerINT_NUMBER   = 25
	IronLangLexerREAL_NUMBER  = 26
	IronLangLexerFN           = 27
	IronLangLexerIF           = 28
	IronLangLexerOR           = 29
	IronLangLexerIN           = 30
	IronLangLexerFOR          = 31
	IronLangLexerNOT          = 32
	IronLangLexerAND          = 33
	IronLangLexerLET          = 34
	IronLangLexerMUT          = 35
	IronLangLexerELSE         = 36
	IronLangLexerWHILE        = 37
	IronLangLexerFOR_EACH     = 38
	IronLangLexerMAP          = 39
	IronLangLexerFILTER       = 40
	IronLangLexerREDUCE       = 41
	IronLangLexerRETURN       = 42
	IronLangLexerPRINT_LN     = 43
	IronLangLexerTYPE_INT     = 44
	IronLangLexerTYPE_FLOAT   = 45
	IronLangLexerTYPE_BOOLEAN = 46
	IronLangLexerIDENTIFIER   = 47
	IronLangLexerWHITE_SPACE  = 48
)
