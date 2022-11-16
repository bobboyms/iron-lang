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
	l.InitForEachFunction()
	l.Visit(ctx.FuncMain())
	l.StrBuilder.WriteString("\n")
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

func (l *ClangPlus) VisitMapFilterReduce(ctx *ironlang.MapFilterReduceContext) {

	if ctx.IDENTIFIER() != nil {

	}

	if ctx.Map() != nil {
		l.Visit(ctx.Map())
	}

	if ctx.Filter() != nil {
		l.Visit(ctx.Filter())
	}

	if ctx.Reduce() != nil {
		l.Visit(ctx.Reduce())
	}

	for _, context := range ctx.AllMapFilterReduce() {
		l.Visit(context)
	}

}

//var localId = 0
//var tempName string
//identifier := ctx.IDENTIFIER().GetText()
//tempName = fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
//
//for _, mpp := range ctx.AllMap() {
//lmp := mpp.(*ironlang.MapContext)
//if localId == 0 {
//localId++
//l.MapId++
//l.StrBuilder.WriteString("auto to_vector = " + identifier + ";\n")
//tempName = fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
//l.StrBuilder.WriteString("auto " + tempName + " = " + identifier + ";\n")
//l.StrBuilder.WriteString("std::transform(" + tempName + ".begin(), " + tempName + ".end(), " + tempName + ".begin(),")
//l.Visit(lmp.AnonimousFunc())
//l.StrBuilder.WriteString(");\n")
//l.LastIdentifier = tempName
//} else {
//l.MapId++
//newName := fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
//l.StrBuilder.WriteString("auto " + newName + " = " + tempName + ";\n")
//l.StrBuilder.WriteString("std::transform(" + newName + ".begin(), " + newName + ".end(), " + newName + ".begin(),")
//l.Visit(lmp.AnonimousFunc())
//l.StrBuilder.WriteString(");\n")
//tempName = newName
//l.LastIdentifier = newName
//}
//}
//
//for _, filter := range ctx.AllFilter() {
//
//flt := filter.(*ironlang.FilterContext)
//
//if localId == 0 {
//localId++
//l.MapId++
//l.StrBuilder.WriteString("auto to_vector = " + identifier + ";\n")
//l.StrBuilder.WriteString("to_vector.clear();\n")
//tempName = fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
//l.StrBuilder.WriteString("auto " + tempName + " = " + identifier + ";\n")
//l.StrBuilder.WriteString("std::copy_if(" + tempName + ".begin(), " + tempName + ".end(), std::back_inserter(to_vector),")
//l.Visit(flt.AnonimousFunc())
//l.StrBuilder.WriteString(");\n")
//l.StrBuilder.WriteString(tempName + " = to_vector;\n")
//l.LastIdentifier = tempName
//} else {
//l.MapId++
//newName := fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
//l.StrBuilder.WriteString("to_vector.clear();\n")
//l.StrBuilder.WriteString("std::copy_if(" + tempName + ".begin(), " + tempName + ".end(), std::back_inserter(to_vector),")
//l.Visit(flt.AnonimousFunc())
//l.StrBuilder.WriteString(");\n")
//l.StrBuilder.WriteString("auto " + newName + " = to_vector;\n")
//tempName = newName
//l.LastIdentifier = newName
//}
//}

func (l *ClangPlus) VisitFuncType(ctx *ironlang.FuncTypeContext) {
	//int (*func)(int x)
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

	//funcName := fmt.Sprintf("anonymous_%v", l.AnonymousCount)
	//l.StrBuilder.WriteString(funcName)

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

	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
	}

	if ctx.AnonimousFunc() != nil {
		l.Visit(ctx.AnonimousFunc())
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
