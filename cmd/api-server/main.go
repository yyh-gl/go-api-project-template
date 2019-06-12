package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/yyh-gl/go-project-template/api/json/requests/books"
	"github.com/yyh-gl/go-project-template/internal/api-server/endpoints"
	"github.com/yyh-gl/go-project-template/internal/api-server/services/implements"
)

func main() {
	svc := implements.BookService{}

	indexHandler := httptransport.NewServer(
		endpoints.UppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	http.Handle("/index", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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

