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
		"')'", "'{'", "'}'", "'++'", "", "", "", "'fn'", "'println'", "'let'",
		"'mut'", "'int'", "'float'",
	}
	staticData.symbolicNames = []string{
		"", "", "DOT", "COMMA", "PLUS", "MINUS", "MULT", "DIV", "EQ", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "PLUS_PLUS", "COMPOP", "INT_NUMBER",
		"REAL_NUMBER", "FN", "PRINT_LN", "LET", "MUT", "TYPE_INT", "TYPE_FLOAT",
		"IDENTIFIER", "WHITE_SPACE",
	}
	staticData.ruleNames = []string{
		"program", "funcMain", "function", "scope", "println", "variable", "functionArgs",
		"funcArg", "dataTypes", "assignment", "mathExpression", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 141, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 5, 0, 33, 8, 0, 10, 0, 12, 0, 36, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		42, 8, 1, 10, 1, 12, 1, 45, 9, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 5, 2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 2, 1, 2, 3, 2, 61, 8, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 71, 8, 3, 10, 3,
		12, 3, 74, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 82, 8, 4, 1,
		4, 1, 4, 1, 5, 3, 5, 87, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 92, 8, 5, 1, 6,
		1, 6, 1, 6, 5, 6, 97, 8, 6, 10, 6, 12, 6, 100, 9, 6, 1, 7, 3, 7, 103, 8,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 112, 8, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 123, 8, 10, 1,
		10, 3, 10, 126, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		134, 8, 10, 10, 10, 12, 10, 137, 9, 10, 1, 11, 1, 11, 1, 11, 0, 1, 20,
		12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 4, 1, 0, 21, 22, 1, 0,
		4, 5, 1, 0, 6, 7, 2, 0, 15, 16, 23, 23, 148, 0, 27, 1, 0, 0, 0, 2, 37,
		1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10,
		86, 1, 0, 0, 0, 12, 93, 1, 0, 0, 0, 14, 102, 1, 0, 0, 0, 16, 107, 1, 0,
		0, 0, 18, 111, 1, 0, 0, 0, 20, 125, 1, 0, 0, 0, 22, 138, 1, 0, 0, 0, 24,
		26, 3, 4, 2, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0,
		0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 34,
		3, 2, 1, 0, 31, 33, 3, 4, 2, 0, 32, 31, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0,
		34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 1, 1, 0, 0, 0, 36, 34, 1, 0,
		0, 0, 37, 38, 5, 17, 0, 0, 38, 39, 5, 1, 0, 0, 39, 43, 5, 9, 0, 0, 40,
		42, 3, 12, 6, 0, 41, 40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0,
		0, 0, 43, 44, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47,
		5, 10, 0, 0, 47, 48, 3, 6, 3, 0, 48, 3, 1, 0, 0, 0, 49, 50, 5, 17, 0, 0,
		50, 51, 5, 23, 0, 0, 51, 55, 5, 9, 0, 0, 52, 54, 3, 12, 6, 0, 53, 52, 1,
		0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56,
		58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 5, 10, 0, 0, 59, 61, 3, 16,
		8, 0, 60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63,
		3, 6, 3, 0, 63, 5, 1, 0, 0, 0, 64, 72, 5, 11, 0, 0, 65, 71, 3, 10, 5, 0,
		66, 71, 3, 18, 9, 0, 67, 71, 3, 4, 2, 0, 68, 71, 3, 6, 3, 0, 69, 71, 3,
		8, 4, 0, 70, 65, 1, 0, 0, 0, 70, 66, 1, 0, 0, 0, 70, 67, 1, 0, 0, 0, 70,
		68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0,
		0, 72, 73, 1, 0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76,
		5, 12, 0, 0, 76, 7, 1, 0, 0, 0, 77, 78, 5, 18, 0, 0, 78, 81, 5, 9, 0, 0,
		79, 82, 3, 10, 5, 0, 80, 82, 5, 23, 0, 0, 81, 79, 1, 0, 0, 0, 81, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 5, 10, 0, 0, 84, 9, 1, 0, 0, 0, 85,
		87, 5, 20, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0,
		0, 0, 88, 89, 5, 19, 0, 0, 89, 91, 5, 23, 0, 0, 90, 92, 3, 16, 8, 0, 91,
		90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 11, 1, 0, 0, 0, 93, 98, 3, 14,
		7, 0, 94, 95, 5, 3, 0, 0, 95, 97, 3, 14, 7, 0, 96, 94, 1, 0, 0, 0, 97,
		100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 13, 1, 0,
		0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 5, 20, 0, 0, 102, 101, 1, 0, 0, 0,
		102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 5, 23, 0, 0, 105,
		106, 3, 16, 8, 0, 106, 15, 1, 0, 0, 0, 107, 108, 7, 0, 0, 0, 108, 17, 1,
		0, 0, 0, 109, 112, 3, 10, 5, 0, 110, 112, 5, 23, 0, 0, 111, 109, 1, 0,
		0, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 5, 8, 0, 0,
		114, 115, 3, 20, 10, 0, 115, 19, 1, 0, 0, 0, 116, 117, 6, 10, -1, 0, 117,
		118, 5, 9, 0, 0, 118, 119, 3, 20, 10, 0, 119, 120, 5, 10, 0, 0, 120, 126,
		1, 0, 0, 0, 121, 123, 7, 1, 0, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 3, 22, 11, 0, 125, 116, 1, 0, 0,
		0, 125, 122, 1, 0, 0, 0, 126, 135, 1, 0, 0, 0, 127, 128, 10, 4, 0, 0, 128,
		129, 7, 2, 0, 0, 129, 134, 3, 20, 10, 5, 130, 131, 10, 3, 0, 0, 131, 132,
		7, 1, 0, 0, 132, 134, 3, 20, 10, 4, 133, 127, 1, 0, 0, 0, 133, 130, 1,
		0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0,
		0, 136, 21, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 7, 3, 0, 0, 139,
		23, 1, 0, 0, 0, 17, 27, 34, 43, 55, 60, 70, 72, 81, 86, 91, 98, 102, 111,
		122, 125, 133, 135,
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
	IronLangParserCOMPOP      = 14
	IronLangParserINT_NUMBER  = 15
	IronLangParserREAL_NUMBER = 16
	IronLangParserFN          = 17
	IronLangParserPRINT_LN    = 18
	IronLangParserLET         = 19
	IronLangParserMUT         = 20
	IronLangParserTYPE_INT    = 21
	IronLangParserTYPE_FLOAT  = 22
	IronLangParserIDENTIFIER  = 23
	IronLangParserWHITE_SPACE = 24
)

// IronLangParser rules.
const (
	IronLangParserRULE_program        = 0
	IronLangParserRULE_funcMain       = 1
	IronLangParserRULE_function       = 2
	IronLangParserRULE_scope          = 3
	IronLangParserRULE_println        = 4
	IronLangParserRULE_variable       = 5
	IronLangParserRULE_functionArgs   = 6
	IronLangParserRULE_funcArg        = 7
	IronLangParserRULE_dataTypes      = 8
	IronLangParserRULE_assignment     = 9
	IronLangParserRULE_mathExpression = 10
	IronLangParserRULE_atom           = 11
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(24)
				p.Function()
			}

		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(30)
		p.FuncMain()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserFN {
		{
			p.SetState(31)
			p.Function()
		}

		p.SetState(36)
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
		p.SetState(37)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(38)
		p.Match(IronLangParserT__0)
	}
	{
		p.SetState(39)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserMUT || _la == IronLangParserIDENTIFIER {
		{
			p.SetState(40)
			p.FunctionArgs()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(46)
		p.Match(IronLangParserR_PAREN)
	}
	{
		p.SetState(47)
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
		p.SetState(49)
		p.Match(IronLangParserFN)
	}
	{
		p.SetState(50)
		p.Match(IronLangParserIDENTIFIER)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(59)
			p.DataTypes()
		}

	}
	{
		p.SetState(62)
		p.Scope()
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

func (s *ScopeContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
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

	return t.(IVariableContext)
}

func (s *ScopeContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *ScopeContext) AllFunction() []IFunctionContext {
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

func (s *ScopeContext) Function(i int) IFunctionContext {
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

func (s *ScopeContext) AllScope() []IScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScopeContext); ok {
			len++
		}
	}

	tst := make([]IScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScopeContext); ok {
			tst[i] = t.(IScopeContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Scope(i int) IScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
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

	return t.(IScopeContext)
}

func (s *ScopeContext) AllPrintln() []IPrintlnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintlnContext); ok {
			len++
		}
	}

	tst := make([]IPrintlnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintlnContext); ok {
			tst[i] = t.(IPrintlnContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Println(i int) IPrintlnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnContext); ok {
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

	return t.(IPrintlnContext)
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
	p.EnterRule(localctx, 6, IronLangParserRULE_scope)
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
		p.SetState(64)
		p.Match(IronLangParserL_CURLY)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IronLangParserL_CURLY)|(1<<IronLangParserFN)|(1<<IronLangParserPRINT_LN)|(1<<IronLangParserLET)|(1<<IronLangParserMUT)|(1<<IronLangParserIDENTIFIER))) != 0 {
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(65)
				p.Variable()
			}

		case 2:
			{
				p.SetState(66)
				p.Assignment()
			}

		case 3:
			{
				p.SetState(67)
				p.Function()
			}

		case 4:
			{
				p.SetState(68)
				p.Scope()
			}

		case 5:
			{
				p.SetState(69)
				p.Println()
			}

		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(IronLangParserR_CURLY)
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
	p.EnterRule(localctx, 8, IronLangParserRULE_println)

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
		p.SetState(77)
		p.Match(IronLangParserPRINT_LN)
	}
	{
		p.SetState(78)
		p.Match(IronLangParserL_PAREN)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(79)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(80)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(83)
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
	p.EnterRule(localctx, 10, IronLangParserRULE_variable)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(85)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(88)
		p.Match(IronLangParserLET)
	}
	{
		p.SetState(89)
		p.Match(IronLangParserIDENTIFIER)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserTYPE_INT || _la == IronLangParserTYPE_FLOAT {
		{
			p.SetState(90)
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
	p.EnterRule(localctx, 12, IronLangParserRULE_functionArgs)
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
		p.SetState(93)
		p.FuncArg()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == IronLangParserCOMMA {
		{
			p.SetState(94)
			p.Match(IronLangParserCOMMA)
		}
		{
			p.SetState(95)
			p.FuncArg()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 14, IronLangParserRULE_funcArg)
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IronLangParserMUT {
		{
			p.SetState(101)
			p.Match(IronLangParserMUT)
		}

	}
	{
		p.SetState(104)
		p.Match(IronLangParserIDENTIFIER)
	}
	{
		p.SetState(105)
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
	p.EnterRule(localctx, 16, IronLangParserRULE_dataTypes)
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
		p.SetState(107)
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

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(IronLangParserEQ, 0)
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

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IronLangParserIDENTIFIER, 0)
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
	p.EnterRule(localctx, 18, IronLangParserRULE_assignment)

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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserLET, IronLangParserMUT:
		{
			p.SetState(109)
			p.Variable()
		}

	case IronLangParserIDENTIFIER:
		{
			p.SetState(110)
			p.Match(IronLangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(113)
		p.Match(IronLangParserEQ)
	}

	{
		p.SetState(114)
		p.mathExpression(0)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, IronLangParserRULE_mathExpression, _p)
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IronLangParserL_PAREN:
		{
			p.SetState(117)
			p.Match(IronLangParserL_PAREN)
		}
		{
			p.SetState(118)
			p.mathExpression(0)
		}
		{
			p.SetState(119)
			p.Match(IronLangParserR_PAREN)
		}

	case IronLangParserPLUS, IronLangParserMINUS, IronLangParserINT_NUMBER, IronLangParserREAL_NUMBER, IronLangParserIDENTIFIER:
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == IronLangParserPLUS || _la == IronLangParserMINUS {
			{
				p.SetState(121)
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
			p.SetState(124)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(128)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserMULT || _la == IronLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(129)
					p.mathExpression(5)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IronLangParserRULE_mathExpression)
				p.SetState(130)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(131)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IronLangParserPLUS || _la == IronLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(132)
					p.mathExpression(4)
				}

			}

		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, IronLangParserRULE_atom)
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
		p.SetState(138)
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
	case 10:
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
