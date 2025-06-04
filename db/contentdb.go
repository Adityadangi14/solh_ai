package db

import (
	"context"
	"encoding/json"

	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func SaveContent(object []*models.Object) ([]models.ObjectsGetResponse, error) {

	batcher := initializers.WeaviateClient.Batch().ObjectsBatcher()

	for _, obj := range object {
		batcher = batcher.WithObjects(obj)
	}

	resp, err := batcher.Do(context.Background())

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NearSearchContent(text string) (string, error) {
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
		return "", err
	}

	result := res.Data

	str, err := json.Marshal(result)

	if err != nil {
		return "", err
	}

	return string(str), nil

}
