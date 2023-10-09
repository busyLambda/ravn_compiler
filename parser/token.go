package parser

import "fmt"

type Token struct {
  kind TokenKind
  literal string
}

func NewToken(kind TokenKind, literal string) Token {
  return Token{kind, literal}
}

func (t *Token) Literal() string {
  return t.literal
}

type TokenKind = int

const (
  EOF TokenKind = iota
  UNKNOWN_TOKEN
  WHITESPACE
  NEWLINE

  LINE_COMMENT
  TERMINATED_BLOCK_COMMENT
  UNTERMINATED_BLOCK_COMMENT

  STRING_LITERAL
  CHAT_LITERAL
  INT_LITERAL
  FLOAT_LITERAL

  IDENT
  
  ADD
  MIN
  DIV
  MUL
  MOD
  REF

  DEC // --
  INC // ++

  ASS // =
  DAS // := (Decleare ASsign)
  COLON // :
  COMMA
  DOT

  L_BRACK
  R_BRACK
  L_CURLY
  R_CURLY

  KW_FN
  KW_LET
  KW_CONST
  KW_STRUCT
  KW_ENUM
  KW_IF
  KW_ELSE
  KW_AND
  KW_OR
  KW_IS
  KW_ISNOT

  RV_T_I128 // Reserved symbols
  RV_T_I64
  RV_T_I32
  RV_T_I16
  RV_T_I8
  RV_T_U128
  RV_T_U64
  RV_T_U32
  RV_T_U16
  RV_T_U8
  RV_T_F64
  RV_T_F32
  RV_T_F16
  RV_T_F8
  RV_T_USIZE

  RV_T_STR
  RV_T_CHAR
  RV_T_BOOL

  RV_V_TRUE
  RV_V_FALSE
)

func (t *Token) String() string {
  switch t.kind {
  case LINE_COMMENT:
    return fmt.Sprintf("// %s", t.literal)
  case TERMINATED_BLOCK_COMMENT:
    return fmt.Sprintf("/* %s */", t.literal)
  case UNTERMINATED_BLOCK_COMMENT:
    return fmt.Sprintf("/* %s !", t.literal)
  case KW_FN:
    return "fn"
  case KW_CONST:
    return "const"
  case KW_LET:
    return "let"
  case KW_STRUCT:
    return "struct"
  case KW_IF:
    return "if"
  case KW_ELSE:
    return "else"
  case KW_AND:
    return "and"
  case KW_OR:
    return "or"
  case KW_IS:
    return "is"
  case KW_ISNOT:
    return "isnot"
  case IDENT:
    return fmt.Sprintf("IDENT -> \"%s\"", t.literal)
  case L_BRACK:
    return "("
  case R_BRACK:
    return ")"
  case L_CURLY:
    return "{"
  case R_CURLY:
    return "}"
  case ADD:
    return "+"
  case MIN:
    return "-"
  case DIV:
    return "/"
  case MUL:
    return "*"
  case ASS:
    return "="
  case DAS:
    return ":="
  case WHITESPACE:
    return "_wh_"
  case NEWLINE:
    return "_\\n_"
  case INT_LITERAL:
    return fmt.Sprintf("INT -> %s", t.literal)
  case FLOAT_LITERAL:
    return fmt.Sprintf("FLOAT -> %s", t.literal)
  }
  return "!"
}
