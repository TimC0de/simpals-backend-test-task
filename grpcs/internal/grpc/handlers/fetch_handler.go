package handlers

import (
	"context"

	pb "github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/data_exchange"
)

type FetchHandlerInterface interface {
	pb.FetchServiceServer
}

type FetchHandler struct {
	pb.UnimplementedFetchServiceServer
}

func (handler *FetchHandler) GetDocuments(pagination *pb.Pagination, fetcher pb.FetchService_GetDocumentsServer) error {
	// TODO: Implement documents fetching as stream
	return nil
}

func (handler *FetchHandler) TitleSearch(titleFilter *pb.TitleFilter, fetcher pb.FetchService_TitleSearchServer) error {
	// TODO: Implement title search as stream
	return nil
}

func (handler *FetchHandler) SubcategoryDocumentAmount(ctx context.Context, categories *pb.Categories) (*pb.DocumentAmount, error) {
	// TODO: Implement document amount fetching by subcategory
	return nil, nil
}

func NewFetchHandler() *FetchHandler {
	handler := new(FetchHandler)
	return handler
}
