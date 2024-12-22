package parser

import (
	"testing"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/lexer"
)

func TestParsingSimpleSectors(t *testing.T) {
	formation := "@A@B"

	tokens, err := lexer.Tokenize(formation)
	if err != nil {
		t.Errorf("Lexer::Tokenize -> %s", err)
	}

	formationObject := Parse(tokens)

	if len(formationObject.Sectors) != 2 {
		t.Errorf("Parser::Parse -> expected 2 sectors, got %d", len(formationObject.Sectors))
		return
	}

	if formationObject.Sectors[0].Name != "A" || formationObject.Sectors[1].Name != "B" {
		t.Errorf("Parser::Parse -> expected sectors to have names A and B, got %s and %s", formationObject.Sectors[0].Name, formationObject.Sectors[1].Name)
	}
}

func TestParsingFullSector(t *testing.T) {
	formation := "@A,F,F,FA:8#VH;NF"

	tokens, err := lexer.Tokenize(formation)
	if err != nil {
		t.Errorf("Lexer::Tokenize -> %s", err)
	}

	formationObject := Parse(tokens)

	if len(formationObject.Sectors) != 1 {
		t.Errorf("Parser::Parse -> expected 1 sector, got %d", len(formationObject.Sectors))
		return
	}

	if formationObject.Sectors[0].Name != "A" {
		t.Errorf("Parser::Parse -> expected sector to have name A, got %s", formationObject.Sectors[0].Name)
	}

	if len(formationObject.Sectors[0].Wagons) != 3 {
		t.Errorf("Parser::Parse -> expected sector to have 3 wagons, got %d", len(formationObject.Sectors[0].Wagons))
	}

	if formationObject.Sectors[0].Wagons[0].Type == "FA" {
		t.Errorf("Parser::Parse -> expected wagon to have type FA, got %s", formationObject.Sectors[0].Wagons[0].Type)
	}
}
