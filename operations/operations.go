package operations

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nitishfy/rest-api/types"
	"io"
	"log"
	"net/http"
	"net/url"
)

const APIUrl = "https://fakestoreapi.com/carts/"

func GetAllCarts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp, err := http.Get(APIUrl)
	if err != nil {
		log.Fatalf("error generating the request: %v", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading the response body: %v", err)
		return
	}
	defer resp.Body.Close()

	var cart []types.Cart
	if err = json.Unmarshal(body, &cart); err != nil {
		log.Fatalf("error unmarshalling the response body: %v", err)
		return
	}

	response, err := json.Marshal(cart)
	if err != nil {
		log.Fatalf("error marshalling the cart struct to json: %v", err)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(response)
}

func GetCart(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	URL, _ := url.JoinPath(APIUrl + id)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("error generating the request: %v", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading the response body: %v", err)
		return
	}

	var cart types.Cart
	if err = json.Unmarshal(body, &cart); err != nil {
		log.Fatalf("error unmarshalling the response body: %v", err)
		return
	}

	response, err := json.Marshal(cart)
	if err != nil {
		log.Fatalf("error marshalling the cart struct to json: %v", err)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(response)
}
