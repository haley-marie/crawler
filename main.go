package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	defaultMaxConcurrency := 5
	defaultMaxPages := 10
	maxConcurrency := defaultMaxConcurrency
	maxPages := defaultMaxPages

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		log.Fatalln("no website provided")
	}

	baseURL := argsWithoutProg[0]
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalln("error parsing baseURL")
	}

	if len(argsWithoutProg) == 2 || len(argsWithoutProg) == 3 {
		maxConcurrency, err = strconv.Atoi(argsWithoutProg[1])

		if err != nil {
			confirmPrompt("Unable to convert max concurrency to int. Would you like to continue with the default of 5?")

			maxConcurrency = defaultMaxConcurrency
		}

	} else {
		log.Println("no max concurrency provided, using default of 5")
	}

	if len(argsWithoutProg) == 3 {
		maxPages, err = strconv.Atoi(argsWithoutProg[2])

		if err != nil {
			confirmPrompt("Unable to convert max pages to int. Would you like to continue with the default of 10?")

			maxPages = defaultMaxPages
		}

	} else {
		log.Println("no max pages provided, using default of 10")
	}

	cfg := Config{
		baseURL:            parsedBaseURL,
		wg:                 &sync.WaitGroup{},
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		pages:              make(map[string]PageData),
		maxPages:           maxPages,
	}

	fmt.Printf("Starting crawl of : %s...\n", baseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	writeCSVReport(cfg.pages, "report.csv")
}
