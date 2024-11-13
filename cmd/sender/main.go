package main

import (
	"context"
	"log"

	"github.com/acronix0/Email-Service-Go/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}

	err = a.Run(ctx)
	if err!= nil {
    log.Fatalf("Failed to run app: %v", err)
  }
}