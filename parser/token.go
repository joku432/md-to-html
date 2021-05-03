package parser

type Kind uint

const (
	// Operators
	PLUS Kind = 0
	MINUS
	MUL
	DIV

	IDENT
	NUMBER
)

type Token struct {
	Kind    Kind
	Content string
}
