package main

import (
	"errors"
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

	if res.StatusCode >= 400 {
		err := fmt.Sprintln("http error: ", res.StatusCode)
		return "", errors.New(err)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", errors.New("Invalid content-type header")
	}

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(d))
	return "", nil
}
