package errors

const (
	VariableNotDeclared = 100
	VariableHasDeclared = 101
	VariableNotDefined  = 102
)

type CodeError int

func GetMessageError(codeError CodeError, identifier string) string {
	if codeError == VariableNotDeclared {
		return "variable << " + identifier + " >> not declared before use"
	}
	if codeError == VariableHasDeclared {
		return "The variable << " + identifier + " >> has already been declared within the scope"
	}
	if codeError == VariableNotDefined {
		return "The type of variable << " + identifier + " >> has not been defined"
	}
	panic("Code error not been defined")
}
