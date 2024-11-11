package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	ConnectToDB()
	
}

func main() {

	http.HandleFunc("/", AddUser)
	http.ListenAndServe(":8080", nil)

}
