package uploader

import (
	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type StreamUploader interface {
	Upload(*pb.Document)
}
