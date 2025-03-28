package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gihanc.dev/web-scraper-app/internal/services"
	"github.com/gihanc.dev/web-scraper-app/internal/utils"
)

func GetWebHTMLByURL(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	if url == "" || !utils.IsValidURL(url) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	siteData, err := services.GetSiteDataByURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(siteData)
}