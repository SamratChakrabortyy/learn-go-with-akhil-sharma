package routes

import (
	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPlayerRoutes = func(basePath string, router *mux.Router) {
	router.HandleFunc(basePath, controllers.GetAllPlayers).Methods("GET")
	router.HandleFunc(basePath+"/{id}", controllers.GetPlayerById).Methods("GET")
	router.HandleFunc(basePath, controllers.AddNewPlayer).Methods("POST")
	router.HandleFunc(basePath+"/{id}", controllers.UpdatePlayer).Methods("PUT")
	router.HandleFunc(basePath+"/{id}", controllers.DeletePlayer).Methods(("DELETE"))
}
