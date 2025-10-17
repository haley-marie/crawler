package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "BunCrawler/1.0")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("http error: %s", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non html response: %s", contentType)
	}

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
