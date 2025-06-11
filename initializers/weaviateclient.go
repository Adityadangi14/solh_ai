package initializers

import (
	"context"
	"fmt"

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
		fmt.Println(err)
	}

	// Check the connection
	ready, err := client.Misc().ReadyChecker().Do(context.Background())
	if err != nil {
		panic(err)
	}

	WeaviateClient = client

	fmt.Printf("Weaviate connection status: %v \n", ready)
}
