package main

import (
	"context"

	"github.com/TimC0de/simpals-backend-test-task/worker/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.NewApp(ctx)
	defer a.Deinitialize()
	a.Run()
}
