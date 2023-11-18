package app

import (
	"context"
	"sync"

	ds "github.com/TimC0de/simpals-backend-test-task/worker/internal/data_source"
	"github.com/TimC0de/simpals-backend-test-task/worker/internal/environment"
	"github.com/TimC0de/simpals-backend-test-task/worker/internal/uploader"
	"github.com/TimC0de/simpals-backend-test-task/worker/internal/worker"
)

var (
	kEnvVariables = [...]string{"DATA_LOCATION", "GRPC_SERVICE_URL"}
)

type App struct {
	environment *environment.Environment
	worker      *worker.Worker
	source      *ds.FileDataSource
	uploader    uploader.StreamUploader
}

func (a *App) Initialize(ctx context.Context) {
	a.environment = environment.NewEnvironment(kEnvVariables[:])
	a.source = ds.NewDataSource(a.environment.GetVal("DATA_LOCATION"))
	a.uploader = uploader.NewUploader(ctx, a.environment.GetVal("GRPC_SERVICE_URL"))
	a.worker = worker.NewWorker(a.source.GetChannel(), a.uploader)
}

func (a *App) waitableWorkerRun(wg *sync.WaitGroup) {
	wg.Add(1)
	go a.worker.Run(wg)
}

func (a *App) Run() {
	var wg sync.WaitGroup
	defer wg.Wait()

	a.waitableWorkerRun(&wg)
	a.source.StartFulfilling()
}

func (a *App) Deinitialize() {
	a.uploader.Deinitialize()
}

func NewApp(ctx context.Context) *App {
	a := new(App)
	a.Initialize(ctx)
	return a
}
