package lexer

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/models"
)

type regexHandler func(*lexer, *regexp.Regexp)

type regexPatterns struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPatterns
	Tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func defaultHandler(tokenType TokenType, value string) regexHandler {
	return func(lexer *lexer, regex *regexp.Regexp) {
		lexer.advanceN(len(value))
		lexer.push(NewToken(tokenType, value))
	}
}

func sectorHandler(lexer *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lexer.remainder())
	lexer.push(NewToken(SECTOR, match[1:]))
	lexer.advanceN(len(match))
}

func wagonTypeHandler(lexer *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lexer.remainder())
	lexer.push(NewToken(WAGON_TYPE, match))
	lexer.advanceN(len(match))
}

func orderNumberHandler(lexer *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lexer.remainder())
	lexer.push(NewToken(ORDER_NUMBER, match[1:]))
	lexer.advanceN(len(match))
}

func offerHandler(lexer *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lexer.remainder())
	// split by semicolon if it exists
	matchWithoutPrefix := match[1:]
	if strings.Contains(matchWithoutPrefix, ";") {
		split := strings.Split(matchWithoutPrefix, ";")
		for _, value := range split {
			lexer.push(NewToken(OFFER, value))
		}
		lexer.advanceN(len(match))
	} else {
		lexer.push(NewToken(OFFER, matchWithoutPrefix))
		lexer.advanceN(len(match))
	}
}

func statusHandler(lexer *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lexer.remainder())
	lexer.push(NewToken(STATUS, match))
	lexer.advanceN(len(match))
}

func createLexer(source string) *lexer {
	return &lexer{
		source: source,
		pos:    0,
		Tokens: make([]Token, 0),
		patterns: []regexPatterns{
			{regexp.MustCompile(`@[A-Z]`), sectorHandler},
			{regexp.MustCompile(models.GetWagonTypesRegex()), wagonTypeHandler},
			{regexp.MustCompile(`\[`), defaultHandler(START_GROUP, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(END_GROUP, "]")},
			{regexp.MustCompile(`\(`), defaultHandler(VEHICLE_BOUNDRY_START, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(VEHICLE_BOUNDRY_END, ")")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`:[0-9]+`), orderNumberHandler},
			{regexp.MustCompile(fmt.Sprintf(`#(%s)(;(%s))*`, models.GetWagonOfferRegex(), models.GetWagonOfferRegex())), offerHandler},
			{regexp.MustCompile(models.GetWagonStatusRegex()), statusHandler},
		},
	}
}

func Tokenize(formationString string) ([]Token, error) {
	lex := createLexer(formationString)

	for !lex.at_eof() {
		matched := false
		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			return nil, errors.New("Lexer::Tokenize -> could not match any pattern near: " + lex.remainder())
		}
	}

	if len(lex.Tokens) == 0 {
		return nil, errors.New("Lexer::Tokenize -> no tokens found")
	}

	lex.push(NewToken(EOF, ""))
	return lex.Tokens, nil
}

func ParseFormation(formationString string) models.TrainFormationAtStop {
	trainFormation := models.NewTrainFormationAtStop(formationString)

	return *trainFormation
}
