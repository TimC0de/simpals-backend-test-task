package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/elastic"
	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type FetchHandlerInterface interface {
	pb.FetchServiceServer
}

type TitleFilter struct {
	titleFilter *pb.TitleFilter
}

func (filter *TitleFilter) GetFilter() string {
	return fmt.Sprintf("title.%s: %s", filter.titleFilter.Language, filter.titleFilter.Subtext)
}

type SubcategoryFilter struct {
	categories *pb.Categories
}

func (filter *SubcategoryFilter) GetFilter() string {
	return fmt.Sprintf("categories.subcategory: %s", filter.categories.Subcategory)
}

type FetchHandler struct {
	ElasticClient *elastic.ElasticSearchClient

	pb.UnimplementedFetchServiceServer
}

func (handler *FetchHandler) Initialize(client *elastic.ElasticSearchClient) {
	handler.ElasticClient = client
}

func (handler *FetchHandler) GetDocuments(pagination *pb.Pagination, fetcher pb.FetchService_GetDocumentsServer) error {
	documents := handler.ElasticClient.Fetch(pagination)
	for _, document := range documents {
		err := fetcher.Send(document)
		if err != nil {
			log.Printf("Failed sending Document: %v", err)
		}
	}
	return nil
}

func (handler *FetchHandler) TitleSearch(titleFilter *pb.TitleFilter, fetcher pb.FetchService_TitleSearchServer) error {
	documents := handler.ElasticClient.Filter(&TitleFilter{titleFilter: titleFilter})
	for _, document := range documents {
		err := fetcher.Send(document)
		if err != nil {
			log.Printf("Failed sending Document: %v", err)
		}
	}
	return nil
}

func (handler *FetchHandler) SubcategoryDocumentAmount(ctx context.Context, categories *pb.Categories) (*pb.DocumentAmount, error) {
	documentAmount := new(pb.DocumentAmount)
	documentAmount.Amount = handler.ElasticClient.Count(&SubcategoryFilter{categories: categories})
	return documentAmount, nil
}

func NewFetchHandler(client *elastic.ElasticSearchClient) *FetchHandler {
	handler := new(FetchHandler)
	handler.Initialize(client)
	return handler
}
