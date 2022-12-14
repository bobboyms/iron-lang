package codegenerator

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	//"github.com/antlr/antlr4/runtime/Go/antlr"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/scopes"
	"strconv"
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

func (l *ClangPlus) InitVectorSlicing() {
	l.StrFuncBuilder.WriteString("template <typename TT>\n")
	l.StrFuncBuilder.WriteString("std::vector<TT> vector_slicing(std::vector<TT> &arr, int X, int Y)\n")
	l.StrFuncBuilder.WriteString("{\n")
	l.StrFuncBuilder.WriteString("auto start = arr.begin() + X;\n")
	l.StrFuncBuilder.WriteString("auto end = arr.begin() + Y + 1;\n")
	l.StrFuncBuilder.WriteString("std::vector<TT> result(Y - X + 1);\n")
	l.StrFuncBuilder.WriteString("copy(start, end, result.begin());\n")
	l.StrFuncBuilder.WriteString("return result;\n")
	l.StrFuncBuilder.WriteString("}\n")
}

func (l *ClangPlus) Visit(tree antlr.ParseTree) {

	switch val := tree.(type) {
	case *ironlang.ProgramContext:
		l.VisitProgram(val)
	case *ironlang.BodyProgramContext:
		l.VisitBodyProgram(val)
	case *ironlang.MathExpressionContext:
		l.VisitMathExpression(val)
	case *ironlang.AssignmentContext:
		l.VisitAssignment(val)
	case *ironlang.LeftAssignmentContext:
		l.VisitLeftAssignment(val)
	case *ironlang.RigthAssignmentContext:
		l.VisitRigthAssignment(val)
	case *ironlang.AtomContext:
		l.VisitAtom(val)
	case *ironlang.VariableContext:
		l.VisitVariable(val)
	case *ironlang.DataTypesContext:
		l.VisitDataTypes(val)
	case *ironlang.FunctionReturnContext:
		l.VisitFunctionReturn(val)
	case *ironlang.FunctionNoReturnContext:
		l.VisitFunctionNoReturn(val)
	case *ironlang.ReturnDefinitionContext:
		l.VisitReturn(val)
	case *ironlang.FuncCallContext:
		l.VisitFuncCall(val)
	case *ironlang.FuncCallArgContext:
		l.VisitFuncCallArg(val)
	case *ironlang.BodyScopeContext:
		l.VisitBodyScope(val)
	case *ironlang.AnonimousFuncWithReturnContext:
		l.VisitAnonymousFuncWithReturn(val)
	case *ironlang.AnonimousFuncNoReturnContext:
		l.VisitAnonymousFuncNoReturn(val)
	case *ironlang.FunctionArgsContext:
		l.VisitFunctionArgs(val)
	//case *ironlang.FuncArgContext:
	//	l.VisitFuncArg(val)
	case *ironlang.FuncTypeContext:
		l.VisitFuncType(val)
	case *ironlang.ArrayContext:
		l.VisitArray(val)
	case *ironlang.ForEachContext:
		l.VisitForEach(val)
	case *ironlang.ExpressionContext:
		l.VisitExpression(val)
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
	case *ironlang.LoopWhileContext:
		l.VisitLoopWhile(val)
	case *ironlang.LoopForInContext:
		l.VisitLoopForIn(val)
	case *ironlang.LoopScopeContext:
		l.VisitLoopScope(val)
	case *ironlang.LoopDoWhileContext:
		l.VisitLoopDoWhile(val)
	case *ironlang.LoopForIContext:
		l.VisitLoopForI(val)
	case *ironlang.SliceContext:
		l.VisitSlice(val)
	case *ironlang.StructDefinitionContext:
		l.VisitStruct(val)
	case *ironlang.DefinitionVariablesContext:
		l.VisitDefinitionVariables(val)
	case *ironlang.InitStructContext:
		l.VisitInitStruct(val)
	case *ironlang.TraitContext:
		l.VisitTrait(val)
	case *ironlang.ImplContext:
		l.VisitImpl(val)
	case *ironlang.ImplConstructorContext:
		l.VisitImplConstructor(val)
	case *ironlang.InitImplContext:
		l.VisitInitImpl(val)

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

func (l *ClangPlus) VisitBodyProgram(ctx *ironlang.BodyProgramContext) {

	if ctx.Trait() != nil {
		l.Visit(ctx.Trait())
	}

	if ctx.StructDefinition() != nil {
		l.Visit(ctx.StructDefinition())
	}

	if ctx.Impl() != nil {
		l.Visit(ctx.Impl())
	}

	if ctx.FunctionReturn() != nil {
		l.Visit(ctx.FunctionReturn())
	}

	if ctx.FunctionNoReturn() != nil {
		l.Visit(ctx.FunctionNoReturn())
	}

}

func (l *ClangPlus) VisitProgram(ctx *ironlang.ProgramContext) {
	l.StrImportBuilder.WriteString("#include <iostream>\n")
	l.StrImportBuilder.WriteString("#include <vector>\n")
	l.StrImportBuilder.WriteString("#include <numeric>\n")

	l.InitForEachFunction()
	l.InitVectorSlicing()
	l.StrBuilder.WriteString("\n")

	for _, bodyProgram := range ctx.AllBodyProgram() {
		l.Visit(bodyProgram)
	}
}

func (l *ClangPlus) VisitImpl(ctx *ironlang.ImplContext) {
	l.StrBuilder.WriteString("class ")
	l.StrBuilder.WriteString(ctx.GetImplName().GetText())

	if ctx.GetTraitName() != nil {
		l.StrBuilder.WriteString(" : public ")
		l.StrBuilder.WriteString(ctx.GetTraitName().GetText())
	}

	l.StrBuilder.WriteString(" {\n")
	l.StrBuilder.WriteString("public:\n")

	for _, variable := range ctx.AllDefinitionVariables() {
		l.Visit(variable)
		l.StrBuilder.WriteString(";\n")
	}

	for _, implConstructor := range ctx.AllImplConstructor() {
		l.Visit(implConstructor)
		l.StrBuilder.WriteString(";\n")
	}

	for _, functionReturn := range ctx.AllFunctionReturn() {
		l.Visit(functionReturn)
		l.StrBuilder.WriteString(";\n")
	}

	for _, functionNoReturn := range ctx.AllFunctionNoReturn() {
		l.Visit(functionNoReturn)
		l.StrBuilder.WriteString(";\n")
	}

	l.StrBuilder.WriteString("};\n")
}

func (l *ClangPlus) VisitImplConstructor(ctx *ironlang.ImplConstructorContext) {

	impl := ctx.GetParent().(*ironlang.ImplContext)
	implName := impl.GetImplName().GetText()

	l.StrBuilder.WriteString(implName)
	l.StrBuilder.WriteString("(")
	for i, functionArg := range ctx.AllFunctionArgs() {
		l.Visit(functionArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}
	}
	l.StrBuilder.WriteString(")")
	l.StrBuilder.WriteString(" {\n")
	for _, scope := range ctx.AllBodyScope() {
		l.Visit(scope)
	}
	l.StrBuilder.WriteString("}")

}

func (l *ClangPlus) VisitTrait(ctx *ironlang.TraitContext) {
	l.StrBuilder.WriteString("class ")
	l.StrBuilder.WriteString(ctx.GetInterfaceName().GetText())
	l.StrBuilder.WriteString(" {\n")
	l.StrBuilder.WriteString("public:\n")
	for _, funcType := range ctx.AllFuncType() {
		l.StrBuilder.WriteString("virtual ")
		l.Visit(funcType)
		l.StrBuilder.WriteString(" = 0;\n")
	}

	l.StrBuilder.WriteString("};\n")
}

func (l *ClangPlus) VisitExpression(ctx *ironlang.ExpressionContext) {

	if ctx.GetLeftExpression() != nil && ctx.GetRigthExpression() != nil {
		l.Visit(ctx.GetLeftExpression())

		if context, ok := ctx.GetLeftExpression().(*ironlang.ExpressionContext); ok {

			if context.THIS() != nil {
				l.StrBuilder.WriteString("->")
			} else {
				l.StrBuilder.WriteString(".")
			}

		} else {
			l.StrBuilder.WriteString(".")
		}

		l.Visit(ctx.GetRigthExpression())
	}

	if ctx.THIS() != nil {
		l.StrBuilder.WriteString("this")
	}

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
	}

}

func (l *ClangPlus) VisitDefinitionVariables(ctx *ironlang.DefinitionVariablesContext) {

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString(ctx.GetGenericType().GetText())
		l.StrBuilder.WriteString(" ")
	}

	l.StrBuilder.WriteString(ctx.GetVariableName().GetText())
}

func (l *ClangPlus) VisitInitImpl(ctx *ironlang.InitImplContext) {
	implName := ctx.GetImplName().GetText()
	l.StrBuilder.WriteString(implName)
	l.StrBuilder.WriteString("(")
	if len(ctx.AllFuncCallArg()) > 0 {
		for i, funcCallArg := range ctx.AllFuncCallArg() {
			l.Visit(funcCallArg)
			if ctx.COMMA(i) != nil {
				l.StrBuilder.WriteString(",")
			}
		}
	}
	l.StrBuilder.WriteString(")")
}

func (l *ClangPlus) VisitInitStruct(ctx *ironlang.InitStructContext) {
	//Employee{"teste", 290}
	structName := ctx.GetStructName().GetText()
	l.StrBuilder.WriteString(structName)
	l.StrBuilder.WriteString("{")
	for i, initStructBody := range ctx.AllInitStructBody() {

		structBody := initStructBody.(*ironlang.InitStructBodyContext)
		//structKey := structBody.GetStructKey().GetText()

		if structBody.InitStruct() != nil {
			l.Visit(structBody.InitStruct())
			//l.StrBuilder.WriteString("}")
		}

		if structBody.INT_NUMBER() != nil {
			l.StrBuilder.WriteString(structBody.INT_NUMBER().GetText())
		}

		if structBody.REAL_NUMBER() != nil {
			l.StrBuilder.WriteString(structBody.REAL_NUMBER().GetText())
		}

		if structBody.BOOLEAN_VALUE() != nil {
			l.StrBuilder.WriteString(structBody.BOOLEAN_VALUE().GetText())
		}

		if structBody.STRING_LITERAL() != nil {
			l.StrBuilder.WriteString(structBody.STRING_LITERAL().GetText())
		}

		if structBody.GetAsIdent() != nil {
			l.StrBuilder.WriteString(structBody.GetAsIdent().GetText())
		}

		if structBody.FuncCall() != nil {
			l.Visit(structBody.FuncCall())
		}

		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(",")
		}

	}
	l.StrBuilder.WriteString("}")
}

func (l *ClangPlus) VisitStruct(ctx *ironlang.StructDefinitionContext) {
	l.StrBuilder.WriteString("struct ")
	l.StrBuilder.WriteString(ctx.GetStructName().GetText())
	l.StrBuilder.WriteString(" {\n")

	for _, definitionVariable := range ctx.AllDefinitionVariables() {
		l.Visit(definitionVariable)
		l.StrBuilder.WriteString(";\n")
	}

	l.StrBuilder.WriteString("};\n")
}

func (l *ClangPlus) VisitLoopForI(ctx *ironlang.LoopForIContext) {

	l.StrBuilder.WriteString(ctx.FOR().GetText())
	l.StrBuilder.WriteString(" (")
	l.Visit(ctx.Assignment())
	l.Visit(ctx.RelExpression())
	l.StrBuilder.WriteString(";")
	l.Visit(ctx.MathExpression())
	l.StrBuilder.WriteString(")")
	l.StrBuilder.WriteString("{\n")
	for _, loopScope := range ctx.AllLoopScope() {
		l.Visit(loopScope)
	}
	l.StrBuilder.WriteString("}\n")

}

func (l *ClangPlus) VisitLoopScope(ctx *ironlang.LoopScopeContext) {

	if ctx.BREAK() != nil {
		l.StrBuilder.WriteString("break;\n")
		return
	}

	if ctx.CONTINUE() != nil {
		l.StrBuilder.WriteString("continue;\n")
		return
	}

	if ctx.BodyScope() != nil {
		l.Visit(ctx.BodyScope())
	}

}

func (l *ClangPlus) VisitLoopDoWhile(ctx *ironlang.LoopDoWhileContext) {

	l.StrBuilder.WriteString(ctx.DO().GetText())
	l.StrBuilder.WriteString(" ")
	l.StrBuilder.WriteString("{\n")
	for _, loopScope := range ctx.AllLoopScope() {
		l.Visit(loopScope)
	}
	l.StrBuilder.WriteString("}")
	l.StrBuilder.WriteString(" ")

	if ctx.WHILE() != nil {
		l.StrBuilder.WriteString(ctx.WHILE().GetText())
		l.StrBuilder.WriteString(" (")
		l.Visit(ctx.RelExpression())
		l.StrBuilder.WriteString(");\n")
	} else {
		l.StrBuilder.WriteString("while")
		l.StrBuilder.WriteString(" (")
		l.StrBuilder.WriteString("true")
		l.StrBuilder.WriteString(");\n")
	}
}

func (l *ClangPlus) VisitLoopForIn(ctx *ironlang.LoopForInContext) {

	if ctx.L_PAREN() != nil {
		vectorName := fmt.Sprintf("vector_s%v_m%v", l.ScopeId, l.MapId)
		init, _ := strconv.ParseInt(ctx.INT_NUMBER(0).GetText(), 0, 32)
		last, _ := strconv.ParseInt(ctx.INT_NUMBER(1).GetText(), 0, 32)

		l.StrBuilder.WriteString("std::vector<int> " + vectorName)
		l.StrBuilder.WriteString(" = ")
		l.StrBuilder.WriteString("{")
		for i := init; i <= last; i++ {
			l.StrBuilder.WriteString(strconv.FormatInt(i, 10) + ",")
		}
		l.StrBuilder.WriteString("};\n")

		l.StrBuilder.WriteString(ctx.FOR().GetText())
		l.StrBuilder.WriteString(" (")
		l.StrBuilder.WriteString("auto ")
		l.StrBuilder.WriteString(ctx.IDENTIFIER(0).GetText())
		l.StrBuilder.WriteString(" : ")
		l.StrBuilder.WriteString(vectorName)
		l.StrBuilder.WriteString(") ")
		l.StrBuilder.WriteString("{\n")
		for _, loopScope := range ctx.AllLoopScope() {
			l.Visit(loopScope)
		}
		l.StrBuilder.WriteString("}\n")

	}

	if ctx.Slice() != nil {
		l.StrBuilder.WriteString(ctx.FOR().GetText())
		l.StrBuilder.WriteString(" (")
		l.StrBuilder.WriteString("auto ")
		l.StrBuilder.WriteString(ctx.IDENTIFIER(0).GetText())
		l.StrBuilder.WriteString(" : ")
		l.Visit(ctx.Slice())
		l.StrBuilder.WriteString(") ")
		l.StrBuilder.WriteString("{\n")
		for _, loopScope := range ctx.AllLoopScope() {
			l.Visit(loopScope)
		}
		l.StrBuilder.WriteString("}\n")
	}

	if ctx.IDENTIFIER(1) != nil {
		l.StrBuilder.WriteString(ctx.FOR().GetText())
		l.StrBuilder.WriteString(" (")
		l.StrBuilder.WriteString("auto ")
		l.StrBuilder.WriteString(ctx.IDENTIFIER(0).GetText())
		l.StrBuilder.WriteString(" : ")
		l.StrBuilder.WriteString(ctx.IDENTIFIER(1).GetText())
		l.StrBuilder.WriteString(") ")
		l.StrBuilder.WriteString("{\n")
		for _, loopScope := range ctx.AllLoopScope() {
			l.Visit(loopScope)
		}
		l.StrBuilder.WriteString("}\n")
	}

}

func (l *ClangPlus) VisitLoopWhile(ctx *ironlang.LoopWhileContext) {
	l.StrBuilder.WriteString(ctx.WHILE().GetText())
	l.StrBuilder.WriteString(" (")
	l.Visit(ctx.RelExpression())
	l.StrBuilder.WriteString(") ")
	l.StrBuilder.WriteString("{\n")
	for _, loopScope := range ctx.AllLoopScope() {
		l.Visit(loopScope)
	}
	l.StrBuilder.WriteString("}\n")
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

	if ctx.FuncType() != nil {
		l.Visit(ctx.FuncType())
	}

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	}

	if ctx.GetIdentType() != nil {
		l.StrBuilder.WriteString(ctx.GetIdentType().GetText())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.GetArgName() != nil {
		l.StrBuilder.WriteString(ctx.GetArgName().GetText())
	}

}

func (l *ClangPlus) VisitMap(ctx *ironlang.MapContext) {
	l.StrBuilder.WriteString("std::transform(" + l.LastIdentifier + ".begin(), " + l.LastIdentifier + ".end(), " + l.LastIdentifier + ".begin(),")
	l.Visit(ctx.AnonimousFuncWithReturn())
	l.StrBuilder.WriteString(");\n")
}

func (l *ClangPlus) VisitFilter(ctx *ironlang.FilterContext) {
	l.StrBuilder.WriteString("std::copy_if(" + l.LastIdentifier + ".begin(), " + l.LastIdentifier + ".end(), std::back_inserter(" + l.VectorTempName + "),")
	l.Visit(ctx.AnonimousFuncWithReturn())
	l.StrBuilder.WriteString(");\n")
}

func (l *ClangPlus) VisitReduce(ctx *ironlang.ReduceContext) {
	reduceResultTempName := fmt.Sprintf("reduce_result_s%v_m%v", l.ScopeId, l.MapId)
	l.StrBuilder.WriteString("auto " + reduceResultTempName + " = ")
	l.StrBuilder.WriteString("std::reduce(" + l.LastIdentifier + ".begin(), " + l.LastIdentifier + ".end(), 0,")
	l.Visit(ctx.AnonimousFuncWithReturn())
	l.StrBuilder.WriteString(");\n")
	l.LastIdentifier = reduceResultTempName
}

//func (l *ClangPlus) VisitExpression(ctx *ironlang.ExpressionContext) {
//
//	if ctx.Array() != nil {
//
//		arr := ctx.Array().(*ironlang.ArrayContext)
//		types := arr.DataTypes().(*ironlang.DataTypesContext)
//
//		arrTempName := fmt.Sprintf("arr_s%v_m%v", l.ScopeId, l.MapId)
//		if types.TYPE_INT() != nil {
//			l.StrBuilder.WriteString("std::vector<int> " + arrTempName + " = ")
//		} else if types.TYPE_FLOAT() != nil {
//			l.StrBuilder.WriteString("std::vector<float> " + arrTempName + " = ")
//		} else {
//			panic("Invalid data type")
//		}
//
//		l.Visit(ctx.Array())
//		l.StrBuilder.WriteString(";\n")
//		l.LastIdentifier = arrTempName
//	}
//
//	if ctx.IDENTIFIER() != nil {
//		l.LastIdentifier = ctx.IDENTIFIER().GetText()
//	}
//
//	if ctx.Slice() != nil {
//		l.MapId++
//		sliceTempName := fmt.Sprintf("slice_s%v_m%v", l.ScopeId, l.MapId)
//		l.StrBuilder.WriteString("auto " + sliceTempName + " = " + l.LastIdentifier)
//		l.Visit(ctx.Slice())
//		l.StrBuilder.WriteString(";\n")
//		l.LastIdentifier = sliceTempName
//	}
//
//	if ctx.Map() != nil {
//		l.MapId++
//		l.Visit(ctx.Map())
//		mapTempName := fmt.Sprintf("map_s%v_m%v", l.ScopeId, l.MapId)
//		l.StrBuilder.WriteString("auto " + mapTempName + " = " + l.LastIdentifier + ";\n")
//		l.LastIdentifier = mapTempName
//	}
//	if ctx.Filter() != nil {
//		l.MapId++
//		l.VectorTempName = fmt.Sprintf("vec_s%v_m%v", l.ScopeId, l.MapId)
//		l.StrBuilder.WriteString("auto " + l.VectorTempName + " = " + l.LastIdentifier + ";\n")
//		l.StrBuilder.WriteString("" + l.VectorTempName + ".clear();\n")
//		l.Visit(ctx.Filter())
//		l.LastIdentifier = l.VectorTempName
//	}
//
//	if ctx.Reduce() != nil {
//		l.MapId++
//		reduceTempName := fmt.Sprintf("reduce_s%v_m%v", l.ScopeId, l.MapId)
//		l.StrBuilder.WriteString("auto " + reduceTempName + " = " + l.LastIdentifier + ";\n")
//		l.Visit(ctx.Reduce())
//		l.LastIdentifier = reduceTempName
//	}
//
//	for _, context := range ctx.AllExpression() {
//		l.Visit(context)
//	}
//
//}

func (l *ClangPlus) VisitFuncType(ctx *ironlang.FuncTypeContext) {
	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("void ")
	}
	l.StrBuilder.WriteString("(")
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

//func (l *ClangPlus) VisitFuncArg(ctx *ironlang.FuncArgContext) {
//	if ctx.DataTypes() != nil {
//		l.Visit(ctx.DataTypes())
//	}
//
//	if ctx.FuncType() != nil {
//		l.Visit(ctx.FuncType())
//	}
//
//	if ctx.IDENTIFIER() != nil {
//		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
//	}
//}

//func (l *ClangPlus) VisitFuncMain(ctx *ironlang.FuncMainContext) {
//	l.StrBuilder.WriteString("int main() ")
//	l.StrBuilder.WriteString("{\n")
//	for _, body := range ctx.AllBodyScope() {
//		l.Visit(body)
//	}
//	l.StrBuilder.WriteString("std::cout << \"Hello world!\";\n")
//	l.StrBuilder.WriteString("return 0;\n")
//	l.StrBuilder.WriteString("}\n")
//}

func (l *ClangPlus) VisitBodyScope(ctx *ironlang.BodyScopeContext) {

	if ctx.Expression() != nil {
		l.Visit(ctx.Expression())
		l.StrBuilder.WriteString(";\n")
		return
	}

	if ctx.Variable() != nil {
		l.Visit(ctx.Variable())
		l.StrBuilder.WriteString(";\n")
		return
	}

	if ctx.Assignment() != nil {
		l.Visit(ctx.Assignment())
		l.StrBuilder.WriteString(";\n")
		return
	}

	if ctx.IfExpression() != nil {
		l.Visit(ctx.IfExpression())
		return
	}

	if ctx.LoopDoWhile() != nil {
		l.Visit(ctx.LoopDoWhile())
		return
	}

	if ctx.LoopWhile() != nil {
		l.Visit(ctx.LoopWhile())
		return
	}

	if ctx.LoopForIn() != nil {
		l.Visit(ctx.LoopForIn())
		return
	}

	if ctx.LoopForI() != nil {
		l.Visit(ctx.LoopForI())
		return
	}

	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
		return
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
		l.StrBuilder.WriteString(";\n")
		return
	}

	if ctx.ForEach() != nil {
		l.Visit(ctx.ForEach())
		return
	}

}

func (l *ClangPlus) EnterScope() {
	l.ScopeId += 1
}

func (l *ClangPlus) ExitScope() {
	l.ScopeId -= 1
}

func (l *ClangPlus) VisitReturn(ctx *ironlang.ReturnDefinitionContext) {

	l.StrBuilder.WriteString("return ")

	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
		//l.StrBuilder.WriteString(";\n")
	}

	if ctx.RelExpression() != nil {
		l.Visit(ctx.RelExpression())
		//l.StrBuilder.WriteString(";\n")
	}

	if ctx.Expression() != nil {
		l.Visit(ctx.Expression())
		//l.StrBuilder.WriteString(";\n")
	}

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
		//l.StrBuilder.WriteString(";\n")
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

	//if _, ok := ctx.GetParent().(*ironlang.MathExpressionContext); !ok {
	//	l.StrBuilder.WriteString(";\n")
	//}

}

func (l *ClangPlus) VisitAnonymousFuncNoReturn(ctx *ironlang.AnonimousFuncNoReturnContext) {
	l.StrBuilder.WriteString("[]")
	if ctx.L_PAREN() != nil {
		l.StrBuilder.WriteString("(")
	}

	for i, functionArg := range ctx.AllFunctionArgs() {
		l.Visit(functionArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(", ")
		}
	}

	if ctx.R_PAREN() != nil {
		l.StrBuilder.WriteString(")")
	}
	l.StrBuilder.WriteString("-> ")
	l.StrBuilder.WriteString("void ")
	l.StrBuilder.WriteString("{\n")

	for _, bodyScope := range ctx.AllBodyScope() {
		l.Visit(bodyScope)
		l.StrBuilder.WriteString(";\n")
	}

	l.StrBuilder.WriteString("}")
}

func (l *ClangPlus) VisitAnonymousFuncWithReturn(ctx *ironlang.AnonimousFuncWithReturnContext) {

	l.StrBuilder.WriteString("[]")
	if ctx.L_PAREN() != nil {
		l.StrBuilder.WriteString("(")
	}

	for i, functionArg := range ctx.AllFunctionArgs() {
		l.Visit(functionArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(", ")
		}
	}

	if ctx.R_PAREN() != nil {
		l.StrBuilder.WriteString(")")
	}

	if ctx.DataTypes() != nil {
		l.StrBuilder.WriteString("-> ")
		l.Visit(ctx.DataTypes())
	}

	l.StrBuilder.WriteString("{\n")

	for _, bodyScope := range ctx.AllBodyScope() {
		l.Visit(bodyScope)
	}

	if ctx.ReturnDefinition() != nil {
		l.Visit(ctx.ReturnDefinition())
		l.StrBuilder.WriteString(";\n")
	}
	l.StrBuilder.WriteString("}")

}

func (l *ClangPlus) VisitFuncCallArg(ctx *ironlang.FuncCallArgContext) {
	if ctx.MathExpression() != nil {
		l.Visit(ctx.MathExpression())
	}

	if ctx.AnonimousFuncWithReturn() != nil {
		l.Visit(ctx.AnonimousFuncWithReturn())
	}

	if ctx.AnonimousFuncNoReturn() != nil {
		l.Visit(ctx.AnonimousFuncNoReturn())
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
	}

	if ctx.InitStruct() != nil {
		l.Visit(ctx.InitStruct())
	}

	if ctx.InitImpl() != nil {
		l.Visit(ctx.InitImpl())
	}

	if ctx.STRING_LITERAL() != nil {
		l.StrBuilder.WriteString(ctx.STRING_LITERAL().GetText())
	}

	if ctx.CHAR_LITERAL() != nil {
		l.StrBuilder.WriteString(ctx.CHAR_LITERAL().GetText())
	}

}

func (l *ClangPlus) VisitFunctionNoReturn(ctx *ironlang.FunctionNoReturnContext) {

	l.StrBuilder.WriteString("void ")
	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString("(")
	for i, functionArg := range ctx.AllFunctionArgs() {
		l.Visit(functionArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(", ")
		}
	}
	l.StrBuilder.WriteString(") ")
	l.StrBuilder.WriteString("{\n")

	for _, body := range ctx.AllBodyScope() {
		l.Visit(body)
	}
	l.StrBuilder.WriteString("}\n")

}

func (l *ClangPlus) VisitFunctionReturn(ctx *ironlang.FunctionReturnContext) {

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("void ")
	}

	l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
	l.StrBuilder.WriteString("(")
	for i, functionArg := range ctx.AllFunctionArgs() {
		l.Visit(functionArg)
		if ctx.COMMA(i) != nil {
			l.StrBuilder.WriteString(", ")
		}
	}
	l.StrBuilder.WriteString(") ")
	l.StrBuilder.WriteString("{\n")

	for _, body := range ctx.AllBodyScope() {
		l.Visit(body)
	}

	if ctx.ReturnDefinition() != nil {
		l.Visit(ctx.ReturnDefinition())
		l.StrBuilder.WriteString(";\n")
	}

	l.StrBuilder.WriteString("}")

}

//func (l *ClangPlus) IsArrayType(ctx *ironlang.VariableContext) (bool, int) {
//	if ctx.GetParent() != nil {
//		assign := ctx.GetParent().(*ironlang.AssignmentContext)
//		if assign.Array() != nil {
//			array := assign.Array().(*ironlang.ArrayContext)
//			dataType := array.DataTypes().(*ironlang.DataTypesContext)
//			var typ = 0
//			if dataType.TYPE_FLOAT() != nil {
//				typ = ironlang.IronLangLexerTYPE_FLOAT
//			} else if dataType.TYPE_INT() != nil {
//				typ = ironlang.IronLangLexerTYPE_INT
//			} else {
//				panic("Data type not found")
//			}
//
//			return true, typ
//		}
//	}
//
//	return false, 0
//}

func (l *ClangPlus) VisitVariable(ctx *ironlang.VariableContext) {

	//if ctx.MUT() == nil {
	//	l.StrBuilder.WriteString("const ")
	//}

	if ctx.DataTypes() != nil {
		l.Visit(ctx.DataTypes())
	} else {
		l.StrBuilder.WriteString("auto ")
	}

	l.StrBuilder.WriteString(ctx.GetVariableName().GetText())

	//if ctx.DataTypes() != nil {
	//	l.Visit(ctx.DataTypes())
	//} else {
	//	isArr, typ := l.IsArrayType(ctx)
	//	if isArr {
	//		if typ == ironlang.IronLangLexerTYPE_INT {
	//			l.StrBuilder.WriteString("std::vector<int> ")
	//		}
	//		if typ == ironlang.IronLangLexerTYPE_FLOAT {
	//			l.StrBuilder.WriteString("std::vector<float> ")
	//		}
	//	} else {
	//		l.StrBuilder.WriteString("auto ")
	//	}
	//
	//}
	//
	//switch ctx.GetParent().(type) {
	//case *ironlang.AssignmentContext:
	//	l.StrBuilder.WriteString(ctx.GetVariableName().GetText())
	//default:
	//	l.StrBuilder.WriteString(ctx.GetVariableName().GetText() + ";\n")
	//}

}

func (l *ClangPlus) VisitDataTypes(ctx *ironlang.DataTypesContext) {

	if ctx.TYPE_INT() != nil {
		l.StrBuilder.WriteString("int ")
		return
	}

	if ctx.TYPE_FLOAT() != nil {
		l.StrBuilder.WriteString("float ")
		return
	}

	if ctx.TYPE_DOUBLE() != nil {
		l.StrBuilder.WriteString("double ")
		return
	}

	if ctx.TYPE_BOOLEAN() != nil {
		l.StrBuilder.WriteString("bool ")
		return
	}

	if ctx.TYPE_CHAR() != nil {
		l.StrBuilder.WriteString("char ")
		return
	}

	if ctx.TYPE_STRING() != nil {
		l.StrBuilder.WriteString("std::string ")
		return
	}

	panic("Invalid data type")
}

func (l *ClangPlus) VisitSlice(ctx *ironlang.SliceContext) {
	identifier := ctx.IDENTIFIER().GetText()
	start := ctx.INT_NUMBER(0).GetText()
	end := ctx.INT_NUMBER(1).GetText()
	l.StrBuilder.WriteString("vector_slicing<int>(" + identifier + ", " + start + "," + end + ")")
}

func (l *ClangPlus) VisitLeftAssignment(ctx *ironlang.LeftAssignmentContext) {
	if ctx.Variable() != nil {
		l.Visit(ctx.Variable())
	}
	if ctx.GetLeftExpression() != nil {
		l.Visit(ctx.GetLeftExpression())
	}
	if ctx.GetVariableName() != nil {
		l.StrBuilder.WriteString(ctx.GetVariableName().GetText())
	}
}

func (l *ClangPlus) VisitRigthAssignment(ctx *ironlang.RigthAssignmentContext) {

	if ctx.GetIdentType() != nil {
		l.StrBuilder.WriteString(ctx.GetIdentType().GetText())
	}

	if ctx.InitStruct() != nil {
		l.Visit(ctx.InitStruct())
	}

	if ctx.InitImpl() != nil {
		l.Visit(ctx.InitImpl())
	}

	if ctx.Slice() != nil {
		l.Visit(ctx.Slice())
	}

	if ctx.GetRightExpression() != nil {
		l.Visit(ctx.GetRightExpression())
	}

	if ctx.AnonimousFuncWithReturn() != nil {
		l.Visit(ctx.AnonimousFuncWithReturn())
	}

	if ctx.AnonimousFuncNoReturn() != nil {
		l.Visit(ctx.AnonimousFuncNoReturn())
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

	if ctx.STRING_LITERAL() != nil {
		l.StrBuilder.WriteString(ctx.STRING_LITERAL().GetText())
	}

	if ctx.CHAR_LITERAL() != nil {
		l.StrBuilder.WriteString(ctx.CHAR_LITERAL().GetText())
	}

	if ctx.BOOLEAN_VALUE() != nil {
		l.StrBuilder.WriteString(ctx.BOOLEAN_VALUE().GetText())
	}
}

func (l *ClangPlus) VisitAssignment(ctx *ironlang.AssignmentContext) {

	l.Visit(ctx.LeftAssignment())
	l.StrBuilder.WriteString(" = ")
	l.Visit(ctx.RigthAssignment())
	//l.StrBuilder.WriteString(";\n")

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
	if ctx.AnonimousFuncNoReturn() != nil {
		l.Visit(ctx.AnonimousFuncNoReturn())
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

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.Atom() != nil {
		l.Visit(ctx.Atom())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.FuncCall() != nil {
		l.Visit(ctx.FuncCall())
		l.StrBuilder.WriteString(" ")
	}

	if ctx.BOOLEAN_VALUE() != nil {
		l.StrBuilder.WriteString(ctx.BOOLEAN_VALUE().GetText())
		l.StrBuilder.WriteString(" ")
	}

	if len(ctx.AllRelExpression()) == 2 {
		l.Visit(ctx.AllRelExpression()[0])
		if ctx.BOOLEAN_VALUE() != nil {
			l.StrBuilder.WriteString(ctx.BOOLEAN_VALUE().GetText() + " ")
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

	if ctx.IDENTIFIER() != nil {
		l.StrBuilder.WriteString(ctx.IDENTIFIER().GetText())
		if ctx.PLUS_PLUS() != nil {
			l.StrBuilder.WriteString(ctx.PLUS_PLUS().GetText())
		}
		if ctx.MINUS_MINUS() != nil {
			l.StrBuilder.WriteString(ctx.MINUS_MINUS().GetText())
		}
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
	}

}
