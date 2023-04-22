package main

import (
	"github.com/Hanasou/flyers/go_services/common/graphql"
	"github.com/Hanasou/flyers/go_services/common/grpc_client"
	"github.com/Hanasou/flyers/go_services/common/shutdown"
)

func main() {
	conn := grpc_client.NewGrpcClientConnection()
	defer conn.Close()
	graphql.StartGraphqlServer()
	shutdown.WaitForInterrupt()
}
