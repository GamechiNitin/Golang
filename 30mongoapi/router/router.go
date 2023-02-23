package router

import (
	"mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie",controller.DeleteAllMovie).Methods("DELETE")
	return router
}
