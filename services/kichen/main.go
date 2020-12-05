package main

import (
	"fmt"
	"log"
	"net"

	"github.com/somen440/saga-sample/proto"
	"github.com/somen440/saga-sample/util"

	"google.golang.org/grpc"
)

const (
	port = "5001"
)

func main() {
	grpcEndpoint := fmt.Sprintf(":%s", port)
	grpcServer := util.NewServer(func(server *grpc.Server) {
		proto.RegisterKitchenService(server, NewService())
	}, true, util.ErrorMap{})
	log.Print("server is ready ", grpcEndpoint)
	listen, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
