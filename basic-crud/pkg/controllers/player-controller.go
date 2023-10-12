package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/datastore"
	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/models"
	"github.com/gorilla/mux"
)

func GetAllPlayers(w http.ResponseWriter, req *http.Request) {
	players := datastore.GetAllPlayers()
	w.Header().Add("Content-Type", "application/json")

	res, err := json.Marshal(players)
	if err != nil {
		log.Panicln("Error: Json serialization error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPlayerById(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error: Failed to get Id from the request")
		fmt.Fprintf(w, "Invalid Id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	player := datastore.GetPlayerById(id)
	resBody, err := json.Marshal(player)
	if err != nil {
		log.Println("Error: Json Parsing error", err)
		fmt.Fprintln(w, "Unknown Exception Occurred")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resBody)
	w.WriteHeader(http.StatusOK)
}

func AddNewPlayer(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var p models.Player
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		log.Println("Failed to read Req Body", err)
		fmt.Fprintln(w, "Invalid Req Body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	datastore.AddPlayer(p)
	fmt.Fprintln(w, "Successfully added player")
	w.WriteHeader(http.StatusOK)
}

func UpdatePlayer(w http.ResponseWriter, req *http.Request) {
	// TODO implement update-player
}

func DeletePlayer(w http.ResponseWriter, req *http.Request) {
	// TODO implement
}
