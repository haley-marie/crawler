package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("Error parsing URL: %w", err)
	}
	normalized := fmt.Sprintf("%s%s", parsedURL.Host, parsedURL.Path)
	normalized = strings.ToLower(normalized)
	normalized = strings.TrimSuffix(normalized, "/")
	return normalized, nil
}
