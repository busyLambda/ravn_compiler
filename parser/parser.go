package parser

import (
	"fmt"
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
			// Get the function args
			t_ype, err := p.parseFuncType()
			if err != nil {
				fmt.Println(err)
			}

			fd.Type = t_ype
			tokenKind, literal = p.ScanSkipWhitespace()
			// token := Token{tokenKind, literal}
			switch tokenKind {
			// Got function body
			case L_CURLY:
				bs, err := p.parseBlockStmt()
				if err != nil {
					fmt.Println(err)
				}
				fd.Body = &bs
			case KW_MATCH:
				//p.parseMatchStmt()
			case IDENT:
				// Get the type
			case L_BLOCK:
				// Get the array type
			}
		default:
			fmt.Printf("Expected `(` after identifier in function declaration instead found -> %s\n", token.String())
		}
	}

	return
}

// Not yet accouting for <> types like Option<T>
func (p *Parser) parseType() (lit string, err error) {
	tokenKind, literal := p.ScanSkipWhitespace()
	token := Token{tokenKind, literal}
	switch tokenKind {
	case IDENT:
		lit = literal
		return
	case L_BLOCK:
		// Get the array type
		tokenKind, literal := p.ScanSkipWhitespace()
		token := Token{tokenKind, literal}
		switch tokenKind {
		case IDENT:
			tokenKind, literal := p.ScanSkipWhitespace()
			token := Token{tokenKind, literal}
			switch tokenKind {
			case R_BLOCK:
				return fmt.Sprintf("[%s]", literal), nil
			default:
				err = fmt.Errorf("Expected `]` after array type instead found -> %s\n", token.String())
				return
			}

		default:
			err = fmt.Errorf("Expected type after `[` instead found -> %s\n", token.String())
			return
		}
	default:
		err = fmt.Errorf("Expected type instead found -> %s\n", token.String())
		return
	}
}

// func (p *Parser) parseLetStmt() (ls LetStmt, err error) {

// }

func (p *Parser) parseBlockStmt() (bs BlockStmt, err error) {
	for {
		tokenKind, _ := p.ScanSkipWhitespace()
		//token := Token{tokenKind, literal}

		switch tokenKind {
		case KW_LET:
			tokenKind, literal := p.ScanSkipWhitespace()
			token := Token{tokenKind, literal}
			switch tokenKind {
			case IDENT:
				tokenKind, literal := p.ScanSkipWhitespace()
				token := Token{tokenKind, literal}
				switch tokenKind {
				case EQ:
					// Expect Expr
				case COLON:
					// Expect Type
				default:
					err = fmt.Errorf("Expected `=` or `:` after `identifier` instead found -> %s\n", token.String())
					return
				}
			default:
				err = fmt.Errorf("Expected `identifier` after `let` instead found -> %s\n", token.String())
				return
			}
		case IDENT:
			// Okay we have a couple of ways to go here

			tokenKind, _ := p.ScanSkipWhitespace()
			// 1. We have a function call
			// 2. We have a variable declaration
			switch tokenKind {
			}
		case R_CURLY:
			return

			// 3. We have a variable assignment
		}
	}
}

func (p *Parser) parseFuncType() (*FuncType, error) {
	var ft FuncType

	ft.Params = []FuncParam{}

	// :3
	for {
		tokenKind, literal := p.ScanSkipWhitespace()
		token := Token{tokenKind, literal}
		switch {
		case tokenKind == EOF:
			return nil, fmt.Errorf("Unclosed function: expected `)` found -> %s\n", token.String())
		case tokenKind == R_BRACK:
			return &ft, nil
		// In case we have multiple args we eat the `,`
		case tokenKind == COMMA && len(ft.Params) > 0:
			tokenKind, literal := p.ScanSkipWhitespace()
			switch tokenKind {
			// Fn param -> `identifier`
			case IDENT:
				var param FuncParam
				param.Ident = NewIdentifier(literal, Span{}, Object{FUNC_PARAM, literal})

				tokenKind, literal := p.ScanSkipWhitespace()
				token := Token{tokenKind, literal}

				// Fn param type
				switch tokenKind {
				case IDENT:
					param.Type = NewIdentifier(literal, Span{}, Object{FUNC_PARAM_TYPE, literal})
					ft.Params = append(ft.Params, param)
				default:
					return nil, fmt.Errorf("Expected type after function parameter identifier, instead found -> `%s`\n", token.String())
				}
			}
		// One arg / first arg
		case tokenKind == IDENT && len(ft.Params) == 0:
			var param FuncParam
			param.Ident = NewIdentifier(literal, Span{}, Object{FUNC_PARAM, literal})

			tokenKind, literal := p.ScanSkipWhitespace()
			token := Token{tokenKind, literal}

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
