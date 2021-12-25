package main

import (
	"fmt"
	"log"
	"strings"
	. "tafl/lib"
)

const mock = `S->Aa|Bb|c
A->Bb|C|E
B->D|C|E
F->aA|@D|E
`

func main() {
	r := strings.NewReader(mock)
	grammar, err := NewGrammar(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Изначальная грамматика:\n%s\n", grammar)

}
