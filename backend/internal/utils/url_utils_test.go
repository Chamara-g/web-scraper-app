package utils_test

import (
	"testing"

	"github.com/gihanc.dev/web-scraper-app/backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

func Test_IsValidURL(t *testing.T) {

	tests := []struct {
		name     string
		siteUrl  string
		expected bool
	}{
		{"Valid URL", "https://example.com", true},
		{"Valid Local URL", "http://localhost:8080", true},
		{"Short form URL", "//example.com", true},
		{"Invalid URL", "invalid_url", false},
		{"Prototype missing URL", "www.missingprotocol.com", false},
		{"Email Address", "mailto:user@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			isValid := utils.IsValidURL(tt.siteUrl)

			assert.Equal(t, tt.expected, isValid)
		})
	}

}