package core

import "testing"

var (
	terminalProduction = NewTerminalProduction(
	NewNonTerminalSymbol(uint('A')),
	NewTerminalSymbol(uint('a')))

	nonTerminalProduction = NewNonTerminalProduction(
		NewNonTerminalSymbol(uint('A')),
		NewNonTerminalSymbol(uint('B')),
		NewNonTerminalSymbol(uint('C')))
)

func TestEmptyPopulationContainsNoProductions(t *testing.T) {
	// Given:
	uut := NewPopulation()

	// When/Then:
	if uut.HasTerminal(&terminalProduction) {
		t.Errorf("Empty population has no terminal productions")
	}

	if uut.HasNonTerminal(&nonTerminalProduction) {
		t.Errorf("Empty population has no non terminal productions")
	}
}

func TestAddingTerminalProduction(t *testing.T) {
	// Given:
	uut := NewPopulation()

	// When/Then:
	uut.AddTerminal(terminalProduction)
	if !uut.HasTerminal(&terminalProduction) {
		t.Errorf("Terminal production has been added!")
	}

	if uut.HasNonTerminal(&nonTerminalProduction) {
		t.Errorf("Population has no non terminal productions")
	}
}

func TestAddingNonTerminalProduction(t *testing.T) {
	// Given:
	uut := NewPopulation()

	// When/Then:
	uut.AddNonTerminal(&nonTerminalProduction)
	if uut.HasTerminal(&terminalProduction) {
		t.Errorf("Population has no terminal productions!")
	}

	if !uut.HasNonTerminal(&nonTerminalProduction) {
		t.Errorf("NonTerminal production has been added!")
	}
}
