package main

import (
	"net/http"
	"log"
	"./controller"
)

func main() {

	err := http.ListenAndServe(":3000", controller.Router())

	log.Println(err)
}