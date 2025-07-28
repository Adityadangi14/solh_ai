package initializers

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

var WeaviateClient *weaviate.Client

func ConnectToWeaviate() {

	cfg := weaviate.Config{
		Scheme: "http",
		Host:   "localhost:8080",
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
	fmt.Println("Connected to Weaviate successfully")

}
