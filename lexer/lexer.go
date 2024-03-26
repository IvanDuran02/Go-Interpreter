package lexer

type Lexer struct {
	input        string
	position     int  // current pos in input (points to current char)
	readPosition int  // current reading pos in input (after char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
