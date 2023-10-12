package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/datastore"
)

func GetAllTeams(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res, err := json.Marshal(datastore.GetAllTeams())
	if err != nil {
		log.Println("ErrOr: Json serialization failed", err)
		fmt.Fprintln(w, "Unknown exception")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}
