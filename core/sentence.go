package core

// Sentence is a sequence of Symbols.
type Sentence []TerminalSymbol

// NewSentence creates new sequence of Symbols.
func NewSentence(args ...TerminalSymbol) Sentence {
	return args
}
