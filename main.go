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
	//date := "2020-03-02T00:00:00.000Z"
	//parsedDate, _ := time.Parse(layout, date)
	//cart := types.Cart{
	//	ID:     1,
	//	UserID: 1,
	//	Date:   parsedDate,
	//	Products: []types.Product{
	//		{
	//			ProductID: 1,
	//			Quantity:  4,
	//		},
	//		{
	//			ProductID: 2,
	//			Quantity:  1,
	//		},
	//	},
	//}
	//
	//JsonData, _ := json.Marshal(cart)
	//fmt.Printf(string(JsonData))
	router := httprouter.New()
	router.GET("/carts/", operations.GetAllCarts)
	router.GET("/carts/:id", operations.GetCart)
	log.Fatal(http.ListenAndServe(":8081", router))
}
