package main

import (
	"fmt"
	"github.com/JohnKucharsky/grpc-gokit/endpoints"
	pb "github.com/JohnKucharsky/grpc-gokit/pb/generated"
	"github.com/JohnKucharsky/grpc-gokit/service"
	transport "github.com/JohnKucharsky/grpc-gokit/transports"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	addService := service.NewService(logger)
	addEndpoint := endpoints.MakeEndpoints(addService)
	grpcServer := transport.NewGRPCServer(addEndpoint, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":666")
	if err != nil {
		_ = logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	baseServer := grpc.NewServer()
	pb.RegisterMathServiceServer(baseServer, grpcServer)
	_ = level.Info(logger).Log("msg", "starting gRPC server")
	_ = baseServer.Serve(grpcListener)

}
