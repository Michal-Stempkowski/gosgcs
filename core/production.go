package core

// TerminalProduction stores representation of grammar terminal production.
type TerminalProduction struct {
	Left  NonTerminalSymbol
	Right TerminalSymbol
}

// NewTerminalProduction creates new instance of TerminalProduction.
func NewTerminalProduction(
	left NonTerminalSymbol, right TerminalSymbol) TerminalProduction {
	return TerminalProduction{Left: left, Right: right}
}

// NonTerminalProduction stores representation of grammar terminal production.
type NonTerminalProduction struct {
	Left  NonTerminalSymbol
	Right NonTerminalSymbolPair
}

// NewNonTerminalProduction creates new instance of NonTerminalProduction.
func NewNonTerminalProduction(
	left, rightA, rightB NonTerminalSymbol) NonTerminalProduction {
	return NonTerminalProduction{
		Left:  left,
		Right: NewNonTerminalSymbolPair(rightA, rightB)}
}
