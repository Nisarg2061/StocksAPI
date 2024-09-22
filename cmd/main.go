package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nisarg2061/StocksAPI/cmd/router"
)

func main() {
	r := router.Router()
	fmt.Println("Listening on port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}
