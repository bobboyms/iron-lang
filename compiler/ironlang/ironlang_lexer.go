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
		"')'", "'{'", "'}'", "'++'", "", "", "", "'fn'", "'println'", "'let'",
		"'mut'", "'return'", "'int'", "'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "COMPOP", "INT_NUMBER",
		"REAL_NUMBER", "FN", "PRINT_LN", "LET", "MUT", "RETURN", "TYPE_INT",
		"TYPE_FLOAT", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "COMPOP", "INT_NUMBER",
		"REAL_NUMBER", "FN", "PRINT_LN", "LET", "MUT", "RETURN", "TYPE_INT",
		"TYPE_FLOAT", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 163, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 92, 8, 13, 1, 14, 4, 14, 95, 8, 14,
		11, 14, 12, 14, 96, 1, 15, 4, 15, 100, 8, 15, 11, 15, 12, 15, 101, 1, 15,
		4, 15, 105, 8, 15, 11, 15, 12, 15, 106, 1, 15, 1, 15, 4, 15, 111, 8, 15,
		11, 15, 12, 15, 112, 3, 15, 115, 8, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 5, 23, 155, 8, 23, 10, 23, 12, 23, 158, 9, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1,
		0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 173, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 56, 1, 0, 0, 0, 5, 58, 1, 0, 0, 0, 7,
		60, 1, 0, 0, 0, 9, 62, 1, 0, 0, 0, 11, 64, 1, 0, 0, 0, 13, 66, 1, 0, 0,
		0, 15, 68, 1, 0, 0, 0, 17, 70, 1, 0, 0, 0, 19, 72, 1, 0, 0, 0, 21, 74,
		1, 0, 0, 0, 23, 76, 1, 0, 0, 0, 25, 78, 1, 0, 0, 0, 27, 91, 1, 0, 0, 0,
		29, 94, 1, 0, 0, 0, 31, 114, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 119,
		1, 0, 0, 0, 37, 127, 1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 135, 1, 0, 0,
		0, 43, 142, 1, 0, 0, 0, 45, 146, 1, 0, 0, 0, 47, 152, 1, 0, 0, 0, 49, 159,
		1, 0, 0, 0, 51, 52, 5, 109, 0, 0, 52, 53, 5, 97, 0, 0, 53, 54, 5, 105,
		0, 0, 54, 55, 5, 110, 0, 0, 55, 2, 1, 0, 0, 0, 56, 57, 5, 46, 0, 0, 57,
		4, 1, 0, 0, 0, 58, 59, 5, 44, 0, 0, 59, 6, 1, 0, 0, 0, 60, 61, 5, 43, 0,
		0, 61, 8, 1, 0, 0, 0, 62, 63, 5, 45, 0, 0, 63, 10, 1, 0, 0, 0, 64, 65,
		5, 42, 0, 0, 65, 12, 1, 0, 0, 0, 66, 67, 5, 47, 0, 0, 67, 14, 1, 0, 0,
		0, 68, 69, 5, 61, 0, 0, 69, 16, 1, 0, 0, 0, 70, 71, 5, 40, 0, 0, 71, 18,
		1, 0, 0, 0, 72, 73, 5, 41, 0, 0, 73, 20, 1, 0, 0, 0, 74, 75, 5, 123, 0,
		0, 75, 22, 1, 0, 0, 0, 76, 77, 5, 125, 0, 0, 77, 24, 1, 0, 0, 0, 78, 79,
		5, 43, 0, 0, 79, 80, 5, 43, 0, 0, 80, 26, 1, 0, 0, 0, 81, 82, 5, 61, 0,
		0, 82, 92, 5, 61, 0, 0, 83, 84, 5, 33, 0, 0, 84, 92, 5, 61, 0, 0, 85, 92,
		5, 60, 0, 0, 86, 87, 5, 60, 0, 0, 87, 92, 5, 61, 0, 0, 88, 92, 5, 62, 0,
		0, 89, 90, 5, 62, 0, 0, 90, 92, 5, 61, 0, 0, 91, 81, 1, 0, 0, 0, 91, 83,
		1, 0, 0, 0, 91, 85, 1, 0, 0, 0, 91, 86, 1, 0, 0, 0, 91, 88, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 92, 28, 1, 0, 0, 0, 93, 95, 7, 0, 0, 0, 94, 93, 1,
		0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97,
		30, 1, 0, 0, 0, 98, 100, 7, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0,
		0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 115, 1, 0, 0, 0,
		103, 105, 7, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110,
		3, 3, 1, 0, 109, 111, 7, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0,
		0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 115, 1, 0, 0, 0,
		114, 99, 1, 0, 0, 0, 114, 104, 1, 0, 0, 0, 115, 32, 1, 0, 0, 0, 116, 117,
		5, 102, 0, 0, 117, 118, 5, 110, 0, 0, 118, 34, 1, 0, 0, 0, 119, 120, 5,
		112, 0, 0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 105, 0, 0, 122, 123, 5,
		110, 0, 0, 123, 124, 5, 116, 0, 0, 124, 125, 5, 108, 0, 0, 125, 126, 5,
		110, 0, 0, 126, 36, 1, 0, 0, 0, 127, 128, 5, 108, 0, 0, 128, 129, 5, 101,
		0, 0, 129, 130, 5, 116, 0, 0, 130, 38, 1, 0, 0, 0, 131, 132, 5, 109, 0,
		0, 132, 133, 5, 117, 0, 0, 133, 134, 5, 116, 0, 0, 134, 40, 1, 0, 0, 0,
		135, 136, 5, 114, 0, 0, 136, 137, 5, 101, 0, 0, 137, 138, 5, 116, 0, 0,
		138, 139, 5, 117, 0, 0, 139, 140, 5, 114, 0, 0, 140, 141, 5, 110, 0, 0,
		141, 42, 1, 0, 0, 0, 142, 143, 5, 105, 0, 0, 143, 144, 5, 110, 0, 0, 144,
		145, 5, 116, 0, 0, 145, 44, 1, 0, 0, 0, 146, 147, 5, 102, 0, 0, 147, 148,
		5, 108, 0, 0, 148, 149, 5, 111, 0, 0, 149, 150, 5, 97, 0, 0, 150, 151,
		5, 116, 0, 0, 151, 46, 1, 0, 0, 0, 152, 156, 7, 1, 0, 0, 153, 155, 7, 2,
		0, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0,
		156, 157, 1, 0, 0, 0, 157, 48, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 160,
		7, 3, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 6, 24, 0, 0, 162, 50, 1, 0,
		0, 0, 8, 0, 91, 96, 101, 106, 112, 114, 156, 1, 6, 0, 0,
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
	IronLangLexerCOMPOP      = 14
	IronLangLexerINT_NUMBER  = 15
	IronLangLexerREAL_NUMBER = 16
	IronLangLexerFN          = 17
	IronLangLexerPRINT_LN    = 18
	IronLangLexerLET         = 19
	IronLangLexerMUT         = 20
	IronLangLexerRETURN      = 21
	IronLangLexerTYPE_INT    = 22
	IronLangLexerTYPE_FLOAT  = 23
	IronLangLexerIDENTIFIER  = 24
	IronLangLexerWHITE_SPACE = 25
)
