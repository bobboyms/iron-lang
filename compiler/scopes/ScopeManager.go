package scopes

type Variable struct {
	Name string
	Type int
}

type ScopeManager struct {
	ScopeVariables map[int][]*Variable
	ActualScopeId  int
}

func NewScopesManager() *ScopeManager {
	return &ScopeManager{
		ScopeVariables: make(map[int][]*Variable),
		ActualScopeId:  -1,
	}
}

func (s *ScopeManager) CreateNewScope() {
	s.ActualScopeId += 1
	s.ScopeVariables[s.ActualScopeId] = make([]*Variable, 0)
}

func (s *ScopeManager) RegisterVariable(name string, tpe int) {

	var variables = s.ScopeVariables[s.ActualScopeId]

	variables = append(variables, &Variable{
		Name: name,
		Type: tpe,
	})

	s.ScopeVariables[s.ActualScopeId] = variables
}

func (s *ScopeManager) GetVariable(name string) *Variable {
	for _, variable := range s.ScopeVariables[s.ActualScopeId] {
		if name == variable.Name {
			return variable
		}
	}
	return nil
}

func (s *ScopeManager) NoHasSymbolTableHigherScopes(name string) bool {

	for i := s.ActualScopeId; i >= 0; i-- {

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
	delete(s.ScopeVariables, s.ActualScopeId)
	s.ActualScopeId -= 1
}
