package uploader

import (
	"context"
	"log"

	"github.com/TimC0de/simpals-backend-test-task/worker/internal/grpc"
	pb "github.com/TimC0de/simpals-backend-test-task/worker/internal/grpc/data_exchange"
)

type StreamUploader interface {
	Upload(*pb.Document) error
	CloseAndRecv() *pb.Status
	Deinitialize()
}

type GrpcDocumentUploader struct {
	GrpcClient *grpc.GrpcClient
	Streamer   pb.UploadService_UploadDataClient

	StreamUploader
}

func (uploader *GrpcDocumentUploader) Initialize(ctx context.Context, grpcServerLocation string) {
	var err error
	uploader.GrpcClient = grpc.NewGrpcClient(grpcServerLocation)
	uploader.Streamer, err = uploader.GrpcClient.Client.UploadData(context.WithoutCancel(ctx))
	if err != nil {
		log.Fatalf("Failed to create 'UploadData' stream: %v", err)
	}
}

func (uploader *GrpcDocumentUploader) Upload(document *pb.Document) error {
	log.Printf("Document to be send: %s", document.Id)
	return uploader.Streamer.Send(document)
}

func (uploader *GrpcDocumentUploader) CloseAndRecv() *pb.Status {
	status, err := uploader.Streamer.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to CloseAndRecv the UploadData: %v", err)
	}
	return status
}

func (uploader *GrpcDocumentUploader) Deinitialize() {
	err := uploader.GrpcClient.Connection.Close()
	if err != nil {
		log.Fatalf("Failed to close gRPC connection: %v\n", err)
	}
}

func NewUploader(ctx context.Context, grpcServerLocation string) StreamUploader {
	uploader := new(GrpcDocumentUploader)
	uploader.Initialize(ctx, grpcServerLocation)
	return uploader
}
