package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joskeiner/go_and_Postgres/router"
)

func main() {
	r := router.Router()

	fmt.Println("strating server on the port : 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
