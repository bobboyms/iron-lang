package compiler

type SymbolTable struct {
	Name  string
	Type  int
	Value string
}

type SymbolTableRecord struct {
	SymbolTables map[string]*SymbolTable
}

func NewSymbolTableRecord() *SymbolTableRecord {
	return &SymbolTableRecord{
		SymbolTables: make(map[string]*SymbolTable),
	}
}

func (s *SymbolTableRecord) Insert(name string, tpe int, value string) {
	s.SymbolTables[name] = &SymbolTable{
		Name:  name,
		Type:  tpe,
		Value: value,
	}
}

func (s *SymbolTableRecord) GetSymbolTable(name string) *SymbolTable {
	return s.SymbolTables[name]
}

type Scopes struct {
	SymbolTableStack []*SymbolTableRecord
}

func NewScopes() *Scopes {
	return &Scopes{}
}

func (s *Scopes) CreateNewScope() {
	s.SymbolTableStack = append(s.SymbolTableStack, NewSymbolTableRecord())
}

func (s *Scopes) GetActualScope() *SymbolTableRecord {

	size := len(s.SymbolTableStack) - 1

	if size < 0 {
		return nil
	}

	return s.SymbolTableStack[size]
}

func (s *Scopes) GetAllScopes() []*SymbolTableRecord {
	return s.SymbolTableStack
}
