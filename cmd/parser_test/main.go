package main

import (
	"fmt"

	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/lexer"
	"github.com/J4yTr1n1ty/FOS-Visualizer/pkg/parser"
)

func main() {
	fmt.Printf("========================\nParser test\n========================\n")
	formation := "@A,F,F,[(FA:8#VH;NF@B,2:7#NF,2:6#NF@C,2:5#NF,%W2:4#NF@D,1:3#NF@E,1:2#NF,1:1#BHP;BZ;NF,LK)]@F,F@G,F"
	tokens, err := lexer.Tokenize(formation)
	parsedFormation := parser.Parse(tokens)
	if parsedFormation == nil {
		panic(err)
	}
	fmt.Println(parsedFormation)
}
