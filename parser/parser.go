package parser

import (
	"fmt"
)

type Parser struct {
	scanner *Scanner
}

func NewParser(input string) *Parser {
	return &Parser{
		scanner: NewScanner(input),
	}
}

func (p *Parser) NextNode() {
	for {
		tokenKind, _ := p.ScanSkipWhitespace()
		if tokenKind == EOF {
			break
		}

		switch tokenKind {
		case KW_FN:
			p.parseFuncDecl()
		default:
			break
		}
	}
}

func (p *Parser) parseFuncDecl() (fd FuncDecl) {
	tokenKind, literal := p.ScanSkipWhitespace()
	token := Token{tokenKind, literal}

	if tokenKind != IDENT {
		panic(fmt.Sprintf("Expected identifier after keyword `fn` instead found -> %s", token.String()))
	}

	return
}

func (p *Parser) ScanSkipWhitespace() (TokenKind, string) {
	for {
		tokenKind, literal := p.scanner.Scan()

		if tokenKind != WHITESPACE {
			return tokenKind, literal
		}
	}
}
