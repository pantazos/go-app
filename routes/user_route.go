package routes

import (
	"mux-mongo-api/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/match/{matchId}", controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/matches", controllers.GetAllUser()).Methods("GET")
}
