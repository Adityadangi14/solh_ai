package initializers

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

var WeaviateClient *weaviate.Client

func ConnectToWeaviate() {

	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		AppLogger.Error("Failed to create Weaviate client", "error", err)
		return
	}

	if client == nil {
		AppLogger.Error("Weaviate client is nil")
		panic("Weaviate client is nil")
	}

	ready, err := client.Misc().ReadyChecker().Do(context.Background())
	if err != nil {
		AppLogger.Error("Weaviate ready check failed", "error", err)
		panic(err)
	}

	WeaviateClient = client
	AppLogger.Info("Connected to Weaviate successfully", "ready", ready)
}
