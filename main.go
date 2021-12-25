package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	. "tafl/lib"
)



func main() {
	fileFlag := flag.String("path", "./grammar.txt", "The way to the file with grammar")
	flag.Parse()

	r, err := os.Open(*fileFlag)
	if err != nil{
		log.Fatal(err)
	}
	grammar, err := NewGrammar(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Изначальная грамматика:\n%s\n", grammar)
	grammar.DeleteUseless()
}
