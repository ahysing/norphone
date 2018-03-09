package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: norphone [inputword]\n")
	os.Exit(2)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	} else {
		word := os.Args[1]
		fmt.Println("processing word ", word)
		soundHash := Norphone(word)
		fmt.Println(soundHash)
	}
}
