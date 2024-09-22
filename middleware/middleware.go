package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Nisarg2061/StocksAPI/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func check(text string, e error) {
	if e != nil {
		log.Fatalf("%v, %v", text, e)
	}
}

func CreateConnection() *sql.DB {
	err3 := godotenv.Load(".env")
	check("Unable to get env file!", err3)

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	check("Unable to connect to postgress!", err)

	err1 := db.Ping()
	check("Database connection failed", err1)

	fmt.Println("Connection Established to POSTGRES....")
	return db
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Stocks API!<h1>"))
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	check("Error while trying to create new stock", err)

	insertid := insertStock(stock)

	res := models.Response{
		ID:      insertid,
		Message: "Stock inserted successfully!",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	check("Unable to fetch the id!", err)

	stock, err := getStock(int64(id))
	check("Unable to retreive the stock!", err)
	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	check("Unable to get all stocks!", err)
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	check("Unable to access the id!", err)

	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	check("Unable to decode the request!", err)

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated Successfully. Total row/records affected %v", updatedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	check("Unable to convert string to int!", err)

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stock deleted Successfully, affected rows/records %v", deletedRows)

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
