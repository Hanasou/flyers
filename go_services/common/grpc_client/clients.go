package grpc_client

import (
	"log"

	"github.com/Hanasou/flyers/go_services/common/proto/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClientConnection() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	return conn
}

func NewExampleServiceGrpcClient(conn *grpc.ClientConn) example.ExampleServiceClient {
	c := example.NewExampleServiceClient(conn)
	return c
}
