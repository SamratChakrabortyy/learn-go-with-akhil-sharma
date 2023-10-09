package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		log.Println("Error: 405 Bad request")
		http.Error(w, "405 Bad Request", http.StatusMethodNotAllowed)
	}
	msg := "Greetings!"
	params := req.URL.Query()

	if params["name"] == nil {
		log.Println("Sending Greetings")
		fmt.Fprintln(w, msg)
	} else {
		log.Println("Sending name specefic Greeting")
		fmt.Fprintln(w, msg[:len(msg)-1], params["name"][0], "!")
	}

}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/greeting", greet)

	fmt.Println("Starting simple Http server at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error: ", err)
	}
}
