package test

import (
	"iron-lang/compiler"
	"testing"
)

func TestScopes(t *testing.T) {
	manager := compiler.NewScopesManager()

	if manager.HasScopes() == true {
		t.Fail()
	}

	manager.CreateNewScope()
	if manager.HasScopes() == false {
		t.Fail()
	}

	manager.GetActualScope().Insert("Test 1", 10)

	if manager.GetActualScope().GetSymbolTable("Test 1") == nil {
		t.Fail()
	}

	manager.DeleteActualScope()
	if manager.HasScopes() == true {
		t.Fail()
	}

	manager.CreateNewScope()
	manager.GetActualScope().Insert("Test 1", 10)
	manager.GetActualScope().Insert("Test 2", 10)

	manager.CreateNewScope()
	manager.GetActualScope().Insert("Test 3", 10)
	manager.GetActualScope().Insert("Test 4", 10)

	if manager.GetActualScope().GetSymbolTable("Test 1") != nil {
		t.Fail()
	}

	if manager.GetActualScope().GetSymbolTable("Test 2") != nil {
		t.Fail()
	}

	if manager.GetActualScope().GetSymbolTable("Test 3") == nil {
		t.Fail()
	}

	if manager.GetActualScope().GetSymbolTable("Test 4") == nil {
		t.Fail()
	}

	manager.DeleteActualScope()
	if manager.GetActualScope().GetSymbolTable("Test 3") != nil {
		t.Fail()
	}

	if manager.GetActualScope().GetSymbolTable("Test 4") != nil {
		t.Fail()
	}

	if manager.GetActualScope().GetSymbolTable("Test 1") == nil {
		t.Fail()
	}

	if manager.GetActualScope().GetSymbolTable("Test 2") == nil {
		t.Fail()
	}

	manager.DeleteActualScope()
	if manager.HasScopes() == true {
		t.Fail()
	}
}
