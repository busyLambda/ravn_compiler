package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Scanner struct {
	reader         *bufio.Reader
	pos            int
	posWithinToken int
	buffer         struct {
		prevToken TokenKind
		literal   string
	}
}

func NewScanner(input string) *Scanner {
	return &Scanner{
		reader:         bufio.NewReader(strings.NewReader(input)),
		pos:            0,
		posWithinToken: 0,
		buffer: struct {
			prevToken int
			literal   string
		}{
			prevToken: EOF,
			literal:   "",
		},
	}
}

func (s *Scanner) read() (ch rune, err error) {
	s.posWithinToken++
	s.pos++
	ch, _, err = s.reader.ReadRune()
	return
}

func (s *Scanner) resetPosWithinToken() {
	s.posWithinToken = 0
}

func (s *Scanner) Scan() (tokenKind TokenKind, literal string) {
	ch, err := s.read()
	if err != nil {
		return EOF, ""
	}

	switch {
	case unicode.IsDigit(ch):
		s.unread()
		return s.ScanNumericLiteral()
	case ch == '\n':
		tokenKind = NEWLINE
	case unicode.IsSpace(ch):
		s.unread()
		return s.ScanWhitespace()
	case unicode.IsLetter(ch) || ch == '_':
		s.unread()
		return s.ScanIdentifierOrKeyword()
	case ch == '.':
		tokenKind = DOT
	case ch == ',':
		tokenKind = COMMA
	case ch == '&':
		tokenKind = REF
	case ch == '(':
		tokenKind = L_BRACK
	case ch == ')':
		tokenKind = R_BRACK
	case ch == '{':
		tokenKind = L_CURLY
	case ch == '}':
		tokenKind = R_CURLY
	case ch == '+':
		tokenKind = ADD
		literal = string(ch)
	case ch == '-':
		tokenKind = MIN
		literal = string(ch)
	case ch == '=':
		tokenKind = EQ
		literal = string(ch)
	case ch == '|':
		tokenKind = PIPE
		literal = string(ch)
	case ch == '/':
		ch2, err := s.read()
		if err != nil {
			s.unread()
			tokenKind = DIV
		} else {
			if ch2 == '/' {
				s.unread()
				return s.ScanCommentLine()
			} else if ch2 == '*' {
				s.unread()
				s.unread()
				return s.ScanCommentBlock()
			} else {
				s.unread()
				tokenKind = DIV
				literal = string(ch)
			}
		}
	case ch == '*':
		tokenKind = MUL
		literal = string(ch)
	case ch == '=':
		tokenKind = ASS
		literal = string(ch)
	case ch == ':':
		ch2, err := s.read()
		if err != nil {
			s.unread()
			tokenKind = COLON
			literal = string(ch)
		} else {
			if ch2 == '=' {
				tokenKind = DAS
				literal = string(ch)
			} else {
				s.unread()
				tokenKind = COLON
				literal = string(ch)
			}
		}
	}

	return
}

func (s *Scanner) unread() {
	s.reader.UnreadRune()
}

func (s *Scanner) ScanIdentifierOrKeyword() (TokenKind, string) {
	var buffer bytes.Buffer
	ch, _ := s.read()
	buffer.WriteRune(ch)

	for {
		ch, err := s.read()

		if err != nil {
			s.unread()
			break
		}

		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
			buffer.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	lit := buffer.String()

	switch lit {
	case "fn":
		return KW_FN, lit
	case "const":
		return KW_CONST, lit
	case "let":
		return KW_LET, lit
	case "if":
		return KW_IF, lit
	case "else":
		return KW_ELSE, lit
	case "and":
		return KW_AND, lit
	case "or":
		return KW_OR, lit
	case "is":
		return KW_IS, lit
	case "isnot":
		return KW_ISNOT, lit
	}

	return IDENT, lit
}

func (s *Scanner) ScanWhitespace() (TokenKind, string) {
	var buffer bytes.Buffer
	ch, _ := s.read()
	buffer.WriteRune(ch)

	for {
		ch, err := s.read()

		if err != nil {
			s.unread()
			fmt.Println("Pull the trigger!")
			break
		}
		if unicode.IsSpace(ch) {
			buffer.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	return WHITESPACE, buffer.String()
}

func (s *Scanner) ScanNumericLiteral() (TokenKind, string) {
	var buffer bytes.Buffer
	ch, _ := s.read()
	buffer.WriteRune(ch)
	var tokenKind TokenKind = INT_LITERAL
	encounteredDot := false

	for {
		ch, err := s.read()
		if err != nil {
			s.unread()
			break
		}

		if unicode.IsDigit(ch) {
			buffer.WriteRune(ch)
		} else if ch == '.' {
			if encounteredDot {
				s.unread()
				break
			}
			tokenKind = FLOAT_LITERAL
			encounteredDot = true
			buffer.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	return tokenKind, buffer.String()
}

func (s *Scanner) ScanCommentLine() (TokenKind, string) {
	var buffer bytes.Buffer
	ch, _ := s.read()
	buffer.WriteRune(ch)
	ch, _ = s.read()
	buffer.WriteRune(ch)

	for {
		ch, err := s.read()
		if err != nil {
			s.unread()
			break
		}
		if ch != '\n' {
			buffer.WriteRune(ch)
		} else {
			s.unread()
			break
		}
	}

	return LINE_COMMENT, buffer.String()
}

func (s *Scanner) ScanCommentBlock() (TokenKind, string) {
	var buffer bytes.Buffer
	ch, _ := s.read()
	buffer.WriteRune(ch)
	ch, _ = s.read()
	buffer.WriteRune(ch)

	depth := 1

	for {
		if depth == 0 {
			break
		}

		ch, err := s.read()
		if err != nil {
			s.unread()
			break
		}

		if ch == '/' {
			buffer.WriteRune(ch)
			ch2, err := s.read()
			if err != nil {
				s.unread()
				break
			}

			if ch2 == '*' {
				depth++
				buffer.WriteRune(ch2)
			} else {
				buffer.WriteRune(ch2)
			}
		} else if ch == '*' {
			buffer.WriteRune(ch)
			ch2, err := s.read()
			if err != nil {
				s.unread()
				break
			}

			if ch2 == '/' {
				depth--
				buffer.WriteRune(ch2)
			} else {
				buffer.WriteRune(ch2)
			}
		} else {
			buffer.WriteRune(ch)
		}
	}

	if depth == 0 {
		return TERMINATED_BLOCK_COMMENT, buffer.String()
	} else {
		return UNTERMINATED_BLOCK_COMMENT, buffer.String()
	}
}
