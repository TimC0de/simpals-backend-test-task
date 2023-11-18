package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	port       string
	listener   net.Listener
	ServerImpl HandlerImplInterface
}

func (serv *GrpcServer) Initialize(serverPort string, impl HandlerImplInterface) {
	serv.port = serverPort
	serv.ServerImpl = impl
}

func (serv *GrpcServer) Run() {
	var err error
	serv.listener, err = net.Listen("tcp", fmt.Sprintf(":%s", serv.port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGrpcServiceServer(s, serv.ServerImpl)
	if err := s.Serve(serv.listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (serv *GrpcServer) GetChannel() *chan *pb.Document {
	return serv.ServerImpl.GetChannel()
}

func NewGrpcServer(serverPort string, impl HandlerImplInterface) *GrpcServer {
	server := new(GrpcServer)
	server.Initialize(serverPort, impl)
	return server
}
