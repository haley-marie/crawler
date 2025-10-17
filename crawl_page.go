package main

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	maxPages           int
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, exists := cfg.pages[normalizedURL]

	if exists {
		return false
	} else {
		cfg.pages[normalizedURL] = PageData{}
		return true
	}
}

func (cfg *Config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}

func (cfg *Config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func (cfg *Config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing currentURL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalizing current URL: ", err)
		return
	}

	if !cfg.addPageVisit(normalizedCurrentURL) {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error getting HTML from current URL:", err)
		return
	}

	d := extractPageData(pageHTML, rawCurrentURL)
	cfg.setPageData(normalizedCurrentURL, d)

	for _, URL := range d.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(URL)
	}
}
