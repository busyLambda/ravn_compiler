package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/busylambda/raven/ast"
	"github.com/busylambda/raven/symtab"
)

type Parser struct {
	scanner *Scanner
	module  string
	ast     *ast.Root
	st      *symtab.SymbolTable
}

func NewParser(input string, module string) *Parser {
	return &Parser{
		scanner: NewScanner(input),
		ast:     ast.NewAstRoot(module),
		st:      symtab.NewSymTabRoot(),
	}
}

func (p *Parser) NextNode() {
	for {
		tok := p.ScanSkipWhitespace()
		if tok.kind == EOF {
			break
		}

		switch tok.kind {
		case KW_FN:
			fd := p.parseFuncDecl()
			p.ast.Decls = append(p.ast.Decls, fd)
		default:
			break
		}
	}
}

func (p *Parser) parseFuncDecl() (fd ast.FuncDecl) {
	tok := p.ScanSkipWhitespace()

	if tok.kind != IDENT {
		fmt.Printf("Expected identifier after keyword `fn` instead found -> %s\n", tok.String())
	} else {
		// Symtab
		p.st.InsertDecl(symtab.NewSymbol(tok.span, symtab.FUNC), tok.literal)
		// :3
		fd.Name = ast.NewIdentifier(tok.literal, tok.span)

		// Scan :3
		tok = p.ScanSkipWhitespace()
		switch tok.kind {
		case L_BRACK:
			// Get the function args
			t_ype, err := p.parseFuncType()
			if err != nil {
				fmt.Println(err)
			}

			fd.Type = t_ype
			tok = p.ScanSkipWhitespace()
			// token := Token{tokenKind, literal}
			switch tok.kind {
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
			fmt.Printf("Expected `(` after identifier in function declaration instead found -> %s\n", tok.String())
		}
	}

	return
}

// Not yet accouting for <> types like Option<T>
func (p *Parser) parseType() (lit string, err error) {
	tok := p.ScanSkipWhitespace()
	switch tok.kind {
	case IDENT:
		lit = tok.literal
		return
	case L_BLOCK:
		// Get the array type
		tok := p.ScanSkipWhitespace()
		switch tok.kind {
		case IDENT:
			tok := p.ScanSkipWhitespace()
			switch tok.kind {
			case R_BLOCK:
				return fmt.Sprintf("[%s]", tok.literal), nil
			default:
				err = fmt.Errorf("Expected `]` after array type instead found -> %s\n", tok.String())
				return
			}

		default:
			err = fmt.Errorf("Expected type after `[` instead found -> %s\n", tok.String())
			return
		}
	default:
		err = fmt.Errorf("Expected type instead found -> %s\n", tok.String())
		return
	}
}

func (p *Parser) parseLetStmt() (ds ast.DeclStmt, err error) {
	tok := p.ScanSkipWhitespace()
	switch tok.kind {
	case IDENT:
		sym := symtab.NewSymbol(tok.span, symtab.VAR)
		p.st.InsertDecl(sym, tok.literal)

		// ds.Left =

		tok := p.ScanSkipWhitespace()
		switch tok.kind {
		case EQ:
			// Expect Expr
		case COLON:
			// Expect Type
		default:
			err = fmt.Errorf("Expected `=` or `:` after `identifier` instead found -> %s\n", tok.String())
			return
		}
	default:
		err = fmt.Errorf("Expected `identifier` after `let` instead found -> %s\n", tok.String())
		return
	}

	return
}

func (p *Parser) parseBlockStmt() (bs ast.BlockStmt, err error) {
	for {
		tok := p.ScanSkipWhitespace()
		//token := Token{tokenKind, literal}

		switch tok.kind {
		case KW_LET:
			ds, perr := p.parseLetStmt()
			if perr != nil {
				err = perr
				return
			}
			bs.List = append(bs.List, ds)
		case IDENT:
			// Okay we have a couple of ways to go here

			tok := p.ScanSkipWhitespace()
			// 1. We have a function call
			// 2. We have a variable declaration
			switch tok.kind {
			}
		case R_CURLY:
			return

			// 3. We have a variable assignment
		}
	}
}

func (p *Parser) parseFuncType() (*ast.FuncType, error) {
	var ft ast.FuncType

	ft.Params = []ast.FuncParam{}

	// :3
	for {
		tok := p.ScanSkipWhitespace()
		switch {
		case tok.kind == EOF:
			return nil, fmt.Errorf("Unclosed function: expected `)` found -> %s\n", tok.String())
		case tok.kind == R_BRACK:
			return &ft, nil
		// In case we have multiple args we eat the `,`
		case tok.kind == COMMA && len(ft.Params) > 0:
			tok := p.ScanSkipWhitespace()
			switch tok.kind {
			// Fn param -> `identifier`
			case IDENT:
				p.st.InsertDecl(symtab.NewSymbol(tok.span, symtab.PARAM), tok.literal)
				var param ast.FuncParam
				param.Ident = ast.NewIdentifier(tok.literal, tok.span)

				tok := p.ScanSkipWhitespace()

				// Fn param type
				switch tok.kind {
				case IDENT:
					param.Type = ast.NewIdentifier(tok.literal, tok.span)
					ft.Params = append(ft.Params, param)
				default:
					return nil, fmt.Errorf("Expected type after function parameter identifier, instead found -> `%s`\n", tok.String())
				}
			}
		// One arg / first arg
		case tok.kind == IDENT && len(ft.Params) == 0:
			var param ast.FuncParam
			p.st.InsertDecl(symtab.NewSymbol(tok.span, symtab.PARAM), tok.literal)
			param.Ident = ast.NewIdentifier(tok.literal, tok.span)

			tok := p.ScanSkipWhitespace()

			switch tok.kind {
			case IDENT:
				param.Type = ast.NewIdentifier(tok.literal, tok.span)
				ft.Params = append(ft.Params, param)
			default:
				return nil, fmt.Errorf("Expected type after function parameter identifier, instead found -> `%s`\n", tok.String())
			}
		}
	}

}

func (p *Parser) ScanSkipWhitespace() Token {
	for {
		tokenKind, literal := p.scanner.Scan()

		// span := symtab.Span{Start: p.scanner.pos - p.scanner.posWithinToken, End: p.scanner.pos}
		span := symtab.NewSpan(p.scanner.pos-p.scanner.posWithinToken, p.scanner.pos)
		p.scanner.resetPosWithinToken()
		if tokenKind != WHITESPACE {
			token := NewToken(tokenKind, literal, span)
			return token
		}
	}
}

func (p *Parser) SymbtabToJsonFile(filename string) error {
	file, err := json.MarshalIndent(p.st, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0644)
}

func (p *Parser) AstToJsonFile(filename string) error {
	file, err := json.MarshalIndent(p.ast, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, file, 0644)
}
