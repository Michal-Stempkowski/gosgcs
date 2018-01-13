package core

const (
	// EmptySymbol represents empty symbol (small epsilon) of  grammar.
	EmptySymbol = TerminalSymbol(-1)
	// StartingSymbol represents starting symbol (big s) of  grammar.
	StartingSymbol = NonTerminalSymbol(-2)
	// UniversalSymbol represents special universal symbol (big s with index u)
	// of  grammar.
	UniversalSymbol = NonTerminalSymbol(-3)
)

// TerminalSymbol is internal sgcs symbol representation for grammar terminal
// symbols.
type TerminalSymbol int

// NewTerminalSymbol creates new instance of TerminalSymbol.
func NewTerminalSymbol(id uint) TerminalSymbol {
	return TerminalSymbol(id)
}

// NonTerminalSymbol is internal sgcs symbol representation for grammar
// non terminal symbols.
type NonTerminalSymbol int

// NewNonTerminalSymbol creates new instance of TerminalSymbol.
func NewNonTerminalSymbol(id uint) NonTerminalSymbol {
	return NonTerminalSymbol(id)
}
