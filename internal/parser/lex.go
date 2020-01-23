package parser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// lexer holds the state of the scanner.
type lexer struct {
	input  string     // the string being scanned
	pos    int        // current position in the input
	start  int        // start position of this item
	width  int        // width of last rune read from input
	depth  int        // nesting depth of [ ] exprs
	tokens chan token // channel of scanned tokens
}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lex creates a new scanner for the input string.
func lex(input string) *lexer {
	l := &lexer{
		input:  input,
		tokens: make(chan token),
	}

	go l.run()

	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for state := lexAction; state != nil; {
		state = state(l)
	}

	close(l.tokens)
}

// nextToken returns the next token from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextToken() token {
	return <-l.tokens
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextToken.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- token{tokenError, fmt.Sprintf(format, args...), l.start, l.depth}
	return nil
}

// emit passes an token back to the client.
func (l *lexer) emit(t TokenType) {
	l.tokens <- token{t, l.input[l.start:l.pos], l.start, l.depth}
	l.start = l.pos
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width

	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()

	return r
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.tokens {
	}
}

// atTerminator reports whether the input is at valid termination character to
// appear after an identifier. Breaks .X.Y into two pieces. Also catches cases
// like "$x+2" not being acceptable without a space, in case we decide one
// day to implement arithmetic.
func (l *lexer) atTerminator() bool {
	r := l.peek()
	if isSpace(r) || isEndOfLine(r) {
		return true
	}

	switch r {
	case eof, '.', ',', '|', ':', ')', '(', '[', ']':
		return true
	}

	return false
}

// lexAction scans the elements inside action delimiters.
func lexAction(l *lexer) stateFn {
	switch r := l.next(); {
	case r == eof:
		l.emit(tokenEOF)
		return nil
	case isEndOfLine(r):
		return l.errorf("unclosed action")
	case isSpace(r):
		l.backup()
		return lexSpace
	case r == '"':
		return lexQuote
	case isAlphaNumeric(r):
		l.backup()
		return lexIdentifier
	case r == '[':
		l.depth++
		l.emit(tokenBracketLeft)

		return lexAction
	case r == ']':
		l.emit(tokenBracketRight)

		l.depth--
		if l.depth < 0 {
			return l.errorf("unexpected right bracket %#U", r)
		}

		return lexAction
	case r <= unicode.MaxASCII && unicode.IsPrint(r):
		return lexAction
	default:
		return l.errorf("unrecognized character in action: %#U", r)
	}
}

// lexSpace scans a run of space characters.
// We have not consumed the first space, which is known to be present.
// Take care if there is a trim-marked right delimiter, which starts with a space.
func lexSpace(l *lexer) stateFn {
	var r rune

	for {
		r = l.peek()
		if !isSpace(r) {
			break
		}

		l.next()
	}

	l.ignore()

	return lexAction
}

// lexQuote scans a quoted string.
func lexQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(tokenIdentifier)

	return lexAction
}

// lexIdentifier scans an alphanumeric.
func lexIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			if !l.atTerminator() {
				return l.errorf("bad character %#U", r)
			}
			switch {
			case key(word) > tokenKeyword:
				l.emit(key(word))
			default:
				l.emit(tokenIdentifier)
			}
			break Loop
		}
	}

	return lexAction
}

func key(s string) TokenType {
	switch strings.ToLower(s) {
	case "sum":
		return TokenSUM
	case "int":
		return TokenINT
	case "dif":
		return TokenDIF
	default:
		return tokenError
	}
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
