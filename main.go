package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nitishfy/rest-api/operations"
	"log"
	"net/http"
)

const (
	layout = "2006-01-02T15:04:05.000Z"
)

func main() {
	router := httprouter.New()
	router.GET("/carts/", operations.GetAllCarts)
	router.GET("/carts/:id", operations.GetCart)
	log.Fatal(http.ListenAndServe(":8081", router))
}
