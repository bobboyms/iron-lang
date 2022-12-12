// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package ironlang // IronLang
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// IronLangListener is a complete listener for a parse tree produced by IronLangParser.
type IronLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBodyProgram is called when entering the bodyProgram production.
	EnterBodyProgram(c *BodyProgramContext)

	// EnterFunctionReturn is called when entering the functionReturn production.
	EnterFunctionReturn(c *FunctionReturnContext)

	// EnterFunctionNoReturn is called when entering the functionNoReturn production.
	EnterFunctionNoReturn(c *FunctionNoReturnContext)

	// EnterFuncType is called when entering the funcType production.
	EnterFuncType(c *FuncTypeContext)

	// EnterImplConstructor is called when entering the implConstructor production.
	EnterImplConstructor(c *ImplConstructorContext)

	// EnterReturnDefinition is called when entering the returnDefinition production.
	EnterReturnDefinition(c *ReturnDefinitionContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncCallArg is called when entering the funcCallArg production.
	EnterFuncCallArg(c *FuncCallArgContext)

	// EnterAnonimousFuncWithReturn is called when entering the anonimousFuncWithReturn production.
	EnterAnonimousFuncWithReturn(c *AnonimousFuncWithReturnContext)

	// EnterAnonimousFuncNoReturn is called when entering the anonimousFuncNoReturn production.
	EnterAnonimousFuncNoReturn(c *AnonimousFuncNoReturnContext)

	// EnterIfScope is called when entering the ifScope production.
	EnterIfScope(c *IfScopeContext)

	// EnterElseExpression is called when entering the elseExpression production.
	EnterElseExpression(c *ElseExpressionContext)

	// EnterElseIfExpression is called when entering the elseIfExpression production.
	EnterElseIfExpression(c *ElseIfExpressionContext)

	// EnterIfExpression is called when entering the ifExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterLoopScope is called when entering the loopScope production.
	EnterLoopScope(c *LoopScopeContext)

	// EnterLoopDoWhile is called when entering the loopDoWhile production.
	EnterLoopDoWhile(c *LoopDoWhileContext)

	// EnterLoopWhile is called when entering the loopWhile production.
	EnterLoopWhile(c *LoopWhileContext)

	// EnterLoopForIn is called when entering the loopForIn production.
	EnterLoopForIn(c *LoopForInContext)

	// EnterLoopForI is called when entering the loopForI production.
	EnterLoopForI(c *LoopForIContext)

	// EnterBodyScope is called when entering the bodyScope production.
	EnterBodyScope(c *BodyScopeContext)

	// EnterTrait is called when entering the trait production.
	EnterTrait(c *TraitContext)

	// EnterImpl is called when entering the impl production.
	EnterImpl(c *ImplContext)

	// EnterInitImpl is called when entering the initImpl production.
	EnterInitImpl(c *InitImplContext)

	// EnterStructDefinition is called when entering the structDefinition production.
	EnterStructDefinition(c *StructDefinitionContext)

	// EnterDefinitionVariables is called when entering the definitionVariables production.
	EnterDefinitionVariables(c *DefinitionVariablesContext)

	// EnterInitStruct is called when entering the initStruct production.
	EnterInitStruct(c *InitStructContext)

	// EnterInitStructBody is called when entering the initStructBody production.
	EnterInitStructBody(c *InitStructBodyContext)

	// EnterPrintln is called when entering the println production.
	EnterPrintln(c *PrintlnContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterDataTypes is called when entering the dataTypes production.
	EnterDataTypes(c *DataTypesContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterLeftAssignment is called when entering the leftAssignment production.
	EnterLeftAssignment(c *LeftAssignmentContext)

	// EnterRigthAssignment is called when entering the rigthAssignment production.
	EnterRigthAssignment(c *RigthAssignmentContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterForEach is called when entering the forEach production.
	EnterForEach(c *ForEachContext)

	// EnterMap is called when entering the map production.
	EnterMap(c *MapContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterReduce is called when entering the reduce production.
	EnterReduce(c *ReduceContext)

	// EnterRelExpression is called when entering the relExpression production.
	EnterRelExpression(c *RelExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMathExpression is called when entering the mathExpression production.
	EnterMathExpression(c *MathExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBodyProgram is called when exiting the bodyProgram production.
	ExitBodyProgram(c *BodyProgramContext)

	// ExitFunctionReturn is called when exiting the functionReturn production.
	ExitFunctionReturn(c *FunctionReturnContext)

	// ExitFunctionNoReturn is called when exiting the functionNoReturn production.
	ExitFunctionNoReturn(c *FunctionNoReturnContext)

	// ExitFuncType is called when exiting the funcType production.
	ExitFuncType(c *FuncTypeContext)

	// ExitImplConstructor is called when exiting the implConstructor production.
	ExitImplConstructor(c *ImplConstructorContext)

	// ExitReturnDefinition is called when exiting the returnDefinition production.
	ExitReturnDefinition(c *ReturnDefinitionContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncCallArg is called when exiting the funcCallArg production.
	ExitFuncCallArg(c *FuncCallArgContext)

	// ExitAnonimousFuncWithReturn is called when exiting the anonimousFuncWithReturn production.
	ExitAnonimousFuncWithReturn(c *AnonimousFuncWithReturnContext)

	// ExitAnonimousFuncNoReturn is called when exiting the anonimousFuncNoReturn production.
	ExitAnonimousFuncNoReturn(c *AnonimousFuncNoReturnContext)

	// ExitIfScope is called when exiting the ifScope production.
	ExitIfScope(c *IfScopeContext)

	// ExitElseExpression is called when exiting the elseExpression production.
	ExitElseExpression(c *ElseExpressionContext)

	// ExitElseIfExpression is called when exiting the elseIfExpression production.
	ExitElseIfExpression(c *ElseIfExpressionContext)

	// ExitIfExpression is called when exiting the ifExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitLoopScope is called when exiting the loopScope production.
	ExitLoopScope(c *LoopScopeContext)

	// ExitLoopDoWhile is called when exiting the loopDoWhile production.
	ExitLoopDoWhile(c *LoopDoWhileContext)

	// ExitLoopWhile is called when exiting the loopWhile production.
	ExitLoopWhile(c *LoopWhileContext)

	// ExitLoopForIn is called when exiting the loopForIn production.
	ExitLoopForIn(c *LoopForInContext)

	// ExitLoopForI is called when exiting the loopForI production.
	ExitLoopForI(c *LoopForIContext)

	// ExitBodyScope is called when exiting the bodyScope production.
	ExitBodyScope(c *BodyScopeContext)

	// ExitTrait is called when exiting the trait production.
	ExitTrait(c *TraitContext)

	// ExitImpl is called when exiting the impl production.
	ExitImpl(c *ImplContext)

	// ExitInitImpl is called when exiting the initImpl production.
	ExitInitImpl(c *InitImplContext)

	// ExitStructDefinition is called when exiting the structDefinition production.
	ExitStructDefinition(c *StructDefinitionContext)

	// ExitDefinitionVariables is called when exiting the definitionVariables production.
	ExitDefinitionVariables(c *DefinitionVariablesContext)

	// ExitInitStruct is called when exiting the initStruct production.
	ExitInitStruct(c *InitStructContext)

	// ExitInitStructBody is called when exiting the initStructBody production.
	ExitInitStructBody(c *InitStructBodyContext)

	// ExitPrintln is called when exiting the println production.
	ExitPrintln(c *PrintlnContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitDataTypes is called when exiting the dataTypes production.
	ExitDataTypes(c *DataTypesContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitLeftAssignment is called when exiting the leftAssignment production.
	ExitLeftAssignment(c *LeftAssignmentContext)

	// ExitRigthAssignment is called when exiting the rigthAssignment production.
	ExitRigthAssignment(c *RigthAssignmentContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitForEach is called when exiting the forEach production.
	ExitForEach(c *ForEachContext)

	// ExitMap is called when exiting the map production.
	ExitMap(c *MapContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitReduce is called when exiting the reduce production.
	ExitReduce(c *ReduceContext)

	// ExitRelExpression is called when exiting the relExpression production.
	ExitRelExpression(c *RelExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMathExpression is called when exiting the mathExpression production.
	ExitMathExpression(c *MathExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}
