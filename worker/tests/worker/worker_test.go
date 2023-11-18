package worker

import (
	"testing"

	ds "github.com/TimC0de/simpals-backend-test-task/worker/internal/data_source"
	pb "github.com/TimC0de/simpals-backend-test-task/worker/internal/grpc/data_exchange"
	upld "github.com/TimC0de/simpals-backend-test-task/worker/internal/uploader"
	wrk "github.com/TimC0de/simpals-backend-test-task/worker/internal/worker"
)

const (
	dataLocation = "C:\\Users\\Me\\backend\\simpals-backend-test-task\\worker\\data\\data.json"
)

type UploaderTest struct {
	T       *testing.T
	Counter int
	upld.StreamUploader
}

func (upld *UploaderTest) Upload(document *pb.Document) error {
	upld.T.Logf("Incoming document Id: %s", document.Id)
	upld.Counter++
	return nil
}

func (upld *UploaderTest) CloseAndRecv() *pb.Status {
	return &pb.Status{Succeded: true}
}

func TestWorkerWithDataSource(t *testing.T) {
	source := ds.NewDataSource(dataLocation)
	sourceChannel := source.GetChannel()

	uploader := UploaderTest{Counter: 0, T: t}
	worker := wrk.NewWorker(sourceChannel, &uploader)

	go worker.Run()
	source.StartFulfilling()

	t.Logf("Documents quantity: %d", uploader.Counter)
}
