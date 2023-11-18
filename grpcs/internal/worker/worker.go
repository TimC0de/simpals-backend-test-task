package worker

import (
	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/uploader"
)

type Worker struct {
	SourceChannel *chan *pb.Document
	Uploader      uploader.StreamUploader
}

func (worker *Worker) Initialize(
	source *chan *pb.Document,
	uploader uploader.StreamUploader,
) {
	worker.SourceChannel = source
	worker.Uploader = uploader
}

func (worker *Worker) Run() {
	document, ok := <-*worker.SourceChannel
	for ok {
		worker.Uploader.Upload(document)
		document, ok = <-*worker.SourceChannel
	}
}

func NewWorker(
	source *chan *pb.Document,
	uploader uploader.StreamUploader,
) *Worker {
	worker := new(Worker)
	worker.Initialize(source, uploader)
	return worker
}
