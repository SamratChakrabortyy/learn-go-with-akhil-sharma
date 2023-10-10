package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterPlayerRoutes("/player", r)
	routes.RegisterTeamRoutes("/team", r)

	http.Handle("/", r)
	fmt.Println("Starting Players Service at 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting http server at 8080", err)
	}
}
