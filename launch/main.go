package main

import (
	"log"
	"net/http"

	"Adlet/core/backend/service/support"

	backend "Adlet/core/backend/service"
)

func main() {
	backend.Jsonhandler()
	support.Assemble()
	http.HandleFunc("/", backend.MainPageHandler)
	http.HandleFunc("/filter", backend.FilterPage)
	http.HandleFunc("/account", backend.AccountPageHandler)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
