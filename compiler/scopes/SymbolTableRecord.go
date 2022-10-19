package scopes

type SymbolTableRecord struct {
	SymbolTables map[string]*SymbolTable
}

func NewSymbolTableRecord() *SymbolTableRecord {
	return &SymbolTableRecord{
		SymbolTables: make(map[string]*SymbolTable),
	}
}

func (s *SymbolTableRecord) insert(name string, tpe int) {
	s.SymbolTables[name] = &SymbolTable{
		Name: name,
		Type: tpe,
	}
}

func (s *SymbolTableRecord) GetSymbolTable(name string) *SymbolTable {
	return s.SymbolTables[name]
}
