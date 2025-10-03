package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		log.Fatalln("no website provided")
	}

	if len(argsWithoutProg) > 1 {
		log.Fatalln("too many arguments provided")
	}

	fmt.Printf("starting crawl of: %v\n", argsWithoutProg[0])
	crawlPage(argsWithoutProg[0], argsWithoutProg[0], make(map[string]int))
}
