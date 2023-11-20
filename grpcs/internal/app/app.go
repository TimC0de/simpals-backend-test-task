package app

import (
	"context"
	"log"

	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/elastic"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/environment"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/handlers"
)

const (
	kIndexName = "data"
)

var (
	kEnvVariables = [...]string{"GRPC_SERVICE_ACCESS_PORT", "ELASTIC_SERVICE_URL", "ELASTIC_INDEX_NAME"}
)

type App struct {
	environment *environment.Environment
	server      *grpc.GrpcServer
	esClient    *elastic.ElasticSearchClient
}

func (a *App) Initialize(ctx context.Context) {
	a.environment = environment.NewEnvironment(kEnvVariables[:])

	a.esClient = elastic.NewElasticSearchClient(
		a.environment.GetVal("ELASTIC_INDEX_NAME"),
		ctx,
		a.environment.GetVal("ELASTIC_SERVICE_URL"),
	)
	a.server = grpc.NewGrpcServer(
		a.environment.GetVal("GRPC_SERVICE_ACCESS_PORT"),
		handlers.NewUploadHandler(a.esClient),
		handlers.NewFetchHandler(a.esClient),
	)

	log.Printf("gRPC Service is up and running...")
}

func (a *App) Run() {
	a.server.Run()
}

func NewApp(ctx context.Context) *App {
	a := new(App)
	a.Initialize(ctx)
	return a
}
