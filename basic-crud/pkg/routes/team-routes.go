package routes

import (
	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTeamRoutes = func(basePath string, router *mux.Router) {
	router.HandleFunc("/get", controllers.GetAllTeams).Methods("GET")
}
