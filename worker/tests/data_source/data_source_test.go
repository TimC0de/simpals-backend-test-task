package data_source

import (
	"testing"

	ds "github.com/TimC0de/simpals-backend-test-task/worker/internal/data_source"
)

const (
	dataLocation = "C:\\Users\\Me\\backend\\simpals-backend-test-task\\worker\\data\\data.json"
)

func TestDataSource(t *testing.T) {
	source := ds.NewDataSource(dataLocation)
	sourceChannel := source.GetChannel()

	t.Logf("Starting fulfilling")
	go source.StartFulfilling()

	document, ok := <-*sourceChannel
	for ok {
		t.Logf("Document Id: %s", document.Id)
		document, ok = <-*sourceChannel
	}
}
