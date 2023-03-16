package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joskeiner/go_and_Postgres/models"
)

type response struct {
	ID      int64  `jsion:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db

}
func GetStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	json.NewDecoder(r.Body).Decode(&stock)

}
func GetAllStock() {

}
func CreateStock() {

}
func UpdateStock() {

}
func DaleteStock() {

}
