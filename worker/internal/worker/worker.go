package worker

import (
	"log"
	"sync"

	pb "github.com/TimC0de/simpals-backend-test-task/worker/internal/grpc/data_exchange"
	"github.com/TimC0de/simpals-backend-test-task/worker/internal/uploader"
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

func (worker *Worker) Run(wg *sync.WaitGroup) *pb.Status {
	defer wg.Done()

	document, ok := <-*worker.SourceChannel
	for ok {
		if err := worker.Uploader.Upload(document); err != nil {
			log.Fatalf("Failed to send document: %v", err)
		}
		document, ok = <-*worker.SourceChannel
	}
	return worker.Uploader.CloseAndRecv()
}

func NewWorker(
	source *chan *pb.Document,
	uploader uploader.StreamUploader,
) *Worker {
	worker := new(Worker)
	worker.Initialize(source, uploader)
	return worker
}
