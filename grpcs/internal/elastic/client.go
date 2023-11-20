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

type ElasticFilter interface {
	GetFilter() string
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

func (esClient *ElasticSearchClient) decodeJson(decoder *json.Decoder) *map[string]interface{} {
	decoded := new(map[string]interface{})
	if err := decoder.Decode(decoded); err != nil {
		log.Printf("Failed to decode Document: %v", err)
		return nil
	}
	return decoded
}

func (esClient *ElasticSearchClient) decodeDocuments(decoder *json.Decoder) []*pb.Document {
	jsonResult := esClient.decodeJson(decoder)
	for _, hit := range (*jsonResult)["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		marshaledSource, _ := json.Marshal(source)
		log.Printf("Marshaled _source from Elastic response: %s", marshaledSource)

		// TODO: Finish implementing documents parsing
	}

	return nil
}

func (esClient *ElasticSearchClient) Fetch(pagination *pb.Pagination) []*pb.Document {
	from := int((pagination.Page - 1) * pagination.Limit)
	size := int(pagination.Limit)

	result, err := esapi.SearchRequest{
		Index: []string{esClient.IndexName},
		From:  &from,
		Size:  &size,
	}.Do(esClient.Context, esClient.Client)

	if err != nil {
		log.Printf("Failed to fetch documents: %v", err)
		return nil
	}
	return esClient.decodeDocuments(json.NewDecoder(result.Body))
}

func (esClient *ElasticSearchClient) Filter(filter ElasticFilter) []*pb.Document {
	result, err := esapi.SearchRequest{
		Index: []string{esClient.IndexName},
		Query: filter.GetFilter(),
	}.Do(esClient.Context, esClient.Client)

	if err != nil {
		log.Printf("Failed to search for documents by filter '%s': %v", filter.GetFilter(), err)
		return nil
	}
	return esClient.decodeDocuments(json.NewDecoder(result.Body))
}

func (esClient *ElasticSearchClient) Count(filter ElasticFilter) int32 {
	result, err := esapi.CountRequest{
		Index: []string{esClient.IndexName},
		Query: filter.GetFilter(),
	}.Do(esClient.Context, esClient.Client)

	if err != nil {
		log.Printf("Failed to count documents by filter '%s': %v", filter.GetFilter(), err)
		return 0
	}
	decodedResult := esClient.decodeJson(json.NewDecoder(result.Body))
	return (*decodedResult)["count"].(int32)
}

func NewElasticSearchClient(indexName string, ctx context.Context, dsUrl string) *ElasticSearchClient {
	client := new(ElasticSearchClient)
	client.Initialize(indexName, ctx, dsUrl)
	return client
}
