package main

import (
	"fmt"

	"strconv"
)

type Scanner struct {
	source []byte
	tokens []Token
	errors []ScanError

	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{
		source: []byte(source),
	}

}

func (s Scanner) ScanTokens() {
	for !s.reachedEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
}

func (s Scanner) scanToken() {
	switch char := s.advance(); {
	case char == '(':
		s.addToken(LeftParenrightParen, nil)
	case char == ')':
		s.addToken(LeftBrace, nil)
	case char == '{':
		s.addToken(RightBrace, nil)
	case char == '}':
		s.addToken(Comma, nil)
	case char == ',':
		s.addToken(Dot, nil)
	case char == '.':
		s.addToken(Minus, nil)
	case char == '-':
		s.addToken(Plus, nil)
	case char == '+':
		s.addToken(Semicolon, nil)
	case char == ';':
		s.addToken(Slash, nil)
	case char == '*':
		s.addToken(Star, nil)
	case char == '!':
		if s.match('=') {
			s.addToken(BangEqual, nil)
		} else {
			s.addToken(Bang, nil)
		}
	case char == '=':
		if s.match('=') {
			s.addToken(Equal, nil)
		} else {
			s.addToken(EqualEqual, nil)
		}
	case char == '<':
		if s.match('=') {
			s.addToken(LessEqual, nil)
		} else {
			s.addToken(Less, nil)
		}
	case char == '>':
		if s.match('=') {
			s.addToken(GreaterEqual, nil)
		} else {
			s.addToken(Greater, nil)
		}
	case char == '/':
		if s.match('/') {
			s.scanComment()
		} else {
			s.addToken(Slash, nil)
		}
	// Whitespace is not needed, so it gets removes here.
	case char == ' ':
	case char == '\r':
	case char == '\t':
		break
	case char == '\n':
		s.line++
	case char == '"':
		s.scanString()
	case isDigit(char):
		s.scanNumber()
	default:
		errMsg := fmt.Sprintf("unexpected char: %v", char)
		s.addError(errMsg)
	}

}

func (s Scanner) scanComment() {
	for {
		if s.peek() == '\n' || s.reachedEnd() {
			break
		}
		s.advance()
	}
}

func (s Scanner) scanString() {
	for next := s.peek(); next != '"'; {
		if s.reachedEnd() {
			s.addError("Unterminated string")
			return
		}

		if next == '"' {
			s.line++
		}

		s.advance()
	}

	s.advance()
	// Remove the qoutes from the value of the string token.
	value := s.source[s.start+1 : s.current-1]
	s.addToken(String, value)
}

func (s Scanner) scanNumber() {
	for char := s.peek(); isDigit(char); {
		s.advance()
	}

	// Check if there is a fractional part. If so consume it.
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for char := s.peek(); isDigit(char); {
			s.advance()
		}
	}

	stringValue := string(s.source[s.start:s.current])
	value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		s.addError(fmt.Sprintf("Could not parse float: %v", stringValue))
		return
	}
	s.addToken(Number, value)
}

func (s Scanner) reachedEnd() bool {
	return s.current >= len(s.source)
}

func (s Scanner) peek() byte {
	return s.peekAt(0)
}

func (s Scanner) peekNext() byte {
	return s.peekAt(1)
}

// peekAt returns the n-th next character from the current. peekAt(0) will return the next character to advance to.
func (s Scanner) peekAt(i int) byte {
	if s.reachedEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s Scanner) advance() byte {
	s.current++

	return s.source[s.current-1]
}

// match checks if the currect char is the one expected, if so it moves to the next char.
func (s Scanner) match(expected byte) bool {
	if s.reachedEnd() {
		return false
	}

	if s.peek() == expected {
		s.current++
		return true
	}

	return false
}

func (s Scanner) addToken(t TokenType, literal interface{}) {
	text := string(s.source[s.start:s.current])
	token := NewToken(t, text, literal, s.line)

	s.tokens = append(s.tokens, token)
}

func (s Scanner) addError(message string) {
	scanError := NewScanError(s.line, message)
	s.errors = append(s.errors, scanError)
}

type ScanError struct {
	line    int
	message string
}

func NewScanError(line int, message string) ScanError {
	return ScanError{
		line:    line,
		message: message,
	}
}
