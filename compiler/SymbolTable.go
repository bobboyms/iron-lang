package compiler

type SymbolTable struct {
	Name string
	Type int
}

type SymbolTableRecord struct {
	SymbolTables map[string]*SymbolTable
}

func NewSymbolTableRecord() *SymbolTableRecord {
	return &SymbolTableRecord{
		SymbolTables: make(map[string]*SymbolTable),
	}
}

func (s *SymbolTableRecord) Insert(name string, tpe int) {
	s.SymbolTables[name] = &SymbolTable{
		Name: name,
		Type: tpe,
	}
}

func (s *SymbolTableRecord) GetSymbolTable(name string) *SymbolTable {
	return s.SymbolTables[name]
}

type ScopesManager struct {
	SymbolTableStack []*SymbolTableRecord
}

func NewScopesManager() *ScopesManager {
	return &ScopesManager{}
}

func (s *ScopesManager) CreateNewScope() {
	s.SymbolTableStack = append(s.SymbolTableStack, NewSymbolTableRecord())
}
func (s *ScopesManager) NoHasSymbolTableHigherScopes(name string) bool {

	size := len(s.SymbolTableStack)

	if size <= 0 {
		return true
	}

	for i := size - 1; i >= 0; i-- {
		if s.SymbolTableStack[i].GetSymbolTable(name) != nil {
			return false
		}
	}

	return true
}

func (s *ScopesManager) GetActualScope() *SymbolTableRecord {

	size := len(s.SymbolTableStack) - 1

	if size < 0 {
		return nil
	}

	return s.SymbolTableStack[size]
}

func (s *ScopesManager) GetAllScopes() []*SymbolTableRecord {
	return s.SymbolTableStack
}

func (s *ScopesManager) HasScopes() bool {
	if len(s.SymbolTableStack) == 0 {
		return false
	}
	return true
}

func (s *ScopesManager) DeleteActualScope() {

	size := len(s.SymbolTableStack)

	if size <= 0 {
		return
	}

	if size == 1 {
		s.SymbolTableStack = make([]*SymbolTableRecord, 0)
		return
	}

	var symbolTableStack []*SymbolTableRecord
	for i := size - 2; i >= 0; i-- {
		symbolTableStack = append(symbolTableStack, s.SymbolTableStack[i])
	}

	s.SymbolTableStack = symbolTableStack

}
