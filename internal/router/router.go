package router

import (
	"github.com/gihanc.dev/web-scraper-app/internal/handler"
	"github.com/gorilla/mux"
)

// Creates all routes with all the handlers
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/scrape", handler.GetWebHTMLByURL).Methods("GET")

	return router
}