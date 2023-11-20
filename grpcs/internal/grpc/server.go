package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/handlers"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	port          string
	listener      net.Listener
	UploadHandler handlers.UploadHandlerInterface
	FetchHandler  handlers.FetchHandlerInterface
}

func (serv *GrpcServer) Initialize(
	serverPort string,
	uploadHdl handlers.UploadHandlerInterface,
	fetchHdl handlers.FetchHandlerInterface,
) {
	serv.port = serverPort
	serv.UploadHandler = uploadHdl
	serv.FetchHandler = fetchHdl
}

func (serv *GrpcServer) Run() {
	var err error
	serv.listener, err = net.Listen("tcp", fmt.Sprintf(":%s", serv.port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUploadServiceServer(s, serv.UploadHandler)
	if err := s.Serve(serv.listener); err != nil {
		log.Fatalf("Failed to register upload handler: %v", err)
	}

	pb.RegisterFetchServiceServer(s, serv.FetchHandler)
	if err := s.Serve(serv.listener); err != nil {
		log.Fatalf("Failed to register fetch handler: %v", err)
	}
}

func NewGrpcServer(
	serverPort string,
	uploadHdl handlers.UploadHandlerInterface,
	fetchHdl handlers.FetchHandlerInterface,
) *GrpcServer {
	server := new(GrpcServer)
	server.Initialize(serverPort, uploadHdl, fetchHdl)
	return server
}
