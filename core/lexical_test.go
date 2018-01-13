package core

import "testing"

const symbolId = int('a')

func TestTerminalSymbol(t *testing.T) {
	uut := NewTerminalWymbol(symbolId)

	if int(uut) != symbolId {
		t.Errorf("TerminalSymbol should be aliased int")
	}
}

func TestSentence(t *testing.T) {
	uut := NewSentence()
	if len(uut) != 0 {
		t.Errorf("Empty Sentence should be aliased slice")
	}

	uut = NewSentence(NewTerminalWymbol(symbolId))

	if len(uut) != 1 || int(uut[0]) != symbolId {
		t.Errorf("Sentence should be aliased slice")
	}

}
