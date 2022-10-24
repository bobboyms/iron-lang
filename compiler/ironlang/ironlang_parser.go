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
		"program", "funcMain", "function", "return", "scope", "funcCall", "funcCallArg",
		"anonimousFunc", "bodyScope", "println", "variable", "functionArgs",
		"funcArg", "dataTypes", "assignment", "mathExpression", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 210, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 5, 0, 36, 8, 0, 10, 0, 12, 0, 39, 9, 0, 1, 0, 1, 0,
		5, 0, 43, 8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 52,
		8, 1, 10, 1, 12, 1, 55, 9, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 66, 8, 2, 10, 2, 12, 2, 69, 9, 2, 3, 2, 71, 8, 2, 1,
		2, 1, 2, 3, 2, 75, 8, 2, 1, 2, 1, 2, 1, 3, 3, 3, 80, 8, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 5, 4, 86, 8, 4, 10, 4, 12, 4, 89, 9, 4, 1, 4, 3, 4, 92, 8,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 101, 8, 5, 10, 5, 12,
		5, 104, 9, 5, 3, 5, 106, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 112, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 118, 8, 7, 10, 7, 12, 7, 121, 9, 7, 3, 7,
		123, 8, 7, 1, 7, 1, 7, 3, 7, 127, 8, 7, 1, 7, 1, 7, 5, 7, 131, 8, 7, 10,
		7, 12, 7, 134, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 143,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 149, 8, 9, 1, 9, 1, 9, 1, 10, 3, 10,
		154, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 159, 8, 10, 1, 11, 1, 11, 1, 11,
		5, 11, 164, 8, 11, 10, 11, 12, 11, 167, 9, 11, 1, 12, 3, 12, 170, 8, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 181,
		8, 14, 1, 14, 3, 14, 184, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 192, 8, 15, 1, 15, 3, 15, 195, 8, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 5, 15, 203, 8, 15, 10, 15, 12, 15, 206, 9, 15, 1, 16,
		1, 16, 1, 16, 0, 1, 30, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 0, 4, 1, 0, 24, 25, 1, 0, 4, 5, 1, 0, 6, 7, 2, 0, 17,
		18, 26, 26, 226, 0, 37, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 59, 1, 0, 0,
		0, 6, 79, 1, 0, 0, 0, 8, 83, 1, 0, 0, 0, 10, 95, 1, 0, 0, 0, 12, 111, 1,
		0, 0, 0, 14, 113, 1, 0, 0, 0, 16, 142, 1, 0, 0, 0, 18, 144, 1, 0, 0, 0,
		20, 153, 1, 0, 0, 0, 22, 160, 1, 0, 0, 0, 24, 169, 1, 0, 0, 0, 26, 174,
		1, 0, 0, 0, 28, 183, 1, 0, 0, 0, 30, 194, 1, 0, 0, 0, 32, 207, 1, 0, 0,
		0, 34, 36, 3, 4, 2, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		40, 44, 3, 2, 1, 0, 41, 43, 3, 4, 2, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1,
		0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 1, 1, 0, 0, 0, 46,
		44, 1, 0, 0, 0, 47, 48, 5, 19, 0, 0, 48, 49, 5, 1, 0, 0, 49, 53, 5, 9,
		0, 0, 50, 52, 3, 22, 11, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53,
		51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0,
		0, 56, 57, 5, 10, 0, 0, 57, 58, 3, 8, 4, 0, 58, 3, 1, 0, 0, 0, 59, 60,
		5, 19, 0, 0, 60, 61, 5, 26, 0, 0, 61, 70, 5, 9, 0, 0, 62, 67, 3, 22, 11,
		0, 63, 64, 5, 3, 0, 0, 64, 66, 3, 22, 11, 0, 65, 63, 1, 0, 0, 0, 66, 69,
		1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0,
		69, 67, 1, 0, 0, 0, 70, 62, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1,
		0, 0, 0, 72, 74, 5, 10, 0, 0, 73, 75, 3, 26, 13, 0, 74, 73, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 3, 8, 4, 0, 77, 5, 1, 0,
		0, 0, 78, 80, 5, 23, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80,
		81, 1, 0, 0, 0, 81, 82, 3, 30, 15, 0, 82, 7, 1, 0, 0, 0, 83, 87, 5, 11,
		0, 0, 84, 86, 3, 16, 8, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87,
		85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0,
		0, 90, 92, 3, 6, 3, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 94, 5, 12, 0, 0, 94, 9, 1, 0, 0, 0, 95, 96, 5, 26, 0, 0,
		96, 105, 5, 9, 0, 0, 97, 102, 3, 12, 6, 0, 98, 99, 5, 3, 0, 0, 99, 101,
		3, 12, 6, 0, 100, 98, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0,
		0, 0, 102, 103, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0,
		105, 97, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108,
		5, 10, 0, 0, 108, 11, 1, 0, 0, 0, 109, 112, 3, 30, 15, 0, 110, 112, 3,
		10, 5, 0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 13, 1, 0, 0,
		0, 113, 122, 5, 9, 0, 0, 114, 119, 3, 22, 11, 0, 115, 116, 5, 3, 0, 0,
		116, 118, 3, 22, 11, 0, 117, 115, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119,
		117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119,
		1, 0, 0, 0, 122, 114, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 126, 5, 10, 0, 0, 125, 127, 3, 26, 13, 0, 126, 125, 1, 0, 0,
		0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 132, 5, 16, 0, 0, 129,
		131, 3, 16, 8, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 15, 1, 0, 0, 0, 134, 132, 1, 0,
		0, 0, 135, 143, 3, 14, 7, 0, 136, 143, 3, 20, 10, 0, 137, 143, 3, 28, 14,
		0, 138, 143, 3, 4, 2, 0, 139, 143, 3, 10, 5, 0, 140, 143, 3, 8, 4, 0, 141,
		143, 3, 18, 9, 0, 142, 135, 1, 0, 0, 0, 142, 136, 1, 0, 0, 0, 142, 137,
		1, 0, 0, 0, 142, 138, 1, 0, 0, 0, 142, 139, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 141, 1, 0, 0, 0, 143, 17, 1, 0, 0, 0, 144, 145, 5, 20, 0, 0,
		145, 148, 5, 9, 0, 0, 146, 149, 3, 20, 10, 0, 147, 149, 5, 26, 0, 0, 148,
		146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151,
		5, 10, 0, 0, 151, 19, 1, 0, 0, 0, 152, 154, 5, 22, 0, 0, 153, 152, 1, 0,
		0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 5, 21, 0, 0,
		156, 158, 5, 26, 0, 0, 157, 159, 3, 26, 13, 0, 158, 157, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 21, 1, 0, 0, 0, 160, 165, 3, 24, 12, 0, 161, 162,
		5, 3, 0, 0, 162, 164, 3, 24, 12, 0, 163, 161, 1, 0, 0, 0, 164, 167, 1,
		0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 23, 1, 0, 0,
		0, 167, 165, 1, 0, 0, 0, 168, 170, 5, 22, 0, 0, 169, 168, 1, 0, 0, 0, 169,
		170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 5, 26, 0, 0, 172, 173,
		3, 26, 13, 0, 173, 25, 1, 0, 0, 0, 174, 175, 7, 0, 0, 0, 175, 27, 1, 0,
		0, 0, 176, 184, 5, 26, 0, 0, 177, 178, 3, 20, 10, 0, 178, 180, 5, 8, 0,
		0, 179, 181, 3, 14, 7, 0, 180, 179, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		184, 1, 0, 0, 0, 182, 184, 3, 30, 15, 0, 183, 176, 1, 0, 0, 0, 183, 177,
		1, 0, 0, 0, 183, 182, 1, 0, 0, 0, 184, 29, 1, 0, 0, 0, 185, 186, 6, 15,
		-1, 0, 186, 187, 5, 9, 0, 0, 187, 188, 3, 30, 15, 0, 188, 189, 5, 10, 0,
		0, 189, 195, 1, 0, 0, 0, 190, 192, 7, 1, 0, 0, 191, 190, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 3, 32, 16, 0, 194, 185,
		1, 0, 0, 0, 194, 191, 1, 0, 0, 0, 195, 204, 1, 0, 0, 0, 196, 197, 10, 4,
		0, 0, 197, 198, 7, 2, 0, 0, 198, 203, 3, 30, 15, 5, 199, 200, 10, 3, 0,
		0, 200, 201, 7, 1, 0, 0, 201, 203, 3, 30, 15, 4, 202, 196, 1, 0, 0, 0,
		202, 199, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204,
		205, 1, 0, 0, 0, 205, 31, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 7,
		3, 0, 0, 208, 33, 1, 0, 0, 0, 28, 37, 44, 53, 67, 70, 74, 79, 87, 91, 102,
		105, 111, 119, 122, 126, 132, 142, 148, 153, 158, 165, 169, 180, 183, 191,
		194, 202, 204,
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
	IronLangParserRULE_return         = 3
	IronLangParserRULE_scope          = 4
	IronLangParserRULE_funcCall       = 5
	IronLangParserRULE_funcCallArg    = 6
	IronLangParserRULE_anonimousFunc  = 7
	IronLangParserRULE_bodyScope      = 8
	IronLangParserRULE_println        = 9
	IronLangParserRULE_variable       = 10
	IronLangParserRULE_functionArgs   = 11
	IronLangParserRULE_funcArg        = 12
	IronLangParserRULE_dataTypes      = 13
	IronLangParserRULE_assignment     = 14
	IronLangParserRULE_mathExpression = 15
	IronLangParserRULE_atom           = 16
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(34)
				p.Function()
			}

		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(40)
		p.FuncMain()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserFN {
		{
			p.SetState(41)
			p.Function()
		}

		p.SetState(46)
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
		p.SetState(47)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(48)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(49)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(50)
			p.FunctionArgs()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(IronLangParserR_PAREN)
	}
	{
		p.SetState(57)
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
		p.SetState(59)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(60)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(61)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(62)
			p.FunctionArgs()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(63)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(64)
				p.FunctionArgs()
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(72)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(73)
			p.DataTypes()
		}

	}
	{
		p.SetState(76)
		p.Scope()
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
	p.EnterRule(localctx, 6, IronLangParserRULE_return)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserRETURN {
		{
			p.SetState(78)
			p.Match(IronLangParserRETURN)
		}

	}
	{
		p.SetState(81)
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
	p.EnterRule(localctx, 8, IronLangParserRULE_scope)
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
		p.SetState(83)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(84)
				p.BodyScope()
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserPLUS)|(1<<IronLangParserMINUS)|(1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserRETURN)|(1<<IronLangParserIDENTIFIER))) != 0 {
		{
			p.SetState(90)
			p.Return()
		}

	}
	{
		p.SetState(93)
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
	p.EnterRule(localctx, 10, IronLangParserRULE_funcCall)
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
		p.SetState(95)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(96)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserPLUS)|(1<<IronLangParserMINUS)|(1<<IronLangParserL_PAREN)|(1<<IronLangParserINT_NUMBER)|(1<<IronLangParserREAL_NUMBER)|(1<<IronLangParserIDENTIFIER))) != 0 {
		{
			p.SetState(97)
			p.FuncCallArg()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(98)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(99)
				p.FuncCallArg()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(107)
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
	p.EnterRule(localctx, 12, IronLangParserRULE_funcCallArg)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.mathExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.FuncCall()
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

func (s *AnonimousFuncContext) AllBodyScope() []IBodyScopeContext {
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

func (s *AnonimousFuncContext) BodyScope(i int) IBodyScopeContext {
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
	p.EnterRule(localctx, 14, IronLangParserRULE_anonimousFunc)
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
		p.SetState(113)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(114)
			p.FunctionArgs()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IronLangParserCOMMA {
			{
				p.SetState(115)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(116)
				p.FunctionArgs()
			}

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(124)
		p.Match(IronLangParserR_PAREN)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(125)
			p.DataTypes()
		}

	}
	{
		p.SetState(128)
		p.Match(IronLangParserARROW)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(129)
				p.BodyScope()
			}

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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

func (s *BodyScopeContext) AnonimousFunc() IAnonimousFuncContext {
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
	p.EnterRule(localctx, 16, IronLangParserRULE_bodyScope)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.AnonimousFunc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Variable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Function()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.FuncCall()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.Scope()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(141)
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
	p.EnterRule(localctx, 18, IronLangParserRULE_println)

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
		p.SetState(144)
		p.Match(IronLangParserPRINT_LN)
	}
	{
		p.SetState(145)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(146)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(147)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(150)
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
	p.EnterRule(localctx, 20, IronLangParserRULE_variable)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(152)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(155)
		p.Match(IronLangParserLET)
	}
	{
		p.SetState(156)
		p.Match(IronLangParserIDENTIFIER)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(157)
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
	p.EnterRule(localctx, 22, IronLangParserRULE_functionArgs)

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
		p.SetState(160)
		p.FuncArg()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(161)
				p.Match(IronLangParserCOMMA)
			}
			{
				p.SetState(162)
				p.FuncArg()
			}

		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 24, IronLangParserRULE_funcArg)
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
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(168)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(171)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(172)
		p.DataTypes()
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
	p.EnterRule(localctx, 26, IronLangParserRULE_dataTypes)
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
		p.SetState(174)
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

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
}

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
	p.EnterRule(localctx, 28, IronLangParserRULE_assignment)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(IronLangParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Variable()
		}
		{
			p.SetState(178)
			p.Match(IronLangParserEQ)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(179)
				p.AnonimousFunc()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.mathExpression(0)
		}

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
	_startState := 30
	p.EnterRecursionRule(localctx, 30, IronLangParserRULE_mathExpression, _p)
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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserL_PAREN:
		{
			p.SetState(186)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(187)
			p.mathExpression(0)
		}
		{
			p.SetState(188)
			p.Match(IronLangParserR_PAREN)
		}

	case IronLangParserPLUS, IronLangParserMINUS, IronLangParserINT_NUMBER, IronLangParserREAL_NUMBER, IronLangParserIDENTIFIER:
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserPLUS || _la == IronLangParserMINUS {
			{
				p.SetState(190)
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
			p.SetState(193)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(197)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserMULT || _la == IronLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(198)
					p.mathExpression(5)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(200)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(201)
					p.mathExpression(4)
				}

			}

		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 32, IronLangParserRULE_atom)
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
		p.SetState(207)
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
	case 15:
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
