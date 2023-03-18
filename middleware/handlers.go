package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/joskeiner/go_and_Postgres/models"
)

// objeto para enviar el response
type response struct {
	ID      int64  `jsion:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// crear connection with postgres db
func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file ")
	}
	// abrir la coneccion
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}
	//	check db
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db
	// devolver la conecion
}

// la funcion crea un stock en la  base de datos de postgres
func CreateStock(w http.ResponseWriter, r *http.Request) {
	//crear un stock vacio de tpi models.Stock
	var stock models.Stock

	//decodificar la peticion para stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body . %v", err)
	}
	// llama a la funcion  insert Stock  y pasa la struct
	insertID := insertStock(stock)

	//formato de respuesta
	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}

	// envio de respuesta
	json.NewEncoder(w).Encode(res)

}

// GetAllStock trera todos los stock
func GetAllStock(w http.ResponseWriter, r *http.Request) {
	//traer todos los stock de la base de datos
	stock, err := getAllStock()

	//en caso  de un error
	if err != nil {
		log.Fatalf("Unable to get all stock %v", err)
	}

	//envio de todos los stocks en formato jsoon
	json.NewEncoder(w).Encode(stock)

}

// GetStock retornara un stock por id
func GetStock(w http.ResponseWriter, r *http.Request) {
	// obtener la variable stockid desde el request
	paramas := mux.Vars(r)

	//convertir the id de string a int
	//para trabajar con el
	id, err := strconv.Atoi(paramas["id"])

	if err != nil {
		log.Fatalf("unable to convert the string int int %v", err)
	}
	//llama a la funcion getStock con el id del strcok a buscar
	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("uncable to get stock .%v", err)
	}

	// enviar el json
	json.NewEncoder(w).Encode(stock)

}

// actualizara un  stock en postgres db
func UpdateStock(w http.ResponseWriter, r *http.Request) {

	// obtener stockid desde r , la llave es 'id'
	paramas := mux.Vars(r)

	//covertir el id de tipo string a int
	id, err := strconv.Atoi(paramas["id"])

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}
	//crear un struc de stock
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v ", err)
	}

	// llamar a la funcion stock para actualizar the stock
	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock update successfully . Total rows/record affected %v", updatedRows)

	// strcut de respose
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// envio de json
	json.NewEncoder(w).Encode(res)
}

// elimina una fila de la base de datos
func DaleteStock(w http.ResponseWriter, r *http.Request) {

	//trae the stockid desde el request
	params := mux.Vars(r)

	// convierte el id de string a int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to covert the string into int . %v", err)
	}

	// llama  al la funcion deleteStock , convierte el int en int64
	deletedRows := deleteStock(int64(id))

	// mensaje
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", deletedRows)

	//struct para  enviar
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	//envio de json

	json.NewEncoder(w).Encode(res)

}

//----------------------------handler funcions ------------------------------------------------

// insetar un stock en la DB
func insertStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

}
