package core

const (
	EmptySymbol = TerminalSymbol(-1)
	StartingSymbol = NonTerminalSymbol(-2)
	UniversalSymbol = NonTerminalSymbol(-3)
)

// TerminalSymbol is internal sgcs symbol representation for grammar terminal
// symbols.
type TerminalSymbol int

// NewTerminalSymbol creates new instance of TerminalSymbol.
func NewTerminalSymbol(id int) TerminalSymbol {
	return TerminalSymbol(id)
}

// NonTerminalSymbol is internal sgcs symbol representation for grammar
// non terminal symbols.
type NonTerminalSymbol int

// NewNonTerminalSymbol creates new instance of TerminalSymbol.
func NewNonTerminalSymbol(id int) NonTerminalSymbol {
	return NonTerminalSymbol(id)
}
