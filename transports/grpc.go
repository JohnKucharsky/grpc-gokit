package transport

import (
	"context"
	"github.com/JohnKucharsky/grpc-gokit/endpoints"
	pb "github.com/JohnKucharsky/grpc-gokit/pb/generated"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type gRPCServer struct {
	pb.UnimplementedMathServiceServer
	add grpcTransport.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, _ log.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: grpcTransport.NewServer(endpoints.Add, decodeMathRequest, encodeMathResponse),
	}
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.MathRequest)

	return endpoints.MathRequest{
		NumA: req.NumA,
		NumB: req.NumB,
	}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResponse)

	return &pb.MathResponse{Result: resp.Result}, nil
}
