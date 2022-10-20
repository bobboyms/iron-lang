package scopes

type Variable struct {
	Name string
	Type int
}

type ScopeManager struct {
	scopeLog       *ScopeLog
	scopeHash      string
	actualScopeId  int
	ScopeVariables map[int][]*Variable
}

func NewScopesManager() *ScopeManager {
	return &ScopeManager{
		ScopeVariables: make(map[int][]*Variable),
		actualScopeId:  -1,
		scopeLog: &ScopeLog{
			ScopeLogs: map[string][]*Variable{},
		},
	}
}

func (s *ScopeManager) CreateNewScope(scopeHash string) {
	s.actualScopeId += 1
	s.scopeHash = scopeHash
	s.ScopeVariables[s.actualScopeId] = make([]*Variable, 0)
}

func (s *ScopeManager) GetScopeLog() *ScopeLog {
	return s.scopeLog
}

func (s *ScopeManager) RegisterVariable(name string, tpe int) {

	var variables = s.ScopeVariables[s.actualScopeId]

	variable := &Variable{
		Name: name,
		Type: tpe,
	}

	variables = append(variables, variable)
	s.scopeLog.registerVariable(variable, s.scopeHash)
	s.ScopeVariables[s.actualScopeId] = variables
}

func (s *ScopeManager) GetVariable(name string) *Variable {
	for _, variable := range s.ScopeVariables[s.actualScopeId] {
		if name == variable.Name {
			return variable
		}
	}
	return nil
}

func (s *ScopeManager) NoHasSymbolTableHigherScopes(name string) bool {

	for i := s.actualScopeId; i >= 0; i-- {

		for _, variable := range s.ScopeVariables[i] {
			if name == variable.Name {
				return false
			}
		}

	}

	return true
}

func (s *ScopeManager) HasScopes() bool {
	if len(s.ScopeVariables) == 0 {
		return false
	}
	return true
}

func (s *ScopeManager) DeleteActualScope() {
	delete(s.ScopeVariables, s.actualScopeId)
	s.actualScopeId -= 1
	s.scopeHash = ""
}
