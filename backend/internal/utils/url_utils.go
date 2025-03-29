package utils

import (
	"net/url"
	"strings"
)

func IsValidURL(link string) bool {

	if strings.HasPrefix(link, "//") {
		link = "https:" + link
	}

	parsedURL, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	return true
}