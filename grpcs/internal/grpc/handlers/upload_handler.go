package handlers

import (
	"io"
	"log"

	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/elastic"
	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type UploadHandlerInterface interface {
	pb.UploadServiceServer
}

type UploadHandler struct {
	ElasticClient *elastic.ElasticSearchClient

	pb.UnimplementedUploadServiceServer
}

func (servImpl *UploadHandler) Initialize(client *elastic.ElasticSearchClient) {
	servImpl.ElasticClient = client
}

func (servImpl *UploadHandler) UploadData(uploader pb.UploadService_UploadDataServer) error {
	for {
		document, err := uploader.Recv()
		if err == io.EOF {
			return uploader.SendAndClose(&pb.Status{Succeded: true})
		}

		for err != nil {
			log.Printf("Failed to receive document: %v", err)
			return err
		}

		go servImpl.ElasticClient.Upload(document)
	}
}

func NewUploadHandler(client *elastic.ElasticSearchClient) *UploadHandler {
	handler := new(UploadHandler)
	handler.Initialize(client)
	return handler
}
