package core

type uniqueTerminals map[TerminalProduction] bool

func (u *uniqueTerminals) Has(t *TerminalProduction) (has bool) {
	_, has = (*u)[*t]
	return
}
type setNonTerminals map[NonTerminalProduction] bool

type nonTerminalsRMap map[NonTerminalSymbolPair]setNonTerminals

func (n *nonTerminalsRMap) add(nt *NonTerminalProduction) {
	s, ok := (*n)[nt.Right]
	if !ok {
		s = make(setNonTerminals)
		(*n)[nt.Right] = s
	}
	s[*nt] = true
}

func (n *nonTerminalsRMap) has(nt *NonTerminalProduction) (result bool) {
	s, ok := (*n)[nt.Right]
	if ok {
		_, ok2 := s[*nt]
		result = ok2
	}
	return
}

type Population struct {
	rightTerminals    map[TerminalSymbol]uniqueTerminals
	rightNonTerminals nonTerminalsRMap
}

func (p *Population) HasTerminal(prod *TerminalProduction) (has bool) {
	if u, ok := p.rightTerminals[prod.Right]; ok {
		has = u.Has(prod)
	}
	return
}

func (p *Population) HasNonTerminal(prod *NonTerminalProduction) bool {
	return p.rightNonTerminals.has(prod)
}

func (p *Population) AddTerminal(t TerminalProduction)  {
	u, ok := p.rightTerminals[t.Right]
	if !ok {
		u = make(uniqueTerminals)
	}
	u[t] = true
	p.rightTerminals[t.Right] = u
}

func (p *Population) AddNonTerminal(n *NonTerminalProduction)  {
	p.rightNonTerminals.add(n)
}

func NewPopulation() *Population {
	return &Population{
		make(map[TerminalSymbol]uniqueTerminals),
		make(nonTerminalsRMap)}
}

