package handlers

import (
	"io"
	"log"

	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type UploadHandlerInterface interface {
	GetChannel() *chan *pb.Document

	pb.UploadServiceServer
}

type UploadHandler struct {
	Source chan *pb.Document

	pb.UnimplementedUploadServiceServer
}

func (servImpl *UploadHandler) Initialize() {
	servImpl.Source = make(chan *pb.Document)
}

func (servImpl *UploadHandler) UploadData(uploader pb.UploadService_UploadDataServer) error {
	for {
		document, err := uploader.Recv()
		if err == io.EOF {
			close(servImpl.Source)
			return uploader.SendAndClose(&pb.Status{Succeded: true})
		}

		for err != nil {
			log.Printf("Failed to receive document: %v", err)
			return err
		}

		servImpl.Source <- document
	}
}

func (servImpl *UploadHandler) GetChannel() *chan *pb.Document {
	return &servImpl.Source
}

func NewUploadHandler() *UploadHandler {
	handler := new(UploadHandler)
	handler.Initialize()
	return handler
}
