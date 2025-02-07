package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Application is running on port 5000")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
