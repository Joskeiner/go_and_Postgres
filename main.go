package main

import (
	"log"
	"net/http"

	"github.com/joskeiner/go_and_Postgres/router"
)

func main() {
	r := router.Router()

	log.Print("strating server on the port : 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
