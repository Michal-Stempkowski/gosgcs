package core

import "testing"

const symbolID = int('a')

func TestTerminalSymbol(t *testing.T) {
	uut := NewTerminalSymbol(symbolID)

	if int(uut) != symbolID {
		t.Errorf("TerminalSymbol should be aliased int")
	}
}

func TestSentence(t *testing.T) {
	uut := NewSentence()
	if len(uut) != 0 {
		t.Errorf("Empty Sentence should be aliased slice")
	}

	uut = NewSentence(NewTerminalSymbol(symbolID))

	if len(uut) != 1 || int(uut[0]) != symbolID {
		t.Errorf("Sentence should be aliased slice")
	}

}
