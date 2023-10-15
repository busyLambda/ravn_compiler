package parser

import (
	"fmt"
	"os"
)

type Parser struct {
	scanner *Scanner
	module  string
	ast     *Root
}

func NewParser(input string, module string) *Parser {
	return &Parser{
		scanner: NewScanner(input),
		ast:     NewAstRoot(module),
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
		fmt.Printf("Expected identifier after keyword `fn` instead found -> %s\n", token.String())
	} else {
		fd.Name = NewIdentifier(literal, Span{}, Object{FUNC, literal})
		tokenKind, literal = p.ScanSkipWhitespace()
		token := Token{tokenKind, literal}
		switch tokenKind {
		case L_BRACK:
			t_ype, err := p.parseFuncType()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fd.Type = t_ype
		default:
			fmt.Printf("Expected `(` after identifier in function declaration instead found -> %s\n", token.String())
		}
	}

	return
}

func (p *Parser) parseFuncType() (*FuncType, error) {
	var ft FuncType

	ft.Params = []FuncParam{}

	for {
		tokenKind, literal := p.ScanSkipWhitespace()
		token := Token{tokenKind, literal}
		switch {
		case tokenKind == EOF:
			return nil, fmt.Errorf("Unclosed function: expected `)` found -> %s\n", token.String())
		case tokenKind == R_BRACK:
			return &ft, nil
		case tokenKind == COMMA && len(ft.Params) > 0:
			tokenKind, literal := p.ScanSkipWhitespace()
			switch tokenKind {
			case IDENT:
				var param FuncParam
				param.Ident = NewIdentifier(literal, Span{}, Object{FUNC_PARAM, literal})

				tokenKind, literal := p.ScanSkipWhitespace()
				token := Token{tokenKind, literal}

				// Get the parameter
				switch tokenKind {
				case IDENT:
					param.Type = NewIdentifier(literal, Span{}, Object{FUNC_PARAM_TYPE, literal})
					ft.Params = append(ft.Params, param)
				default:
					return nil, fmt.Errorf("Expected type after function parameter identifier, instead found -> `%s`\n", token.String())
				}
			}
		case tokenKind == IDENT && len(ft.Params) == 0:
			var param FuncParam
			param.Ident = NewIdentifier(literal, Span{}, Object{FUNC_PARAM, literal})

			tokenKind, literal := p.ScanSkipWhitespace()
			token := Token{tokenKind, literal}

			// Get the parameter
			switch tokenKind {
			case IDENT:
				param.Type = NewIdentifier(literal, Span{}, Object{FUNC_PARAM_TYPE, literal})
				ft.Params = append(ft.Params, param)
			default:
				return nil, fmt.Errorf("Expected type after function parameter identifier, instead found -> `%s`\n", token.String())
			}
		}
	}

}

func (p *Parser) ScanSkipWhitespace() (TokenKind, string) {
	for {
		tokenKind, literal := p.scanner.Scan()

		if tokenKind != WHITESPACE {
			return tokenKind, literal
		}
	}
}
