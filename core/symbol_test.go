package core

import "testing"

const symbolID = int('a')

func TestTerminalSymbol(t *testing.T) {
	uut := NewTerminalSymbol(symbolID)

	if int(uut) != symbolID {
		t.Errorf("TerminalSymbol should be aliased int")
	}

	if //noinspection GoBoolExpressions
	TerminalSymbol(-1) != EmptySymbol {
		t.Errorf("EmptySymbol is represented by -1")
	}
}

func TestNonTerminalSymbol(t *testing.T) {
	uut := NewNonTerminalSymbol(symbolID)

	if int(uut) != symbolID {
		t.Errorf("NewNonTerminalSymbol should be aliased int")
	}

	if //noinspection GoBoolExpressions
	NonTerminalSymbol(-2) != StartingSymbol {
		t.Errorf("StartingSymbol is represented by -2")
	}

	if //noinspection GoBoolExpressions
	NonTerminalSymbol(-3) != UniversalSymbol {
		t.Errorf("UniversalSymbol is represented by -3")
	}
}
