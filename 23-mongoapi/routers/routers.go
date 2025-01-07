package routers

import (
	"github.com/gorilla/mux"
	"github.com/piyushpanchal114/mongoapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies/", controllers.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movies/", controllers.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/movies/{id}/", controllers.MarkAsWatchedController).Methods("PUT")
	router.HandleFunc("/api/movies/{id}/", controllers.DeleteOneMovieController).Methods("DELETE")
	router.HandleFunc("/api/movies/", controllers.DeleteAllMoviesController).Methods("DELETE")

	return router
}
