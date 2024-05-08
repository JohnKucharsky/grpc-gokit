package endpoints

import (
	"context"
	"github.com/JohnKucharsky/grpc-gokit/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

type MathRequest struct {
	NumA, NumB float32
}

type MathResponse struct {
	Result float32
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MathRequest)
		result, _ := s.Add(ctx, req.NumA, req.NumB)

		return &MathResponse{result}, nil
	}
}
