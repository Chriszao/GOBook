package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.Generate()

	fmt.Println("✨ Listening att http://localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", router))
}
