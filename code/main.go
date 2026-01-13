package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
type Log_entry struct {
	App string
	Action string
	TimeStamp time.Time // The input will increment this.
}

var log []Log_entry

type data struct {
	Application string `json:"application"`
	Action_code string `json:"action_code"`
}

type Action_Code_map struct {}

func postDataHandler(w http.ResponseWriter, r *http.Request) {
	var d data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	Err := http.ListenAndServe(":8080", nil)
	if Err != nil {
		fmt.Println("Error starting server:", Err)
	}
}
