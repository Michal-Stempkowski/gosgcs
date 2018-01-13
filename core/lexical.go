package core

// Internal sgcs symbol representation.
type TerminalSymbol int

// Creates new instance of TerminalSymbol.
func NewTerminalWymbol(id int) TerminalSymbol {
	return TerminalSymbol(id)
}

// Sentence is a sequence of Symbols.
type Sentence []TerminalSymbol

// Creates new sequence of Symbols.
func NewSentence(args ...TerminalSymbol) Sentence {
	return args
}
