package main

import (
	"fmt"
	"log"
	"net"

	"github.com/somen440/saga-sample/proto"
	"github.com/somen440/saga-sample/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	port = "5010"
)

func main() {
	grpcEndpoint := fmt.Sprintf(":%s", port)
	grpcServer := util.NewServer(func(server *grpc.Server) {
		proto.RegisterSagaService(server, NewService())
	}, true, util.ErrorMap{
		ErrNotFound: codes.NotFound,
	})
	log.Print("server is ready ", grpcEndpoint)
	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
