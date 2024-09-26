package middleware

import (
	"database/sql"
	"fmt"

	"github.com/Nisarg2061/StocksAPI/cmd/models"
)

func insertStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	check("Unable to execute query", err)

	fmt.Println("Insterted a single record", id)

	return id
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT * FROM stocks where stockid=$1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

	fmt.Println("Selected record", id)

	switch err {
	case sql.ErrNoRows:
		return stock, nil
	case nil:
		return stock, nil
	default:
		check("Unable to Scan the Row", err)
	}

	return stock, err
}

func getAllStocks() (map[int]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()

	stocks := make(map[int]models.Stock)

	sqlStatement := `SELECT * FROM stocks`

	row, err := db.Query(sqlStatement)

	check("Unable to execute the query.", err)
	defer row.Close()

	for row.Next() {
		var stock models.Stock
		err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
		check("Unable to Scan the Row", err)
		stocks[int(stock.StockId)] = stock
	}

	return stocks, err
}

func deleteStock(id int64) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stocks where stockid=$1`

	res, err := db.Exec(sqlStatement, id)
	check("Unable to execute the query", err)

	rowsAffected, err := res.RowsAffected()

	check("Error while checking affected rows", err)

	fmt.Println("Rows/records affected:", rowsAffected)
	return rowsAffected
}

func updateStock(id int64, stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 where stockid=$1`

	res, err := db.Exec(sqlStatement, id, &stock.Name, &stock.Price, &stock.Company)
	check("Unable to execute the query", err)

	rowsAffected, err := res.RowsAffected()

	check("Error while checking affected rows", err)

	fmt.Println("Rows/records affected:", rowsAffected)
	return rowsAffected
}
