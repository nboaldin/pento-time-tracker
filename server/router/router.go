package router

import (
	"github.com/gorilla/mux"
	"github.com/nboaldin/pento-time-tracker/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.GetHome)
	router.HandleFunc("/api/user", middleware.CreateUser)
	router.HandleFunc("/api/user/{id}/session_start", middleware.SessionStart)
	router.HandleFunc("/api/user/{id}/session_stop", middleware.SessionStop)
	return router
}
