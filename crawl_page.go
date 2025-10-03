package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Error parsing base URL: ", err)
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Error parsing current URL: ", err)
		return
	}

	if parsedBaseURL.Host != parsedCurrentURL.Host {
		fmt.Println("base and current URL hosts do not match.")
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Error normalizing current URL: ", err)
	}

	page, exists := pages[normalizedCurrentURL]
	if exists {
		page += 1
	} else {
		pages[normalizedCurrentURL] = 1
	}

	fmt.Printf("crawling URL: %v\n", rawCurrentURL)
	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("Error getting HTML from current URL:", err)
		return
	}

	d := extractPageData(pageHTML, normalizedCurrentURL)
	URLs := d.OutgoingLinks
	for _, URL := range URLs {
		crawlPage(rawBaseURL, URL, pages)
	}

}
