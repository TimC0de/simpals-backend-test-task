package app

import (
	"context"
	"log"

	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/elastic"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/environment"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/grpc/handlers"
	"github.com/TimC0de/simpals-backend-test-task/grpcs/internal/worker"
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
	uploader    *elastic.ElasticSearchClient
	worker      *worker.Worker
}

func (a *App) Initialize(ctx context.Context) {
	a.environment = environment.NewEnvironment(kEnvVariables[:])

	a.server = grpc.NewGrpcServer(
		a.environment.GetVal("GRPC_SERVICE_ACCESS_PORT"),
		handlers.NewUploadHandler(),
		handlers.NewFetchHandler(),
	)

	a.uploader = elastic.NewElasticSearchClient(
		a.environment.GetVal("ELASTIC_INDEX_NAME"),
		ctx,
		a.environment.GetVal("ELASTIC_SERVICE_URL"),
	)
	a.worker = worker.NewWorker(a.server.GetUploadChannel(), a.uploader)

	log.Printf("gRPC Service is up and running...")
}

func (a *App) Run() {
	go a.worker.Run()
	a.server.Run()
}

func (a *App) Deinitialize() {
	// TODO: Fill with instructions, if needed
}

func NewApp(ctx context.Context) *App {
	a := new(App)
	a.Initialize(ctx)
	return a
}
