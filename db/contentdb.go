package db

import (
	"context"

	"github.com/Adityadangi14/solh_ai/initializers"
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
