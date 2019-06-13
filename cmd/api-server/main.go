package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"go-api-project-template/api/json/requests/books"
	"go-api-project-template/internal/api-server/endpoints"
	"go-api-project-template/internal/api-server/services/implements"
)

func main() {
	svc := implements.BookService{}

	indexHandler := httptransport.NewServer(
		endpoints.UppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	http.Handle("/index", indexHandler)

	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("========================")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.IndexRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

