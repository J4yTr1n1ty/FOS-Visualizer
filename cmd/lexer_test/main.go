package main

import (
	"fmt"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/lexer"
)

func main() {
	fmt.Printf("========================\nLexer test\n========================\n")
	formation := "@A,F,F,[(FA:8#VH;NF@B,2:7#NF,2:6#NF@C,2:5#NF,%W2:4#NF@D,1:3#NF@E,1:2#NF,1:1#BHP;BZ;NF,LK)]@F,F@G,F"

	_, err := lexer.Tokenize(formation)
	if err != nil {
		panic(err)
	}

	second_formation := "@H,[(1:1#NF,1:2#KW;NF,2:3#BHP;NF,2:4#VH;KW;NF@J,2:5#VH;KW;NF,2:6#VH;KW;NF,2:7#VH;KW;NF@K,2):8#VH;KW;FZ;NF],F@L,F,F@M,F,F,F@N,F,F,F"

	tokens, err := lexer.Tokenize(second_formation)
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		token.Debug()
	}
}
