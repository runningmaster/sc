package parser

import (
	"errors"
)

// Parse makes AST
func Parse(input string) (*Node, error) {
	lex := lex(input)

	var tree, n *Node

	for {
		token := lex.nextToken()
		if token.typ == tokenEOF {
			break
		}

		if token.typ == tokenError {
			return nil, errors.New(token.val)
		}

		switch token.typ {
		case tokenBracketLeft:
			n = &Node{prev: n, depth: token.depth - 1}

			if tree == nil {
				tree = n
			} else if n.prev != nil {
				n.prev.next = append(n.prev.next, n)
			}

		case tokenBracketRight:
			if n == nil {
				return nil, errors.New("syntax error n is nil")
			}

			if n.prev != nil {
				n = n.prev
			}

		case TokenSUM, TokenINT, TokenDIF:
			if n == nil {
				return nil, errors.New("syntax error n is nil")
			}

			n.typ = token.typ

		case tokenIdentifier:
			if n == nil {
				return nil, errors.New("syntax error n is nil")
			}

			n.vals = append(n.vals, token.val)
		}
	}

	lex.drain()

	return tree, nil
}
