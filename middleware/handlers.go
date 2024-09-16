package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CreateConnection() *sql.DB {
	err3 := godotenv.Load(".env")
	check(err3)

	db, err2 := sql.Open("postgress", os.Getenv("POSTGRESS_URL"))
	check(err2)

	err1 := db.Ping()
	check(err1)

	fmt.Println("Connection Established to POSTGRES....")
	return db
}

func Home() {
}
