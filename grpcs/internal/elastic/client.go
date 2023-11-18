package elastic

import (
	"bytes"
	"context"
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	json "github.com/goccy/go-json"

	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type ElasticSearchClient struct {
	IndexName string
	Context   context.Context
	Client    *elasticsearch.Client
}

func (esClient *ElasticSearchClient) initializeClient(dsUrl string) {
	cfg := elasticsearch.Config{
		Addresses: []string{dsUrl},
	}

	var err error
	esClient.Client, err = elasticsearch.NewClient(cfg)
	for err != nil {
		log.Printf("Failed to connect to ElasticSearch: %v", err)
	}
}

func (esClient *ElasticSearchClient) indexExists() bool {
	result, err := esapi.IndicesExistsRequest{
		Index: []string{esClient.IndexName},
	}.Do(esClient.Context, esClient.Client)

	if err != nil {
		log.Fatalf("Failed to request whether '%s' index exists: %v", esClient.IndexName, err)
	}
	return result.StatusCode != 404
}

func (esClient *ElasticSearchClient) createIndex() {
	_, err := esClient.Client.Indices.Create(esClient.IndexName)
	if err != nil {
		log.Fatalf("Failed to create '%s' index: %v", esClient.IndexName, err)
	}
}

func (esClient *ElasticSearchClient) Initialize(indexName string, ctx context.Context, dsUrl string) {
	esClient.IndexName = indexName
	esClient.Context = ctx

	esClient.initializeClient(dsUrl)
	if !esClient.indexExists() {
		esClient.createIndex()
	}
}

func (esClient *ElasticSearchClient) Upload(document *pb.Document) {
	data, _ := json.Marshal(document)

	result, err := esapi.IndexRequest{
		Index:      esClient.IndexName,
		DocumentID: document.Id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}.Do(esClient.Context, esClient.Client)

	if err != nil {
		log.Printf("Failed to add document '%s': %s", document.Id, err.Error())
	} else if result.IsError() {
		log.Printf("Failed to add document '%s': %s", document.Id, result.String())
	} else {
		log.Printf("Document successfully added: %s", document.Id)
	}
}

func NewElasticSearchClient(indexName string, ctx context.Context, dsUrl string) *ElasticSearchClient {
	client := new(ElasticSearchClient)
	client.Initialize(indexName, ctx, dsUrl)
	return client
}
