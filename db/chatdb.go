package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Adityadangi14/solh_ai/constants"
	"github.com/Adityadangi14/solh_ai/initializers"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/fault"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func SaveData(obj map[string]any) (*data.ObjectWrapper, error) {

	created, err := initializers.WeaviateClient.Data().Creator().
		WithClassName(constants.ClassChat.String()).
		WithProperties(obj).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return created, nil
}

func GetPreviousChat() (*models.GraphQLResponse, error) {
	// Define the fields to retrieve from the Chat class
	fields := []graphql.Field{
		{Name: "query"},
		{Name: "answer"},
		{Name: "userID"},
		{Name: "timestamp"},
	}

	// Build the GraphQL query
	resp, err := initializers.WeaviateClient.GraphQL().Get().
		WithClassName(constants.ClassChat.String()). // e.g., "Chat"
		WithFields(fields...).
		WithLimit(10).
		Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to get previous chat: %w", err)
	}

	return resp, nil
}

func DeleteAllChat() error {

	if err := initializers.WeaviateClient.Schema().ClassDeleter().WithClassName(constants.ClassChat.String()).Do(context.Background()); err != nil {

		if status, ok := err.(*fault.WeaviateClientError); ok && status.StatusCode != http.StatusBadRequest {
			return err
		}
	}
	return nil
}

func DeleteChatByUserId(userId string) error {
	_, err := initializers.WeaviateClient.Batch().ObjectsBatchDeleter().
		WithClassName(constants.ClassChat.String()).
		WithWhere(filters.Where().
			WithPath([]string{"userID"}).
			WithOperator(filters.Equal).
			WithValueText(userId),
		).
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil

}

func ReadChatsByUserId(userId string) (*models.GraphQLResponse, error) {
	fields := []graphql.Field{
		{Name: "query"},
		{Name: "answer"},
		{Name: "userID"},
		{Name: "timestamp"},
	}

	response, err := initializers.WeaviateClient.GraphQL().Get().
		WithClassName("Chat").
		WithFields(
			fields...,
		).
		WithWhere(filters.Where().
			WithPath([]string{"userID"}).
			WithOperator(filters.Equal).
			WithValueText(userId),
		).
		WithSort(
			graphql.Sort{
				Path:  []string{"timestamp"},
				Order: graphql.Desc,
			},
		).
		WithLimit(90).
		Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("GraphQL query failed: %w", err)
	}

	// Optional: Log or inspect raw data
	if raw, err := json.Marshal(response.Data); err == nil {
		fmt.Println("GraphQL Raw Response:", string(raw))
	}

	return response, nil
}
