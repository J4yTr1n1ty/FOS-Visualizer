package lexer

import (
	"errors"
	"regexp"

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

func unknownHandler(lexer *lexer, regex *regexp.Regexp) {
	lexer.advanceN(len(regex.String()))
	lexer.push(NewToken(UNKNOWN, regex.String()))
}

func createLexer(source string) *lexer {
	return &lexer{
		source: source,
		pos:    0,
		Tokens: make([]Token, 0),
		patterns: []regexPatterns{
			// TODO: Will probably remove most of these types just because they are so interlinked
			{regexp.MustCompile(`@`), defaultHandler(AT, "@")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`#`), defaultHandler(HASHTAG, "#")},
			{regexp.MustCompile(`.*`), unknownHandler},

			// define all regex patterns
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
