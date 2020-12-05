package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/somen440/saga-sample/proto"
	"github.com/somen440/saga-sample/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	port = "5002"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	InitializeWorker()
	defer CloseWorker()

	grpcEndpoint := fmt.Sprintf(":%s", port)
	grpcServer := util.NewServer(func(server *grpc.Server) {
		proto.RegisterOrderService(server, NewService())
	}, true, util.ErrorMap{
		ErrNotFound: codes.NotFound,
	})
	log.Print("server is ready ", grpcEndpoint)
	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		return err
	}
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}
	return nil
}
