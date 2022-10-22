package llvm

import (
	"fmt"
	"iron-lang/compiler/ironlang"
	"iron-lang/compiler/utils"
	"strconv"
	"strings"
)

type Expression struct {
	TypeIdentifier int
	Identifier     string
	Operation      int
	DataType       int
}

type ExpressionStack struct {
	Stack        []*Expression
	TempVarName  string
	TempVarType  int
	PreviousName string
	PreviousType int
	Counter      int
}

func NewExpressionStack() *ExpressionStack {
	return &ExpressionStack{
		Counter:      -1,
		PreviousType: 0,
		TempVarType:  0,
	}
}

func (e *ExpressionStack) InitExpression() {
	e.Counter += 1
	e.Stack = make([]*Expression, 0)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func (e *ExpressionStack) Register(typeIdentifier int, identifier string, operation int, dataType int) {
	e.Stack = append(e.Stack, &Expression{
		TypeIdentifier: typeIdentifier,
		Identifier:     identifier,
		Operation:      operation,
		DataType:       dataType,
	})

	println(len(e.Stack))
}

func (e *ExpressionStack) GetStackId() string {
	return fmt.Sprintf("%v", e.Counter)
}

func (e *ExpressionStack) CreateVar(expression *Expression, refId string) string {
	build := &strings.Builder{}
	if expression.TypeIdentifier == ironlang.IronLangLexerIDENTIFIER {

		if expression.DataType == ironlang.IronLangLexerTYPE_INT {
			e.TempVarType = ironlang.IronLangLexerTYPE_INT
			build.WriteString(e.TempVarName + " = load i32, i32* %" + expression.Identifier + ", align 4\n")
		} else {
			e.TempVarType = ironlang.IronLangLexerTYPE_FLOAT
			build.WriteString(e.TempVarName + " = load float, float* %" + expression.Identifier + ", align 4\n")
		}

	} else {

		local := "%st" + e.GetStackId() + ".xp." + refId
		//Se for um número
		if expression.DataType == ironlang.IronLangLexerTYPE_INT {
			e.TempVarType = ironlang.IronLangLexerTYPE_INT
			build.WriteString(local + " = alloca i32, align 4\n")
			build.WriteString("store i32 " + expression.Identifier + ", i32* " + local + ", align 4\n")
			build.WriteString(e.TempVarName + " = load i32, i32* " + local + ", align 4\n")

		} else {
			e.TempVarType = ironlang.IronLangLexerTYPE_FLOAT
			value, _ := strconv.ParseFloat(expression.Identifier, 64)
			vScience := utils.GetScienceValue(value)
			build.WriteString(local + " = alloca float, align 4\n")
			build.WriteString("store float " + vScience + ", float* " + local + ", align 4\n")
			build.WriteString(e.TempVarName + " = load float, float* " + local + ", align 4\n")
		}
	}
	return build.String()
}

func (e *ExpressionStack) GenerateLLVMCode() string {
	build := &strings.Builder{}

	for i := len(e.Stack) - 1; i >= 0; i-- {
		expression := e.Stack[i]
		e.TempVarName = "%st" + e.GetStackId() + ".load." + fmt.Sprintf("%v", i)

		if expression.Operation == 0 {

			build.WriteString(e.CreateVar(expression, fmt.Sprintf("%v", i)))
			e.PreviousName = e.TempVarName
			e.PreviousType = e.TempVarType

		} else if expression.Operation == ironlang.IronLangLexerPLUS {
			addTemp := "%st" + e.GetStackId() + ".add." + fmt.Sprintf("%v", i)
			build.WriteString(e.CreateVar(expression, fmt.Sprintf("%v", i)))

			if e.PreviousType == ironlang.IronLangLexerTYPE_INT && e.TempVarType == ironlang.IronLangLexerTYPE_INT {
				build.WriteString(addTemp + " = add nsw i32 " + e.TempVarName + ", " + e.PreviousName + "\n")
				e.PreviousType = ironlang.IronLangLexerTYPE_INT
			}

			if e.PreviousType == ironlang.IronLangLexerTYPE_FLOAT && e.TempVarType == ironlang.IronLangLexerTYPE_INT {
				castName := "%st" + e.GetStackId() + ".cst." + fmt.Sprintf("%v", i)
				build.WriteString(castName + " = fptosi i32 " + e.TempVarName + " to double\n")
				build.WriteString(addTemp + " = fadd float " + castName + ", " + e.PreviousName + "\n")
				e.PreviousType = ironlang.IronLangLexerTYPE_FLOAT
			}

			if e.PreviousType == ironlang.IronLangLexerTYPE_INT && e.TempVarType == ironlang.IronLangLexerTYPE_FLOAT {
				castName := "%st" + e.GetStackId() + ".cst." + fmt.Sprintf("%v", i)
				build.WriteString(castName + " = fptosi i32 " + e.PreviousName + " to double\n")
				build.WriteString(addTemp + " = fadd float " + e.TempVarName + ", " + castName + "\n")
				e.PreviousType = ironlang.IronLangLexerTYPE_FLOAT
			}

			e.PreviousName = addTemp

			//} else if expression.Operation == ironlang.IronLangLexerMINUS {
			//	addTemp := "%st" + e.GetStackId() + ".min." + fmt.Sprintf("%v", i)
			//	build.WriteString(e.CreateVar(expression, fmt.Sprintf("%v", i)))
			//	if expression.DataType == ironlang.IronLangLexerTYPE_INT {
			//		build.WriteString(addTemp + " = sub nsw i32 " + e.TempVarName + ", " + previous + "\n")
			//	} else {
			//		build.WriteString(addTemp + " = fsub float " + e.TempVarName + ", " + previous + "\n")
			//	}
			//	previous = addTemp
			//} else if expression.Operation == ironlang.IronLangLexerMULT {
			//	addTemp := "%st" + e.GetStackId() + ".mul." + fmt.Sprintf("%v", i)
			//	build.WriteString(e.CreateVar(expression, fmt.Sprintf("%v", i)))
			//	if expression.DataType == ironlang.IronLangLexerTYPE_INT {
			//		build.WriteString(addTemp + " = mul nsw i32 " + e.TempVarName + ", " + previous + "\n")
			//	} else {
			//		build.WriteString(addTemp + " = fmul float " + e.TempVarName + ", " + previous + "\n")
			//	}
			//	previous = addTemp
			//} else if expression.Operation == ironlang.IronLangLexerDIV {
			//	addTemp := "%st" + e.GetStackId() + "div." + fmt.Sprintf("%v", i)
			//	build.WriteString(e.CreateVar(expression, fmt.Sprintf("%v", i)))
			//	if expression.DataType == ironlang.IronLangLexerTYPE_INT {
			//		build.WriteString(addTemp + " = sdiv i32 " + e.TempVarName + ", " + previous + "\n")
			//	} else {
			//		build.WriteString(addTemp + " = fdiv float " + e.TempVarName + ", " + previous + "\n")
			//	}
			//	previous = addTemp
		} else {
			panic("Operação matemática não reconhecida")
		}

	}

	return build.String()
}

func (e *ExpressionStack) GetExpressions() []*Expression {
	return e.Stack
}
