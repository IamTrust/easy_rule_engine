package handler

import (
	"github.com/trust/easy_rule_engine/compiler"
	"github.com/trust/easy_rule_engine/executor"
)

func Compiler(exp string) (*executor.Node, error) {
	tokenScanner := compiler.NewScanner(exp)
	tokens, err := tokenScanner.Lexer()
	if err != nil {
		return nil, err
	}

	parser := compiler.NewParser(tokens)
	err = parser.ParseSyntax()
	if err != nil {
		return nil, err
	}

	astBuilder := compiler.NewBuilder(parser)
	ast, err := astBuilder.Build()
	if err != nil {
		return nil, err
	}

	return ast, nil
}
