package codegenerator

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
	"iron-lang/compiler/utils"
	"strings"
)

type ClangPlus struct {
	StrBuilder       *StringBuilder
	StrImportBuilder *StringBuilder
	StrFuncBuilder   *StringBuilder
	TempArrCount     int
	ScopeLog         *scopes.ScopeLog
	ScopeId          int
	MapId            int
	LastIdentifier   string
	//MapTempName      string
	VectorTempName string
}

func NewClang(scopeLog *scopes.ScopeLog) *ClangPlus {
	return &ClangPlus{
		StrBuilder:       NewStringBuilder(),
		StrImportBuilder: NewStringBuilder(),
		StrFuncBuilder:   NewStringBuilder(),
		TempArrCount:     -1,
		ScopeLog:         scopeLog,
		ScopeId:          0,
		MapId:            0,
	}
}

func (l *ClangPlus) InitForEachFunction() {
	l.StrFuncBuilder.WriteString("\n")
	l.StrFuncBuilder.WriteString("template <typename TTX>\n")
	l.StrFuncBuilder.WriteString("void arrForEach(void (*func)(TTX), TTX val) {\n")
	l.StrFuncBuilder.WriteString("func(val);\n")
	l.StrFuncBuilder.WriteString("}\n")
}

func (l *ClangPlus) Visit(tree antlr.ParseTree) {
	switch val := tree.(type) {
	case *ironlang.ProgramContext:
		l.VisitProgram(val)
	case *ironlang.FuncMainContext:
		l.VisitFuncMain(val)
	case *ironlang.ScopeContext:
		l.VisitScope(val)
	case *ironlang.MathExpressionContext:
		l.VisitMathExpression(val)
	case *ironlang.AssignmentContext:
		l.VisitAssignment(val)
	case *ironlang.AtomContext:
		l.VisitAtom(val)
	case *ironlang.VariableContext:
		l.VisitVariable(val)
	case *ironlang.DataTypesContext:
		l.VisitDataTypes(val)
	case *ironlang.FunctionContext:
		l.VisitFunction(val)
	case *ironlang.ReturnContext:
		l.VisitReturn(val)
	case *ironlang.FuncCallContext:
		l.VisitFuncCall(val)
	case *ironlang.FuncCallArgContext:
		l.VisitFuncCallArg(val)
	case *ironlang.BodyScopeContext:
		l.VisitBodyScope(val)
	case *ironlang.AnonimousFuncContext:
		l.VisitAnonymousFunc(val)
	case *ironlang.FunctionArgsContext:
		l.VisitFunctionArgs(val)
	case *ironlang.FuncArgContext:
		l.VisitFuncArg(val)
	case *ironlang.FuncTypeContext:
		l.VisitFuncType(val)
	case *ironlang.ArrayContext:
		l.VisitArray(val)
	case *ironlang.ForEachContext:
		l.VisitForEach(val)
	case *ironlang.MapFilterReduceContext:
		l.VisitMapFilterReduce(val)
	case *ironlang.MapContext:
		l.VisitMap(val)
	case *ironlang.FilterContext:
		l.VisitFilter(val)
	case *ironlang.ReduceContext:
		l.VisitReduce(val)
	case *ironlang.RelExpressionContext:
		l.VisitRelExpression(val)
	case *ironlang.IfScopeContext:
		l.VisitIfScope(val)
	case *ironlang.IfExpressionContext:
		l.VisitIfExpression(val)
	case *ironlang.ElseExpressionContext:
		l.VisitElseExpression(val)
	case *ironlang.ElseIfExpressionContext:
		l.VisitElseIfExpression(val)

	default:
		panic("Unknown context")
	}
}

func (l *ClangPlus) GetBuilder() *strings.Builder {
	builder := &strings.Builder{}
	builder.WriteString(l.StrImportBuilder.GetBuilder().String())
	if l.StrFuncBuilder != nil {
		builder.WriteString(l.StrFuncBuilder.GetBuilder().String())
	}
	builder.WriteString(l.StrBuilder.GetBuilder().String())
	return builder
}

func (l *ClangPlus) VisitProgram(ctx *ironlang.ProgramContext) {
	l.StrImportBuilder.WriteString("#include <iostream>\n")
	l.StrImportBuilder.WriteString("#include <vector>\n")
	l.StrImportBuilder.WriteString("#include <numeric>\n")
	l.InitForEachFunction()
	l.Visit(ctx.FuncMain())
	l.StrBuilder.WriteString("\n")
}

func (l *ClangPlus) VisitElseIfExpression(ctx *ironlang.ElseIfExpressionContext) {
	if ctx.ELSE() != nil {
		l.StrBuilder.WriteString(ctx.ELSE().GetText())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.IF() != nil {
		l.StrBuilder.WriteString(ctx.IF().GetText())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.RelExpression() != nil {
		l.StrBuilder.WriteString("(")
		l.Visit(ctx.RelExpression())
		l.StrBuilder.WriteString(")")
	}

	if ctx.IfScope() != nil {
		l.Visit(ctx.IfScope())
	}

	if ctx.ElseExpression() != nil {
		l.Visit(ctx.ElseExpression())
	}
}

func (l *ClangPlus) VisitElseExpression(ctx *ironlang.ElseExpressionContext) {
	if ctx.ELSE() != nil {
		l.StrBuilder.WriteString(ctx.ELSE().GetText())
	}
	if ctx.IfScope() != nil {
		l.Visit(ctx.IfScope())
	}
}

func (l *ClangPlus) VisitIfExpression(ctx *ironlang.IfExpressionContext) {

	l.StrBuilder.WriteString("if")
	if ctx.RelExpression() != nil {
		l.StrBuilder.WriteString("(")
		l.Visit(ctx.RelExpression())
		l.StrBuilder.WriteString(")")
	}
	if ctx.IfScope() != nil {
		l.Visit(ctx.IfScope())
	}

	if ctx.ElseIfExpression() != nil {
		l.Visit(ctx.ElseIfExpression())
	}

	if ctx.ElseExpression() != nil {
		l.Visit(ctx.ElseExpression())
	}

}

func (l *ClangPlus) VisitIfScope(ctx *ironlang.IfScopeContext) {
	l.StrBuilder.WriteString("{\n")
	for _, bodyScope := range ctx.AllBodyScope() {
		l.Visit(bodyScope)
	}
	l.StrBuilder.WriteString("}\n")
}

func (l *ClangPlus) VisitArray(ctx *ironlang.ArrayContext) {
	l.StrBuilder.WriteString("{")
	for i, number := range ctx.AllINT_NUMBER() {
		l.StrBuilder.WriteString(number.GetText())
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}
	}
	l.StrBuilder.WriteString("}")
}

func (l *ClangPlus) VisitFunctionArgs(ctx *ironlang.FunctionArgsContext) {

	for i, funcArg := range ctx.AllFuncArg() {
		l.Visit(funcArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}
	}
}

func (l *ClangPlus) VisitMap(ctx *ironlang.MapContext) {
	l.StrBuilder.WriteString("std::transform(" + l.LastIdentifier + ".begin(), " + l.LastIdentifier + ".end(), " + l.LastIdentifier + ".begin(),")
	l.Visit(ctx.AnonimousFunc())
	l.StrBuilder.WriteString(");\n")
}

func (l *ClangPlus) VisitFilter(ctx *ironlang.FilterContext) {
	l.StrBuilder.WriteString("std::copy_if(" + l.LastIdentifier + ".begin(), " + l.LastIdentifier + ".end(), std::back_inserter(" + l.VectorTempName + "),")
	l.Visit(ctx.AnonimousFunc())
	l.StrBuilder.WriteString(");\n")
}

func (l *ClangPlus) VisitReduce(ctx *ironlang.ReduceContext) {
	reduceResultTempName := fmt.Sprintf("reduce_result_s%v_m%v", l.ScopeId, l.MapId)
	l.StrBuilder.WriteString("auto " + reduceResultTempName + " = ")
	l.StrBuilder.WriteString("std::reduce(" + l.LastIdentifier + ".begin(), " + l.LastIdentifier + ".end(), 0,")
	l.Visit(ctx.AnonimousFunc())
	l.StrBuilder.WriteString(");\n")
	l.LastIdentifier = reduceResultTempName
}

func (l *ClangPlus) VisitMapFilterReduce(ctx *ironlang.MapFilterReduceContext) {

	if ctx.Array() != nil {

		arr := ctx.Array().(*ironlang.ArrayContext)
		types := arr.DataTypes().(*ironlang.DataTypesContext)

		arrTempName := fmt.Sprintf("arr_s%v_m%v", l.ScopeId, l.MapId)
		if types.TYPE_INT() != nil {
			l.StrBuilder.WriteString("std::vector<int> " + arrTempName + " = ")
		} else if types.TYPE_FLOAT() != nil {
			l.StrBuilder.WriteString("std::vector<float> " + arrTempName + " = ")
		} else {
			panic("Invalid data type")
		}

		l.Visit(ctx.Array())
		l.StrBuilder.WriteString(";\n")
		l.LastIdentifier = arrTempName
	}

	if ctx.IDENTIFIER() != nil {
		l.LastIdentifier = ctx.IDENTIFIER().GetText()
	}

	if ctx.Map() != nil {
		l.MapId++
		mapTempName := fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
		l.StrBuilder.WriteString("auto " + mapTempName + " = " + l.LastIdentifier + ";\n")
		l.Visit(ctx.Map())
		l.LastIdentifier = mapTempName
	}
	if ctx.Filter() != nil {
		l.MapId++
		l.VectorTempName = fmt.Sprintf("vec_s%v_m%v", l.ScopeId, l.MapId)
		l.StrBuilder.WriteString("auto " + l.VectorTempName + " = " + l.LastIdentifier + ";\n")
		l.StrBuilder.WriteString("" + l.VectorTempName + ".clear();\n")
		l.Visit(ctx.Filter())
		l.LastIdentifier = l.VectorTempName
	}

	if ctx.Reduce() != nil {
		l.MapId++
		reduceTempName := fmt.Sprintf("reduce_s%v_m%v", l.ScopeId, l.MapId)
		l.StrBuilder.WriteString("auto " + reduceTempName + " = " + l.LastIdentifier + ";\n")
		l.Visit(ctx.Reduce())
		l.LastIdentifier = reduceTempName
	}

	for _, context := range ctx.AllMapFilterReduce() {
		l.Visit(context)
	}

}

func (l *ClangPlus) VisitFuncType(ctx *ironlang.FuncTypeContext) {
	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("void ")
	}
	l.StrBuilder.WriteString("(*")
	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString(")")
	l.StrBuilder.WriteString("(")
	for i, functionArgs := range ctx.AllFunctionArgs() {
		l.Visit(functionArgs)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}
	}
	l.StrBuilder.WriteString(")")
}

func (l *ClangPlus) VisitFuncArg(ctx *ironlang.FuncArgContext) {
	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	}

	if ctx.FuncType() != nil {
		l.Visit(ctx.FuncType())
	}

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	}
}

func (l *ClangPlus) VisitFuncMain(ctx *ironlang.FuncMainContext) {
	l.StrBuilder.WriteString("int main() ")
	l.StrBuilder.WriteString("{\n")
	l.Visit(ctx.Scope())
	l.StrBuilder.WriteString("std::cout << \"Hello world!\";\n")
	l.StrBuilder.WriteString("return 0;\n")
	l.StrBuilder.WriteString("}\n")
}

func (l *ClangPlus) VisitBodyScope(ctx *ironlang.BodyScopeContext) {

	if ctx.Variable() != nil {
		l.Visit(ctx.Variable())
	}

	if ctx.Assignment() != nil {
		l.Visit(ctx.Assignment())
	}

	if ctx.IfExpression() != nil {
		l.Visit(ctx.IfExpression())
	}

	var builder = NewStringBuilder()
	if ctx.Function() != nil {
		c := NewClang(l.ScopeLog)
		c.Visit(ctx.Function())
		builder.WriteString(c.GetBuilder().String())
	}

	if len(builder.GetBuilder().String()) > 0 {
		l.StrFuncBuilder = NewStringBuilder()
		l.StrFuncBuilder.WriteString(builder.GetBuilder().String())
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
	}

	if ctx.ForEach() != nil {
		l.Visit(ctx.ForEach())
	}

	if ctx.Scope() != nil {
		l.Visit(ctx.Scope())
	}

}

func (l *ClangPlus) VisitScope(ctx *ironlang.ScopeContext) {

	l.EnterScope()
	l.ScopeLog.EnterScope(utils.GetMD5Hash(ctx.GetText()))

	for _, body := range ctx.AllBodyScope() {
		l.Visit(body)
	}

	if ctx.Return() != nil {
		l.Visit(ctx.Return())
	}

	l.ExitScope()
	l.ScopeLog.ExitScope()

}

func (l *ClangPlus) EnterScope() {
	l.ScopeId += 1
}

func (l *ClangPlus) ExitScope() {
	l.ScopeId -= 1
}

func (l *ClangPlus) VisitReturn(ctx *ironlang.ReturnContext) {

	if ctx.RETURN() == nil {
		if v, ok := ctx.GetParent().(*ironlang.AnonimousFuncContext); ok {
			if v.DataTypes() != nil {
				l.StrBuilder.WriteString("return ")
			}
		}

		if v, ok := ctx.GetParent().GetParent().GetParent().(*ironlang.AnonimousFuncContext); ok {
			if v.DataTypes() != nil {
				l.StrBuilder.WriteString("return ")
			}
		}
	} else {
		l.StrBuilder.WriteString("return ")
	}

	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
		l.StrBuilder.WriteString(";\n")
	}
}

func (l *ClangPlus) VisitFuncCall(ctx *ironlang.FuncCallContext) {
	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString("(")
	for i, funcCallArg := range ctx.AllFuncCallArg() {
		l.Visit(funcCallArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}
	}
	l.StrBuilder.WriteString(")")

	if _, ok := ctx.GetParent().(*ironlang.MathExpressionContext); !ok {
		l.StrBuilder.WriteString(";\n")
	}

}

func (l *ClangPlus) VisitAnonymousFunc(ctx *ironlang.AnonimousFuncContext) {

	l.StrBuilder.WriteString("[]")
	if ctx.L_PAREN() != nil {
		l.StrBuilder.WriteString("(")
	}

	for _, funcArg := range ctx.AllFunctionArgs() {
		l.Visit(funcArg)
	}

	if ctx.R_PAREN() != nil {
		l.StrBuilder.WriteString(")")
	}

	if ctx.DataTypes() != nil {
		l.StrBuilder.WriteString("-> ")
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("-> ")
		l.StrBuilder.WriteString("void ")
	}

	l.StrBuilder.WriteString("{\n")
	if ctx.BodyScope() != nil {
		l.Visit(ctx.BodyScope())
	}

	if ctx.Return() != nil {
		l.Visit(ctx.Return())
	}
	l.StrBuilder.WriteString("}")

}

func (l *ClangPlus) VisitFuncCallArg(ctx *ironlang.FuncCallArgContext) {
	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
	}

	if ctx.AnonimousFunc() != nil {
		l.Visit(ctx.AnonimousFunc())
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
	}

}

func (l *ClangPlus) VisitFunction(ctx *ironlang.FunctionContext) {

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("void ")
	}

	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString("(")
	l.StrBuilder.WriteString(") ")
	l.StrBuilder.WriteString("{\n")
	l.Visit(ctx.Scope())
	l.StrBuilder.WriteString("};\n")
}

func (l *ClangPlus) IsArrayType(ctx *ironlang.VariableContext) (bool, int) {
	if ctx.GetParent() != nil {
		assign := ctx.GetParent().(*ironlang.AssignmentContext)
		if assign.Array() != nil {
			array := assign.Array().(*ironlang.ArrayContext)
			dataType := array.DataTypes().(*ironlang.DataTypesContext)
			var typ = 0
			if dataType.TYPE_FLOAT() != nil {
				typ = ironlang.IronLangLexerTYPE_FLOAT
			} else if dataType.TYPE_INT() != nil {
				typ = ironlang.IronLangLexerTYPE_INT
			} else {
				panic("Data type not found")
			}

			return true, typ
		}
	}

	return false, 0
}

func (l *ClangPlus) VisitVariable(ctx *ironlang.VariableContext) {

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		isArr, typ := l.IsArrayType(ctx)
		if isArr {
			if typ == ironlang.IronLangLexerTYPE_INT {
				l.StrBuilder.WriteString("std::vector<int> ")
			}
			if typ == ironlang.IronLangLexerTYPE_FLOAT {
				l.StrBuilder.WriteString("std::vector<float> ")
			}
		} else {
			l.StrBuilder.WriteString("auto ")
		}

	}

	switch ctx.GetParent().(type) {
	case *ironlang.AssignmentContext:
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	default:
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText() + ";\n")
	}

}

func (l *ClangPlus) VisitDataTypes(ctx *ironlang.DataTypesContext) {
	if ctx.TYPE_INT() != nil {
		l.StrBuilder.WriteString("int ")
	} else {
		l.StrBuilder.WriteString("float ")
	}
}

func (l *ClangPlus) VisitAssignment(ctx *ironlang.AssignmentContext) {

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText() + " = ")
	}

	if ctx.MapFilterReduce() == nil {
		if ctx.Variable() != nil {
			l.Visit(ctx.Variable())
			l.StrBuilder.WriteString(" = ")
		}
	}

	if ctx.AnonimousFunc() != nil {
		l.Visit(ctx.AnonimousFunc())
	}

	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
	}

	if ctx.RelExpression() != nil {
		l.Visit(ctx.RelExpression())
	}

	if ctx.Array() != nil {
		l.Visit(ctx.Array())
	}

	if ctx.MapFilterReduce() != nil {
		l.Visit(ctx.MapFilterReduce())
	}

	if ctx.MapFilterReduce() != nil {
		if ctx.Variable() != nil {
			l.Visit(ctx.Variable())
			l.StrBuilder.WriteString(" = ")
			l.StrBuilder.WriteString(l.LastIdentifier)
		}
	}

	l.StrBuilder.WriteString(";\n")

}

func (l *ClangPlus) VisitForEach(ctx *ironlang.ForEachContext) {

	var identifier string
	var dataType string
	if ctx.IDENTIFIER() != nil {
		identifier = ctx.IDENTIFIER().GetText()
		//TODO: acessar o identificador para obter o tipo correto
		dataType = "int"
	} else {
		array := ctx.Array().(*ironlang.ArrayContext)
		lType := array.DataTypes().(*ironlang.DataTypesContext)

		if lType.TYPE_INT() != nil {
			dataType = "int"
		}

		if lType.TYPE_FLOAT() != nil {
			dataType = "float"
		}

		l.TempArrCount += 1
		identifier = "temp_arr_" + fmt.Sprintf("%v", l.TempArrCount)
		l.StrBuilder.WriteString("std::vector<" + dataType + "> " + identifier)
		l.StrBuilder.WriteString(" = ")
		l.Visit(ctx.Array())
		l.StrBuilder.WriteString(";\n")

	}

	l.StrBuilder.WriteString("for (auto localVar : " + identifier + ") {\n")
	l.StrBuilder.WriteString("arrForEach<" + dataType + ">(")
	if ctx.AnonimousFunc() != nil {
		l.Visit(ctx.AnonimousFunc())
	}
	l.StrBuilder.WriteString(",localVar);\n")
	l.StrBuilder.WriteString("};\n")

}

func (l *ClangPlus) VisitRelExpression(ctx *ironlang.RelExpressionContext) {

	if ctx.NOT() != nil {
		l.StrBuilder.WriteString(ctx.NOT().GetText() + " ")
		l.Visit(ctx.AllRelExpression()[0])
	}

	if ctx.L_PAREN() != nil && ctx.R_PAREN() != nil {
		l.StrBuilder.WriteString("(")
		l.Visit(ctx.AllRelExpression()[0])
		l.StrBuilder.WriteString(")")
	}

	if ctx.Atom() != nil {
		l.Visit(ctx.Atom())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.TYPE_BOOLEAN() != nil {
		l.StrBuilder.WriteString(ctx.TYPE_BOOLEAN().GetText())
		l.StrBuilder.WriteString(" ")
	}

	if len(ctx.AllRelExpression()) == 2 {
		l.Visit(ctx.AllRelExpression()[0])
		if ctx.TYPE_BOOLEAN() != nil {
			l.StrBuilder.WriteString(ctx.TYPE_BOOLEAN().GetText() + " ")
		}

		if ctx.EQEQ() != nil {
			l.StrBuilder.WriteString(ctx.EQEQ().GetText() + " ")
		}

		if ctx.DIF() != nil {
			l.StrBuilder.WriteString(ctx.DIF().GetText() + " ")
		}

		if ctx.LT() != nil {
			l.StrBuilder.WriteString(ctx.LT().GetText() + " ")
		}

		if ctx.GT() != nil {
			l.StrBuilder.WriteString(ctx.GT().GetText() + " ")
		}

		if ctx.LTEQ() != nil {
			l.StrBuilder.WriteString(ctx.LTEQ().GetText() + " ")
		}

		if ctx.GTEQ() != nil {
			l.StrBuilder.WriteString(ctx.GTEQ().GetText() + " ")
		}

		if ctx.AND() != nil {
			l.StrBuilder.WriteString(ctx.AND().GetText() + " ")
		}

		if ctx.OR() != nil {
			l.StrBuilder.WriteString(ctx.OR().GetText() + " ")
		}

		l.Visit(ctx.AllRelExpression()[1])
	}

}

func (l *ClangPlus) VisitMathExpression(ctx *ironlang.MathExpressionContext) {

	if ctx.L_PAREN() != nil && ctx.R_PAREN() != nil {
		l.StrBuilder.WriteString("(")
		l.Visit(ctx.AllMathExpression()[0])
		l.StrBuilder.WriteString(")")
	}

	if ctx.Atom() != nil {
		l.Visit(ctx.Atom())
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
	}

	if len(ctx.AllMathExpression()) == 2 {
		l.Visit(ctx.AllMathExpression()[0])
		if ctx.PLUS() != nil {
			l.StrBuilder.WriteString(" + ")
		}

		if ctx.MINUS() != nil {
			l.StrBuilder.WriteString(" - ")
		}

		if ctx.DIV() != nil {
			l.StrBuilder.WriteString(" / ")
		}

		if ctx.MULT() != nil {
			l.StrBuilder.WriteString(" * ")
		}

		l.Visit(ctx.AllMathExpression()[1])

	}

}

func (l *ClangPlus) VisitAtom(ctx *ironlang.AtomContext) {

	if ctx.INT_NUMBER() != nil {
		l.StrBuilder.WriteString(ctx.INT_NUMBER().GetText())
	} else if ctx.REAL_NUMBER() != nil {
		l.StrBuilder.WriteString(ctx.REAL_NUMBER().GetText())
	} else {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	}

}
