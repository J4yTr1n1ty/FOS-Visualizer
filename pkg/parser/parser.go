package parser

import (
	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/lexer"
	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/models"
)

type parser struct {
	errors []error // Not yet implemented
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens: tokens,
	}
}

func Parse(tokens []lexer.Token) *models.TrainFormationAtStop {
	formation := &models.TrainFormationAtStop{}
	p := createParser(tokens)

	for p.hasTokens() {
		// TODO: Figure out how to parse the train formation
		// formation.Sectors = append(formation.Sectors, p.parseSector())
		break
	}

	return formation
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenType() lexer.TokenType {
	return p.currentToken().Type
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenType() != lexer.EOF
}
