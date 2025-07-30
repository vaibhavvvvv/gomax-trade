package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting broker...")
	http.ListenAndServe(":8080", nil)
}
