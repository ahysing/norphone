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
		for _, word := range os.Args {
			soundHash := Norphone(word)
			fmt.Print(word)
			fmt.Print(" => ")
			fmt.Println(soundHash)
		}
	}
}
