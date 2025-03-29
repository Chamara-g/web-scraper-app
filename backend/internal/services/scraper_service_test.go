package services_test

import (
	"errors"
	"testing"

	"github.com/gihanc.dev/web-scraper-app/backend/internal/models"
	"github.com/gihanc.dev/web-scraper-app/backend/internal/services"
	"github.com/stretchr/testify/assert"
)

func Test_GetSiteDataByURL(t *testing.T) {

	tests := []struct {
		name string
		siteUrl string
		expected *models.SiteData
		expectError bool
		errorBody error
	}{
		{"Invalid URL", "this+is+invalid+url", nil, true, errors.New("Invalid URL")},
		{"URL not found", "", nil, true, errors.New("URL not found")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			siteData, err := services.GetSiteDataByURL(tt.siteUrl)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, siteData)

				assert.Equal(t, tt.errorBody, err)
			} else{
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, siteData)
			}
		})
	}

}

func Test_ContainsLoginForm(t *testing.T) {

	tests := []struct {
		name string
		siteUrl string
		expected bool
	}{
		{"Found Login Form", "https://gihan.orizel.com/", true},
		{"Not Found Form", "https://www.adidas.com/us", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			loginFormFound := services.ContainsLoginForm(tt.siteUrl)

			assert.Equal(t, tt.expected, loginFormFound)
		})
	}

}