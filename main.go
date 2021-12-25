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
`

func main(){
	r := strings.NewReader(mock)
	grammar, err := NewGrammar(r)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(grammar.String())
	grammar.DeleteBarren()
	fmt.Println(grammar.String())
}
