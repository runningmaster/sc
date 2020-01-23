package parser

import (
	"fmt"
)

// tokenType identifies the type of lex tokens.
type TokenType int

const (
	tokenError TokenType = iota // error occurred; value is text of error
	tokenEOF
	tokenBracketLeft  // '[' inside action
	tokenBracketRight // ']' inside action
	tokenIdentifier   // alphanumeric identifier not starting with '.'
	tokenKeyword      // used only to delimit the keywords
	TokenSUM
	TokenINT
	TokenDIF
)

const eof = -1

// Name makes the types prettyprint.
func (t TokenType) Name() string {
	switch t {
	case tokenError:
		return "error"
	case tokenEOF:
		return "EOF"
	case tokenBracketLeft:
		return "["
	case tokenBracketRight:
		return "]"
	case tokenIdentifier:
		return "identifier"
	case tokenKeyword:
		return "keyword"
	case TokenSUM:
		return "SUM"
	case TokenINT:
		return "INT"
	case TokenDIF:
		return "DIF"
	default:
		return fmt.Sprintf("token%d", int(t))
	}
}

func (t TokenType) String() string {
	return t.Name()
}

// token represents a token or text string returned from the scanner.
type token struct {
	typ   TokenType // The type of this token.
	val   string    // The value of this token.
	pos   int       // The starting position, in bytes, of this token in the input string.
	depth int
}

func (t token) String() string {
	const lenVal = 10

	switch {
	case t.typ == tokenEOF:
		return tokenEOF.Name()
	case t.typ == tokenError:
		return t.val
	case t.typ > tokenKeyword:
		return fmt.Sprintf("[%s]", t.val)
	case len(t.val) > lenVal:
		return fmt.Sprintf("%.10q...", t.val)
	}

	return fmt.Sprintf("%q", t.val)
}
