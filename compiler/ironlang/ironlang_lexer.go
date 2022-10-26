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
		"", "'main'", "'.'", "','", "'+'", "'-'", "'*'", "'/'", "'='", "'('",
		"')'", "'{'", "'['", "']'", "'}'", "'++'", "'|'", "", "'->'", "", "",
		"'fn'", "'println'", "'let'", "'mut'", "'forEach'", "'return'", "'int'",
		"'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "L_BRACKET", "R_BRACKET", "R_CURLY", "PLUS_PLUS",
		"PIPE", "COMPOP", "ARROW", "INT_NUMBER", "REAL_NUMBER", "FN", "PRINT_LN",
		"LET", "MUT", "FOR_EACH", "RETURN", "TYPE_INT", "TYPE_FLOAT", "IDENTIFIER",
		"WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "L_BRACKET", "R_BRACKET", "R_CURLY", "PLUS_PLUS",
		"PIPE", "COMPOP", "ARROW", "INT_NUMBER", "REAL_NUMBER", "FN", "PRINT_LN",
		"LET", "MUT", "FOR_EACH", "RETURN", "TYPE_INT", "TYPE_FLOAT", "IDENTIFIER",
		"WHITE_SPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 190, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		108, 8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 4, 18, 114, 8, 18, 11, 18, 12,
		18, 115, 1, 19, 4, 19, 119, 8, 19, 11, 19, 12, 19, 120, 1, 19, 4, 19, 124,
		8, 19, 11, 19, 12, 19, 125, 1, 19, 1, 19, 4, 19, 130, 8, 19, 11, 19, 12,
		19, 131, 3, 19, 134, 8, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 5, 28, 182, 8,
		28, 10, 28, 12, 28, 185, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 0, 0, 30, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 200, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61, 1, 0, 0, 0, 3, 66,
		1, 0, 0, 0, 5, 68, 1, 0, 0, 0, 7, 70, 1, 0, 0, 0, 9, 72, 1, 0, 0, 0, 11,
		74, 1, 0, 0, 0, 13, 76, 1, 0, 0, 0, 15, 78, 1, 0, 0, 0, 17, 80, 1, 0, 0,
		0, 19, 82, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23, 86, 1, 0, 0, 0, 25, 88,
		1, 0, 0, 0, 27, 90, 1, 0, 0, 0, 29, 92, 1, 0, 0, 0, 31, 95, 1, 0, 0, 0,
		33, 107, 1, 0, 0, 0, 35, 109, 1, 0, 0, 0, 37, 113, 1, 0, 0, 0, 39, 133,
		1, 0, 0, 0, 41, 135, 1, 0, 0, 0, 43, 138, 1, 0, 0, 0, 45, 146, 1, 0, 0,
		0, 47, 150, 1, 0, 0, 0, 49, 154, 1, 0, 0, 0, 51, 162, 1, 0, 0, 0, 53, 169,
		1, 0, 0, 0, 55, 173, 1, 0, 0, 0, 57, 179, 1, 0, 0, 0, 59, 186, 1, 0, 0,
		0, 61, 62, 5, 109, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64, 5, 105, 0, 0, 64,
		65, 5, 110, 0, 0, 65, 2, 1, 0, 0, 0, 66, 67, 5, 46, 0, 0, 67, 4, 1, 0,
		0, 0, 68, 69, 5, 44, 0, 0, 69, 6, 1, 0, 0, 0, 70, 71, 5, 43, 0, 0, 71,
		8, 1, 0, 0, 0, 72, 73, 5, 45, 0, 0, 73, 10, 1, 0, 0, 0, 74, 75, 5, 42,
		0, 0, 75, 12, 1, 0, 0, 0, 76, 77, 5, 47, 0, 0, 77, 14, 1, 0, 0, 0, 78,
		79, 5, 61, 0, 0, 79, 16, 1, 0, 0, 0, 80, 81, 5, 40, 0, 0, 81, 18, 1, 0,
		0, 0, 82, 83, 5, 41, 0, 0, 83, 20, 1, 0, 0, 0, 84, 85, 5, 123, 0, 0, 85,
		22, 1, 0, 0, 0, 86, 87, 5, 91, 0, 0, 87, 24, 1, 0, 0, 0, 88, 89, 5, 93,
		0, 0, 89, 26, 1, 0, 0, 0, 90, 91, 5, 125, 0, 0, 91, 28, 1, 0, 0, 0, 92,
		93, 5, 43, 0, 0, 93, 94, 5, 43, 0, 0, 94, 30, 1, 0, 0, 0, 95, 96, 5, 124,
		0, 0, 96, 32, 1, 0, 0, 0, 97, 98, 5, 61, 0, 0, 98, 108, 5, 61, 0, 0, 99,
		100, 5, 33, 0, 0, 100, 108, 5, 61, 0, 0, 101, 108, 5, 60, 0, 0, 102, 103,
		5, 60, 0, 0, 103, 108, 5, 61, 0, 0, 104, 108, 5, 62, 0, 0, 105, 106, 5,
		62, 0, 0, 106, 108, 5, 61, 0, 0, 107, 97, 1, 0, 0, 0, 107, 99, 1, 0, 0,
		0, 107, 101, 1, 0, 0, 0, 107, 102, 1, 0, 0, 0, 107, 104, 1, 0, 0, 0, 107,
		105, 1, 0, 0, 0, 108, 34, 1, 0, 0, 0, 109, 110, 5, 45, 0, 0, 110, 111,
		5, 62, 0, 0, 111, 36, 1, 0, 0, 0, 112, 114, 7, 0, 0, 0, 113, 112, 1, 0,
		0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0,
		116, 38, 1, 0, 0, 0, 117, 119, 7, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 134, 1, 0,
		0, 0, 122, 124, 7, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		129, 3, 3, 1, 0, 128, 130, 7, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0,
		0, 0, 133, 118, 1, 0, 0, 0, 133, 123, 1, 0, 0, 0, 134, 40, 1, 0, 0, 0,
		135, 136, 5, 102, 0, 0, 136, 137, 5, 110, 0, 0, 137, 42, 1, 0, 0, 0, 138,
		139, 5, 112, 0, 0, 139, 140, 5, 114, 0, 0, 140, 141, 5, 105, 0, 0, 141,
		142, 5, 110, 0, 0, 142, 143, 5, 116, 0, 0, 143, 144, 5, 108, 0, 0, 144,
		145, 5, 110, 0, 0, 145, 44, 1, 0, 0, 0, 146, 147, 5, 108, 0, 0, 147, 148,
		5, 101, 0, 0, 148, 149, 5, 116, 0, 0, 149, 46, 1, 0, 0, 0, 150, 151, 5,
		109, 0, 0, 151, 152, 5, 117, 0, 0, 152, 153, 5, 116, 0, 0, 153, 48, 1,
		0, 0, 0, 154, 155, 5, 102, 0, 0, 155, 156, 5, 111, 0, 0, 156, 157, 5, 114,
		0, 0, 157, 158, 5, 69, 0, 0, 158, 159, 5, 97, 0, 0, 159, 160, 5, 99, 0,
		0, 160, 161, 5, 104, 0, 0, 161, 50, 1, 0, 0, 0, 162, 163, 5, 114, 0, 0,
		163, 164, 5, 101, 0, 0, 164, 165, 5, 116, 0, 0, 165, 166, 5, 117, 0, 0,
		166, 167, 5, 114, 0, 0, 167, 168, 5, 110, 0, 0, 168, 52, 1, 0, 0, 0, 169,
		170, 5, 105, 0, 0, 170, 171, 5, 110, 0, 0, 171, 172, 5, 116, 0, 0, 172,
		54, 1, 0, 0, 0, 173, 174, 5, 102, 0, 0, 174, 175, 5, 108, 0, 0, 175, 176,
		5, 111, 0, 0, 176, 177, 5, 97, 0, 0, 177, 178, 5, 116, 0, 0, 178, 56, 1,
		0, 0, 0, 179, 183, 7, 1, 0, 0, 180, 182, 7, 2, 0, 0, 181, 180, 1, 0, 0,
		0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		58, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 7, 3, 0, 0, 187, 188, 1,
		0, 0, 0, 188, 189, 6, 29, 0, 0, 189, 60, 1, 0, 0, 0, 8, 0, 107, 115, 120,
		125, 131, 133, 183, 1, 6, 0, 0,
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
	IronLangLexerT__0        = 1
	IronLangLexerDOT         = 2
	IronLangLexerCOMMA       = 3
	IronLangLexerPLUS        = 4
	IronLangLexerMINUS       = 5
	IronLangLexerMULT        = 6
	IronLangLexerDIV         = 7
	IronLangLexerEQ          = 8
	IronLangLexerL_PAREN     = 9
	IronLangLexerR_PAREN     = 10
	IronLangLexerL_CURLY     = 11
	IronLangLexerL_BRACKET   = 12
	IronLangLexerR_BRACKET   = 13
	IronLangLexerR_CURLY     = 14
	IronLangLexerPLUS_PLUS   = 15
	IronLangLexerPIPE        = 16
	IronLangLexerCOMPOP      = 17
	IronLangLexerARROW       = 18
	IronLangLexerINT_NUMBER  = 19
	IronLangLexerREAL_NUMBER = 20
	IronLangLexerFN          = 21
	IronLangLexerPRINT_LN    = 22
	IronLangLexerLET         = 23
	IronLangLexerMUT         = 24
	IronLangLexerFOR_EACH    = 25
	IronLangLexerRETURN      = 26
	IronLangLexerTYPE_INT    = 27
	IronLangLexerTYPE_FLOAT  = 28
	IronLangLexerIDENTIFIER  = 29
	IronLangLexerWHITE_SPACE = 30
)
