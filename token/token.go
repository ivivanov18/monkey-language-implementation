package token

type TokenType string

type Token struct {
	Type TokenType
	Litteral string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + litterals
)