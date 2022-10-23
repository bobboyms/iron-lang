package codegenerator

import (
	"strings"
)

const (
	FuncMainMode = 1
	FuncMode     = 2
	ImportMode   = 3
)

type StringBuilder struct {
	MainStrBuilder *strings.Builder
	ModeOperation  int
	Counter        int
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{
		MainStrBuilder: &strings.Builder{},
		Counter:        -1,
	}
}

//func (s *StringBuilder) ChangeScopeToFuncMode() {
//	s.ModeOperation = FuncMode
//}
//
//func (s *StringBuilder) ChangeScopeToMainMode() {
//	s.ModeOperation = FuncMainMode
//}
//
//func (s *StringBuilder) ChangeScopeToImportMode() {
//	s.ModeOperation = ImportMode
//}

func (s *StringBuilder) WriteString(value string) {

	s.MainStrBuilder.WriteString(value)

	//if s.ModeOperation == FuncMode {
	//	s.FuncStrBuilder.WriteString(value)
	//} else if s.ModeOperation == FuncMainMode {
	//	s.MainStrBuilder.WriteString(value)
	//} else if s.ModeOperation == ImportMode {
	//	s.ImportStrBuilder.WriteString(value)
	//} else {
	//	panic("Tipo de operação de escrita não definido")
	//}
}

func (s *StringBuilder) GetBuilder() *strings.Builder {
	builder := &strings.Builder{}
	//builder.WriteString(s.ImportStrBuilder.String())
	////builder.WriteString(s.FuncStrBuilder.String())
	builder.WriteString(s.MainStrBuilder.String())
	return builder
}
