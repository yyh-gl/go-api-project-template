package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/yyh-gl/go-api-project-template/api/json/requests/books"
	"github.com/yyh-gl/go-api-project-template/api/json/responses/books"
	"github.com/yyh-gl/go-api-project-template/internal/api-server/services/interfaces"
)

func UppercaseEndpoint(svc interfaces.BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.IndexRequest)
		v, err := svc.Index(req.S)
		if err != nil {
			return responses.IndexResponse{v, err.Error()}, nil
		}
		return responses.IndexResponse{v, ""}, nil
	}
}
