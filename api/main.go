package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ToLoad()
	fmt.Println("Application is running on port 5000")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
