package lexer

import "fmt"

type TokenType int

const (
	// Special Tokens
	EOF TokenType = -1

	// Separators
	COMMA TokenType = iota

	// OPERATORS
	START_GROUP
	END_GROUP
	VEHICLE_BOUNDRY_START
	VEHICLE_BOUNDRY_END

	// Value Tokens
	SECTOR
	STATUS
	WAGON_TYPE
	ORDER_NUMBER
	OFFER
)

type Token struct {
	Type  TokenType
	Value string
}

func (token Token) IsOneOf(tokenTypes ...TokenType) bool {
	for _, expected := range tokenTypes {
		if token.Type == expected {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.IsOneOf(SECTOR, STATUS, WAGON_TYPE, ORDER_NUMBER, OFFER) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Type), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Type))
	}
}

func NewToken(tokenType TokenType, value string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}

func TokenKindString(tokenType TokenType) string {
	switch tokenType {
	case EOF:
		return "EOF"
	case COMMA:
		return "COMMA"
	case START_GROUP:
		return "START_GROUP"
	case END_GROUP:
		return "END_GROUP"
	case VEHICLE_BOUNDRY_START:
		return "VEHICLE_BOUNDRY_START"
	case VEHICLE_BOUNDRY_END:
		return "VEHICLE_BOUNDRY_END"
	case SECTOR:
		return "SECTOR"
	case STATUS:
		return "STATUS"
	case WAGON_TYPE:
		return "WAGON_TYPE"
	case ORDER_NUMBER:
		return "ORDER_NUMBER"
	case OFFER:
		return "OFFER"
	default:
		return "UNKNOWN"
	}
}
