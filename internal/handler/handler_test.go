package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gihanc.dev/web-scraper-app/internal/handler"
)

func Test_GetWebHTMLData(t *testing.T){
	testCases := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			
			// Act
			handler.GetWebHTMLData()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}