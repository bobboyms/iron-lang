package scopes

type ScopeLog struct {
	ScopeLogs map[string][]*Variable
	ScopeHash string
}

//func NewScopeLogManager() *scopeLog {
//	return &scopeLog{
//		ScopeLogs: make(map[string][]*Variable),
//	}
//}

func (s *ScopeLog) EnterScope(scopeHash string) {
	s.ScopeHash = scopeHash
}

func (s *ScopeLog) ExitScope() {
	s.ScopeHash = ""
}

func (s *ScopeLog) GetVariable(name string) *Variable {
	for _, variable := range s.ScopeLogs[s.ScopeHash] {
		if variable.Name == name {
			return variable
		}
	}
	return nil
}

func (s *ScopeLog) registerVariable(variable *Variable, scopeHash string) {
	var variables = s.ScopeLogs[scopeHash]
	variables = append(variables, variable)
	s.ScopeLogs[scopeHash] = variables
}
