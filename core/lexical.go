package core

// TerminalSymbol is internal sgcs symbol representation for grammar terminal
// symbols.
type TerminalSymbol int

// NewTerminalSymbol creates new instance of TerminalSymbol.
func NewTerminalSymbol(id int) TerminalSymbol {
	return TerminalSymbol(id)
}

// Sentence is a sequence of Symbols.
type Sentence []TerminalSymbol

// NewSentence creates new sequence of Symbols.
func NewSentence(args ...TerminalSymbol) Sentence {
	return args
}
