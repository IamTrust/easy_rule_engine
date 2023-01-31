// 词法定义

package token

const NoPos = 0

// Token 保存词法分析的结果.
type Token struct {
	Kind     Kind
	Value    interface{}
	Position int
}

var keywords = map[string]Kind{
	"true":  BoolLiteral,
	"false": BoolLiteral,
}

func LookupOperator(op string) Kind {
	if kind, exist := operatorToKind[op]; exist {
		return kind
	}
	return Illegal
}

// Lookup maps an identifier to its keyword token or Identifier (if not a keyword).
func Lookup(ident string) Kind {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}

	return Identifier
}
