package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gihanc.dev/web-scraper-app/backend/internal/handler"
)

func Test_GetWebHTMLData(t *testing.T){
	
	testCases := []struct{
		name string
		url string
		expectedStatus int
		responseBody string
	}{
		{
			name: "check Forbidden request",
			url: "https://www.stories.com/en/index.html",
			expectedStatus: http.StatusBadRequest,
			responseBody: `{"error":"Forbidden","code":400}`,
		},
		{
			name: "check Invalid URL",
			url: "sample_invalid_test",
			expectedStatus: http.StatusBadRequest,
			responseBody: `{"error":"Invalid URL","code":400}`,
		},
		{
			name: "success request",
			url: "https://gihan.orizel.com/",
			expectedStatus: http.StatusOK,
			responseBody: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/scrape?url="+tc.url, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}
		
			rr := httptest.NewRecorder()
		
			handler := http.HandlerFunc(handler.GetWebHTMLByURL)
			handler.ServeHTTP(rr, req)
		
			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status %d but got %d", tc.expectedStatus, rr.Code)
			}

			if tc.responseBody != "" && (rr.Body.String() != tc.responseBody+"\n") {
				t.Errorf("Unexpected response body: got %s want %s", rr.Body.String(), tc.responseBody)
			}
		})
	}
}