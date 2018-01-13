package core

import "testing"

func TestTerminalProduction(t *testing.T) {
	uut := NewTerminalProduction(
		NewNonTerminalSymbol(uint('A')),
		NewTerminalSymbol(uint('a')))

	if uut.Left != NewNonTerminalSymbol(uint('A')) {
		t.Errorf("Problem with storing TerminalProduction.Left")
	}

	if uut.Right != NewTerminalSymbol(uint('a')) {
		t.Errorf("Problem with storing TerminalProduction.Right")
	}
}

func TestNonTerminalProduction(t *testing.T) {
	uut := NewNonTerminalProduction(
		NewNonTerminalSymbol(uint('A')),
		NewNonTerminalSymbol(uint('B')),
		NewNonTerminalSymbol(uint('C')))

	if uut.Left != NewNonTerminalSymbol(uint('A')) {
		t.Errorf("Problem with storing NonTerminalProduction.Left")
	}

	if uut.Right.First != NewNonTerminalSymbol(uint('B')) {
		t.Errorf(
			"Problem with storing NonTerminalProduction.Right.First")
	}

	if uut.Right.Second != NewNonTerminalSymbol(uint('C')) {
		t.Errorf(
			"Problem with storing NonTerminalProduction.Right.Second")
	}
}
