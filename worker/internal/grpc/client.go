package grpc

import (
	"log"

	pb "github.com/TimC0de/simpals-backend-test-task/worker/internal/grpc/data_exchange"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	Valid      bool
	Connection *grpc.ClientConn
	Client     pb.GrpcServiceClient
}

func (gc *GrpcClient) initializeConnection(serverLocation string) {
	var err error
	gc.Connection, err = grpc.Dial(serverLocation, grpc.WithTransportCredentials(insecure.NewCredentials()))
	gc.Valid = true
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service: %v", err)
		gc.Valid = false
	}
}

func (gc *GrpcClient) Initialize(serverLocation string) {
	gc.initializeConnection(serverLocation)
	gc.Client = pb.NewGrpcServiceClient(gc.Connection)
}

func (gc *GrpcClient) Close() {
	gc.Connection.Close()
	gc.Valid = false
}

func NewGrpcClient(serverLocation string) *GrpcClient {
	grpcClient := new(GrpcClient)
	grpcClient.Initialize(serverLocation)
	return grpcClient
}
