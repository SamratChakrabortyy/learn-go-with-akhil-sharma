package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/datastore"
)

func GetAllPlayers(w http.ResponseWriter, req *http.Request) {
	players := datastore.GetAllPlayers()
	w.Header().Add("Content-Type", "pkglication/json")

	res, err := json.Marshal(players)
	if err != nil {
		log.Panicln("Error: Json serialization error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
