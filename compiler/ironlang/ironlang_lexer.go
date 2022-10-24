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
		"T__0", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "PIPE", "COMPOP", "ARROW",
		"INT_NUMBER", "REAL_NUMBER", "FN", "PRINT_LN", "LET", "MUT", "RETURN",
		"TYPE_INT", "TYPE_FLOAT", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 172, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3,
		14, 98, 8, 14, 1, 15, 1, 15, 1, 15, 1, 16, 4, 16, 104, 8, 16, 11, 16, 12,
		16, 105, 1, 17, 4, 17, 109, 8, 17, 11, 17, 12, 17, 110, 1, 17, 4, 17, 114,
		8, 17, 11, 17, 12, 17, 115, 1, 17, 1, 17, 4, 17, 120, 8, 17, 11, 17, 12,
		17, 121, 3, 17, 124, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		5, 25, 164, 8, 25, 10, 25, 12, 25, 167, 9, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 182, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0,
		3, 60, 1, 0, 0, 0, 5, 62, 1, 0, 0, 0, 7, 64, 1, 0, 0, 0, 9, 66, 1, 0, 0,
		0, 11, 68, 1, 0, 0, 0, 13, 70, 1, 0, 0, 0, 15, 72, 1, 0, 0, 0, 17, 74,
		1, 0, 0, 0, 19, 76, 1, 0, 0, 0, 21, 78, 1, 0, 0, 0, 23, 80, 1, 0, 0, 0,
		25, 82, 1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 97, 1, 0, 0, 0, 31, 99, 1,
		0, 0, 0, 33, 103, 1, 0, 0, 0, 35, 123, 1, 0, 0, 0, 37, 125, 1, 0, 0, 0,
		39, 128, 1, 0, 0, 0, 41, 136, 1, 0, 0, 0, 43, 140, 1, 0, 0, 0, 45, 144,
		1, 0, 0, 0, 47, 151, 1, 0, 0, 0, 49, 155, 1, 0, 0, 0, 51, 161, 1, 0, 0,
		0, 53, 168, 1, 0, 0, 0, 55, 56, 5, 109, 0, 0, 56, 57, 5, 97, 0, 0, 57,
		58, 5, 105, 0, 0, 58, 59, 5, 110, 0, 0, 59, 2, 1, 0, 0, 0, 60, 61, 5, 46,
		0, 0, 61, 4, 1, 0, 0, 0, 62, 63, 5, 44, 0, 0, 63, 6, 1, 0, 0, 0, 64, 65,
		5, 43, 0, 0, 65, 8, 1, 0, 0, 0, 66, 67, 5, 45, 0, 0, 67, 10, 1, 0, 0, 0,
		68, 69, 5, 42, 0, 0, 69, 12, 1, 0, 0, 0, 70, 71, 5, 47, 0, 0, 71, 14, 1,
		0, 0, 0, 72, 73, 5, 61, 0, 0, 73, 16, 1, 0, 0, 0, 74, 75, 5, 40, 0, 0,
		75, 18, 1, 0, 0, 0, 76, 77, 5, 41, 0, 0, 77, 20, 1, 0, 0, 0, 78, 79, 5,
		123, 0, 0, 79, 22, 1, 0, 0, 0, 80, 81, 5, 125, 0, 0, 81, 24, 1, 0, 0, 0,
		82, 83, 5, 43, 0, 0, 83, 84, 5, 43, 0, 0, 84, 26, 1, 0, 0, 0, 85, 86, 5,
		124, 0, 0, 86, 28, 1, 0, 0, 0, 87, 88, 5, 61, 0, 0, 88, 98, 5, 61, 0, 0,
		89, 90, 5, 33, 0, 0, 90, 98, 5, 61, 0, 0, 91, 98, 5, 60, 0, 0, 92, 93,
		5, 60, 0, 0, 93, 98, 5, 61, 0, 0, 94, 98, 5, 62, 0, 0, 95, 96, 5, 62, 0,
		0, 96, 98, 5, 61, 0, 0, 97, 87, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0, 97, 91,
		1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 97, 94, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0,
		98, 30, 1, 0, 0, 0, 99, 100, 5, 45, 0, 0, 100, 101, 5, 62, 0, 0, 101, 32,
		1, 0, 0, 0, 102, 104, 7, 0, 0, 0, 103, 102, 1, 0, 0, 0, 104, 105, 1, 0,
		0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 34, 1, 0, 0, 0,
		107, 109, 7, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 124, 1, 0, 0, 0, 112, 114,
		7, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0,
		0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 3, 3, 1, 0,
		118, 120, 7, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 108,
		1, 0, 0, 0, 123, 113, 1, 0, 0, 0, 124, 36, 1, 0, 0, 0, 125, 126, 5, 102,
		0, 0, 126, 127, 5, 110, 0, 0, 127, 38, 1, 0, 0, 0, 128, 129, 5, 112, 0,
		0, 129, 130, 5, 114, 0, 0, 130, 131, 5, 105, 0, 0, 131, 132, 5, 110, 0,
		0, 132, 133, 5, 116, 0, 0, 133, 134, 5, 108, 0, 0, 134, 135, 5, 110, 0,
		0, 135, 40, 1, 0, 0, 0, 136, 137, 5, 108, 0, 0, 137, 138, 5, 101, 0, 0,
		138, 139, 5, 116, 0, 0, 139, 42, 1, 0, 0, 0, 140, 141, 5, 109, 0, 0, 141,
		142, 5, 117, 0, 0, 142, 143, 5, 116, 0, 0, 143, 44, 1, 0, 0, 0, 144, 145,
		5, 114, 0, 0, 145, 146, 5, 101, 0, 0, 146, 147, 5, 116, 0, 0, 147, 148,
		5, 117, 0, 0, 148, 149, 5, 114, 0, 0, 149, 150, 5, 110, 0, 0, 150, 46,
		1, 0, 0, 0, 151, 152, 5, 105, 0, 0, 152, 153, 5, 110, 0, 0, 153, 154, 5,
		116, 0, 0, 154, 48, 1, 0, 0, 0, 155, 156, 5, 102, 0, 0, 156, 157, 5, 108,
		0, 0, 157, 158, 5, 111, 0, 0, 158, 159, 5, 97, 0, 0, 159, 160, 5, 116,
		0, 0, 160, 50, 1, 0, 0, 0, 161, 165, 7, 1, 0, 0, 162, 164, 7, 2, 0, 0,
		163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 52, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 7,
		3, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 6, 26, 0, 0, 171, 54, 1, 0, 0,
		0, 8, 0, 97, 105, 110, 115, 121, 123, 165, 1, 6, 0, 0,
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
	IronLangLexerR_CURLY     = 12
	IronLangLexerPLUS_PLUS   = 13
	IronLangLexerPIPE        = 14
	IronLangLexerCOMPOP      = 15
	IronLangLexerARROW       = 16
	IronLangLexerINT_NUMBER  = 17
	IronLangLexerREAL_NUMBER = 18
	IronLangLexerFN          = 19
	IronLangLexerPRINT_LN    = 20
	IronLangLexerLET         = 21
	IronLangLexerMUT         = 22
	IronLangLexerRETURN      = 23
	IronLangLexerTYPE_INT    = 24
	IronLangLexerTYPE_FLOAT  = 25
	IronLangLexerIDENTIFIER  = 26
	IronLangLexerWHITE_SPACE = 27
)
