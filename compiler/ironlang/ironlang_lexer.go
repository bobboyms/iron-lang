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
		"", "'main'", "','", "'.'", "'+'", "'-'", "'*'", "'/'", "'='", "'('",
		"')'", "'{'", "'}'", "'++'", "", "", "'fn'", "'let'", "'mut'", "'int'",
		"'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "DOT", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "INT_NUMBER", "REAL_NUMBER",
		"FN", "LET", "MUT", "TYPE_INT", "TYPE_FLOAT", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "DOT", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "INT_NUMBER", "REAL_NUMBER",
		"FN", "LET", "MUT", "TYPE_INT", "TYPE_FLOAT", "IDENTIFIER", "WHITE_SPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 130, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 4,
		13, 77, 8, 13, 11, 13, 12, 13, 78, 1, 14, 4, 14, 82, 8, 14, 11, 14, 12,
		14, 83, 1, 14, 4, 14, 87, 8, 14, 11, 14, 12, 14, 88, 1, 14, 1, 14, 4, 14,
		93, 8, 14, 11, 14, 12, 14, 94, 3, 14, 97, 8, 14, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 122,
		8, 20, 10, 20, 12, 20, 125, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 0, 0, 22,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 1, 0, 4, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 135, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 45, 1, 0, 0, 0, 3, 50,
		1, 0, 0, 0, 5, 52, 1, 0, 0, 0, 7, 54, 1, 0, 0, 0, 9, 56, 1, 0, 0, 0, 11,
		58, 1, 0, 0, 0, 13, 60, 1, 0, 0, 0, 15, 62, 1, 0, 0, 0, 17, 64, 1, 0, 0,
		0, 19, 66, 1, 0, 0, 0, 21, 68, 1, 0, 0, 0, 23, 70, 1, 0, 0, 0, 25, 72,
		1, 0, 0, 0, 27, 76, 1, 0, 0, 0, 29, 96, 1, 0, 0, 0, 31, 98, 1, 0, 0, 0,
		33, 101, 1, 0, 0, 0, 35, 105, 1, 0, 0, 0, 37, 109, 1, 0, 0, 0, 39, 113,
		1, 0, 0, 0, 41, 119, 1, 0, 0, 0, 43, 126, 1, 0, 0, 0, 45, 46, 5, 109, 0,
		0, 46, 47, 5, 97, 0, 0, 47, 48, 5, 105, 0, 0, 48, 49, 5, 110, 0, 0, 49,
		2, 1, 0, 0, 0, 50, 51, 5, 44, 0, 0, 51, 4, 1, 0, 0, 0, 52, 53, 5, 46, 0,
		0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 43, 0, 0, 55, 8, 1, 0, 0, 0, 56, 57, 5,
		45, 0, 0, 57, 10, 1, 0, 0, 0, 58, 59, 5, 42, 0, 0, 59, 12, 1, 0, 0, 0,
		60, 61, 5, 47, 0, 0, 61, 14, 1, 0, 0, 0, 62, 63, 5, 61, 0, 0, 63, 16, 1,
		0, 0, 0, 64, 65, 5, 40, 0, 0, 65, 18, 1, 0, 0, 0, 66, 67, 5, 41, 0, 0,
		67, 20, 1, 0, 0, 0, 68, 69, 5, 123, 0, 0, 69, 22, 1, 0, 0, 0, 70, 71, 5,
		125, 0, 0, 71, 24, 1, 0, 0, 0, 72, 73, 5, 43, 0, 0, 73, 74, 5, 43, 0, 0,
		74, 26, 1, 0, 0, 0, 75, 77, 7, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 28, 1, 0, 0, 0, 80,
		82, 7, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0,
		0, 83, 84, 1, 0, 0, 0, 84, 97, 1, 0, 0, 0, 85, 87, 7, 0, 0, 0, 86, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 92, 3, 5, 2, 0, 91, 93, 7, 0, 0, 0, 92, 91, 1,
		0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		97, 1, 0, 0, 0, 96, 81, 1, 0, 0, 0, 96, 86, 1, 0, 0, 0, 97, 30, 1, 0, 0,
		0, 98, 99, 5, 102, 0, 0, 99, 100, 5, 110, 0, 0, 100, 32, 1, 0, 0, 0, 101,
		102, 5, 108, 0, 0, 102, 103, 5, 101, 0, 0, 103, 104, 5, 116, 0, 0, 104,
		34, 1, 0, 0, 0, 105, 106, 5, 109, 0, 0, 106, 107, 5, 117, 0, 0, 107, 108,
		5, 116, 0, 0, 108, 36, 1, 0, 0, 0, 109, 110, 5, 105, 0, 0, 110, 111, 5,
		110, 0, 0, 111, 112, 5, 116, 0, 0, 112, 38, 1, 0, 0, 0, 113, 114, 5, 102,
		0, 0, 114, 115, 5, 108, 0, 0, 115, 116, 5, 111, 0, 0, 116, 117, 5, 97,
		0, 0, 117, 118, 5, 116, 0, 0, 118, 40, 1, 0, 0, 0, 119, 123, 7, 1, 0, 0,
		120, 122, 7, 2, 0, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 42, 1, 0, 0, 0, 125, 123, 1,
		0, 0, 0, 126, 127, 7, 3, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 6, 21, 0,
		0, 129, 44, 1, 0, 0, 0, 7, 0, 78, 83, 88, 94, 96, 123, 1, 6, 0, 0,
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
	IronLangLexerT__1        = 2
	IronLangLexerDOT         = 3
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
	IronLangLexerINT_NUMBER  = 14
	IronLangLexerREAL_NUMBER = 15
	IronLangLexerFN          = 16
	IronLangLexerLET         = 17
	IronLangLexerMUT         = 18
	IronLangLexerTYPE_INT    = 19
	IronLangLexerTYPE_FLOAT  = 20
	IronLangLexerIDENTIFIER  = 21
	IronLangLexerWHITE_SPACE = 22
)
