package test

import (
	"iron-lang/compiler/scopes"
	"testing"
)

func TestScopes(t *testing.T) {
	manager := scopes.NewScopesManager()

	if manager.HasScopes() == true {
		t.Fail()
	}

	manager.CreateNewScope()
	if manager.HasScopes() == false {
		t.Fail()
	}

	manager.RegisterVariable("Test 1", 10)

	if manager.GetVariable("Test 1") == nil {
		t.Fail()
	}

	manager.DeleteActualScope()
	if manager.HasScopes() == true {
		t.Fail()
	}

	manager.CreateNewScope()
	manager.RegisterVariable("Test 1", 10)
	manager.RegisterVariable("Test 2", 10)

	manager.CreateNewScope()
	manager.RegisterVariable("Test 3", 10)
	manager.RegisterVariable("Test 4", 10)

	if manager.GetVariable("Test 1") != nil {
		t.Fail()
	}

	if manager.GetVariable("Test 2") != nil {
		t.Fail()
	}

	if manager.GetVariable("Test 3") == nil {
		t.Fail()
	}

	if manager.GetVariable("Test 4") == nil {
		t.Fail()
	}

	manager.DeleteActualScope()
	if manager.GetVariable("Test 3") != nil {
		t.Fail()
	}

	if manager.GetVariable("Test 4") != nil {
		t.Fail()
	}

	if manager.GetVariable("Test 1") == nil {
		t.Fail()
	}

	if manager.GetVariable("Test 2") == nil {
		t.Fail()
	}

	manager.DeleteActualScope()
	if manager.HasScopes() == true {
		t.Fail()
	}
}
