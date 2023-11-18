package data_source

import (
	"log"
	"os"

	json "github.com/goccy/go-json"

	pb "github.com/TimC0de/simpals-backend-test-task/worker/internal/grpc/data_exchange"
)

type FileDataSource struct {
	Decoder *json.Decoder
	Sink    chan *pb.Document
}

func (source *FileDataSource) Initialize(dataLocation string) {
	source.Sink = make(chan *pb.Document)
	fileHandler, err := os.Open(dataLocation)
	if err != nil {
		log.Fatalf("Failed to open file with data: %v", err)
		return
	}
	source.Decoder = json.NewDecoder(fileHandler)
}

func (source *FileDataSource) GetChannel() *chan *pb.Document {
	return &source.Sink
}

func (source *FileDataSource) processToken() {
	_, err := source.Decoder.Token()
	if err != nil {
		log.Fatalf("Failed to skip JSON token: %v", err)
	}
}

func (source *FileDataSource) StartFulfilling() {
	source.processToken()

	for source.Decoder.More() {
		var document pb.Document
		err := source.Decoder.Decode(&document)
		if err != nil {
			log.Fatalf("Failed to decode document: %v\n", err)
		}

		source.Sink <- &document
	}

	close(source.Sink)
}

func NewDataSource(dataLocation string) *FileDataSource {
	source := new(FileDataSource)
	source.Initialize(dataLocation)
	return source
}
