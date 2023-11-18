package grpc

import (
	"io"
	"log"

	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type HandlerImplInterface interface {
	GetChannel() *chan *pb.Document

	pb.GrpcServiceServer
}

type GrpcHandlerImpl struct {
	Source chan *pb.Document

	pb.UnimplementedGrpcServiceServer
}

func (servImpl *GrpcHandlerImpl) Initialize() {
	servImpl.Source = make(chan *pb.Document)
}

func (servImpl *GrpcHandlerImpl) UploadData(uploader pb.GrpcService_UploadDataServer) error {
	for {
		document, err := uploader.Recv()
		if err == io.EOF {
			close(servImpl.Source)
			return uploader.SendAndClose(&pb.Status{Succeded: true})
		}

		for err != nil {
			close(servImpl.Source)
			log.Printf("Failed to receive document: %v", err)
			return err
		}

		servImpl.Source <- document
	}
}

func (servImpl *GrpcHandlerImpl) GetChannel() *chan *pb.Document {
	return &servImpl.Source
}

func NewGrpcHandlerImpl() *GrpcHandlerImpl {
	handler := new(GrpcHandlerImpl)
	handler.Initialize()
	return handler
}
