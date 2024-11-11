package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type reponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	res := reponse{}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println("Error in decoding the request body ")
		w.Write([]byte("error in decoding the request"))
		return
	}
	err = InsertToDB(res)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error in inserting to DB %v\n", err)))
		return
	}
	w.Write([]byte("Data inserted successfully"))

}
