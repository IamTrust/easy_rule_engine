package handler

import (
	"github.com/trust/easy_rule_engine/compiler"
	"github.com/trust/easy_rule_engine/executor"
)

// Compiler 编译表达式, 生成语法树, 返回语法树的根节点
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
