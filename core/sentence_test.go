package core

import "testing"

func TestSentence(t *testing.T) {
	uut := NewSentence()
	if len(uut) != 0 {
		t.Errorf("Empty Sentence should be aliased slice")
	}

	uut = NewSentence(NewTerminalSymbol(symbolID))

	if len(uut) != 1 || int(uut[0]) != int(symbolID) {
		t.Errorf("Sentence should be aliased slice")
	}

}
