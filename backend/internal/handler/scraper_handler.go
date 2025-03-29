package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gihanc.dev/web-scraper-app/backend/internal/middleware"
	"github.com/gihanc.dev/web-scraper-app/backend/internal/services"
	"github.com/gihanc.dev/web-scraper-app/backend/internal/utils"
)

func GetWebHTMLByURL(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	if url == "" || !utils.IsValidURL(url) {
		middleware.WriteErrorResponse(w, http.StatusBadRequest, "Invalid URL")
		
		return
	}
	
	siteData, err := services.GetSiteDataByURL(url)
	if err != nil {
	
		middleware.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(siteData)
}