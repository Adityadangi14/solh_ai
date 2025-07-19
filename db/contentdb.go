package db

import (
	"context"
	"encoding/json"

	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func SaveContent(objects []*models.Object) ([]models.ObjectsGetResponse, error) {
	batcher := initializers.WeaviateClient.Batch().ObjectsBatcher()

	for _, obj := range objects {
		batcher = batcher.WithObjects(obj)
	}

	resp, err := batcher.Do(context.Background())
	if err != nil {
		initializers.AppLogger.Error("Failed to batch save content", "error", err, "count", len(objects))
		return nil, err
	}

	initializers.AppLogger.Info("Batch content saved successfully", "count", len(resp))
	return resp, nil
}

func NearSearchContent(text string) (string, error) {
	initializers.AppLogger.Info("Performing near text search on Content", "query", text)

	query := initializers.WeaviateClient.GraphQL().Get().
		WithClassName("Content").
		WithFields(
			graphql.Field{Name: "title"},
			graphql.Field{Name: "description"},
			graphql.Field{Name: "url"},
			graphql.Field{Name: "contentType"},
		).
		WithNearText(initializers.WeaviateClient.GraphQL().NearTextArgBuilder().WithConcepts([]string{text}).WithCertainty(0.7)).
		WithLimit(5)

	res, err := query.Do(context.Background())
	if err != nil {
		initializers.AppLogger.Error("Near search failed", "error", err)
		return "", err
	}

	result := res.Data
	str, err := json.Marshal(result)
	if err != nil {
		initializers.AppLogger.Error("Failed to marshal near search result", "error", err)
		return "", err
	}

	initializers.AppLogger.Info("Near search completed successfully")
	return string(str), nil
}

func GetUrlObject(url string) (*models.GraphQLResponse, error) {
	initializers.AppLogger.Info("Fetching object by URL", "url", url)

	whereFilter := filters.Where().
		WithPath([]string{"url"}).
		WithOperator(filters.Equal).
		WithValueText(url)

	resp, err := initializers.WeaviateClient.GraphQL().Get().
		WithClassName("Content").
		WithWhere(whereFilter).
		WithFields(
			graphql.Field{Name: "title"},
			graphql.Field{Name: "contentType"},
			graphql.Field{Name: "description"},
			graphql.Field{Name: "url"},
			graphql.Field{Name: "image"},
		).
		WithLimit(1).
		Do(context.Background())

	if err != nil {
		initializers.AppLogger.Error("Failed to get URL object", "url", url, "error", err)
		return nil, err
	}

	initializers.AppLogger.Info("URL object fetched successfully", "url", url)
	return resp, nil
}
