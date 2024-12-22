package lexer

import (
	"testing"
)

func TestLexingTwoSectors(t *testing.T) {
	formation := "@A@B"

	tokens, err := Tokenize(formation)
	if err != nil {
		t.Errorf("Lexer::Tokenize -> %s", err)
	}

	if len(tokens) != 3 {
		t.Errorf("Lexer::Tokenize -> expected 3 tokens, got %d", len(tokens))
	}

	if tokens[0].Type != SECTOR || tokens[1].Type != SECTOR {
		t.Errorf("Lexer::Tokenize -> expected tokens to be of type SECTOR, got %s and %s", TokenKindString(tokens[0].Type), TokenKindString(tokens[1].Type))
	}

	if tokens[0].Value != "A" || tokens[1].Value != "B" {
		t.Errorf("Lexer::Tokenize -> expected tokens to have value A and B, got %s and %s", tokens[0].Value, tokens[1].Value)
	}

	if tokens[2].Type != EOF {
		t.Errorf("Lexer::Tokenize -> expected last token to be EOF, got %s", TokenKindString(tokens[2].Type))
	}
}

func TestLexingFullSector(t *testing.T) {
	formation := "@A,F,F,FA:8#VH;NF"

	tokens, err := Tokenize(formation)
	if err != nil {
		t.Errorf("Lexer::Tokenize -> %s", err)
	}

	if len(tokens) != 11 {
		t.Errorf("Lexer::Tokenize -> expected 11 tokens, got %d", len(tokens))
	}

	if tokens[0].Type != SECTOR {
		t.Errorf("Lexer::Tokenize -> expected first token to be of type SECTOR, got %s", TokenKindString(tokens[0].Type))
	}

	if tokens[0].Value != "A" {
		t.Errorf("Lexer::Tokenize -> expected first token to have value A, got %s", tokens[0].Value)
	}

	if tokens[len(tokens)-1].Type != EOF {
		t.Errorf("Lexer::Tokenize -> expected last token to be EOF, got %s", TokenKindString(tokens[10].Type))
	}
}
