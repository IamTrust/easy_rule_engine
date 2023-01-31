package compiler

import (
	"errors"
	"fmt"
	"github.com/trust/easy_rule_engine/token"
)

type Parser struct {
	tokens      []token.Token
	index       int
	tokenLength int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:      tokens,
		tokenLength: len(tokens),
	}
}

func (p *Parser) rewind() {
	p.index -= 1
}

func (p *Parser) next() token.Token {
	var tok token.Token
	tok = p.tokens[p.index]
	p.index += 1
	return tok
}

func (p *Parser) hasNext() bool {
	return p.index < p.tokenLength
}

// checkBalance Checks the balance of tokens which have multiple parts, such as parenthesis.
func (p *Parser) checkBalance() error {
	var parens int

	for p.hasNext() {
		tok := p.next()
		if tok.Kind == token.OpenParen {
			parens++
			continue
		}
		if tok.Kind == token.CloseParen {
			parens--
			continue
		}
	}

	if parens != 0 {
		return errors.New("unbalanced parenthesis")
	}
	p.Reset()
	return nil
}

func (p *Parser) ParseSyntax() error {
	// '(a + (b > c)' is illegal
	err := p.checkBalance()
	if err != nil {
		return err
	}

	// 'param1 + 100 param2' is illegal
	var lastTok token.Token
	state, err := lastTok.Kind.GetLexerState()
	for p.hasNext() {
		tok := p.next()
		if !state.CanTransitionTo(tok.Kind) {
			return fmt.Errorf("cannot transition token types from %s [%v] to %s [%v]",
				lastTok.Kind.String(), lastTok.Value, tok.Kind.String(), tok.Value)
		}

		state, err = tok.Kind.GetLexerState()
		if err != nil {
			return err
		}

		lastTok = tok
	}

	// 'a + b +' is illegal
	if !state.IsEOF() {
		return errors.New("unexpected end of expression")
	}
	p.Reset()
	return nil
}

func (p *Parser) Print() {
	for i, tok := range p.tokens {
		fmt.Printf("%3d: kind=[%s], val=[%v], pos=[%d]\n", i, tok.Kind, tok.Value, tok.Position)
	}
}

func (p *Parser) Reset() {
	p.index = 0
}
