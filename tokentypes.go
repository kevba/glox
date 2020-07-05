package main

import "fmt"

// Token tracks anyhting for a token. Tokens contains things like te lexeme and its location.
//   A token gets created by the scanner.
type Token struct {
	ttype   TokenType
	lexeme  string
	literal interface{}
	line    int
	column  int // not yet used
}

// NewToken creates a new token.
func NewToken(ttype TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		ttype:   ttype,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%v %v %v", t.ttype, t.lexeme, t.literal)
}

// TokenType represents the type of token.
type TokenType int

// Single-character tokens.
const (
	LeftParenrightParen TokenType = iota
	LeftBrace           TokenType = iota
	RightBrace          TokenType = iota
	Comma               TokenType = iota
	Dot                 TokenType = iota
	Minus               TokenType = iota
	Plus                TokenType = iota
	Semicolon           TokenType = iota
	Slash               TokenType = iota
	Star                TokenType = iota

	// One or two character tokens.
	Bang         TokenType = iota
	BangEqual    TokenType = iota
	Equal        TokenType = iota
	EqualEqual   TokenType = iota
	Greater      TokenType = iota
	GreaterEqual TokenType = iota
	Less         TokenType = iota
	LessEqual    TokenType = iota

	// Literals.
	Identifier TokenType = iota
	String     TokenType = iota
	Number     TokenType = iota

	// Keywords.
	And    TokenType = iota
	Class  TokenType = iota
	Else   TokenType = iota
	False  TokenType = iota
	Fun    TokenType = iota
	For    TokenType = iota
	If     TokenType = iota
	Nil    TokenType = iota
	Or     TokenType = iota
	Print  TokenType = iota
	Return TokenType = iota
	Super  TokenType = iota
	This   TokenType = iota
	True   TokenType = iota
	Var    TokenType = iota
	While  TokenType = iota

	EOF TokenType = iota
)

var keywords = map[string]TokenType{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"fun":    Fun,
	"for":    For,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}
